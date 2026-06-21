package comment_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/interaction/utility/interactionRedis"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func Create(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.CommentInfo, consts.CreateFail)

	rootId := 0
	replyUserId := 0
	if req.ParentId > 0 {
		parent := entity.CommentInfo{}
		err := dao.CommentInfo.Ctx(ctx).
			Where(dao.CommentInfo.Columns().Id, req.ParentId).
			Where(dao.CommentInfo.Columns().ObjectId, req.ObjectId).
			Where(dao.CommentInfo.Columns().Type, req.Type).
			WhereNull(dao.CommentInfo.Columns().DeletedAt).
			Scan(&parent)
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
		if parent.Id == 0 {
			return nil, gerror.NewCode(gcode.CodeValidationFailed, "父级评论不存在")
		}
		if parent.RootId == 0 {
			rootId = parent.Id
		} else {
			rootId = parent.RootId
		}
		replyUserId = parent.UserId
	}

	result, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(g.Map{
		"object_id":     req.ObjectId,
		"type":          req.Type,
		"parent_id":     req.ParentId,
		"content":       req.Content,
		"root_id":       rootId,
		"reply_user_id": replyUserId,
		"user_id":       req.UserID,
	})
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 若他不是根评论，增加回复数
	if req.ParentId > 0 {
		_, err = dao.CommentInfo.Ctx(ctx).
			Where(dao.CommentInfo.Columns().Id, rootId).
			Increment(dao.CommentInfo.Columns().ReplyCount, 1)
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
	}
	// 返回创建成功响应，包含新创建的ID
	return &v1.CommentInfoCreateRes{Id: uint32(result)}, nil
}

// GetList 分页查询某个对象（如商品/文章）的评论列表。
//
// 评论展示为两段：顶层评论（parent_id=0）和它们下面的回复列表。
// 当请求查询顶层评论（req.ParentId==0）时，本方法会额外把这一页每条顶层评论
// 的所有后代回复一并查出并挂到 Replies 字段；查询某条评论下的回复时（req.ParentId!=0）
// 则只返回该层数据，不再向下展开。
//
// 回复采用「一次性批量查询 + 内存分组」的方式挂载，避免对每条顶层评论各查一次（N+1）。
func GetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
	// 初始化响应结构：List 预置为空切片，避免无数据时返回 null
	response := &v1.CommentInfoListResponse{
		List:  make([]*pbentity.CommentInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.CommentInfo, consts.GetListFail)

	// 基础查询条件：限定对象、评论类型、父级（parent_id），并排除软删除的记录。
	// req.ParentId==0 表示查顶层评论，非 0 表示查某条评论下的回复。
	query := dao.CommentInfo.Ctx(ctx).
		Where(dao.CommentInfo.Columns().ObjectId, req.ObjectId).
		Where(dao.CommentInfo.Columns().Type, req.Type).
		Where(dao.CommentInfo.Columns().ParentId, req.ParentId).
		WhereNull(dao.CommentInfo.Columns().DeletedAt)

	// 查询总数（用于前端分页，反映满足条件的全部条数，与本页返回条数无关）
	total, err := query.Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据，按 id 倒序（最新的评论排在前面）
	commentRecords, err := query.
		Page(int(req.Page), int(req.Size)).
		OrderDesc(dao.CommentInfo.Columns().Id).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 逐条转换为 pb 结构填入响应；同时收集本页顶层评论的 id，供后续批量查回复用。
	rootIds := make([]int, 0, len(commentRecords))
	for _, record := range commentRecords {
		var comment entity.CommentInfo
		// 单条转换失败只跳过该条，不中断整批
		if err := record.Struct(&comment); err != nil {
			continue
		}

		pbComment := convertCommentEntity(ctx, comment)
		response.List = append(response.List, pbComment)
		// 只有在查顶层评论时才需要继续查它们的回复
		if req.ParentId == 0 {
			rootIds = append(rootIds, comment.Id)
		}
	}

	// 查顶层评论时，批量加载这一页所有顶层评论的回复并挂载到对应评论上
	if req.ParentId == 0 && len(rootIds) > 0 {
		// 用 root_id IN (...) 一次查出本页所有顶层评论的回复，避免 N+1 查询。
		// 回复按 id 升序（同一条评论下按时间正序展示）。
		replyRecords, err := dao.CommentInfo.Ctx(ctx).
			Where(dao.CommentInfo.Columns().ObjectId, req.ObjectId).
			Where(dao.CommentInfo.Columns().Type, req.Type).
			WhereIn(dao.CommentInfo.Columns().RootId, rootIds).
			WhereNull(dao.CommentInfo.Columns().DeletedAt).
			OrderAsc(dao.CommentInfo.Columns().Id).
			All()
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}

		// 按 root_id 把回复分组到 map：root_id -> 该顶层评论的回复列表
		replyMap := make(map[int32][]*pbentity.CommentInfo)
		for _, record := range replyRecords {
			var reply entity.CommentInfo
			if err := record.Struct(&reply); err != nil {
				continue
			}

			pbReply := convertCommentEntity(ctx, reply)
			replyMap[pbReply.RootId] = append(replyMap[pbReply.RootId], pbReply)
		}

		// 把分组好的回复回填到对应顶层评论的 Replies 字段
		for _, comment := range response.List {
			if replies, ok := replyMap[comment.Id]; ok {
				comment.Replies = replies
			}
		}
	}

	return &v1.CommentInfoGetListRes{Data: response}, nil
}

func convertCommentEntity(ctx context.Context, comment entity.CommentInfo) *pbentity.CommentInfo {
	likeCount := uint32(comment.LikeCount)
	if cachedCount, ok, err := interactionRedis.GetCommentLikeCount(ctx, uint32(comment.Id)); err == nil && ok {
		likeCount = cachedCount
	} else if err == nil {
		_ = interactionRedis.SetCommentLikeCount(ctx, uint32(comment.Id), likeCount)
	} else {
		g.Log().Warningf(ctx, "读取评论点赞 Redis 计数失败: comment_id=%d, err=%v", comment.Id, err)
	}

	return &pbentity.CommentInfo{
		Id:          int32(comment.Id),
		ParentId:    int32(comment.ParentId),
		RootId:      int32(comment.RootId),
		UserId:      int32(comment.UserId),
		ObjectId:    int32(comment.ObjectId),
		Type:        int32(comment.Type),
		ReplyUserId: int32(comment.ReplyUserId),
		LikeCount:   int32(likeCount),
		ReplyCount:  int32(comment.ReplyCount),
		Content:     comment.Content,
		CreatedAt:   utility.SafeConvertTime(comment.CreatedAt),
		UpdatedAt:   utility.SafeConvertTime(comment.UpdatedAt),
		DeletedAt:   utility.SafeConvertTime(comment.DeletedAt),
		Replies:     make([]*pbentity.CommentInfo, 0),
	}
}
