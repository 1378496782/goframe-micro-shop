package bargain_info

import (
	"context"
	"fmt"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "shop-goframe-micro-service-refacotor/app/goods/api/bargain_info/v1"
	vlogic "shop-goframe-micro-service-refacotor/app/goods/internal/logic/bargain_check"
)

type Controller struct {
	v1.UnimplementedBargainInfoServer
}

// cmd注册
func Register(s *grpcx.GrpcServer) {
	v1.RegisterBargainInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.BargainInfoGetListReq) (res *v1.BargainInfoGetListRes, err error) {
<<<<<<< HEAD
	infoError := consts.InfoError(consts.BargainInfo, consts.GetDetailFile) //构建错误信息
=======
	infoError := consts.InfoError(consts.BargainInfo, consts.GetDetailFail) //构建错误信息
>>>>>>> d924eccf78c04f3d02c9def3e16d8c975f2b0fca

	//验证参数合法性
	if req.Id <= 0 || req.UserId <= 0 || req.GoodsId <= 0 {
		return nil, fmt.Errorf("参数非法")
	}

	// 构建查询条件
	query := dao.BargainInfo.Ctx(ctx).
		Where("id", req.Id).
		Where("user_id", req.UserId).
		Where("goods_id", req.GoodsId).
		Where("deleted_time", nil) //过滤deleted_time字段

	record, err := query.One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 检查记录是否存在
	if record.IsEmpty() {
		g.Log().Errorf(ctx, "%v 砍价信息不存在, ID: %d, UserId: %d, GoodsId: %d", infoError, req.Id, req.UserId, req.GoodsId)
		return nil, gerror.WrapCode(gcode.CodeNotFound, err, "砍价信息不存在")
	}

	// 获取时间并转换为google.protobuf.Timestamp
	createdTime := record["created_time"].Time()
	updatedTime := record["updated_time"].Time()
	expiredTime := record["expired_time"].Time()
	createdTimestamp := timestamppb.New(createdTime)
	updatedTimestamp := timestamppb.New(updatedTime)
	expiredTimestamp := timestamppb.New(expiredTime)

	// 构建响应数据
	res = &v1.BargainInfoGetListRes{
		Id:          record["id"].Int32(),
		UserId:      record["user_id"].Int32(),
		GoodsId:     record["goods_id"].Int32(),
		Counts:      record["counts"].Int32(),
		CreatedTime: createdTimestamp, // 需要根据实际字段类型转换
		UpdatedTime: updatedTimestamp, // 需要根据实际字段类型转换
		ExpiredTime: expiredTimestamp, // 需要根据实际字段类型转换
	}
	return res, nil
}

// 创建
func (*Controller) Create(ctx context.Context, req *v1.BargainInfoCreateReq) (res *v1.BargainInfoCreateRes, err error) {
	//错误类型
	infoError := consts.InfoError(consts.BargainInfo, consts.CreateFail)
	//获取当前时间
	nowtime := time.Now()
	//加入过期 time计算时间
	expiredTime := vlogic.GenerateRandomTime(nowtime, 10)

	//创建数据结构
	data := g.Map{
		"user_id":      req.UserID,
		"goods_id":     req.GoodsID,
		"counts":       req.Counts,
		"expired_time": expiredTime,
		"created_time": nowtime,
		"updated_time": nowtime,
	}

	result, err := dao.BargainInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil || result <= 0 {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 构建响应数据
	res = &v1.BargainInfoCreateRes{
		Id:          int32(result),                // 使用插入返回的ID
		UserId:      req.UserID,                   // 来自请求参数
		GoodsId:     req.GoodsID,                  // 来自请求参数
		Counts:      req.Counts,                   // 来自请求参数
		CreatedTime: timestamppb.New(nowtime),     // 转换创建时间 todo 目前只能转化为时间戳格式 等在h5网关层 再转化为标准的格式
		ExpiredTime: timestamppb.New(expiredTime), // 转换过期时间
	}

	return res, nil
}

// 删除
func (*Controller) Delete(ctx context.Context, req *v1.BargainInfoDeleteReq) (res *v1.BargainInfoDeleteRes, err error) {
	//先查询要删除的记录信息
	bargainInfo, err := dao.BargainInfo.Ctx(ctx).
		Fields("id", "user_id", "goods_id", "counts", "created_time").
		Where("user_id", req.UserId).
		Where("goods_id", req.GoodsId).
		Where("Id", req.Id).
		One()

	if err != nil {
<<<<<<< HEAD
		infoError := consts.InfoError(consts.BargainInfo, consts.GetDetailFile)
=======
		infoError := consts.InfoError(consts.BargainInfo, consts.GetDetailFail)
>>>>>>> d924eccf78c04f3d02c9def3e16d8c975f2b0fca
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	if bargainInfo.IsEmpty() {
		infoError := consts.InfoError(consts.BargainInfo, consts.Empty)
		g.Log().Error(ctx, infoError)
		return nil, gerror.NewCode(gcode.CodeNotFound, infoError)
	}

	now := time.Now()
	//根据传入id 删除对应数据
	_, err = dao.BargainInfo.Ctx(ctx).
		Where("user_id", req.UserId).
		Where("goods_id", req.GoodsId).
		Where("Id", req.Id).
		Data(g.Map{"deleted_time": now}).Update()

	//错误类型
	infoError := consts.InfoError(consts.BargainInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 构建并返回完整的响应对象
	res = &v1.BargainInfoDeleteRes{
		Id:      bargainInfo["id"].Int32(),
		GoodsId: bargainInfo["goods_id"].Int32(),
		UserId:  bargainInfo["user_id"].Int32(),
		CreatedTime: &timestamppb.Timestamp{
			Seconds: bargainInfo["created_time"].Time().Unix(),
			Nanos:   int32(bargainInfo["created_time"].Time().Nanosecond()),
		},
		DeletedTime: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	return res, nil
}
