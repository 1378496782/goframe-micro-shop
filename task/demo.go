/*
* @desc:测试定时任务
* @company:云南省奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2021/7/16 15:52
 */

package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func Test1(ctx context.Context) {
	fmt.Println("无参测试")
	service.SysJobLog().Add(ctx, &do.SysJobLog{
		TargetName: "test1",
		CreatedAt:  gtime.Now(),
		Result:     "无参测试运行成功",
	})
}

func Test2(ctx context.Context) {
	//获取参数
	t := service.TaskList().GetByName("test2")
	if t == nil {
		return
	}
	for _, v := range t.Param {
		fmt.Printf("参数:%s;  ", v)
	}
	fmt.Println()
	service.SysJobLog().Add(ctx, &do.SysJobLog{
		TargetName: "test2",
		CreatedAt:  gtime.Now(),
		Result:     "有参测试运行成功",
	})
}
