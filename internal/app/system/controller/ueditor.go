/*
* @desc:UEditor处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/11/17 14:40
 */

package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonConsts "github.com/tiger1103/gfast/v3/internal/app/common/consts"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"math"
)

var UEditor = new(uEditorController)

type uEditorController struct {
	BaseController
}

func (c *uEditorController) GetConfig(ctx context.Context, req *system.UEditorConfigReq) (res *system.UEditorRes, err error) {
	return c.action(ctx, &req.UEditorReq)
}

func (c *uEditorController) UpFile(ctx context.Context, req *system.UEditorUpFileReq) (res *system.UEditorRes, err error) {
	if req.File == nil {
		err = gerror.New("请选择文件上传")
		return
	}
	return c.action(ctx, &req.UEditorReq)
}

func (c *uEditorController) action(ctx context.Context, req *system.UEditorReq) (res *system.UEditorRes, err error) {
	r := g.RequestFromCtx(ctx)
	switch req.Action {
	//获取配置
	case "config":
		c.ueditorConfig(ctx, req.Callback)
	//上传图片  上传涂鸦
	case "uploadimage", "uploadscrawl":
		c.uEditorUpload(ctx, req.File, "image")
	//上传视频  上传文件
	case "uploadvideo", "uploadfile":
		c.uEditorUpload(ctx, req.File, "file")
	//列出图片
	case "listimage":
		var (
			list  g.List
			total interface{} = 0
		)
		list, total, err = c.getFileList(ctx, req.Start, req.Size, "image")
		if err != nil {
			r.Response.WriteJson(g.Map{
				"state": err.Error(),
			})
			r.Exit()
		}
		r.Response.WriteJson(g.Map{
			"state": "SUCCESS",
			"total": total,
			"start": req.Start,
			"list":  list,
		})
	//列出文件
	case "listfile":
		var (
			list  g.List
			total interface{} = 0
		)
		list, total, err = c.getFileList(ctx, req.Start, req.Size, "")
		if err != nil {
			r.Response.WriteJson(g.Map{
				"state": err.Error(),
			})
			r.Exit()
		}
		r.Response.WriteJson(g.Map{
			"state": "SUCCESS",
			"total": total,
			"start": req.Start,
			"list":  list,
		})
	//抓取远端图片
	case "catchimage":
		r.Response.WriteJson(g.Map{
			"state": "抓取失败",
		})
	default:
		r.Response.WriteJson(g.Map{
			"state": "请求地址出错",
		})
	}
	r.Exit()
	return
}

