/*
* @desc:角色
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:11
 */

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type EffectiveTimeInfo struct {
	EffectiveType int           `json:"effectiveType"`
	WeekDay       []int         `json:"weekDay"`
	DayRange      []*gtime.Time `json:"dayRange"`
	DateRange     []*gtime.Time `json:"dateRange"`
}

type RoleInfoRes struct {
	*entity.SysRole
	*EffectiveTimeInfo
}
