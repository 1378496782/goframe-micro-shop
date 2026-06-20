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

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedCommentInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCommentInfoServer(s.Server, &Controller{})
}

func convertCommentEntity(comment entity.CommentInfo) *pbentity.CommentInfo {
	return &pbentity.CommentInfo{
		Id:        int32(comment.Id),
		ParentId:  int32(comment.ParentId),
		UserId:    int32(comment.UserId),
		ObjectId:  int32(comment.ObjectId),
		Type:      int32(comment.Type),
		Content:   comment.Content,
		CreatedAt: utility.SafeConvertTime(comment.CreatedAt),
		UpdatedAt: utility.SafeConvertTime(comment.UpdatedAt),
		DeletedAt: utility.SafeConvertTime(comment.DeletedAt),
		Replies:   make([]*pbentity.CommentInfo, 0),
	}
}

// GetList 列表
func (*Controller) GetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
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

// Create 创建
func (*Controller) Create(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CommentInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.CommentInfoCreateRes{Id: uint32(result)}, nil
}

// Delete 删除
func (*Controller) Delete(ctx context.Context, req *v1.CommentInfoDeleteReq) (res *v1.CommentInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.CommentInfo.Ctx(ctx).Where(g.Map{"id": req.Id, "user_id": req.UserId}).Delete()
	infoError := consts.InfoError(consts.CommentInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.CommentInfoDeleteRes{Id: req.Id}, nil // 返回空结构体
}
