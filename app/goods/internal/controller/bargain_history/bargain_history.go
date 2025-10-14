package bargain_history

import (
	"context"
	"fmt"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/bargain_history/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	c1 "shop-goframe-micro-service-refacotor/app/goods/internal/logic/bargain_check"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Controller struct {
	v1.UnimplementedBargainHistoryServer
}

// CMD注册
func Register(s *grpcx.GrpcServer) {
	v1.RegisterBargainHistoryServer(s.Server, &Controller{})
}

// 查询
func (*Controller) GetList(ctx context.Context, req *v1.BargainHistoryGetListReq) (res *v1.BargainHistoryGetListRes, err error) {
	infoError := consts.InfoError(consts.BargainHistoryInfo, consts.GetListFail) //构建错误信息
	//验证参数合法性
	if req.UserId <= 0 || req.Id <= 0 {
		return nil, fmt.Errorf("参数非法")
	}

	query := dao.BargainHistory.Ctx(ctx).
		Where("user_id", req.UserId).
		Where("deleted_time", nil) // 只根据UserId查询用户的帮砍记录,并过滤deleted_time字段

	// 如果提供了具体的记录Id，则额外添加Id条件
	if req.Id > 0 {
		query = query.Where("id", req.Id)
	}

	record, err := query.One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 检查记录是否存在
	if record.IsEmpty() {
		g.Log().Errorf(ctx, "%v 帮砍价信息不存在, ID: %d, UserId: %d", infoError, req.Id, req.UserId)
		return nil, gerror.WrapCode(gcode.CodeNotFound, err, "帮砍价信息不存在")
	}

	// 获取时间并转换为google.protobuf.Timestamp
	createdTime := record["created_time"].Time()
	protoTimestamp := timestamppb.New(createdTime)

	res = &v1.BargainHistoryGetListRes{
		Id:          record["id"].Int32(),
		BargainId:   record["bargain_id"].Int32(),
		Amount:      record["amount"].Int32(),
		UserId:      record["user_id"].Int32(),
		CreatedTime: protoTimestamp, // 需要根据实际字段类型转换
	}

	return res, nil
}

// 创建
func (*Controller) Create(ctx context.Context, req *v1.BargainHistoryCreateReq) (res *v1.BargainHistoryCreateRes, err error) {
	//错误类型
	infoError := consts.InfoError(consts.BargainHistoryInfo, consts.CreateFail)

	//调用Logic逻辑 先检查 时间有没有到期
	bool1, _, err1 := c1.CheckBaseInfo(ctx, req.BargainId)
	if !bool1 || err1 != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err1, infoError)
	}

	//调用logic逻辑 检查 帮砍记录到没到上限
	bool2, err2 := c1.CheckHelpInfo(ctx, req.BargainId)
	if !bool2 || err2 != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err2, infoError)
	}

	//计算本次金额的逻辑 逻辑写在Logic\bargain_help_amount中
	core_id := req.BargainId
	core_goods, core_counts, err3 := c1.CountsCal(ctx, core_id) //调用Logic去获取goodid和counts次数
	if err3 != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err3, infoError)
	}

	diff_price, err4 := c1.Price_Set(ctx, core_goods) //调用logic 去查goodsinfo表 计算price
	if err4 != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err4, infoError)
	}
	amount, err5 := c1.Range_amount(diff_price, int(core_counts)) //调用Logic 最后去获取随机砍价金额
	if err5 != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err5, infoError)
	}

	// 修改Create方法中的数据映射部分
	createdTime := time.Now()
	data := g.Map{
		"id":           nil,
		"bargain_id":   req.BargainId, // 改为小写字段名
		"amount":       amount,        // 改为小写字段名
		"user_id":      req.UserId,    // 改为小写字段名
		"created_time": createdTime,
	}

	//事务内部的变量无法在事务外调用，提前声明result
	var result int64
	// 插入BH表数据 与更新Bi表updatetime 要开始开启事务，避免Bh表与Bi表只有单方面更新 导致数据不一致
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 在事务中插入BH表数据
		var errs1 error
		result, errs1 = dao.BargainHistory.Ctx(ctx).TX(tx).Data(data).InsertAndGetId()
		if errs1 != nil {
			return errs1
		}

		// 在事务中更新Bi表的updatetime时间
		result1, errs2 := c1.Bi_Update_async_TX(ctx, tx, createdTime, int(req.BargainId))
		if errs2 != nil || !result1 {
			return errs2
		}

		return nil
	})
	//检查事务是否执行成功
	if err != nil {
		g.Log().Errorf(ctx, "%v %v,事务执行失败", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	//事务结束

	return &v1.BargainHistoryCreateRes{
		Id:          int32(result),
		BargainId:   req.BargainId,                // 补充BargainId
		Amount:      int32(amount),                // 补充Amount
		UserId:      req.UserId,                   // 补充UserId
		CreatedTime: timestamppb.New(createdTime), // 补充CreatedTime
	}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.BargainHistoryDeleteReq) (res *v1.BargainHistoryDeleteRes, err error) {
	// 先查询要删除的记录信息
	historyInfo, err := dao.BargainHistory.Ctx(ctx).
		Fields("id", "bargain_id", "user_id", "created_time").
		Where("user_id", req.UserId).
		Where("bargain_id", req.BargainId).
		Where("Id", req.Id).
		Where("deleted_time", nil).
		One()
	if err != nil {
<<<<<<< HEAD
		infoError := consts.InfoError(consts.BargainHistoryInfo, consts.GetDetailFile)
=======
		infoError := consts.InfoError(consts.BargainHistoryInfo, consts.GetDetailFail)
>>>>>>> d924eccf78c04f3d02c9def3e16d8c975f2b0fca
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	if historyInfo.IsEmpty() {
		infoError := consts.InfoError(consts.BargainHistoryInfo, consts.Empty)
		g.Log().Error(ctx, infoError)
		return nil, gerror.NewCode(gcode.CodeNotFound, infoError)
	}

	now := time.Now()
	//根据传入id 从数据库中删除对应信息，即强制更新
	_, err = dao.BargainHistory.Ctx(ctx).
		Where("user_id", req.UserId).
		Where("bargain_id", req.BargainId).
		Where("Id", req.Id).
		Data(g.Map{"deleted_time": now}).Update()

	//错误类型
	infoError := consts.InfoError(consts.BargainHistoryInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 构建并返回完整的响应对象
	res = &v1.BargainHistoryDeleteRes{
		Id:        historyInfo["id"].Int32(),
		BargainId: historyInfo["bargain_id"].Int32(),
		UserId:    historyInfo["user_id"].Int32(),
		CreatedTime: &timestamppb.Timestamp{
			Seconds: historyInfo["created_time"].Time().Unix(),
			Nanos:   int32(historyInfo["created_time"].Time().Nanosecond()),
		},
		DeletedTime: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	return res, nil

}
