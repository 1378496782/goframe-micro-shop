package file

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"shop-goframe-micro-service-refacotor/app/gateway-admin/internal/logic/file_info"
	"shop-goframe-micro-service-refacotor/app/gateway-admin/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"shop-goframe-micro-service-refacotor/utility/middleware"

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
	url, fileName, err := utility.UploadToQiniu(ctx, fileContent, req.File.Filename)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "上传到七牛云失败")
	}

	// 5. 从上下文中获取上传者ID
	userId := ctx.Value(middleware.CtxUserId)
	if userId == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户ID")
	}

	// 错误类型
	infoError := consts.InfoError(consts.FileInfo, consts.UploadImageFail)
	err = file_info.UploadImage(ctx, url, fileName, userId.(int))
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 5. 返回结果
	return &v1.UploadImageRes{Url: url}, nil
}
