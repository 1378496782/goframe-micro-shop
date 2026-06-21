package comment_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/logic/comment_info"
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

// GetList 列表
func (*Controller) GetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
	return comment_info.GetList(ctx, req)
}

// Create 创建
func (*Controller) Create(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	return comment_info.Create(ctx, req)
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
