/*
* @desc:代码生成字段处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/10/28 10:10
 */

package toolsGenTableColumn

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterToolsGenTableColumn(New())
}

func New() service.IToolsGenTableColumn {
	return &sToolsGenTableColumn{
		ColumnTypeStr:       []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext", "character", "character varying", "json", "jsonb"},
		ColumnTypeTime:      []string{"datetime", "time", "date", "timestamp"},
		ColumnTypeNumber:    []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal", "real", "numeric"},
		ColumnNameNotEdit:   []string{"created_by", "created_at", "updated_by", "updated_at", "deleted_at", "dept_id"},
		ColumnNameNotList:   []string{"updated_by", "updated_at", "deleted_at"},
		ColumnNameNotDetail: []string{"updated_at", "deleted_at"},
		ColumnNameNotQuery:  []string{"updated_by", "updated_at", "deleted_at", "remark"},
	}
}

type sToolsGenTableColumn struct {
	ColumnTypeStr       []string //数据库字符串类型
	ColumnTypeTime      []string //数据库时间类型
	ColumnTypeNumber    []string //数据库数字类型
	ColumnNameNotEdit   []string //页面不需要编辑字段
	ColumnNameNotList   []string //页面不需要显示的列表字段
	ColumnNameNotDetail []string //页面不需要显示的详情字段
	ColumnNameNotQuery  []string //页面不需要查询字段
}

// 自定义错误类型
type MyError struct {
	Code    int
	Message string
}

// 实现 error 接口的 Error() 方法
func (e MyError) Error() string {
	return fmt.Sprintf("错误代码：%d，错误信息：%s", e.Code, e.Message)
}

