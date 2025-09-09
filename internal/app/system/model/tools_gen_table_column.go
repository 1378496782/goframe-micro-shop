/*
* @desc:代码生成表字段数据模型
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/10/28 11:41
 */

package model

import "github.com/gogf/gf/v2/container/gset"

// ToolsGenTableColumnEx 扩展
type ToolsGenTableColumnEx struct {
	ColumnId               int64                  // 编号
	TableId                int64                  // 归属表编号
	ColumnName             string                 // 列名称
	ColumnComment          string                 // 列描述
	ColumnType             string                 // 列类型
	GoType                 string                 // Go类型
	TsType                 string                 // Ts类型
	GoField                string                 // Go字段名
	HtmlField              string                 // html字段名
	IsPk                   bool                   // 是否主键
	IsIncrement            bool                   // 是否自增
	IsRequired             bool                   // 是否必填
	IsEdit                 bool                   // 是否编辑字段
	IsList                 bool                   // 是否列表字段
	IsDetail               bool                   // 是否详情字段
	IsQuery                bool                   // 是否查询字段
	IsStatus               bool                   // 是否状态字段
	SortOrderEdit          int                    // 插入编辑显示顺序
	SortOrderList          int                    // 列表显示顺序
	SortOrderDetail        int                    // 详情显示顺序
	SortOrderQuery         int                    // 查询显示顺序
	QueryType              string                 // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType               string                 // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType               string                 // 字典类型
	LinkTableName          string                 // 关联表名
	LinkTableClass         string                 // 关联表类名
	LinkTableModuleName    string                 // 关联表业务名
	LinkTableBusinessName  string                 // 关联表业务名
	LinkTablePackage       string                 // 关联表包名
	LinkLabelId            string                 // 关联表键名
	LinkLabelName          string                 // 关联表字段值
	ColSpan                int                    // 占多少列
	RowSpan                int                    // 占多少行
	IsRowStart             bool                   // 是否另起一行
	MinWidth               int                    // 表格最小宽度
	IsFixed                bool                   // 是否表格列左固定
	IsOverflowTooltip      bool                   // 是否过长自动隐藏
	FieldConversion        string                 // 字段转换方法
	FieldValidation        string                 // 字段校验规则
	IsCascade              bool                   // 是否级联查询
	ParentColumnName       string                 // 上级字段名
	CascadeColumnName      string                 // 级联查询字段
	IsCascadeParent        bool                   // 是否为级联父字段
	CascadeParent          *ToolsGenTableColumnEx // 级联父字段指针
	CascadeChildrenColumns *gset.StrSet           // 所有级联子字段名（按级联顺序）
}
