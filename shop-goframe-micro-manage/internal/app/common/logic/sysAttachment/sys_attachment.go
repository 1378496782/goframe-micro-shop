// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/logic/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package sysAttachment

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/common/dao"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

var fileKind = map[string]*gset.StrSet{
	//图片
	"image": gset.NewStrSetFrom([]string{
		"jpeg",
		"jpg",
		"png",
		"gif",
		"webp",
		"cr2",
		"tif",
		"bmp",
		"heif",
		"jxr",
		"psd",
		"ico",
		"dwg",
	}),
	//文档
	"doc": gset.NewStrSetFrom([]string{
		"doc",
		"docx",
		"dot",
		"xls",
		"xlt",
		"xlsx",
		"xltx",
		"ppt",
		"pptx",
		"pdf",
		"txt",
		"csv",
		"html",
		"xml",
		"pptm",
		"pot",
		"wpd",
		"md",
		"json",
		"yaml",
		"markdown",
		"asciidoc",
		"xsl",
		"wps",
		"sxi",
		"sti",
		"odp",
	}),
	//视频
	"video": gset.NewStrSetFrom([]string{
		"mp4",
		"m4v",
		"mkv",
		"webm",
		"mov",
		"avi",
		"wmv",
		"mpg",
		"flv",
		"3gp",
	}),
	//音频
	"audio": gset.NewStrSetFrom([]string{
		"mid",
		"mp3",
		"m4a",
		"ogg",
		"flac",
		"wav",
		"amr",
		"aac",
		"aiff",
	}),
	//压缩包
	"zip": gset.NewStrSetFrom([]string{
		"zip",
		"rar",
		"tar",
		"gz",
		"7z",
		"tar.gz",
	}),
	//其它
	"other": gset.NewStrSetFrom([]string{
		"dwf",
		"ics",
		"vcard",
		"apk",
		"ipa",
	}),
}

func init() {
	service.RegisterSysAttachment(New())
}

func New() service.ISysAttachment {
	return &sSysAttachment{}
}

type sSysAttachment struct{}
type AddHandler = func(ctx context.Context) (err error)

func (s *sSysAttachment) List(ctx context.Context, req *model.SysAttachmentSearchReq) (listRes *model.SysAttachmentSearchRes, err error) {
	listRes = new(model.SysAttachmentSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAttachment.Ctx(ctx).WithAll()
		if req.AppId != "" {
			m = m.Where(dao.SysAttachment.Columns().AppId+" = ?", req.AppId)
		}
		if req.Drive != "" {
			m = m.Where(dao.SysAttachment.Columns().Drive+" = ?", req.Drive)
		}
		if req.Name != "" {
			m = m.Where(dao.SysAttachment.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Kind != "" {
			m = m.Where(dao.SysAttachment.Columns().Kind+" = ?", req.Kind)
		}
		if req.MimeType != "" {
			m = m.Where(dao.SysAttachment.Columns().MimeType+" like ?", "%"+req.MimeType+"%")
		}
		if req.Status != "" {
			m = m.Where(dao.SysAttachment.Columns().Status+" = ?", gconv.Bool(req.Status))
		}
		if req.CreatedAt != nil && len(req.CreatedAt) > 0 {
			if req.CreatedAt[0] != "" {
				m = m.Where(dao.SysAttachment.Columns().UpdatedAt+" >= ?", gconv.Time(req.CreatedAt[0]))
			}
			if len(req.CreatedAt) > 1 && req.CreatedAt[1] != "" {
				m = m.Where(dao.SysAttachment.Columns().UpdatedAt+" < ?", gconv.Time(req.CreatedAt[1]))
			}
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "updated_at desc,id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysAttachmentListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysAttachmentListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysAttachmentListRes{
				Id:        v.Id,
				AppId:     v.AppId,
				Drive:     v.Drive,
				Name:      v.Name,
				Kind:      v.Kind,
				Path:      v.Path,
				Size:      v.Size,
				Ext:       v.Ext,
				Status:    v.Status,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}
		}
	})
	return
}

func (s *sSysAttachment) GetById(ctx context.Context, id int64) (res *model.SysAttachmentInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAttachment.Ctx(ctx).WithAll().Where(dao.SysAttachment.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysAttachment) GetByMd5(ctx context.Context, md5 string) (res *model.SysAttachmentInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAttachment.Ctx(ctx).WithAll().Where(dao.SysAttachment.Columns().Md5, md5).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
		if res != nil {
			_, _ = dao.SysAttachment.Ctx(ctx).Unscoped().WherePri(res.Id).Update(do.SysAttachment{
				UpdatedAt: gtime.Now(),
			})
		}
	})
	return
}

func (s *sSysAttachment) AddUpload(ctx context.Context, req *model.UploadResponse, attr *model.SysAttachmentAddAttribute) (err error) {
	ext := gstr.SubStrRune(req.Name, gstr.PosRRune(req.Name, ".")+1, gstr.LenRune(req.Name)-1)
	err = s.Add(ctx, &model.SysAttachmentAddReq{
		AppId:     attr.AppId,
		Drive:     attr.Driver,
		Name:      req.Name,
		Kind:      s.getFileKind(ext),
		MimeType:  req.Type,
		Path:      req.Path,
		Size:      req.Size,
		Ext:       ext,
		Md5:       attr.Md5,
		Status:    true,
		CreatedBy: attr.UserId,
	})
	return
}

func (s *sSysAttachment) Add(ctx context.Context, req *model.SysAttachmentAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).Insert(do.SysAttachment{
			AppId:     req.AppId,
			Drive:     req.Drive,
			Name:      req.Name,
			Kind:      req.Kind,
			MimeType:  req.MimeType,
			Path:      req.Path,
			Size:      req.Size,
			Ext:       req.Ext,
			Md5:       req.Md5,
			Status:    req.Status,
			CreatedBy: req.CreatedBy,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSysAttachment) Edit(ctx context.Context, req *model.SysAttachmentEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).WherePri(req.Id).Update(do.SysAttachment{
			AppId:    req.AppId,
			Drive:    req.Drive,
			Name:     req.Name,
			Kind:     req.Kind,
			MimeType: req.MimeType,
			Path:     req.Path,
			Size:     req.Size,
			Ext:      req.Ext,
			Md5:      req.Md5,
			Status:   req.Status,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysAttachment) Delete(ctx context.Context, ids []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).Delete(dao.SysAttachment.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 附件管理状态修改（状态）
func (s *sSysAttachment) ChangeStatus(ctx context.Context, id int64, status bool) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).WherePri(id).
			Update(do.SysAttachment{
				Status: status,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysAttachment) getFileKind(ext string) string {
	for k, v := range fileKind {
		if v.ContainsI(ext) {
			return k
		}
	}
	return "other"
}
