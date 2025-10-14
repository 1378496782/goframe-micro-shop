// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BargainInfo is the golang structure for table bargain_info.
type BargainInfo struct {
	Id          int         `json:"id"          orm:"id"           description:""` //
	UserId      int         `json:"userId"      orm:"user_id"      description:""` //
	GoodsId     int         `json:"goodsId"     orm:"goods_id"     description:""` //
	Counts      int         `json:"counts"      orm:"counts"       description:""` //
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" description:""` //
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" description:""` //
	DeletedTime *gtime.Time `json:"deletedTime" orm:"deleted_time" description:""` //
	ExpiredTime *gtime.Time `json:"expiredTime" orm:"expired_time" description:""` //
}