// SelectDbTableColumnsByName 根据表名称查询列信息
func (s *sToolsGenTableColumn) SelectDbTableColumnsByName(ctx context.Context, tableName string, group string) ([]*entity.ToolsGenTableColumn, error) {
	var res []*entity.ToolsGenTableColumn
	err := g.Try(ctx, func(ctx context.Context) {
		db := g.DB(group)
		namespace := db.GetConfig().Namespace
		if namespace == "" {
			namespace = "public"
		}
		var sql string
		if service.ToolsGenTable().IsPg(group) {
			//pg数据库
			sql = "SELECT" +
				"     a.attname AS column_name," +
				"     (NOT a.attnotnull OR i.indisprimary::int > 0)::text AS is_required," +
				"     i.indisprimary::text AS is_pk," +
				"     a.attnum AS sort_order_edit," +
				"     des.description AS column_comment," +
				"     (a.atthasdef OR (pg_get_expr(def.adbin, def.adrelid) ~ '^nextval'))::text AS is_increment," +
				"     format_type(a.atttypid, a.atttypmod) AS column_type" +
				" FROM" +
				"     pg_class t" +
				" JOIN" +
				"     pg_namespace n ON n.oid = t.relnamespace" +
				" JOIN" +
				"     pg_attribute a ON a.attrelid = t.oid" +
				" LEFT JOIN" +
				"     pg_description des ON des.objoid = t.oid AND des.objsubid = a.attnum" +
				" LEFT JOIN" +
				"     pg_attrdef def ON def.adrelid = t.oid AND def.adnum = a.attnum" +
				" LEFT JOIN (" +
				"     SELECT " +
				"         indrelid, " +
				"         UNNEST(indkey) AS colidx, " +
				"         indisprimary" +
				"     FROM " +
				"         pg_index" +
				" ) i ON i.indrelid = t.oid AND i.colidx = a.attnum "
			sql += " WHERE t.relkind = 'r' AND a.attnum > 0    AND NOT a.attisdropped AND" +
				gdb.FormatSqlWithArgs(" t.relname = ?  AND n.nspname = ? ", []interface{}{tableName, namespace}) + " order by  a.attnum"
		} else if service.ToolsGenTable().IsDM(group) {
			dbName := g.DB().GetSchema()
			//达梦数据库
			sql = "SELECT" +
				"    A.COLUMN_NAME," +
				"    (CASE WHEN A.NULLABLE = 'N' THEN '1' ELSE '0' END) AS IS_REQUIRED," +
				"    (CASE WHEN B.CONSTRAINT_TYPE = 'P' THEN '1' ELSE '0' END) AS IS_PK," +
				"    A.COLUMN_ID AS SORT_ORDER_EDIT," +
				"    C.COMMENTS AS COLUMN_COMMENT," +
				"    (CASE WHEN D.IS_INCREMENT = 1 THEN '1' ELSE '0' END) AS IS_INCREMENT," +
				"    A.DATA_TYPE || '(' || A.DATA_LENGTH || ')' AS COLUMN_TYPE " +
				"    FROM" +
				"    ALL_TAB_COLUMNS A" +
				"    LEFT JOIN (" +
				"    SELECT" +
				"        C.OWNER, " +
				"        C.TABLE_NAME," +
				"        CC.COLUMN_NAME," +
				"        C.CONSTRAINT_TYPE" +
				"    FROM" +
				"        ALL_CONSTRAINTS C" +
				"    JOIN ALL_CONS_COLUMNS CC" +
				"        ON C.OWNER = CC.OWNER  " +
				"        AND C.CONSTRAINT_NAME = CC.CONSTRAINT_NAME" +
				"        AND C.TABLE_NAME = CC.TABLE_NAME" +
				"    WHERE" +
				"        C.CONSTRAINT_TYPE = 'P'" +
				"        AND C.OWNER = ?  " +
				") B " +
				"    ON A.OWNER = B.OWNER " +
				"    AND A.TABLE_NAME = B.TABLE_NAME" +
				"    AND A.COLUMN_NAME = B.COLUMN_NAME" +
				"    LEFT JOIN ALL_COL_COMMENTS C" +
				"    ON A.OWNER = C.SCHEMA_NAME " +
				"    AND A.TABLE_NAME = C.TABLE_NAME" +
				"    AND A.COLUMN_NAME = C.COLUMN_NAME" +
				"    LEFT JOIN (" +
				"    SELECT" +
				"        sf_get_schema_name_by_id(st.schid) AS TABLE_OWNER," +
				"        st.name AS TABLE_NAME," +
				"        sco.name AS COLUMN_NAME," +
				"        CASE WHEN bitand(sco.info2, 0x0001) = 1 THEN 1 ELSE 0 END AS IS_INCREMENT" +
				"    FROM" +
				"        syscolumns sco" +
				"    JOIN sysobjects st" +
				"        ON sco.id = st.id" +
				"        AND st.subtype$ = 'UTAB'" +
				") D " +
				"    ON A.OWNER = D.TABLE_OWNER " +
				"    AND A.TABLE_NAME = D.TABLE_NAME" +
				"    AND A.COLUMN_NAME = D.COLUMN_NAME" +
				"    WHERE" +
				"    A.OWNER = ? " +
				"    AND A.TABLE_NAME = ?" +
				"    ORDER BY" +
				"    A.COLUMN_ID ASC"
			sql = gdb.FormatSqlWithArgs(sql, []interface{}{dbName, dbName, tableName})
		} else {
			sql = "select column_name, (case when (is_nullable = 'YES' || is_nullable = 'NO' && column_default is not null) then '0' else '1' end) as is_required, " +
				"(case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort_order_edit, column_comment," +
				" (case when extra = 'auto_increment' then '1' else '0' end) as is_increment, column_type from information_schema.columns" +
				" where table_schema = (select database()) "
			sql += " and " + gdb.FormatSqlWithArgs(" table_name=? ", []interface{}{tableName}) + " order by ordinal_position ASC "

		}

		err := db.GetScan(ctx, &res, sql)
		liberr.ErrIsNil(ctx, err, "查询列信息失败")
	})
	return res, err
}

