// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/logic/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	sysDao "github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libWebsocket"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func init() {
	service.RegisterSysNotice(New())
}

func New() service.ISysNotice {
	return &sSysNotice{}
}

type sSysNotice struct{}

func (s *sSysNotice) List(ctx context.Context, req *model.SysNoticeSearchReq) (listRes *model.SysNoticeSearchRes, err error) {
	listRes = new(model.SysNoticeSearchRes)
	currentUserId := service.Context().GetUserId(ctx)
	if currentUserId <= 0 {
		err = errors.New("用户信息查询失败")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysNotice.Ctx(ctx).WithAll().As("n")
		if req.Id != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Id+" = ?", req.Id)
		}
		if req.Title != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Title+" = ?", req.Title)
		}
		if req.Type != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Type+" = ?", gconv.Int64(req.Type))
		}
		if req.Tag != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Tag+" = ?", gconv.Int(req.Tag))
		}
		if req.Status != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if len(req.DateRange) != 0 {
			m = m.Where("n."+dao.SysNotice.Columns().CreatedAt+" >=? AND "+dao.SysNotice.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		order := "n.id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		m = m.LeftJoin("sys_notice_read as nr", "nr.notice_id=n.id")
		//Where("nr.user_id=?", currentUserId)
		var res []*model.SysNoticeListRes

		err = m.Page(req.PageNum, req.PageSize).Fields("" +
			"n.id,n.title,n.type,n.tag,n.remark,n.sort,n.status,n.created_by,n.created_at," +
			"SUM(nr.clicks) as clickNumber" +
			//"(nr.user_id=" + strconv.FormatUint(currentUserId, 10) + ") as isRead" +
			"").Order(order).Group("n.id").Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = res

	})
	return
}
func (s *sSysNotice) ListShow(ctx context.Context, req *model.SysNoticeSearchReq) (listRes *model.SysNoticeSearchRes, err error) {
	listRes = new(model.SysNoticeSearchRes)
	currentUserId := service.Context().GetUserId(ctx)
	if currentUserId <= 0 {
		err = errors.New("用户信息查询失败")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysNotice.Ctx(ctx).WithAll().As("n")
		m = m.LeftJoin("sys_notice_read as nr", fmt.Sprintf("nr.notice_id=n.id AND nr.user_id=%d", currentUserId))
		if req.Id != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Id+" = ?", req.Id)
		}
		if req.Title != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Title+" = ?", req.Title)
		}
		if req.Type != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Type+" = ?", gconv.Int64(req.Type))
			if gconv.Int(req.Type) == consts.SysLetterType {
				if service.ToolsGenTable().IsMysql(dao.SysNotice.Group()) {
					m = m.Where("JSON_CONTAINS(n.receiver,?)", gconv.String(currentUserId))
				} else if service.ToolsGenTable().IsDM(dao.SysNotice.Group()) {
					m = m.Where("INSTR(n.receiver,?)>?", currentUserId, 0)
				} else {
					m = m.Where(fmt.Sprintf("receiver::jsonb @> '%d'::jsonb", currentUserId))
				}
			}
		}
		if req.Tag != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Tag+" = ?", gconv.Int(req.Tag))
		}
		if req.Status != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if len(req.DateRange) != 0 {
			m = m.Where("n."+dao.SysNotice.Columns().CreatedAt+" >=? AND "+dao.SysNotice.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		order := "n.id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysNoticeListRes
		err = m.Page(req.PageNum, req.PageSize).Fields("n.*,CASE WHEN nr.id IS NOT NULL THEN 1 ELSE 0 END AS isRead").Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		if req.IsTrim {
			for k, v := range res {
				v.Content = s.extractTextFromHTML(v.Content)
				if gstr.LenRune(v.Content) > 90 {
					res[k].Content = gstr.SubStrRune(v.Content, 0, 90) + "..."
				}
			}
		}
		listRes.List = res
	})
	return
}

