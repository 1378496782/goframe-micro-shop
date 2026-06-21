// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: internal/app/shop/controller/user_info.go
// 生成人：gfast
// desc:用户
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/api/v1/shop"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/xuri/excelize/v2"
)

type userInfoController struct {
	systemController.BaseController
}

var UserInfo = new(userInfoController)

// List 列表
func (c *userInfoController) List(ctx context.Context, req *shop.UserInfoSearchReq) (res *shop.UserInfoSearchRes, err error) {
	res = new(shop.UserInfoSearchRes)
	res.UserInfoSearchRes, err = service.UserInfo().List(ctx, &req.UserInfoSearchReq)
	return
}

// Export 导出excel
func (c *userInfoController) Export(ctx context.Context, req *shop.UserInfoExportReq) (res *shop.UserInfoExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.UserInfoInfoRes
		//表头
		tableHead = []interface{}{"", "用户名", "头像", "", "加密盐 生成密码用", "1男 2女", "1正常 2拉黑冻结", "个性签名", "密保问题的答案", "", "", ""}
		excelData [][]interface{}
		//字典选项处理
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.UserInfo().GetExportData(ctx, &req.UserInfoSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var ()
			dt := []interface{}{
				v.Id,
				v.Name,
				v.Avatar,
				v.Password,
				v.UserSalt,
				v.Sex,
				v.Status,
				v.Sign,
				v.SecretAnswer,
				v.CreatedAt.Format("Y-m-d H:i:s"),
				v.UpdatedAt.Format("Y-m-d H:i:s"),
				v.DeletedAt.Format("Y-m-d H:i:s"),
			}
			excelData = append(excelData, dt)
		}
		req.PageNum++
	}
	//创建excel处理对象
	excel := new(libUtils.ExcelHelper).CreateFile()
	defer excel.Close()
	excel.ArrToExcel("Sheet1", "A1", excelData)
	col, _ := excelize.ColumnNumberToName(len(tableHead))
	row := len(excelData)
	cr, _ := excelize.JoinCellName(col, row)
	excel.SetCellBorder("Sheet1", "A1", cr)
	_, err = excel.WriteTo(r.Response.BufferWriter)
	if err != nil {
		return
	}
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Expose-Headers", "*")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("用户")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}

// Get 获取用户
func (c *userInfoController) Get(ctx context.Context, req *shop.UserInfoGetReq) (res *shop.UserInfoGetRes, err error) {
	res = new(shop.UserInfoGetRes)
	res.UserInfoInfoRes, err = service.UserInfo().GetById(ctx, req.Id)
	return
}

// Add 添加用户
func (c *userInfoController) Add(ctx context.Context, req *shop.UserInfoAddReq) (res *shop.UserInfoAddRes, err error) {
	err = service.UserInfo().Add(ctx, req.UserInfoAddReq)
	return
}

// Edit 修改用户
func (c *userInfoController) Edit(ctx context.Context, req *shop.UserInfoEditReq) (res *shop.UserInfoEditRes, err error) {
	err = service.UserInfo().Edit(ctx, req.UserInfoEditReq)
	return
}

// Delete 删除用户
func (c *userInfoController) Delete(ctx context.Context, req *shop.UserInfoDeleteReq) (res *shop.UserInfoDeleteRes, err error) {
	err = service.UserInfo().Delete(ctx, req.Ids)
	return
}
