/*
* @desc:雪花ID生成
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/6/2 14:52
 */

package snowIDGen

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sony/sonyflake"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
)

var machineID uint16 = 1

func init() {
	service.RegisterSnowID(New())
}

func New() service.ISnowID {
	return &sSnowID{
		sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: gtime.NewFromStr("2010-05-01").Time,
			MachineID: GetMachineId,
		}),
	}
}

type sSnowID struct {
	*sonyflake.Sonyflake
}

func (s *sSnowID) GenID() (uint64, error) {
	return s.NextID()
}

func GetMachineId() (uint16, error) {
	return machineID, nil
}
