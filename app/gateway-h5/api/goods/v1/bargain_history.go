package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/gogf/gf/v2/frame/g"
)

//创建 查询 删除

type Bargain_history_CreateReq struct {
	g.Meta     `path:"/goods/bargain_history/Create" method:"post" tags:"帮砍信息管理" sm:"帮砍信息创建" `
	User_id    int32 `json:"user_Id"   v:"required" dc:"帮砍用户id"`
	Bargain_id int32 `json:"bargain_Id" d:"0" v:"required" dc:"帮砍的砍价信息id"` //对应bargain_info的id
}

type Bargain_history_CreateRes struct {
	Id          int32  `json:"Id"    dc:"帮砍信息id"`
	Bargain_id  int32  `json:"bargain_Id"  dc:"帮砍的砍价信息id"`
	Amount      int32  `json:"amount"  dc:"本次砍到的优惠额度"`
	User_id     int32  `json:"user_Id" dc:"帮砍用户id"`
	Create_time string `json:"create_time"  dc:"创建时间"`
}

type Bargain_history_GetReq struct {
	g.Meta  `path:"/goods/bargain_history/Get" method:"get" tags:"帮砍信息管理" sm:"帮砍信息查询" `
	Id      int32 `json:"Id"    dc:"帮砍信息id"`
	User_id int32 `json:"user_Id" dc:"帮砍用户id"`
}

type Bargain_history_GetRes struct {
	Id          int32  `json:"Id"    dc:"帮砍信息id"`
	Bargain_id  int32  `json:"bargain_Id"  dc:"帮砍的砍价信息id"`
	Amount      int32  `json:"amount"  dc:"本次砍到的优惠额度"`
	User_id     int32  `json:"user_Id" dc:"帮砍用户id"`
	Create_time string `json:"create_time"  dc:"创建时间"`
}

type Bargain_history_DeleteReq struct {
	g.Meta     `path:"/goods/bargain_history/Delete" method:"post" tags:"帮砍信息管理" sm:"帮砍信息删除" `
	Id         int32 `json:"Id"   v:"required" dc:"帮砍的信息id"`
	Bargain_id int32 `json:"bargain_Id"  dc:"帮砍的砍价信息id"`
	User_id    int32 `json:"user_Id" dc:"帮砍用户id"`
}

type Bargain_history_DeleteRes struct {
	Id          int32  `json:"Id"    dc:"帮砍信息id"`
	Bargain_id  int32  `json:"bargain_Id"  dc:"帮砍的砍价信息id"`
	User_id     int32  `json:"user_Id" dc:"帮砍用户id"`
	Create_time string `json:"create_time"  dc:"创建时间"`
	Delete_time string `json:"delete_time"  dc:"删除时间"`
}
