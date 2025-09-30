package file

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/internal/logic/file_info"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"strings"

	"shop-goframe-micro-service-refacotor/app/gateway-resource/api/file/v1"
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

	// 处理 URL，去掉可能的换行或空格，保证前端可用
	url = strings.ReplaceAll(url, "\n", "")
	url = strings.TrimSpace(url)

	var fileData *entity.FileInfo
	// 将请求参数 req 转换为实体对象 fileData
	if err := gconv.Struct(req, &fileData); err != nil {
		return nil, err
	}

	// 错误类型
	infoError := consts.InfoError(consts.FileInfo, consts.UploadImageFail)
	err = file_info.UploadImage(ctx, url, fileName, fileData)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	fmt.Println("上传七牛云 URL:", url)

	// 5. 返回结果
	return &v1.UploadImageRes{Url: url}, nil
}