// InitColumnField 初始化列属性字段
func (s *sToolsGenTableColumn) InitColumnField(column *entity.ToolsGenTableColumn, table *entity.ToolsGenTable) {
	dataType := s.GetDbType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.TableId
	//设置字段名
	column.GoField = gstr.CaseCamel(columnName)
	column.HtmlField = gstr.CaseCamelLower(columnName)

	//字段类型
	column.GoType, column.TsType, column.HtmlType = s.SetFieldType(dataType, column.ColumnType)

	// 编辑字段
	if s.IsNotEdit(columnName) {
		column.IsEdit = false
	} else if column.IsPk {
		column.IsEdit = false
	} else {
		column.IsEdit = true
	}
	// 列表字段
	if s.IsNotList(columnName) {
		column.IsList = false
	} else {
		column.IsList = true
	}
	// 详情字段
	if s.IsNotDetail(columnName) {
		column.IsDetail = false
	} else {
		column.IsDetail = true
	} // 查询字段
	if s.IsNotQuery(columnName) {
		column.IsQuery = false
	} else {
		column.IsQuery = true
	}

	// 查询字段类型
	if s.CheckNameColumn(columnName) {
		column.QueryType = "LIKE"
	} else {
		column.QueryType = "EQ"
	}

	// todo 最好不要用字段名来推断是否必填
	if strings.Index(columnName, "name") >= 0 || strings.Index(columnName, "status") >= 0 {
		column.IsRequired = true
	}

	// 状态字段设置单选框
	if s.CheckStatusColumn(columnName) {
		column.HtmlType = "radio"
		//column.IsStatus = true
	} else if s.CheckTypeColumn(columnName) || s.CheckSexColumn(columnName) {
		// 类型&性别字段设置下拉框
		column.HtmlType = "select"
	}

	column.SortOrderList = column.SortOrderEdit
	column.SortOrderDetail = column.SortOrderEdit
	column.SortOrderQuery = column.SortOrderEdit
	column.ColSpan = 1
	column.RowSpan = 1
	column.IsRowStart = false
	column.MinWidth = 150
	column.IsFixed = false
	column.IsOverflowTooltip = false
	column.IsCascade = false
	column.ParentColumnName = ""
	column.CascadeColumnName = ""
}

func (s *sToolsGenTableColumn) SetFieldType(sqlType, columnType string) (goType, tsType, htmlType string) {
	if s.IsStringObject(sqlType) {
		//字段为字符串类型
		goType = "string"
		tsType = "string"
		columnLength := s.GetColumnLength(columnType)
		if columnLength >= 500 {
			htmlType = "textarea"
		} else {
			htmlType = "input"
		}
	} else if s.IsTimeObject(sqlType) {
		//字段为时间类型
		goType = "Time"
		tsType = "string"
		htmlType = "datetime"
	} else if s.IsNumberObject(sqlType) {
		//字段为数字类型
		htmlType = "input"
		t, _ := gregex.ReplaceString(`\(.+\)`, "", columnType)
		t = gstr.Split(gstr.Trim(t), " ")[0]
		t = gstr.ToLower(t)
		// 如果是浮点型
		switch t {
		case "float", "double", "decimal", "numeric", "real":
			goType = "float64"
		case "bit", "int", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "integer":
			if gstr.ContainsI(columnType, "unsigned") {
				goType = "uint"
			} else {
				goType = "int"
			}
		case "big_int", "bigint":
			if gstr.ContainsI(columnType, "unsigned") {
				goType = "uint64"
			} else {
				goType = "int64"
			}
		}
		tsType = "number"
	} else if sqlType == "bit" || sqlType == "boolean" || sqlType == "bool" {
		goType = "bool"
		tsType = "boolean"
	}

	return
}

// GetDbType 获取数据库类型字段
func (s *sToolsGenTableColumn) GetDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else if strings.Index(columnType, " ") > 0 { // 类似 int unsigned， 即 int 类型后面并没有标注长度
		return columnType[0:strings.Index(columnType, " ")]
	} else {
		return columnType
	}
}

// IsExistInArray 判断 value 是否存在在切片array中
func (s *sToolsGenTableColumn) IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// IsStringObject 判断是否是数据库字符串类型
func (s *sToolsGenTableColumn) IsStringObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeStr)
}

