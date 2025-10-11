package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImageReq struct {
	g.Meta       `path:"/upload/image" tags:"文件上传" method:"post" summary:"上传图片"`
	File         *ghttp.UploadFile `json:"-" dc:"上传的文件" v:"required#请选择上传文件"`
	UploaderId   uint              `json:"uploader_id" dc:"上传者ID（根据uploader_type区分是用户ID还是管理员ID" v:"required#上传者id不能为空"`
	UploaderType uint              `json:"uploader_type" dc:"上传者类型：1-H5用户，2-管理员" v:"required#上传者类型不能为空"`
	FileType     uint              `json:"file_type" dc:"文件类型：1-图片，2-视频，3-其他" v:"required#文件类型不能为空"`
}

type UploadImageRes struct {
	Key        string `json:"key" dc:"唯一文件名"`
	Url        string `json:"url" dc:"图片访问URL"`
	ExpireTime int64  `json:"expireTime" dc:"过期时间"`
}

type GetAvatarImageReq struct {
	g.Meta `path:"/image/url" tags:"获取文件 url" method:"get" summary:"获取文件 url"`
	Key    string `json:"key" v:"required#key 不能为空" dc:"文件唯一标识"`
}

type GetAvatarImageRes struct {
	Url        string `json:"url" dc:"图片访问URL"`
	ExpireTime int64  `json:"expireTime" dc:"过期时间"`
}
