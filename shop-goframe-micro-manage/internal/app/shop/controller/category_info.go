// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/controller/category_info.go
// 生成人：王中阳
// desc:商品分类
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

type categoryInfoController struct {
	systemController.BaseController
}

var CategoryInfo = new(categoryInfoController)

// List 列表
func (c *categoryInfoController) List(ctx context.Context, req *shop.CategoryInfoSearchReq) (res *shop.CategoryInfoSearchRes, err error) {
	res = new(shop.CategoryInfoSearchRes)
	res.CategoryInfoSearchRes, err = service.CategoryInfo().List(ctx, &req.CategoryInfoSearchReq)
	return
}

// Export 导出excel
func (c *categoryInfoController) Export(ctx context.Context, req *shop.CategoryInfoExportReq) (res *shop.CategoryInfoExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.CategoryInfoInfoRes
		//表头
		tableHead = []interface{}{"ID", "父级id", "名称", "等级", "排序", "创建时间", "", ""}
		excelData [][]interface{}
		//字典选项处理
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.CategoryInfo().GetExportData(ctx, &req.CategoryInfoSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-父级id
				parentId string
			)
			//关联表-单选-父级id
			if v.LinkedParentId != nil {
				parentId = gconv.String(v.LinkedParentId.Name)
			}
			dt := []interface{}{
				v.Id,
				//关联表-单选-父级id
				parentId,
				v.Name,
				v.Level,
				v.Sort,
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("商品分类")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *categoryInfoController) ExcelTemplate(ctx context.Context, req *shop.CategoryInfoExcelTemplateReq) (res *shop.CategoryInfoExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"父级id", "名称", "等级", "排序", "创建时间", "", ""}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("商品分类模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *categoryInfoController) Import(ctx context.Context, req *shop.CategoryInfoImportReq) (res *shop.CategoryInfoImportRes, err error) {
	err = service.CategoryInfo().Import(ctx, req.File)
	return
}

// LinkedCategoryInfoDataSearch 相关连表查询数据
func (c *categoryInfoController) LinkedCategoryInfoDataSearch(ctx context.Context, req *shop.LinkedCategoryInfoDataSearchReq) (res *shop.LinkedCategoryInfoDataSearchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/shop/categoryInfo/list") {
		err = errors.New("没有访问权限")
		return
	}
	res = new(shop.LinkedCategoryInfoDataSearchRes)
	res.LinkedCategoryInfoDataSearchRes, err = service.CategoryInfo().LinkedCategoryInfoDataSearch(ctx)
	return
}

// Get 获取商品分类
func (c *categoryInfoController) Get(ctx context.Context, req *shop.CategoryInfoGetReq) (res *shop.CategoryInfoGetRes, err error) {
	res = new(shop.CategoryInfoGetRes)
	res.CategoryInfoInfoRes, err = service.CategoryInfo().GetById(ctx, req.Id)
	return
}

// Add 添加商品分类
func (c *categoryInfoController) Add(ctx context.Context, req *shop.CategoryInfoAddReq) (res *shop.CategoryInfoAddRes, err error) {
	err = service.CategoryInfo().Add(ctx, req.CategoryInfoAddReq)
	return
}

// Edit 修改商品分类
func (c *categoryInfoController) Edit(ctx context.Context, req *shop.CategoryInfoEditReq) (res *shop.CategoryInfoEditRes, err error) {
	err = service.CategoryInfo().Edit(ctx, req.CategoryInfoEditReq)
	return
}

// Delete 删除商品分类
func (c *categoryInfoController) Delete(ctx context.Context, req *shop.CategoryInfoDeleteReq) (res *shop.CategoryInfoDeleteRes, err error) {
	err = service.CategoryInfo().Delete(ctx, req.Ids)
	return
}
