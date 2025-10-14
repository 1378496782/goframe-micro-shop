package bargain_check

import (
	"context"
	"math/rand"
	"time"

	"fmt"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// todo
// 用于计算帮砍时砍下的金额
// 业务逻辑，根据传入的goods_id 去查询相关price和bargain_price
// 再去获取counts 数量 三个参数计算，随机出一个本次帮砍数值

// 思路 拿 Bh表 historyReq的bargainid 作为id去查 Bi表的goodsid并获取counts数据
func CountsCal(ctx context.Context, bargainId int32) (goodsID int32, counts int32, err error) {
	infoError := consts.InfoError(consts.BargainHistoryInfo, consts.CreateFail)
	if bargainId <= 0 {
		return 0, 0, fmt.Errorf("参数非法")
	}

	//开始查询
	query := dao.BargainInfo.Ctx(ctx).
		Where("id", bargainId)

	record, err := query.One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return 0, 0, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	goods_ID := record["goods_id"].Int32()
	count := record["counts"].Int32()

	if record["goods_id"].IsNil() || record["counts"].IsNil() || goods_ID <= 0 || count <= 0 {
		return 0, 0, fmt.Errorf("查询为空 请检查")

	}

	return goods_ID, count, nil
}

// 拿到goodsid 跨表查询 去goods_info表拿price与bargain_price 计算差价
func Price_Set(ctx context.Context, goodsid int32) (diffprice int, err error) {

	//创建查询
	query := dao.GoodsInfo.Ctx(ctx).Where("id", goodsid)
	record, err := query.One()

	origin_price := record["price"].Int()
	bargain_price := record["bargain_price"].Int()

	if record["price"].IsNil() || record["bargain_price"].IsNil() || origin_price <= 0 || bargain_price <= 0 || origin_price <= bargain_price {
		return 0, fmt.Errorf("价格设置异常，请检查商品信息表")

	}
	diff_price := origin_price - bargain_price

	if diff_price <= 0 {
		return 0, fmt.Errorf("价格设置异常，请检查商品信息表")
	}

	return diff_price, nil
}

// 拿到以上数据，进行计算 随机一个数值 计算结果进行一个强制转换
func Range_amount(diffprice int, count int) (amount int, err error) {
	if diffprice <= 0 || count <= 0 {
		return 0, fmt.Errorf("参数必须大于0")
	}
	var max_discount int
	min_discount := 10                     //最低有效帮砍优惠10
	max_discount = diffprice / (count + 2) //控制最大优惠，避免真砍到底线优惠 增加计算工作与逻辑

	remainder := diffprice % (count + 2)

	// 如果有余数，向上取整
	if remainder > 0 {
		max_discount++
	}
	if max_discount < min_discount {
		max_discount = min_discount
	} //防止出现负数

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(max_discount-min_discount+1) + min_discount

	return randomNum, nil
}

//为了方便计算 以防万一 加一个统计已砍下的优惠 todo

//为了方便订单系统 同时结算优惠卷和砍价，再加入一个计算当前已砍价格的方法，实现在"logic\bargain_help_check"
//todo 为了增加技术性 考虑加入消息队列
//帮砍的操作进队列 队列消费
