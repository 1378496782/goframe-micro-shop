package comment_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

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
	// 初始化响应结构
	response := &v1.CommentInfoListResponse{
		List:  make([]*pbentity.CommentInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.CommentInfo, consts.GetListFail)
	// 查询总数
	total, err := dao.CommentInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	commentRecords, err := dao.CommentInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range commentRecords {
		var comment entity.CommentInfo
		if err := record.Struct(&comment); err != nil {
			continue
		}

		var pbComment pbentity.CommentInfo
		if err := gconv.Struct(comment, &pbComment); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbComment.CreatedAt = utility.SafeConvertTime(comment.CreatedAt)
		pbComment.UpdatedAt = utility.SafeConvertTime(comment.UpdatedAt)

		response.List = append(response.List, &pbComment)
	}

	return &v1.CommentInfoGetListRes{Data: response}, nil
}

// Create 创建
func (*Controller) Create(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	// 定义一个实体对象，用于接收转换后的请求数据
	var commentInfo *entity.CommentInfo
	// 将请求参数req转换为实体对象commentInfo
	if err := gconv.Struct(req, &commentInfo); err != nil {
		return nil, err
	}
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
	_, err = dao.CommentInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CommentInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.CommentInfoDeleteRes{}, nil // 返回空结构体
}
