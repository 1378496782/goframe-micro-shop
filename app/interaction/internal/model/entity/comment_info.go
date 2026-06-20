// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure for table comment_info.
type CommentInfo struct {
	Id          int         `json:"id"          orm:"id"            description:""`             //
	ParentId    int         `json:"parentId"    orm:"parent_id"     description:"父级评论id"`       // 父级评论id
	RootId      int         `json:"rootId"      orm:"root_id"       description:"一级评论id"`       // 一级评论id
	UserId      int         `json:"userId"      orm:"user_id"       description:""`             //
	ReplyUserId int         `json:"replyUserId" orm:"reply_user_id" description:"被回复用户id"`      // 被回复用户id
	ObjectId    int         `json:"objectId"    orm:"object_id"     description:""`             //
	Type        int         `json:"type"        orm:"type"          description:"评论类型：1商品 2文章"` // 评论类型：1商品 2文章
	Content     string      `json:"content"     orm:"content"       description:"评论内容"`         // 评论内容
	LikeCount   int         `json:"likeCount"   orm:"like_count"    description:"点赞数"`          // 点赞数
	ReplyCount  int         `json:"replyCount"  orm:"reply_count"   description:"回复数"`          // 回复数
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`             //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:""`             //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"    description:""`             //
}
