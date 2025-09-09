// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-09 15:39:41
// 生成路径: internal/app/shop/controller/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/shop"
	systemApi "github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/xuri/excelize/v2"
)

type userCouponInfoController struct {
	systemController.BaseController
}

var UserCouponInfo = new(userCouponInfoController)

// List 列表
func (c *userCouponInfoController) List(ctx context.Context, req *shop.UserCouponInfoSearchReq) (res *shop.UserCouponInfoSearchRes, err error) {
	res = new(shop.UserCouponInfoSearchRes)
	res.UserCouponInfoSearchRes, err = service.UserCouponInfo().List(ctx, &req.UserCouponInfoSearchReq)
	return
}

// Export 导出excel
func (c *userCouponInfoController) Export(ctx context.Context, req *shop.UserCouponInfoExportReq) (res *shop.UserCouponInfoExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.UserCouponInfoInfoRes
		//表头
		tableHead = []interface{}{"", "用户id", "优惠券id", "状态", "优惠金额（元）", "创建时间", "更新时间", "删除时间（软删除）"}
		excelData [][]interface{}
		//字典选项处理
		shopCouponStatus    *systemApi.GetDictRes
		shopCouponStatusMap = gmap.New()
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	shopCouponStatus, err = commonService.SysDictData().GetDictWithDataByType(ctx, "shop_coupon_status", "")
	if err != nil {
		return
	}
	for _, v := range shopCouponStatus.Values {
		shopCouponStatusMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.UserCouponInfo().GetExportData(ctx, &req.UserCouponInfoSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-优惠券id
				couponId string
				//单选-状态
				status interface{}
			)
			//关联表-单选-优惠券id
			if v.LinkedCouponId != nil {
				couponId = gconv.String(v.LinkedCouponId.Name)
			}
			//字典-单选-状态
			status = shopCouponStatusMap.Get(gconv.String(v.Status))
			dt := []interface{}{
				v.Id,
				v.UserId,
				//关联表-单选-优惠券id
				couponId,
				//字典-单选-状态
				status,
				v.Amount,
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
	r.Response.Header().Set("Content-Type", "application/vnd.DEMO_WECHAT_OPEN_ID.spreadsheetml.sheet")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Expose-Headers", "*")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("用户优惠券管理")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *userCouponInfoController) ExcelTemplate(ctx context.Context, req *shop.UserCouponInfoExcelTemplateReq) (res *shop.UserCouponInfoExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"用户id", "优惠券id", "状态", "优惠金额（元）", "创建时间", "更新时间", "删除时间（软删除）"}
		excelData = [][]interface{}{tableHead}
	)
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
	r.Response.Header().Set("Content-Type", "application/vnd.DEMO_WECHAT_OPEN_ID.spreadsheetml.sheet")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Expose-Headers", "*")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("用户优惠券管理模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *userCouponInfoController) Import(ctx context.Context, req *shop.UserCouponInfoImportReq) (res *shop.UserCouponInfoImportRes, err error) {
	err = service.UserCouponInfo().Import(ctx, req.File)
	return
}

// LinkedUserCouponInfoDataSearch 相关连表查询数据
func (c *userCouponInfoController) LinkedUserCouponInfoDataSearch(ctx context.Context, req *shop.LinkedUserCouponInfoDataSearchReq) (res *shop.LinkedUserCouponInfoDataSearchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/shop/userCouponInfo/list") {
		err = errors.New("没有访问权限")
		return
	}
	res = new(shop.LinkedUserCouponInfoDataSearchRes)
	res.LinkedUserCouponInfoDataSearchRes, err = service.UserCouponInfo().LinkedUserCouponInfoDataSearch(ctx)
	return
}

// Get 获取用户优惠券
func (c *userCouponInfoController) Get(ctx context.Context, req *shop.UserCouponInfoGetReq) (res *shop.UserCouponInfoGetRes, err error) {
	res = new(shop.UserCouponInfoGetRes)
	res.UserCouponInfoInfoRes, err = service.UserCouponInfo().GetById(ctx, req.Id)
	return
}

// Add 添加用户优惠券
func (c *userCouponInfoController) Add(ctx context.Context, req *shop.UserCouponInfoAddReq) (res *shop.UserCouponInfoAddRes, err error) {
	err = service.UserCouponInfo().Add(ctx, req.UserCouponInfoAddReq)
	return
}

// Edit 修改用户优惠券
func (c *userCouponInfoController) Edit(ctx context.Context, req *shop.UserCouponInfoEditReq) (res *shop.UserCouponInfoEditRes, err error) {
	err = service.UserCouponInfo().Edit(ctx, req.UserCouponInfoEditReq)
	return
}

// Delete 删除用户优惠券
func (c *userCouponInfoController) Delete(ctx context.Context, req *shop.UserCouponInfoDeleteReq) (res *shop.UserCouponInfoDeleteRes, err error) {
	err = service.UserCouponInfo().Delete(ctx, req.Ids)
	return
}
