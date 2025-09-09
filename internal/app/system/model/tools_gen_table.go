/*
* @desc:代码生成表数据模型
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/10/28 11:43
 */

package model

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// ToolsGenTableEx 实体扩展
type ToolsGenTableEx struct {
	TableId        int64                    // 编号
	Group          string                   // 数据库group
	TableName      string                   // 表名称
	TableComment   string                   // 表描述
	ClassName      string                   // 实体类名称
	TplCategory    string                   // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string                   // 生成包路径
	ModuleName     string                   // 生成模块名
	BusinessName   string                   // 生成业务名
	FunctionName   string                   // 生成功能名
	FunctionAuthor string                   // 生成功能作者
	Options        string                   // 其它生成选项
	CreateBy       string                   // 创建者
	CreateTime     *gtime.Time              // 创建时间
	UpdateBy       string                   // 更新者
	UpdateTime     *gtime.Time              // 更新时间
	Remark         string                   // 备注
	Overwrite      bool                     // 是否覆盖原有文件
	OverwriteInfo  []*entity.OverwriteInfo  //覆盖文件信息
	SortColumn     string                   // 缺省排序字段
	SortType       string                   // 缺省排序方式 (asc顺序 desc倒序)
	ShowDetail     bool                     // 是否有查看详情功能
	ExcelPort      bool                     // 是否有导出excel功能
	ExcelImp       bool                     // 是否有导入excel功能
	UseSnowId      bool                     // 主键是否使用雪花ID
	UseVirtual     bool                     // 树表是否使用虚拟表
	TreeCode       string                   // 树编码字段
	TreeParentCode string                   // 树父编码字段
	TreeName       string                   // 树名称字段
	IsPkInsertable bool                     // 主键是否应该在新增时填写
	IsPkListable   bool                     // 主键是否出现在列表中
	HasTimeColumn  bool                     // 是否有时间类型
	HasCreatedBy   bool                     // 是否有created_by字段
	HasDeptId      bool                     // 是否右dept_id字段
	HasUpdatedBy   bool                     // 是否有updated_by字段
	HasConversion  bool                     // service中是否有转换函数
	Columns        []*ToolsGenTableColumnEx // 表列信息
	EditColumns    []*ToolsGenTableColumnEx // 编辑界面列信息，主键单独判断不在此列
	ListColumns    []*ToolsGenTableColumnEx // 列表界面列信息
	DetailColumns  []*ToolsGenTableColumnEx // 详情界面列信息
	QueryColumns   []*ToolsGenTableColumnEx // 查询表单列信息
	PkColumn       *ToolsGenTableColumnEx   // 主键列信息
	LinkedTables   []*ToolsGenTableLinked   // 关联表信息
	MenuPid        uint                     // 父级菜单ID
}

// ToolsGenTableLinked 关联表
type ToolsGenTableLinked struct {
	// 这里不加group默认和主表同一个group
	TableName     string                 // 表名称
	TableComment  string                 // 表描述
	ClassName     string                 // 实体类名称
	PackageName   string                 // 生成包路径
	ModuleName    string                 // 生成模块名
	BusinessName  string                 // 生成业务名
	TplCategory   string                 // 表类型列表或树表
	Options       string                 // 表相关选项配置
	OptionsStruct ToolsGenTableOptions   // 表相关选项配置结构
	RefColumns    *gmap.ListMap          // 要被查询的所有数据列信息
	LinkedTables  []*ToolsGenTableLinked // 嵌套关联表信息
}

// ToolsGenTableOptions 表相关选项配置
type ToolsGenTableOptions struct {
	TreeCode       string `json:"treeCode"`
	TreeName       string `json:"treeName"`
	TreeParentCode string `json:"treeParentCode"`
	ColumnAttr     *ToolsGenTableColumnEx
}

// ToolsGenTableEditData 编辑页面表数据详情
type ToolsGenTableEditData struct {
	*entity.ToolsGenTable
	TreeCode       interface{} `json:"treeCode"`
	TreeParentCode interface{} `json:"treeParentCode"`
	TreeName       interface{} `json:"treeName"`
}

// ToolsGenTableColumnsData 表与字段组合数据
type ToolsGenTableColumnsData struct {
	*entity.ToolsGenTable
	Columns []*entity.ToolsGenTableColumn `json:"columns"`
}