// extractParagraphsFromHTML extracts text from all <p> tags in the HTML, ignoring <img> and other non-text nodes, and decodes HTML entities.
func (s *sSysNotice) extractTextFromHTML(htmlStr string) string {
	// Parse the HTML
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return htmlStr
	}

	// Create a slice to store the extracted paragraphs
	var paragraphs []string

	// Helper function to recursively traverse nodes and collect text content into a buffer
	var collectText func(*html.Node, *bytes.Buffer)
	collectText = func(n *html.Node, buf *bytes.Buffer) {
		switch n.Type {
		case html.TextNode:
			// Append the text content to the buffer
			buf.WriteString(n.Data)
		case html.ElementNode:
			// Recursively traverse child nodes, but ignore <img> and other non-text-producing elements
			// Note: In a more sophisticated implementation, you might want to handle other elements like <br>, <strong>, etc.
			if n.DataAtom != atom.Img && n.DataAtom != atom.Script && n.DataAtom != atom.Style {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					collectText(c, buf)
				}
			}
		}
	}

	// Define a function to recursively traverse nodes and collect text from <p> tags
	var traverseNodes func(*html.Node)
	traverseNodes = func(n *html.Node) {
		// Check the type of the node
		switch n.Type {
		case html.ElementNode:
			// Check if the node is a <p> tag
			if n.DataAtom == atom.P {
				// Create a buffer to store the text content of the <p> tag
				var textBuf bytes.Buffer
				collectText(n, &textBuf)
				// Decode HTML entities in the collected text and add it to the paragraphs slice
				paragraphs = append(paragraphs, html.UnescapeString(strings.TrimSpace(textBuf.String())))
			} else {
				// For other tags, just recursively traverse their children
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					traverseNodes(c)
				}
			}
		}
	}

	for n := doc; n != nil; n = n.NextSibling {
		if n.Type == html.DocumentNode {
			traverseNodes(n.FirstChild) // Start traversal from the first child of the document node (which is usually the <html> tag)
			break
		}
	}
	return strings.Join(paragraphs, "")
}

func (s *sSysNotice) GetById(ctx context.Context, id int64) (res *model.SysNoticeInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysNotice.Ctx(ctx).WithAll().Where(dao.SysNotice.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysNotice) Add(ctx context.Context, req *model.SysNoticeAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := do.SysNotice{
			Title:     req.Title,
			Type:      req.Type,
			Tag:       req.Tag,
			Content:   req.Content,
			Remark:    req.Remark,
			Receiver:  req.Receiver,
			Sort:      req.Sort,
			Status:    req.Status,
			CreatedBy: req.CreatedBy,
		}
		var rows sql.Result
		rows, err = dao.SysNotice.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, err, "添加失败")
		data.Id, err = rows.LastInsertId()
		liberr.ErrIsNil(ctx, err, "获取ID失败")
		response := &libWebsocket.WResponse{
			Event: consts.WebsocketTypeNotice,
			Data:  data,
		}
		if req.Type == consts.SysLetterType {
			//系统私信
			if len(req.Receiver) > 0 {
				for _, id := range req.Receiver {
					libWebsocket.SendToUser(id, response)
				}
			}
		} else {
			//系统通知
			libWebsocket.SendToAll(response)
		}
	})
	return
}

func (s *sSysNotice) Edit(ctx context.Context, req *model.SysNoticeEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := do.SysNotice{
			Title:     req.Title,
			Tag:       req.Tag,
			Content:   req.Content,
			Remark:    req.Remark,
			Sort:      req.Sort,
			Status:    req.Status,
			Receiver:  req.Receiver,
			UpdatedBy: req.UpdatedBy,
		}
		_, err = dao.SysNotice.Ctx(ctx).WherePri(req.Id).Update(data)
		liberr.ErrIsNil(ctx, err, "修改失败")
		data.Id = req.Id
		data.Type = req.Type
		response := &libWebsocket.WResponse{
			Event: consts.WebsocketTypeNotice,
			Data:  data,
		}
		if req.Type == consts.SysLetterType {
			//系统私信
			if len(req.Receiver) > 0 {
				for _, id := range req.Receiver {
					libWebsocket.SendToUser(id, response)
				}
			}
		} else {
			//系统通知
			libWebsocket.SendToAll(response)
		}
	})
	return
}

