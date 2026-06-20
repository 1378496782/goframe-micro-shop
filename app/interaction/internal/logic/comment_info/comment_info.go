package comment_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
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

	// 返回创建成功响应，包含新创建的ID
	return &v1.CommentInfoCreateRes{Id: uint32(result)}, nil
}

func GetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.CommentInfoListResponse{
		List:  make([]*pbentity.CommentInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.CommentInfo, consts.GetListFail)

	query := dao.CommentInfo.Ctx(ctx).
		Where(dao.CommentInfo.Columns().ObjectId, req.ObjectId).
		Where(dao.CommentInfo.Columns().Type, req.Type).
		Where(dao.CommentInfo.Columns().ParentId, req.ParentId).
		WhereNull(dao.CommentInfo.Columns().DeletedAt)

	// 查询总数
	total, err := query.Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	commentRecords, err := query.
		Page(int(req.Page), int(req.Size)).
		OrderDesc(dao.CommentInfo.Columns().Id).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	parentIds := make([]int, 0, len(commentRecords))
	for _, record := range commentRecords {
		var comment entity.CommentInfo
		if err := record.Struct(&comment); err != nil {
			continue
		}

		pbComment := convertCommentEntity(comment)
		response.List = append(response.List, pbComment)
		if req.ParentId == 0 {
			parentIds = append(parentIds, comment.Id)
		}
	}

	if req.ParentId == 0 && len(parentIds) > 0 {
		replyRecords, err := dao.CommentInfo.Ctx(ctx).
			Where(dao.CommentInfo.Columns().ObjectId, req.ObjectId).
			Where(dao.CommentInfo.Columns().Type, req.Type).
			WhereIn(dao.CommentInfo.Columns().ParentId, parentIds).
			WhereNull(dao.CommentInfo.Columns().DeletedAt).
			OrderAsc(dao.CommentInfo.Columns().Id).
			All()
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}

		replyMap := make(map[int32][]*pbentity.CommentInfo)
		for _, record := range replyRecords {
			var reply entity.CommentInfo
			if err := record.Struct(&reply); err != nil {
				continue
			}

			pbReply := convertCommentEntity(reply)
			replyMap[pbReply.ParentId] = append(replyMap[pbReply.ParentId], pbReply)
		}

		for _, comment := range response.List {
			if replies, ok := replyMap[comment.Id]; ok {
				comment.Replies = replies
			}
		}
	}

	return &v1.CommentInfoGetListRes{Data: response}, nil
}

func convertCommentEntity(comment entity.CommentInfo) *pbentity.CommentInfo {
	return &pbentity.CommentInfo{
		Id:          int32(comment.Id),
		ParentId:    int32(comment.ParentId),
		RootId:      int32(comment.RootId),
		UserId:      int32(comment.UserId),
		ObjectId:    int32(comment.ObjectId),
		Type:        int32(comment.Type),
		ReplyUserId: int32(comment.ReplyUserId),
		Content:     comment.Content,
		CreatedAt:   utility.SafeConvertTime(comment.CreatedAt),
		UpdatedAt:   utility.SafeConvertTime(comment.UpdatedAt),
		DeletedAt:   utility.SafeConvertTime(comment.DeletedAt),
		Replies:     make([]*pbentity.CommentInfo, 0),
	}
}
