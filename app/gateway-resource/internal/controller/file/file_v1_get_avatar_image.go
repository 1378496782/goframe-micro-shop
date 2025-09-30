package file

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"shop-goframe-micro-service-refacotor/app/gateway-resource/api/file/v1"
)

func (c *ControllerV1) GetAvatarImage(ctx context.Context, req *v1.GetAvatarImageReq) (res *v1.GetAvatarImageRes, err error) {
	url, err := utility.GetFileUrl(ctx, req.Key)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取文件 url 失败")
	}
	return &v1.GetAvatarImageRes{
		Url: url,
	}, nil
}
