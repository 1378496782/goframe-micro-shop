package bargain_check

import (
	"context"
	"math/rand"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// todo 重写随机生成时间的逻辑
// 设置一个本文件内的全局变量，方便传参
var current_time = time.Now() // time.Time

// 生成随机时间函数
func GenerateRandomTime(etime time.Time, num uint) time.Time {

	var random_time time.Time
	random := rand.Intn(int(num)) + 10 //生成一个随机数 在10到10+num之间

	random_time = etime.Add(time.Duration(random) * time.Hour) // 利用time

	return random_time
}

// 支持事务的版本 获取一个时间变量，更新到bi表的update_time上 加入事务处理
func Bi_Update_async_TX(ctx context.Context, tx gdb.TX, time1 time.Time, id int) (result bool, err error) {
	//根据id 查询bargain_info表的相关数据，并将time1更新到bargain_info表的update_time字段上
	_, err = dao.BargainInfo.Ctx(ctx).TX(tx).
		Where("id", id).
		Update(g.Map{"updated_time": time1})

	if err != nil {
		return false, err
	}
	return true, nil
}
