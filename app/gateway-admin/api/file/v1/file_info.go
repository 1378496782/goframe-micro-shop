package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImageReq struct {
	g.Meta `path:"/upload/image" tags:"文件上传" method:"post" summary:"上传图片"`
	File   *ghttp.UploadFile `json:"-" dc:"上传的文件" v:"required#请选择上传文件"` // 注意：实际文件数据不会在JSON中传输
}

type UploadImageRes struct {
	Url string `json:"url" dc:"图片访问URL"`
}
