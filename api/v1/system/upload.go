package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// 单图上传
type UploadSingleImgReq struct {
	g.Meta `path:"/upload/singleImg" tags:"系统后台/后台文件上传" method:"post" summary:"上传图片"`
	File   *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件" v:"required#上传文件必须"`
}

// 单文件上传
type UploadSingleFileReq struct {
	g.Meta `path:"/upload/singleFile" tags:"系统后台/后台文件上传" method:"post" summary:"上传文件"`
	File   *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}

type UploadSingleRes struct {
	g.Meta `mime:"application/json"`
	*model.UploadResponse
}

// 多图上传
type UploadMultipleImgReq struct {
	g.Meta `path:"/upload/multipleImg" tags:"系统后台/后台文件上传" method:"post" summary:"上传多图片"`
	File   ghttp.UploadFiles `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}

// 多文件上传
type UploadMultipleFileReq struct {
	g.Meta `path:"/upload/multipleFile" tags:"系统后台/后台文件上传" method:"post" summary:"上传多文件"`
	File   ghttp.UploadFiles `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}

type UploadMultipleRes []*model.UploadResponse

type UploadResponse struct {
	*model.UploadResponse
}

type CheckMultipartReq struct {
	g.Meta `path:"/upload/checkMultipart" tags:"系统后台/后台文件上传" method:"post" summary:"检查分片"`
	common.Author
	*model.CheckMultipartReq
}

type CheckMultipartRes struct {
	g.Meta `mime:"application/json"`
	*model.CheckMultipartRes
}

type UploadPartReq struct {
	g.Meta `path:"/upload/uploadPart" tags:"系统后台/后台文件上传" method:"post" summary:"分片上传"`
	common.Author
	*model.UploadPartReq
}

type UploadPartRes struct {
	g.Meta `mime:"application/json"`
	*model.UploadPartRes
}
