// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-05 12:04:34
// 生成路径: internal/app/shop/controller/goods_info.go
// 生成人：王中阳
// desc:商品表
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/api/v1/shop"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/xuri/excelize/v2"
)

type goodsInfoController struct {
	systemController.BaseController
}

var GoodsInfo = new(goodsInfoController)

// List 列表
func (c *goodsInfoController) List(ctx context.Context, req *shop.GoodsInfoSearchReq) (res *shop.GoodsInfoSearchRes, err error) {
	res = new(shop.GoodsInfoSearchRes)
	res.GoodsInfoSearchRes, err = service.GoodsInfo().List(ctx, &req.GoodsInfoSearchReq)
	return
}

// Export 导出excel
func (c *goodsInfoController) Export(ctx context.Context, req *shop.GoodsInfoExportReq) (res *shop.GoodsInfoExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.GoodsInfoInfoRes
		//表头
		tableHead = []interface{}{"ID", "名称", "价格(分)", "一级分类", "二级分类", "三级分类", "品牌", "库存", "销量", "标签", "", "", ""}
		excelData [][]interface{}
		//字典选项处理
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.GoodsInfo().GetExportData(ctx, &req.GoodsInfoSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-一级分类
				level1CategoryId string
				//单选-二级分类
				level2CategoryId string
				//单选-三级分类
				level3CategoryId string
			)
			//关联表-单选-一级分类
			if v.LinkedLevel1CategoryId != nil {
				level1CategoryId = gconv.String(v.LinkedLevel1CategoryId.Name)
			}
			//关联表-单选-二级分类
			if v.LinkedLevel2CategoryId != nil {
				level2CategoryId = gconv.String(v.LinkedLevel2CategoryId.Name)
			}
			//关联表-单选-三级分类
			if v.LinkedLevel3CategoryId != nil {
				level3CategoryId = gconv.String(v.LinkedLevel3CategoryId.Name)
			}
			dt := []interface{}{
				v.Id,
				v.Name,
				v.Price,
				//关联表-单选-一级分类
				level1CategoryId,
				//关联表-单选-二级分类
				level2CategoryId,
				//关联表-单选-三级分类
				level3CategoryId,
				v.Brand,
				v.Stock,
				v.Sale,
				v.Tags,
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("商品表")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *goodsInfoController) ExcelTemplate(ctx context.Context, req *shop.GoodsInfoExcelTemplateReq) (res *shop.GoodsInfoExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"名称", "价格(分)", "一级分类", "二级分类", "三级分类", "品牌", "库存", "销量", "标签", "商品详情", "", "", ""}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("商品表模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *goodsInfoController) Import(ctx context.Context, req *shop.GoodsInfoImportReq) (res *shop.GoodsInfoImportRes, err error) {
	err = service.GoodsInfo().Import(ctx, req.File)
	return
}

// LinkedGoodsInfoDataSearch 相关连表查询数据
func (c *goodsInfoController) LinkedGoodsInfoDataSearch(ctx context.Context, req *shop.LinkedGoodsInfoDataSearchReq) (res *shop.LinkedGoodsInfoDataSearchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/shop/goodsInfo/list") {
		err = errors.New("没有访问权限")
		return
	}
	res = new(shop.LinkedGoodsInfoDataSearchRes)
	res.LinkedGoodsInfoDataSearchRes, err = service.GoodsInfo().LinkedGoodsInfoDataSearch(ctx)
	return
}

// Get 获取商品表
func (c *goodsInfoController) Get(ctx context.Context, req *shop.GoodsInfoGetReq) (res *shop.GoodsInfoGetRes, err error) {
	res = new(shop.GoodsInfoGetRes)
	res.GoodsInfoInfoRes, err = service.GoodsInfo().GetById(ctx, req.Id)
	return
}

// Add 添加商品表
func (c *goodsInfoController) Add(ctx context.Context, req *shop.GoodsInfoAddReq) (res *shop.GoodsInfoAddRes, err error) {
	err = service.GoodsInfo().Add(ctx, req.GoodsInfoAddReq)
	return
}

// Edit 修改商品表
func (c *goodsInfoController) Edit(ctx context.Context, req *shop.GoodsInfoEditReq) (res *shop.GoodsInfoEditRes, err error) {
	err = service.GoodsInfo().Edit(ctx, req.GoodsInfoEditReq)
	return
}

// Delete 删除商品表
func (c *goodsInfoController) Delete(ctx context.Context, req *shop.GoodsInfoDeleteReq) (res *shop.GoodsInfoDeleteRes, err error) {
	err = service.GoodsInfo().Delete(ctx, req.Ids)
	return
}
