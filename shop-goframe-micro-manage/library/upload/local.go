package upload

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/tiger1103/gfast/v3/internal/app/common/consts"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Local struct {
}

func (s *Local) Upload(ctx context.Context, file *ghttp.UploadFile) (result *model.UploadResponse, err error) {
	if file == nil {
		err = errors.New("文件必须!")
		return
	}
	urlPerfix := libUtils.GetDomain(ctx, true)
	p := strings.Trim(consts.UploadPath, "/")
	sp := s.getStaticPath(ctx)
	if sp != "" {
		sp = strings.TrimRight(sp, "/")
	}
	nowData := time.Now().Format("2006-01-02")
	// 包含静态文件夹的路径
	fullDirPath := sp + "/" + p + "/" + nowData
	fileName, err := file.Save(fullDirPath, true)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath := p + "/" + nowData + "/" + fileName
	result = &model.UploadResponse{
		Size:     file.Size,
		Path:     fullPath,
		FullPath: urlPerfix + "/" + fullPath,
		Name:     file.Filename,
		Type:     file.Header.Get("Content-type"),
	}
	return
}

// 静态文件夹目录
func (s *Local) getStaticPath(ctx context.Context) string {
	value, _ := g.Cfg().Get(ctx, "server.serverRoot")
	if !value.IsEmpty() {
		return value.String()
	}
	return ""
}

func (s *Local) CheckMultipart(ctx context.Context, req *model.CheckMultipartReq) (res *model.CheckMultipartRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		res = new(model.CheckMultipartRes)
		//计算分片数
		for i := 0; i < req.ShardsCount; i++ {
			res.WaitUploadIndex = append(res.WaitUploadIndex, i+1)
		}
		var progress *model.MultipartProgress
		progress, err = s.GetMultipartProgress(ctx, req)
		liberr.ErrIsNil(ctx, err)
		if progress != nil && len(progress.UploadedIndex) > 0 {
			res.WaitUploadIndex = libUtils.DiffSlice(progress.UploadedIndex, res.WaitUploadIndex)
		}
	})
	return
}

func (s *Local) UploadPart(ctx context.Context, req *model.UploadPartReq) (res *model.UploadPartRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		pq := &model.CheckMultipartReq{
			UploadType:  req.UploadType,
			FileName:    req.FileName,
			Size:        req.Size,
			Md5:         req.Md5,
			ShardsCount: req.ShardsCount,
			DriverType:  req.DriverType,
			UserId:      req.UserId,
			AppId:       req.AppId,
		}
		var progress *model.MultipartProgress
		progress, err = s.GetMultipartProgress(ctx, pq)
		liberr.ErrIsNil(ctx, err)
		if progress == nil {
			liberr.ErrIsNil(ctx, errors.New("分片上传信息不存在"))
		}
		for _, i := range progress.UploadedIndex {
			if req.Index == i {
				liberr.ErrIsNil(ctx, errors.New("分片已上传"))
			}
		}
		uploadId := s.GenUploadId(req.CheckMultipartReq)
		fullPath := s.getPartPath(ctx, uploadId)
		//保存分片
		req.File.Filename = gconv.String(req.Index) + ".part"
		_, err = req.File.Save(fullPath)
		liberr.ErrIsNil(ctx, err, "写入分片文件内容失败")

		res = new(model.UploadPartRes)
		//已上传完毕
		if req.ShardsCount == req.Index {
			res.Finish = true
			//合并文件
			sp := s.getStaticPath(ctx)
			if sp != "" {
				sp = strings.TrimRight(sp, "/")
			}
			nowData := time.Now().Format("2006-01-02")
			// 包含静态文件夹的路径
			sp = sp + "/" + strings.Trim(consts.UploadPath, "/") + "/" + nowData
			fileName := s.GenNewFileName(req.FileName)
			err = s.MergerPath(fullPath, sp+"/"+fileName)
			liberr.ErrIsNil(ctx, err, "合并分片失败")
			//删除临时文件
			gfile.Remove(fullPath)
			path := gstr.SubStr(sp, strings.Index(sp, "/"+consts.UploadPath+"/")+1)

			res.Attachment = &model.UploadResponse{
				Size:     req.Size,
				Path:     path + "/" + fileName,
				FullPath: libUtils.GetDomain(ctx, true) + "/" + path + "/" + fileName,
				Name:     req.FileName,
				Type:     req.File.FileHeader.Header.Get("Content-type"),
			}
		}
	})
	return
}
func (s *Local) GenNewFileName(oldFileName string) string {
	ext := gfile.Ext(oldFileName)
	fileName := strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6)
	return strings.ToLower(fileName + ext)
}
func (s *Local) MergerPath(src string, dst string) error {
	parts, err := gfile.ScanDirFile(src, "*.part")
	if err != nil {
		return err
	}
	sort.Slice(parts, func(i, j int) bool {
		fileI := gfile.Basename(parts[i])
		fileJ := gfile.Basename(parts[j])
		return gconv.Int(gstr.TrimRight(fileI, ".part")) < gconv.Int(gstr.TrimRight(fileJ, ".part"))
	})
	for _, file := range parts {
		if err = gfile.PutBytesAppend(dst, gfile.GetBytes(file)); err != nil {
			return err
		}
	}
	return nil
}

// GenUploadId 生成上传ID
func (s *Local) GenUploadId(req *model.CheckMultipartReq) string {
	return fmt.Sprintf("%s_%d_%s_%d", req.Md5, req.UserId, req.AppId, req.DriverType)
}

func (s *Local) getPartPath(ctx context.Context, uploadId string) string {
	p := strings.Trim(consts.UploadPath, "/")
	sp := s.getStaticPath(ctx)
	if sp != "" {
		sp = strings.TrimRight(sp, "/")
	}
	// 包含静态文件夹的路径
	return sp + "/" + p + "/tmp/" + uploadId
}

// GetMultipartProgress 获取或创建分片上传事件进度
func (s *Local) GetMultipartProgress(ctx context.Context, req *model.CheckMultipartReq) (res *model.MultipartProgress, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		uploadId := s.GenUploadId(req)
		fullDirPath := s.getPartPath(ctx, uploadId)
		res = &model.MultipartProgress{
			UploadId:      uploadId,
			CreatedAt:     gtime.Now(),
			ShardCount:    req.ShardsCount,
			UploadedIndex: []int{},
		}
		//路径不存在说明不存在分片信息
		if !gfile.Exists(fullDirPath) {
			return
		}
		//读取路径下的分片
		var filePath []string
		filePath, err = gfile.ScanDirFile(fullDirPath, "*.part")
		liberr.ErrIsNil(ctx, err, "获取分片文件失败")
		for _, v := range filePath {
			index := gfile.Basename(v)
			index = strings.TrimSuffix(index, ".part")
			i := gconv.Int(index)
			res.UploadedIndex = append(res.UploadedIndex, i)
		}
	})
	return
}