// IsTimeObject 判断是否是数据库时间类型
func (s *sToolsGenTableColumn) IsTimeObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeTime)
}

// IsNumberObject 是否数字类型
func (s *sToolsGenTableColumn) IsNumberObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeNumber)
}

// IsNotEdit 是否不可编辑
func (s *sToolsGenTableColumn) IsNotEdit(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotEdit)
}

// IsNotList 不在列表显示
func (s *sToolsGenTableColumn) IsNotList(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotList)
}

// IsNotDetail 不在详情显示
func (s *sToolsGenTableColumn) IsNotDetail(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotDetail)
}

// IsNotQuery 不可用于查询
func (s *sToolsGenTableColumn) IsNotQuery(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotQuery)
}

// CheckNameColumn 检查字段名后4位是否是name
func (s *sToolsGenTableColumn) CheckNameColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4
		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]
		if tmp == "name" {
			return true
		}
	}
	return false
}

// CheckStatusColumn 检查字段名后6位是否是status
func (s *sToolsGenTableColumn) CheckStatusColumn(columnName string) bool {
	if len(columnName) >= 6 {
		end := len(columnName)
		start := end - 6

		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]

		if tmp == "status" {
			return true
		}
	}

	return false
}

// CheckTypeColumn 检查字段名后4位是否是type
func (s *sToolsGenTableColumn) CheckTypeColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "type" {
			return true
		}
	}
	return false
}

// CheckSexColumn 检查字段名后3位是否是sex
func (s *sToolsGenTableColumn) CheckSexColumn(columnName string) bool {
	if len(columnName) >= 3 {
		end := len(columnName)
		start := end - 3
		if start <= 0 {
			start = 0
		}
		if columnName[start:end] == "sex" {
			return true
		}
	}
	return false
}

// SelectGenTableColumnListByTableId 查询业务字段列表
func (s *sToolsGenTableColumn) SelectGenTableColumnListByTableId(ctx context.Context, tableId int64) (list []*entity.ToolsGenTableColumn, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ToolsGenTableColumn.Ctx(ctx).
			Where(dao.ToolsGenTableColumn.Columns().TableId, tableId).
			Order(dao.ToolsGenTableColumn.Columns().SortOrderEdit + " asc, " +
				dao.ToolsGenTableColumn.Columns().ColumnId + " asc").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取字段信息出错")
	})
	return list, nil
}

// GetAllTableColumns 获取所有字段信息
func (s *sToolsGenTableColumn) GetAllTableColumns(ctx context.Context) (list []*entity.ToolsGenTableColumn, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ToolsGenTableColumn.Ctx(ctx).
			Order(dao.ToolsGenTableColumn.Columns().SortOrderEdit + " asc, " +
				dao.ToolsGenTableColumn.Columns().ColumnId + " asc").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取字段信息出错")
	})
	return
}

func (s *sToolsGenTableColumn) SetPkColumn(table *model.ToolsGenTableEx, columns []*model.ToolsGenTableColumnEx) {
	for _, column := range columns {
		if column.IsPk {
			table.PkColumn = column
			break
		}
	}
	if table.PkColumn == nil {
		table.PkColumn = columns[0]
	}
}

// GetColumnLength 获取字段长度
func (s *sToolsGenTableColumn) GetColumnLength(columnType string) int {
	start := strings.Index(columnType, "(")
	end := strings.Index(columnType, ")")
	result := ""
	if start >= 0 && end >= 0 {
		result = columnType[start+1 : end-1]
	}
	return gconv.Int(result)
}

func (s *sToolsGenTableColumn) SelectDbTableColumnMapByName(ctx context.Context, tableName string, group string) (res map[string]*entity.ToolsGenTableColumn, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		list, err := s.SelectDbTableColumnsByName(ctx, tableName, group)
		liberr.ErrIsNil(ctx, err)
		res = make(map[string]*entity.ToolsGenTableColumn, len(list))
		for _, item := range list {
			res[item.ColumnName] = item
		}
	})
	return res, err
}