// 获取ueditor配置
func (c *uEditorController) ueditorConfig(ctx context.Context, callback string) {
	config := g.Map{
		/* 上传图片配置项 */
		"imageActionName":     "uploadimage",                                             /* 执行上传图片的action名称 */
		"imageFieldName":      "upfile",                                                  /* 提交的图片表单名称 */
		"imageMaxSize":        2048000,                                                   /* 上传大小限制，单位B */
		"imageAllowFiles":     g.Slice{".png", ".jpg", ".jpeg", ".gif", ".bmp"},          /* 上传图片格式显示 */
		"imageCompressEnable": true,                                                      /* 是否压缩图片,默认是true */
		"imageCompressBorder": 1600,                                                      /* 图片压缩最长边限制 */
		"imageInsertAlign":    "none",                                                    /* 插入的图片浮动方式 */
		"imageUrlPrefix":      "",                                                        /* 图片访问路径前缀 */
		"imagePathFormat":     "/ueditor/php/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}", /* 上传保存路径,可以自定义保存路径和文件名格式 */
		/* {filename} 会替换成原文件名,配置这项需要注意中文乱码问题 */
		/* {rand:6} 会替换成随机数,后面的数字是随机数的位数 */
		/* {time} 会替换成时间戳 */
		/* {yyyy} 会替换成四位年份 */
		/* {yy} 会替换成两位年份 */
		/* {mm} 会替换成两位月份 */
		/* {dd} 会替换成两位日期 */
		/* {hh} 会替换成两位小时 */
		/* {ii} 会替换成两位分钟 */
		/* {ss} 会替换成两位秒 */
		/* 非法字符 \ : * ? " < > | */
		/* 具请体看线上文档: fex.baidu.com/ueditor/#use-format_upload_filename */

		/* 涂鸦图片上传配置项 */
		"scrawlActionName":  "uploadscrawl",                                            /* 执行上传涂鸦的action名称 */
		"scrawlFieldName":   "upfile",                                                  /* 提交的图片表单名称 */
		"scrawlPathFormat":  "/ueditor/php/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}", /* 上传保存路径,可以自定义保存路径和文件名格式 */
		"scrawlMaxSize":     2048000,                                                   /* 上传大小限制，单位B */
		"scrawlUrlPrefix":   "",                                                        /* 图片访问路径前缀 */
		"scrawlInsertAlign": "none",

		/* 截图工具上传 */
		"snapscreenActionName":  "uploadimage",                                             /* 执行上传截图的action名称 */
		"snapscreenPathFormat":  "/ueditor/php/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}", /* 上传保存路径,可以自定义保存路径和文件名格式 */
		"snapscreenUrlPrefix":   "",                                                        /* 图片访问路径前缀 */
		"snapscreenInsertAlign": "none",                                                    /* 插入的图片浮动方式 */

		/* 抓取远程图片配置 */
		"catcherLocalDomain": g.Slice{"127.0.0.1", "localhost", "img.baidu.com"},
		"catcherActionName":  "catchimage",                                              /* 执行抓取远程图片的action名称 */
		"catcherFieldName":   "source",                                                  /* 提交的图片列表表单名称 */
		"catcherPathFormat":  "/ueditor/php/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}", /* 上传保存路径,可以自定义保存路径和文件名格式 */
		"catcherUrlPrefix":   "",                                                        /* 图片访问路径前缀 */
		"catcherMaxSize":     2048000,                                                   /* 上传大小限制，单位B */
		"catcherAllowFiles":  g.Slice{".png", ".jpg", ".jpeg", ".gif", ".bmp"},          /* 抓取图片格式显示 */

		/* 上传视频配置 */
		"videoActionName": "uploadvideo",                                             /* 执行上传视频的action名称 */
		"videoFieldName":  "upfile",                                                  /* 提交的视频表单名称 */
		"videoPathFormat": "/ueditor/php/upload/video/{yyyy}{mm}{dd}/{time}{rand:6}", /* 上传保存路径,可以自定义保存路径和文件名格式 */
		"videoUrlPrefix":  "",                                                        /* 视频访问路径前缀 */
		"videoMaxSize":    1024000000,                                                /* 上传大小限制，单位B，默认1000MB */
		"videoAllowFiles": g.Slice{
			".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
			".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid"}, /* 上传视频格式显示 */

		/* 上传文件配置 */
		"fileActionName": "uploadfile",                                             /* controller里,执行上传视频的action名称 */
		"fileFieldName":  "upfile",                                                 /* 提交的文件表单名称 */
		"filePathFormat": "/ueditor/php/upload/file/{yyyy}{mm}{dd}/{time}{rand:6}", /* 上传保存路径,可以自定义保存路径和文件名格式 */
		"fileUrlPrefix":  "",                                                       /* 文件访问路径前缀 */
		"fileMaxSize":    51200000,                                                 /* 上传大小限制，单位B，默认50MB */
		"fileAllowFiles": g.Slice{
			".png", ".jpg", ".jpeg", ".gif", ".bmp",
			".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
			".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid",
			".rar", ".zip", ".tar", ".gz", ".7z", ".bz2", ".cab", ".iso",
			".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".xml",
		}, /* 上传文件格式显示 */

		/* 列出指定目录下的图片 */
		"imageManagerActionName":  "listimage",                                      /* 执行图片管理的action名称 */
		"imageManagerListPath":    "/ueditor/php/upload/image/",                     /* 指定要列出图片的目录 */
		"imageManagerListSize":    20,                                               /* 每次列出文件数量 */
		"imageManagerUrlPrefix":   "",                                               /* 图片访问路径前缀 */
		"imageManagerInsertAlign": "none",                                           /* 插入的图片浮动方式 */
		"imageManagerAllowFiles":  g.Slice{".png", ".jpg", ".jpeg", ".gif", ".bmp"}, /* 列出的文件类型 */

		/* 列出指定目录下的文件 */
		"fileManagerActionName": "listfile",                  /* 执行文件管理的action名称 */
		"fileManagerListPath":   "/ueditor/php/upload/file/", /* 指定要列出文件的目录 */
		"fileManagerUrlPrefix":  "",                          /* 文件访问路径前缀 */
		"fileManagerListSize":   20,                          /* 每次列出文件数量 */
		"fileManagerAllowFiles": g.Slice{
			".png", ".jpg", ".jpeg", ".gif", ".bmp",
			".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
			".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid",
			".rar", ".zip", ".tar", ".gz", ".7z", ".bz2", ".cab", ".iso",
			".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".xml",
		}, /* 列出的文件类型 */
	}
	r := g.RequestFromCtx(ctx)
	r.Response.Write(fmt.Sprintf("%s(%s)", callback, gconv.String(config)))
	r.Exit()
}

// ueditor上传图片
func (c *uEditorController) uEditorUpload(ctx context.Context, upFile *ghttp.UploadFile, fType string) {
	var (
		info *model.UploadResponse
		err  error
		r    = g.RequestFromCtx(ctx)
	)
	v, _ := g.Cfg().Get(ctx, "upload.default")
	if fType == "image" {
		info, err = commonService.Upload().UploadFile(ctx, upFile, commonConsts.CheckFileTypeImg, v.Int(), service.Context().Get(ctx).User.Id, consts.UploadAppId)
	} else if fType == "file" {
		info, err = commonService.Upload().UploadFile(ctx, upFile, commonConsts.CheckFileTypeFile, v.Int(), service.Context().Get(ctx).User.Id, consts.UploadAppId)
	}

	if err != nil {
		r.Response.WriteJson(g.Map{
			"state": err.Error(),
		})
	} else {
		r.Response.WriteJson(g.Map{
			"state":    "SUCCESS",
			"url":      info.FullPath,
			"title":    info.Name,
			"original": info.Name,
		})
	}
	r.Exit()
}

func (c *uEditorController) getFileList(ctx context.Context, start int, size int, kind string) (list g.List, total interface{}, err error) {
	var res *model.SysAttachmentSearchRes
	pageNum := gconv.Int(math.Ceil(gconv.Float64(start)/gconv.Float64(size)) + 1)
	res, err = commonService.SysAttachment().List(ctx, &model.SysAttachmentSearchReq{
		PageReq: model.PageReq{
			PageNum:  pageNum,
			PageSize: size,
		},
		Kind:   kind,
		Status: "true",
	})
	if err != nil {
		return
	}
	if res != nil {
		total = res.Total
		list = make(g.List, len(res.List))
		for k, v := range res.List {
			path := ""
			if v.Drive == commonConsts.UploadDriverLocal {
				path = libUtils.GetDomain(ctx, true) + "/" + v.Path
			} else {
				path = v.Path
			}
			list[k] = g.Map{
				"url":   path,
				"mtime": v.UpdatedAt.Timestamp(),
			}
		}

	}
	return
}
