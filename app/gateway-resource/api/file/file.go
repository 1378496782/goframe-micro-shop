// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package file

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/gateway-resource/api/file/v1"
)

type IFileV1 interface {
	UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error)
	GetAvatarImage(ctx context.Context, req *v1.GetAvatarImageReq) (res *v1.GetAvatarImageRes, err error)
}
