// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/controller/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
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
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/xuri/excelize/v2"
)

type couponInfoController struct {
	systemController.BaseController
}

var CouponInfo = new(couponInfoController)

// List 列表
func (c *couponInfoController) List(ctx context.Context, req *shop.CouponInfoSearchReq) (res *shop.CouponInfoSearchRes, err error) {
	res = new(shop.CouponInfoSearchRes)
	res.CouponInfoSearchRes, err = service.CouponInfo().List(ctx, &req.CouponInfoSearchReq)
	return
}

// Export 导出excel
func (c *couponInfoController) Export(ctx context.Context, req *shop.CouponInfoExportReq) (res *shop.CouponInfoExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.CouponInfoInfoRes
		//表头
		tableHead = []interface{}{"ID", "关联商品id（0表示全场通用）", "名称", "类型", "优惠金额（元）", "过期时间", "创建时间", "更新时间", "删除时间（软删除）"}
		excelData [][]interface{}
		//字典选项处理
		shopCoupon    *systemApi.GetDictRes
		shopCouponMap = gmap.New()
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	shopCoupon, err = commonService.SysDictData().GetDictWithDataByType(ctx, "shop_coupon", "")
	if err != nil {
		return
	}
	for _, v := range shopCoupon.Values {
		shopCouponMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.CouponInfo().GetExportData(ctx, &req.CouponInfoSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-类型
				coupon_type interface{}
			)
			//字典-单选-类型
			coupon_type = shopCouponMap.Get(gconv.String(v.Type))
			dt := []interface{}{
				v.Id,
				v.GoodsId,
				v.Name,
				//字典-单选-类型
				coupon_type,
				v.Amount,
				v.Deadline.Format("Y-m-d H:i:s"),
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("优惠券")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *couponInfoController) ExcelTemplate(ctx context.Context, req *shop.CouponInfoExcelTemplateReq) (res *shop.CouponInfoExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"关联商品id（0表示全场通用）", "名称", "类型", "优惠金额（元）", "过期时间", "创建时间", "更新时间", "删除时间（软删除）"}
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
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Expose-Headers", "*")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("优惠券模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *couponInfoController) Import(ctx context.Context, req *shop.CouponInfoImportReq) (res *shop.CouponInfoImportRes, err error) {
	err = service.CouponInfo().Import(ctx, req.File)
	return
}

// Get 获取优惠券
func (c *couponInfoController) Get(ctx context.Context, req *shop.CouponInfoGetReq) (res *shop.CouponInfoGetRes, err error) {
	res = new(shop.CouponInfoGetRes)
	res.CouponInfoInfoRes, err = service.CouponInfo().GetById(ctx, req.Id)
	return
}

// Add 添加优惠券
func (c *couponInfoController) Add(ctx context.Context, req *shop.CouponInfoAddReq) (res *shop.CouponInfoAddRes, err error) {
	err = service.CouponInfo().Add(ctx, req.CouponInfoAddReq)
	return
}

// Edit 修改优惠券
func (c *couponInfoController) Edit(ctx context.Context, req *shop.CouponInfoEditReq) (res *shop.CouponInfoEditRes, err error) {
	err = service.CouponInfo().Edit(ctx, req.CouponInfoEditReq)
	return
}

// Delete 删除优惠券
func (c *couponInfoController) Delete(ctx context.Context, req *shop.CouponInfoDeleteReq) (res *shop.CouponInfoDeleteRes, err error) {
	err = service.CouponInfo().Delete(ctx, req.Ids)
	return
}
