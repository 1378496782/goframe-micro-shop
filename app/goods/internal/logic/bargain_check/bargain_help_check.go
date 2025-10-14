package bargain_check

import (
	"context"
	_ "context"
	"fmt"
	_ "fmt"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

//用于检查本次帮砍能否帮忙 是否有效
//先检查 当前时间与deleted time比较 超了就直接返回
//如果当前时间不到deletedtime 再去查询计算全部的帮砍记录 确保能和counts比较
//counts 比较完 返回一个参数，确认可以进行后续的流程

// 先查询bargainInfo 确认时间上可以 并获取counts 逻辑到时候是在bargain_histroy调用bargain_info 参数需要转化
// 关键字 Bh的bargainid 与Bi 的id

func CheckBaseInfo(ctx context.Context, bargainID int32) (ok_time bool, count int32, err error) {
	infoError := consts.InfoError(consts.BargainInfo, consts.GetDetailFail)

	if bargainID <= 0 {
		return false, 0, fmt.Errorf("参数非法")
	}
	//使用unscoped 无视软删除
	query := dao.BargainInfo.Ctx(ctx).
		Where("id", bargainID)

	record, err := query.One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return false, 0, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 检查记录是否存在
	if record.IsEmpty() {
		g.Log().Errorf(ctx, "%v 砍价信息不存在, bargainid: %d", infoError, bargainID)
		return false, 0, gerror.WrapCode(gcode.CodeNotFound, err, "砍价信息不存在")
	}

	//转化 查询到的expiretime并进行比较
	current_time := time.Now()
	expttime := record["expired_time"].Time()
	is_before := current_time.Before(expttime)
	if !is_before {
		return false, 0, fmt.Errorf("该砍价信息已过期，请寻找下一个需要帮助的顾客")
	}

	// 转换查询到的counts字段
	current_counts := record["counts"].Int32()

	// 返回最终结果 时间是否正确、允许砍价几次
	return is_before, current_counts, nil
}

// 再检查全部有效的当前帮砍记录 确保砍价记录还能再增加
func CheckHelpInfo(ctx context.Context, bargainID int32) (ok_counts bool, err error) {
	truth, origin_counts, error1 := CheckBaseInfo(ctx, bargainID)
	if !truth || error1 != nil {
		return false, fmt.Errorf("当前砍价信息无效，帮砍无效 请另寻目标")
	}

	counts, error2 := dao.BargainHistory.Ctx(ctx).
		Where("bargain_id", bargainID).Count()

	if error2 != nil {
		return false, fmt.Errorf("当前查询失败，请检查数据库链接与配置")
	}

	if int32(counts) >= origin_counts {
		return false, fmt.Errorf("当前帮砍人数已到上限，请另寻目标")
	}

	return true, nil
}

//理论上应该有检查金额是否超出的计算，但是预想的砍价逻辑中，绝对不可能砍到底线 因此待定
