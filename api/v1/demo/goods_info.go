// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-05 11:24:19
// 生成路径: api/v1/demo/goods_info.go
// 生成人：王中阳
// desc:商品表相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package demo

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/demo/model"
)

// GoodsInfoSearchReq 分页请求参数
type GoodsInfoSearchReq struct {
	g.Meta `path:"/list" tags:"商品表" method:"get" summary:"商品表列表"`
	commonApi.Author
	model.GoodsInfoSearchReq
}

// GoodsInfoSearchRes 列表返回结果
type GoodsInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.GoodsInfoSearchRes
}

// GoodsInfoExportReq 导出请求
type GoodsInfoExportReq struct {
	g.Meta `path:"/export" tags:"商品表" method:"get" summary:"商品表导出"`
	commonApi.Author
	model.GoodsInfoSearchReq
}

// GoodsInfoExportRes 导出响应
type GoodsInfoExportRes struct {
	commonApi.EmptyRes
}
type GoodsInfoExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"商品表" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type GoodsInfoExcelTemplateRes struct {
	commonApi.EmptyRes
}
type GoodsInfoImportReq struct {
	g.Meta `path:"/import" tags:"商品表" method:"post" summary:"商品表导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type GoodsInfoImportRes struct {
	commonApi.EmptyRes
}

// 相关连表查询数据
type LinkedGoodsInfoDataSearchReq struct {
	g.Meta `path:"/linkedData" tags:"商品表" method:"get" summary:"商品表关联表数据"`
	commonApi.Author
}

// 相关连表查询数据
type LinkedGoodsInfoDataSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.LinkedGoodsInfoDataSearchRes
}

// GoodsInfoAddReq 添加操作请求参数
type GoodsInfoAddReq struct {
	g.Meta `path:"/add" tags:"商品表" method:"post" summary:"商品表添加"`
	commonApi.Author
	*model.GoodsInfoAddReq
}

// GoodsInfoAddRes 添加操作返回结果
type GoodsInfoAddRes struct {
	commonApi.EmptyRes
}

// GoodsInfoEditReq 修改操作请求参数
type GoodsInfoEditReq struct {
	g.Meta `path:"/edit" tags:"商品表" method:"put" summary:"商品表修改"`
	commonApi.Author
	*model.GoodsInfoEditReq
}

// GoodsInfoEditRes 修改操作返回结果
type GoodsInfoEditRes struct {
	commonApi.EmptyRes
}

// GoodsInfoGetReq 获取一条数据请求
type GoodsInfoGetReq struct {
	g.Meta `path:"/get" tags:"商品表" method:"get" summary:"获取商品表信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// GoodsInfoGetRes 获取一条数据结果
type GoodsInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.GoodsInfoInfoRes
}

// GoodsInfoDeleteReq 删除数据请求
type GoodsInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"商品表" method:"delete" summary:"删除商品表"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// GoodsInfoDeleteRes 删除数据返回
type GoodsInfoDeleteRes struct {
	commonApi.EmptyRes
}
