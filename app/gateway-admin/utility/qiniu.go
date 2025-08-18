package utility

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// UploadToQiniu 上传文件到七牛云
func UploadToQiniu(ctx context.Context, fileContent []byte, filename string) (string, error) {
	// 读取配置
	cfg := g.Cfg().MustGet(ctx, "qiniu")
	if cfg.IsEmpty() {
		return "", errors.New("七牛云配置缺失")
	}

	// 解析配置
	qiniuCfg := cfg.Map()
	accessKey := qiniuCfg["accessKey"].(string)
	secretKey := qiniuCfg["secretKey"].(string)
	bucket := qiniuCfg["bucket"].(string)
	domain := qiniuCfg["domain"].(string)

	// 生成上传凭证
	putPolicy := storage.PutPolicy{Scope: bucket}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 配置上传参数（华南区）
	cfgUpload := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      true,
		UseCdnDomains: true,
	}

	// 获取文件扩展名
	fileExt := gfile.ExtName(filename)
	if fileExt == "" {
		fileExt = "jpg" // 默认扩展名
	}

	// 生成保留原始文件名的唯一文件名
	key := generateUniqueFilename(filename)

	// 创建表单上传器
	formUploader := storage.NewFormUploader(&cfgUpload)
	ret := storage.PutRet{}

	// 上传文件 - 使用标准库的bytes.Reader
	err := formUploader.Put(
		context.Background(),
		&ret,
		upToken,
		key,
		bytes.NewReader(fileContent), // 使用标准库的bytes.Reader
		int64(len(fileContent)),
		nil,
	)

	if err != nil {
		return "", err
	}

	// 返回完整访问URL
	return domain + "/" + key, nil
}

// generateUniqueFilename 生成保留原始文件名的唯一文件名
func generateUniqueFilename(originalName string) string {
	// 获取文件扩展名
	ext := gfile.ExtName(originalName)
	if ext == "" {
		ext = "jpg" // 默认扩展名
	}

	// 获取原始文件名（不含扩展名）
	baseName := gfile.Name(originalName)
	if baseName == "" {
		baseName = "file"
	}

	// 添加时间戳和随机数确保唯一性
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS格式
	random := grand.S(4)                             // 4位随机字符串

	// 清理文件名中的特殊字符
	safeName := cleanFilename(baseName)

	// 组合唯一文件名
	return safeName + "_" + timestamp + "_" + random + "." + ext
}

// cleanFilename 清理文件名中的特殊字符
func cleanFilename(name string) string {
	// 替换空格为下划线
	name = strings.ReplaceAll(name, " ", "_")

	// 移除非法字符
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		name = strings.ReplaceAll(name, char, "")
	}

	return name
}
