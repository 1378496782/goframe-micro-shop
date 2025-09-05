/*
* @desc:上传文件model
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/12/6 15:29
 */

package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type UpFile struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	FileType string `json:"fileType"`
	Size     uint64 `json:"size"`
}

type UploadResponse struct {
	Size     int64  `json:"size"   dc:"文件大小"`
	Path     string `json:"path" dc:"文件相对路径"`
	FullPath string `json:"fullPath" dc:"文件绝对路径"`
	Name     string `json:"name" dc:"文件名称"`
	Type     string `json:"type" dc:"文件类型"`
}

// CheckMultipartReq 检查文件分片
type CheckMultipartReq struct {
	UploadType  string `p:"uploadType"  dc:"文件类型"`
	FileName    string `p:"fileName"    dc:"文件名称"`
	Size        int64  `p:"size"        dc:"文件大小"`
	Md5         string `p:"md5"         dc:"文件md5值"`
	ShardsCount int    `p:"shardsCount"  dc:"分片数量"`
	DriverType  int    //上传驱动
	UserId      uint64
	AppId       string
}

type CheckMultipartRes struct {
	Attachment      *UploadResponse `json:"attachment"      dc:"附件"`
	WaitUploadIndex []int           `json:"waitUploadIndex" dc:"等待上传的分片索引"`
	Progress        float64         `json:"progress"        dc:"上传进度"`
}

// UploadPartReq 分片上传
type UploadPartReq struct {
	*CheckMultipartReq
	Index int               `p:"index"       dc:"分片索引"`
	File  *ghttp.UploadFile `p:"file" type:"file" dc:"分片文件"`
}

type UploadPartRes struct {
	Attachment *UploadResponse `json:"attachment" dc:"附件"`
	Finish     bool            `json:"finish"     dc:"是否完成"`
}

// MultipartProgress 分片进度
type MultipartProgress struct {
	UploadId      string      `json:"uploadId"`      // 上传事件ID
	ShardCount    int         `json:"shardCount"`    // 分片数量
	UploadedIndex []int       `json:"uploadedIndex"` // 已上传的分片索引
	CreatedAt     *gtime.Time `json:"createdAt"`     // 创建时间
}
