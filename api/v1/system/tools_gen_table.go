/*
* @desc:代码生成
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/10/26 16:44
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// ToolsGenDBReq 获取数据库列表请求参数
type ToolsGenDBReq struct {
	g.Meta `path:"/tools/gen/dbList" tags:"系统后台/代码生成" method:"get" summary:"获取数据库"`
}

// ToolsGenDBRes 获取数据库列表返回参数
type ToolsGenDBRes struct {
	g.Meta `mime:"application/json"`
	List   []string `json:"list"`
}

// ToolsGenTableSearchReq 列表分页请求参数
type ToolsGenTableSearchReq struct {
	g.Meta       `path:"/tools/gen/tableList" tags:"系统后台/代码生成" method:"get" summary:"获取数据表"`
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	commonApi.PageReq
	commonApi.Author
}

// ToolsGenTableSearchRes 表列表数据
type ToolsGenTableSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.ToolsGenTable `json:"list"`
}

// ToolsGenTableImportSearchReq 要导入的表列表数据
type ToolsGenTableImportSearchReq struct {
	g.Meta       `path:"/tools/gen/dataList" tags:"系统后台/代码生成" method:"get" summary:"获取需要导入的数据表"`
	Group        string `p:"group"`        //数据库group
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	commonApi.PageReq
	commonApi.Author
}

// ToolsGenTableImportTableReq 导入表数据操作
type ToolsGenTableImportTableReq struct {
	g.Meta `path:"/tools/gen/importTableSave" tags:"系统后台/代码生成" method:"post" summary:"导入表结构操作"`
	commonApi.Author
	Group  string   `p:"group"`
	Tables []string `p:"tables" v:"required#表名必须指定"`
}

type ToolsGenTableDeleteReq struct {
	g.Meta `path:"/tools/gen/tableDelete" tags:"系统后台/代码生成" method:"delete" summary:"删除已导入的表"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#删除的id必须"`
}

type ToolsGenTableEditReq struct {
	g.Meta `path:"/tools/gen/columnList" tags:"系统后台/代码生成" method:"get" summary:"生成数据编辑"`
	commonApi.Author
	TableId int64 `p:"tableId" v:"required#tableId字段必须"`
}

// ToolsGenTableEditRes 生成数据编辑页面返回数据
type ToolsGenTableEditRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.ToolsGenTableColumn `json:"list"`
	Info   *model.ToolsGenTableEditData  `json:"info"`
}

// ToolsGenRelationTableReq 获取关联表数据
type ToolsGenRelationTableReq struct {
	g.Meta       `path:"/tools/gen/relationTable" tags:"系统后台/代码生成" method:"get" summary:"获取关联表数据"`
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	commonApi.PageReq
	commonApi.Author
}

// ToolsGenRelationTableRes 关联表返回数据
type ToolsGenRelationTableRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	Data []*model.ToolsGenTableColumnsData `json:"data"`
}

// ToolsGenTableColumnsEditReq 生成信息修改参数
type ToolsGenTableColumnsEditReq struct {
	g.Meta `path:"/tools/gen/editSave" tags:"系统后台/代码生成" method:"post" summary:"生成信息修改保存"`
	commonApi.Author
	TableId        int64                         `p:"tableId" v:"required#主键ID不能为空"`
	Group          string                        `p:"group"`
	TableName      string                        `p:"tableName"  v:"required#表名称不能为空"`
	TableComment   string                        `p:"tableComment"  v:"required#表描述不能为空"`
	ClassName      string                        `p:"className" v:"required#实体类名称不能为空"`
	FunctionAuthor string                        `p:"functionAuthor"  v:"required#作者不能为空"`
	TplCategory    string                        `p:"tplCategory"`
	PackageName    string                        `p:"packageName" v:"required#生成包路径不能为空"`
	ModuleName     string                        `p:"moduleName"`
	BusinessName   string                        `p:"businessName" v:"required#生成业务名不能为空"`
	FunctionName   string                        `p:"functionName" v:"required#生成功能名不能为空"`
	Remark         string                        `p:"remark"`
	Overwrite      string                        `p:"overwrite" v:"required#是否覆盖原有文件不能为空"`
	SortColumn     string                        `p:"sort_column" v:"required#缺省排序字段不能为空"`
	SortType       string                        `p:"sort_type" v:"required#缺省排序方式不能为空"`
	ShowDetail     string                        `p:"show_detail" v:"required#是否有查看详情功能不能为空"`
	Params         string                        `p:"params"`
	Columns        []*entity.ToolsGenTableColumn `p:"columns"`
	TreeCode       string                        `p:"tree_code"`
	TreeParentCode string                        `p:"tree_parent_code"`
	TreeName       string                        `p:"tree_name"`
	ExcelPort      string                        `p:"excelPort"`
	ExcelImp       string                        `p:"excelImp"`
	UseSnowId      string                        `p:"useSnowId"`
	UseVirtual     string                        `p:"useVirtual"`
	OverwriteInfo  []*entity.OverwriteInfo       `p:"overwriteInfo"`
	MenuPid        uint                          `p:"menuPid"`
}

type ToolsGenTableColumnsEditRes struct {
	commonApi.EmptyRes
}

type ToolsGenTablePreviewReq struct {
	g.Meta  `path:"/tools/gen/preview" tags:"系统后台/代码生成" method:"get" summary:"代码预览"`
	TableId int64 `p:"tableId" v:"required#表ID必须"`
}

type ToolsGenTablePreviewRes struct {
	g.Meta `mime:"application/json"`
	Data   g.MapStrStr `json:"data"`
}

type ToolsGenTableBatchGenCodeReq struct {
	g.Meta `path:"/tools/gen/batchGenCode" tags:"系统后台/代码生成" method:"post" summary:"代码生成"`
	Ids    []int `p:"ids" v:"required#ids必须且不能为空"`
}

type ToolsGenTableBatchGenCodeRes struct {
	commonApi.EmptyRes
}

type ToolsGenTableSyncTableReq struct {
	g.Meta `path:"/tools/gen/syncTable" tags:"系统后台/代码生成" method:"post" summary:"同步表结构"`
	commonApi.Author
	TableId int64 `p:"tableId" v:"required#表ID必须"`
}

type ToolsGenTableSyncTableRes struct {
	commonApi.EmptyRes
}