func (s *sSysNotice) Delete(ctx context.Context, ids []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysNotice.Ctx(ctx).Delete(dao.SysNotice.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

/*
	func (s *sSysNotice) IndexData(ctx context.Context) (res *model.SysNoticeIndexRes, err error) {
		err = g.Try(ctx, func(ctx context.Context) {
			userInfo := service.Context().GetLoginUser(ctx)
			readIdMap, err := dao.SysNoticeRead.Ctx(ctx).Where(dao.SysNoticeRead.Columns().UserId, userInfo.Id).Fields(dao.SysNoticeRead.Columns().NoticeId).All()
			fmt.Println(readIdMap, err)
			dao.SysNotice.Ctx(ctx).Where(dao.SysNotice.Columns().Type, "1").Order("id desc").Scan(res.Type1List)
			dao.SysNotice.Ctx(ctx)

		})
		return
	}
*/
//未读消息列表
func (s *sSysNotice) UnReadList(ctx context.Context) (res *model.SysNoticeListRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sSysNotice) UnReadCount(ctx context.Context, currentUserId uint64) (sysNoticeUnreadCount *model.SysNoticeUnreadCount, err error) {
	if currentUserId <= 0 {
		err = errors.New("获取用户信息失败")
		return
	}
	countFunc := func(t int) (count int) {
		noticeIds, serr := s.CurrentUseWithIds(ctx, currentUserId, t)
		if serr != nil {
			return 0
		}
		all, serr := dao.SysNotice.Ctx(ctx).
			WhereIn(dao.SysNotice.Columns().Id, noticeIds).
			Count()
		if serr != nil {
			return 0
		}
		read, serr := dao.SysNoticeRead.Ctx(ctx).
			WhereIn(dao.SysNoticeRead.Columns().NoticeId, noticeIds).
			Where(dao.SysNoticeRead.Columns().UserId, currentUserId).
			Count()
		if serr != nil {
			return 0
		}
		count = all - read
		return
	}
	sysNoticeUnreadCount = new(model.SysNoticeUnreadCount)
	sysNoticeUnreadCount.NotifyCount = countFunc(consts.SysNoticeType) //获取未读通知数量
	sysNoticeUnreadCount.NoticeCount = countFunc(consts.SysLetterType) //获取未读私信数量
	return
}

func (s *sSysNotice) ReadAll(ctx context.Context, nType string) (err error) {
	currentUserId := service.Context().GetUserId(ctx)
	if currentUserId <= 0 {
		err = errors.New("获取当前用户信息失败")
		return
	}
	subQuery := g.Model("sys_notice_read").Where("user_id=?", currentUserId).Fields("id")
	/*	dao.SysNoticeRead.Ctx(ctx).Where(dao.SysNoticeRead.Columns().UserId, currentUserId)*/
	unReadIdMap, err := dao.SysNotice.Ctx(ctx).WhereNotIn("id", subQuery).Fields("id").All()
	if unReadIdMap != nil {
		insertList := make([]*model.SysNoticeReadAddReq, 0)
		for _, record := range unReadIdMap {
			insertList = append(insertList, &model.SysNoticeReadAddReq{
				NoticeId:  record["id"].Int64(),
				UserId:    int64(currentUserId),
				CreatedAt: gtime.Now(),
			})
		}
		_, err = dao.SysNoticeRead.Ctx(ctx).Data(insertList).Insert()
		if err != nil {
			return
		}
	}
	return
}

func (s *sSysNotice) GetUserNameList(ctx context.Context, search string) (res []*model.SysNoticeUserNickname, err error) {
	res = make([]*model.SysNoticeUserNickname, 0)
	err = sysDao.SysUser.Ctx(ctx).Where("user_nickname like ?", "%"+search+"%").Fields("id,user_nickname").Page(1, 6).Scan(&res)
	return
}

func (s *sSysNotice) NoticeReadAddUserId(ctx context.Context, req *model.SysNoticeReadAddUserReq) (err error) {

	return
}

func (s *sSysNotice) CurrentUseWithIds(ctx context.Context, currentUserId uint64, noticeType int) (ids []int64, err error) {
	m := dao.SysNotice.Ctx(ctx)
	m = m.Where(dao.SysNotice.Columns().Status, 1).
		Where(dao.SysNotice.Columns().Type, noticeType)
	if noticeType == consts.SysLetterType {
		if service.ToolsGenTable().IsMysql(dao.SysNotice.Group()) {
			m = m.Where("JSON_CONTAINS(`receiver`,'" + fmt.Sprintf("%d", currentUserId) + "')")
		} else {
			m = m.Where("receiver::jsonb @> '" + gconv.String(currentUserId) + "'::jsonb")
		}
	}
	all, err := m.Fields("id").All()
	if err != nil {
		return
	}
	for _, record := range all {
		ids = append(ids, record["id"].Int64())
	}
	return
}
