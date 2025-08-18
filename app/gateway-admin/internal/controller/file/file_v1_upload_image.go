package file

import (
	"context"
	"io"
	"shop-goframe-micro-service-refacotor/app/gateway-admin/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/file/v1"
)

func (c *ControllerV1) UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error) {
	// 1. 获取上传文件
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择上传文件")
	}

	// 2. 打开上传文件
	file, err := req.File.Open()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "打开上传文件失败")
	}
	defer file.Close()

	// 3. 读取文件内容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "读取文件内容失败")
	}

	// 4. 上传到七牛云
	url, err := utility.UploadToQiniu(ctx, fileContent, req.File.Filename)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "上传到七牛云失败")
	}

	// 5. 返回结果
	return &v1.UploadImageRes{Url: url}, nil
}
