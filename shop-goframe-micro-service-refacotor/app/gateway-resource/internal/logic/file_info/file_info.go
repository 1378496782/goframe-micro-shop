package file_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/internal/dao"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/internal/model/do"
	"shop-goframe-micro-service-refacotor/app/gateway-resource/internal/model/entity"
)

func UploadImage(ctx context.Context, url, fileName string, req *entity.FileInfo) (err error) {
	// 创建DO对象
	fileRecord := &do.FileInfo{
		Name:         fileName,
		Url:          url,
		UploaderId:   req.UploaderId,
		UploaderType: req.UploaderType,
		FileType:     req.FileType,
	}
	_, err = dao.FileInfo.Ctx(ctx).Insert(fileRecord)
	if err != nil {
		return gerror.Wrap(err, "创建文件记录失败")
	}
	return nil
}
