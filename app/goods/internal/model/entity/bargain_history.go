// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BargainHistory is the golang structure for table bargain_history.
type BargainHistory struct {
	Id          int         `json:"id"          orm:"id"           description:""` //
	BargainId   int         `json:"bargainId"   orm:"bargain_id"   description:""` //
	Amount      int         `json:"amount"      orm:"amount"       description:""` //
	UserId      int         `json:"userId"      orm:"user_id"      description:""` //
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" description:""` //
	DeletedTime *gtime.Time `json:"deletedTime" orm:"deleted_time" description:""` //
}
