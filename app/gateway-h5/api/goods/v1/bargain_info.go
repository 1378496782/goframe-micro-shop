package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/gogf/gf/v2/frame/g"
)

// 创建 查询 删除

// 时间类 为string类型 到时候 获取goods微服务返回的stamp类型 转换成传统的yyy-mmm-ddd 再换成字符串

type Bargain_info_CreateReq struct {
	g.Meta   `path:"/goods/bargain_info/Create" method:"post" tags:"砍价信息管理" sm:"砍价信息创建" `
	User_id  int32 `json:"user_Id"   v:"required" dc:"创建砍价用户id"`
	Goods_id int32 `json:"goods_Id" d:"0" v:"required" dc:"砍价商品id"`
	Counts   int32 `json:"counts" d:"0" v:"required" dc:"最大帮砍次数"`
} //三个入参默认值 全是0

type Bargain_info_CreateRes struct {
	Id          int32  `json:"Id"    dc:"砍价信息id"`
	User_id     int32  `json:"user_Id"  dc:"创建砍价用户id"`
	Goods_id    int32  `json:"goods_Id"  dc:"砍价商品id"`
	Counts      int32  `json:"counts"  dc:"最大帮砍次数"`
	Create_time string `json:"create_time"  dc:"创建时间"`
	Expire_time string `json:"expire_time"  dc:"过期时间"`
}

type Bargain_info_GetReq struct {
	g.Meta   `path:"/goods/bargain_info/Get" method:"get" tags:"砍价信息管理" sm:"砍价信息查询" `
	Id       int32 `json:"Id"   v:"required" dc:"砍价信息id"`
	User_id  int32 `json:"user_Id" d:"0"  v:"required" dc:"创建砍价用户id"`
	Goods_id int32 `json:"goods_Id" d:"0" v:"required" dc:"砍价商品id"`
}

type Bargain_info_GetRes struct {
	Id          int32  `json:"Id"    dc:"砍价信息id"`
	User_id     int32  `json:"user_Id"  dc:"创建砍价用户id"`
	Goods_id    int32  `json:"goods_Id"  dc:"砍价商品id"`
	Counts      int32  `json:"counts"  dc:"最大帮砍次数"`
	Create_time string `json:"create_time"  dc:"创建时间"`
	Update_time string `json:"update_time"  dc:"更新时间"`
	Expire_time string `json:"expire_time"  dc:"过期时间"`
}

type Bargain_info_DeleteReq struct {
	g.Meta   `path:"/goods/bargain_info/delete" method:"post" tags:"砍价信息管理" sm:"砍价信息删除" `
	Id       int32 `json:"Id"   v:"required" dc:"砍价信息id"`
	User_id  int32 `json:"user_Id"   v:"required" dc:"创建砍价信息用户id"`
	Goods_id int32 `json:"goods_Id"   v:"required" dc:"商品id"`
}

type Bargain_info_DeleteRes struct {
	Id          int32  `json:"Id"    dc:"砍价信息id"`
	User_id     int32  `json:"user_Id"  dc:"创建砍价用户id"`
	Create_time string `json:"create_time"  dc:"创建时间"`
	Delete_time string `json:"delete_time"  dc:"删除时间"`
}
