/*
* @desc:代码生成表数据处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/10/26 17:00
 */

package toolsGenTable

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"golang.org/x/tools/imports"
)

func init() {
	service.RegisterToolsGenTable(New())
}

func New() service.IToolsGenTable {
	return new(sToolsGenTable)
}

type sToolsGenTable struct{}

// DBList 获取数据库列表
func (s *sToolsGenTable) DBList(ctx context.Context, req *system.ToolsGenDBReq) (res *system.ToolsGenDBRes, err error) {
	res = new(system.ToolsGenDBRes)
	res.List = make([]string, 0)
	val, err := g.Cfg().Get(ctx, "database")
	liberr.ErrIsNil(ctx, err, "获取数据库失败")
	databases := val.Map()
	for k := range databases {
		if k == "logger" {
			continue
		}
		res.List = append(res.List, k)
	}
	sort.Stable(sort.StringSlice(res.List))
	return
}

// List 代码生成页面数据列表
func (s *sToolsGenTable) List(ctx context.Context, req *system.ToolsGenTableSearchReq) (res *system.ToolsGenTableSearchRes, err error) {
	res = new(system.ToolsGenTableSearchRes)
	m := dao.ToolsGenTable.Ctx(ctx)
	if req != nil {
		if req.TableName != "" {
			m = m.Where(dao.ToolsGenTable.Columns().TableName+" like ?", "%"+req.TableName+"%")
		}
		if req.TableComment != "" {
			m = m.Where(dao.ToolsGenTable.Columns().TableComment+" like ?", "%"+req.TableComment+"%")
		}
		if len(req.DateRange) > 0 {
			m = m.Where(dao.ToolsGenTable.Columns().CreateTime+" >=? AND "+dao.ToolsGenTable.Columns().CreateTime+" <=?", req.DateRange[0], req.DateRange[1])
		}
		err = g.Try(ctx, func(ctx context.Context) {
			res.Total, err = m.Count()
			liberr.ErrIsNil(ctx, err, "获取总行数失败")
			if req.PageNum == 0 {
				req.PageNum = 1
			}
			if req.PageSize == 0 {
				req.PageSize = consts.PageSize
			}
			err = m.Page(req.PageNum, req.PageSize).Order(dao.ToolsGenTable.Columns().TableId + " asc").Scan(&res.List)
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
	}
	return
}

// SelectDbTableList 查询据库表
func (s *sToolsGenTable) SelectDbTableList(ctx context.Context, req *system.ToolsGenTableImportSearchReq) (res *system.ToolsGenTableSearchRes, err error) {
	res = new(system.ToolsGenTableSearchRes)
	db := g.DB(req.Group)
	err = g.Try(ctx, func(ctx context.Context) {
		if !s.IsMysql(req.Group) && !s.IsPg(req.Group) && !s.IsDM(req.Group) {
			liberr.ErrIsNil(ctx, gerror.New("代码生成暂时只支持mysql、postgresql及达梦数据库"))
		}
		var sqlStr string
		if s.IsMysql(req.Group) {
			sqlStr = " from information_schema.tables where table_schema = (select database())"
			if s.IsDefault(req.Group) {
				sqlStr += " and table_name NOT LIKE 'tools_gen_%' and table_name NOT IN (select table_name from " + dao.ToolsGenTable.Table() + ") "
			} else {
				sqlStr += " and table_name NOT LIKE 'tools_gen_%'"
			}

			if req != nil {
				if req.TableName != "" {
					sqlStr += gdb.FormatSqlWithArgs(" and lower(table_name) like lower(?)", []interface{}{"%" + req.TableName + "%"})
				}

				if req.TableComment != "" {
					sqlStr += gdb.FormatSqlWithArgs(" and lower(table_comment) like lower(?)", []interface{}{"%" + req.TableComment + "%"})
				}
				if len(req.DateRange) > 0 {
					sqlStr += gdb.FormatSqlWithArgs(" and date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", []interface{}{req.DateRange[0]})
					sqlStr += gdb.FormatSqlWithArgs(" and date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", []interface{}{req.DateRange[1]})
				}
			}
			countSql := "select count(1) " + sqlStr
			res.Total, err = db.GetCount(ctx, countSql)
			liberr.ErrIsNil(ctx, err, "读取总表数失败")
			sqlStr = "table_name, table_comment, create_time, update_time " + sqlStr
			if req.PageNum == 0 {
				req.PageNum = 1
			}
			if req.PageSize == 0 {
				req.PageSize = consts.PageSize
			}
			page := (req.PageNum - 1) * req.PageSize
			sqlStr += " order by create_time desc,table_name asc limit  " + gconv.String(page) + "," + gconv.String(req.PageSize)
			g.Log().Debug(ctx, "select "+sqlStr)
			err = db.GetScan(ctx, &res.List, "select "+sqlStr)
			liberr.ErrIsNil(ctx, err, "读取数据失败")
		} else if s.IsPg(req.Group) {
			sqlStr = " from information_schema.tables JOIN pg_class ON pg_class.relname = table_name "
			if s.IsDefault(req.Group) {
				sqlStr += " WHERE table_schema = current_schema() AND table_name NOT LIKE 'tools_gen_%' and table_name NOT IN (select table_name from " + dao.ToolsGenTable.Table() + ") "
			} else {
				sqlStr += " WHERE table_schema = current_schema() AND table_name NOT LIKE 'tools_gen_%'"
			}

			if req != nil {
				if req.TableName != "" {
					sqlStr += gdb.FormatSqlWithArgs(" and lower(table_name) like lower(?)", []interface{}{"%" + req.TableName + "%"})
				}

				if req.TableComment != "" {
					sqlStr += gdb.FormatSqlWithArgs(" and lower(obj_description(pg_class.oid)) like lower(?)", []interface{}{"%" + req.TableComment + "%"})
				}
				if len(req.DateRange) > 0 {
					sqlStr += gdb.FormatSqlWithArgs(" and pg_stat_get_last_analyze_time(pg_class.oid) >= date_format(?,'%y%m%d') ", []interface{}{req.DateRange[0]})
					sqlStr += gdb.FormatSqlWithArgs(" and pg_stat_get_last_analyze_time(pg_class.oid) <= date_format(?,'%y%m%d') ", []interface{}{req.DateRange[1]})
				}
			}
			countSql := "select count(1) " + sqlStr
			res.Total, err = db.GetCount(ctx, countSql)
			liberr.ErrIsNil(ctx, err, "读取总表数失败")
			sqlStr = "table_name, obj_description(pg_class.oid) AS table_comment, " +
				"pg_stat_get_last_analyze_time(pg_class.oid) AS create_time, " +
				"pg_stat_get_last_autoanalyze_time(pg_class.oid) AS update_time " + sqlStr
			if req.PageNum == 0 {
				req.PageNum = 1
			}
			if req.PageSize == 0 {
				req.PageSize = consts.PageSize
			}
			page := (req.PageNum - 1) * req.PageSize
			sqlStr += " order by pg_stat_get_last_analyze_time(pg_class.oid) desc,table_name asc LIMIT  " + gconv.String(req.PageSize) + " OFFSET " + gconv.String(page)
			err = db.GetScan(ctx, &res.List, "select "+sqlStr)
			liberr.ErrIsNil(ctx, err, "读取数据失败")
		} else if s.IsDM(req.Group) {
			dbName := gstr.ToUpper(g.DB().GetConfig().Name)
			sqlStr = " FROM  ALL_TABLES T " +
				"LEFT JOIN  ALL_TAB_COMMENTS C  ON T.OWNER = C.OWNER AND T.TABLE_NAME = C.TABLE_NAME " +
				"LEFT JOIN  DBA_OBJECTS O ON T.OWNER = O.OWNER AND T.TABLE_NAME = O.OBJECT_NAME " +
				" AND O.OBJECT_TYPE = 'TABLE' "
			sqlStr += " WHERE T.OWNER = '" + dbName + "'"
			sqlStr += " AND T.TABLE_NAME NOT LIKE 'TOOLS_GEN_%' "
			if s.IsDefault(req.Group) {
				sqlStr += " AND T.TABLE_NAME NOT IN (SELECT TABLE_NAME FROM TOOLS_GEN_TABLE) "
			}
			if req != nil {
				if req.TableName != "" {
					sqlStr += gdb.FormatSqlWithArgs(" AND T.TABLE_NAME like ?", []interface{}{"%" + req.TableName + "%"})
				}

				if req.TableComment != "" {
					sqlStr += gdb.FormatSqlWithArgs(" AND C.COMMENTS like ?", []interface{}{"%" + req.TableComment + "%"})
				}
				if len(req.DateRange) > 0 {
					sqlStr += gdb.FormatSqlWithArgs(" AND O.CREATED >= date_format(?,'%y%m%d') ", []interface{}{req.DateRange[0]})
					sqlStr += gdb.FormatSqlWithArgs(" AND O.CREATED <= date_format(?,'%y%m%d') ", []interface{}{req.DateRange[1]})
				}
			}
			countSql := "select count(1) " + sqlStr
			res.Total, err = db.GetCount(ctx, countSql)
			liberr.ErrIsNil(ctx, err, "读取总表数失败")
			sqlStr = `T.TABLE_NAME,C.COMMENTS AS TABLE_COMMENT, O.CREATED AS CREATE_TIME, O.LAST_DDL_TIME AS UPDATE_TIME ` + sqlStr
			if req.PageNum == 0 {
				req.PageNum = 1
			}
			if req.PageSize == 0 {
				req.PageSize = consts.PageSize
			}
			page := (req.PageNum - 1) * req.PageSize
			sqlStr += " ORDER BY O.CREATED desc,T.TABLE_NAME asc LIMIT  " + gconv.String(req.PageSize) + " OFFSET " + gconv.String(page)
			err = db.GetScan(ctx, &res.List, "SELECT "+sqlStr)
			liberr.ErrIsNil(ctx, err, "读取数据失败")
		}
	})
	return
}

// SelectDbTableListByNames 查询数据库中对应的表数据
func (s *sToolsGenTable) SelectDbTableListByNames(ctx context.Context, tableNames []string, group ...string) ([]*entity.ToolsGenTable, error) {
	var result []*entity.ToolsGenTable
	err := g.Try(ctx, func(ctx context.Context) {
		if !s.IsMysql(group...) && !s.IsPg(group...) && !s.IsDM(group...) {
			liberr.ErrIsNil(ctx, gerror.New("代码生成暂时只支持mysql、postgresql及达梦数据库"))
		}
		db := g.DB(group...)
		var sqlStr string
		if s.IsMysql(group...) {
			sqlStr = "select * from information_schema.tables where table_name NOT LIKE 'tools_gen_%' " +
				" and table_schema = (select database()) "
			if len(tableNames) > 0 {
				in := gstr.TrimRight(gstr.Repeat("?,", len(tableNames)), ",")
				sqlStr += " and " + gdb.FormatSqlWithArgs("table_name in ("+in+")", gconv.SliceAny(tableNames))
			}
			err := db.GetScan(ctx, &result, sqlStr)
			liberr.ErrIsNil(ctx, err, "获取表格信息失败")
		} else if s.IsPg(group...) {
			sqlStr = "select information_schema.tables.*,obj_description(pg_class.oid) AS table_comment " +
				" from information_schema.tables JOIN pg_class ON pg_class.relname = table_name " +
				" WHERE table_schema = current_schema() " +
				"AND table_name NOT LIKE 'tools_gen_%' "
			if len(tableNames) > 0 {
				in := gstr.TrimRight(gstr.Repeat("?,", len(tableNames)), ",")
				sqlStr += " and " + gdb.FormatSqlWithArgs("table_name in ("+in+")", gconv.SliceAny(tableNames))
			}
			err := db.GetScan(ctx, &result, sqlStr)
			liberr.ErrIsNil(ctx, err, "获取表格信息失败")
		} else if s.IsDM(group...) {
			dbName := gstr.ToUpper(g.DB().GetConfig().Name)
			sqlStr = "SELECT T.TABLE_NAME,C.COMMENTS AS TABLE_COMMENT," +
				"O.CREATED AS CREATE_TIME," +
				"O.LAST_DDL_TIME AS UPDATE_TIME " +
				"FROM ALL_TABLES T LEFT JOIN ALL_TAB_COMMENTS C " +
				"ON T.OWNER = C.OWNER AND T.TABLE_NAME = C.TABLE_NAME " +
				"LEFT JOIN DBA_OBJECTS O " +
				"ON T.OWNER = O.OWNER " +
				"AND T.TABLE_NAME = O.OBJECT_NAME " +
				"AND O.OBJECT_TYPE = 'TABLE' " +
				gdb.FormatSqlWithArgs("WHERE T.OWNER = ? ", []interface{}{dbName}) +
				gdb.FormatSqlWithArgs("AND T.TABLE_NAME in(?) ", gconv.SliceAny(tableNames)) +
				"ORDER BY O.CREATED DESC,T.TABLE_NAME ASC"
			err := db.GetScan(ctx, &result, sqlStr)
			liberr.ErrIsNil(ctx, err, "获取表格信息失败")
		}
		dbGroup := "default"
		if len(group) > 0 {
			dbGroup = group[0]
		}
		for i := range result {
			result[i].Group = dbGroup
		}
	})
	return result, err
}

// ImportGenTable 导入表结构
func (s *sToolsGenTable) ImportGenTable(ctx context.Context, tableList []*entity.ToolsGenTable) error {
	if tableList != nil {
		err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err := g.Try(ctx, func(ctx context.Context) {
				for _, table := range tableList {
					tableName := table.TableName
					// 保存列信息
					genTableColumns, err := service.ToolsGenTableColumn().SelectDbTableColumnsByName(ctx, tableName, table.Group)
					liberr.ErrIsNil(ctx, err, "获取列数据失败")
					if len(genTableColumns) <= 0 {
						liberr.ErrIsNil(ctx, gerror.New("获取列数据失败"))
					}
					err = s.InitTable(ctx, table, genTableColumns)
					liberr.ErrIsNil(ctx, err)
					result, err1 := tx.Model(dao.ToolsGenTable.Table()).
						Data(do.ToolsGenTable{
							Group:          table.Group,
							TableName:      table.TableName,
							TableComment:   table.TableComment,
							ClassName:      table.ClassName,
							TplCategory:    table.TplCategory,
							PackageName:    table.PackageName,
							ModuleName:     table.ModuleName,
							BusinessName:   table.BusinessName,
							FunctionName:   table.FunctionName,
							FunctionAuthor: table.FunctionAuthor,
							Options:        table.Options,
							CreateTime:     table.CreateTime,
							UpdateTime:     table.UpdateTime,
							Remark:         table.Remark,
							Overwrite:      table.Overwrite,
							SortColumn:     table.SortColumn,
							SortType:       table.SortType,
							ShowDetail:     table.ShowDetail,
							ExcelPort:      table.ExcelPort,
							ExcelImp:       table.ExcelImp,
							UseSnowId:      table.UseSnowId,
							UseVirtual:     table.UseVirtual,
							OverwriteInfo:  table.OverwriteInfo,
							MenuPid:        table.MenuPid,
						}).
						Insert()
					liberr.ErrIsNil(ctx, err1)
					tmpId, err2 := result.LastInsertId()
					liberr.ErrIsNil(ctx, err2, "保存数据失败")
					if tmpId <= 0 {
						liberr.ErrIsNil(ctx, gerror.New("保存数据失败"))
					}
					table.TableId = tmpId
					for _, column := range genTableColumns {
						service.ToolsGenTableColumn().InitColumnField(column, table)
						_, err3 := tx.Model(dao.ToolsGenTableColumn.Table()).FieldsEx("column_id").
							Insert(column)
						liberr.ErrIsNil(ctx, err3, "保存列数据失败")
					}
				}
			})
			return err
		})
		return err
	} else {
		return gerror.New("参数错误")
	}
}

// InitTable 初始化表信息
func (s *sToolsGenTable) InitTable(ctx context.Context, table *entity.ToolsGenTable, columns []*entity.ToolsGenTableColumn) error {
	table.ClassName = s.ConvertClassName(ctx, table.TableName)
	table.PackageName = g.Cfg().MustGet(ctx, "gen.packageName").String()
	table.ModuleName = g.Cfg().MustGet(ctx, "gen.moduleName").String()
	table.BusinessName = s.GetBusinessName(ctx, table.TableName)
	table.FunctionName = table.TableComment
	table.FunctionAuthor = g.Cfg().MustGet(ctx, "gen.author").String()
	table.TplCategory = "crud"
	pkColumn, err := s.getPkColumn(columns)
	if err != nil {
		return err
	}
	table.Overwrite = false
	table.SortColumn = pkColumn.ColumnName
	table.SortType = "asc"
	table.ShowDetail = false
	table.CreateTime = gtime.Now()
	table.UpdateTime = table.CreateTime
	return nil
}

// ConvertClassName 表名转换成类名
func (s *sToolsGenTable) ConvertClassName(ctx context.Context, tableName string) string {
	return gstr.CaseCamel(s.removeTablePrefix(ctx, tableName))
}

// GetBusinessName 获取业务名
func (s *sToolsGenTable) GetBusinessName(ctx context.Context, tableName string) string {
	return s.removeTablePrefix(ctx, tableName)
}

// 找到主键 column
func (s *sToolsGenTable) getPkColumn(columns []*entity.ToolsGenTableColumn) (*entity.ToolsGenTableColumn, error) {
	for _, column := range columns {
		if column.IsPk {
			return column, nil
		}
	}
	return nil, gerror.New("无法找到主键")
}

// 删除表前缀
func (s *sToolsGenTable) removeTablePrefix(ctx context.Context, tableName string) string {
	autoRemovePre := g.Cfg().MustGet(ctx, "gen.autoRemovePre").Bool()
	tablePrefix := g.Cfg().MustGet(ctx, "gen.tablePrefix").String()
	if autoRemovePre && tablePrefix != "" {
		searchList := strings.Split(tablePrefix, ",")
		for _, str := range searchList {
			if strings.HasPrefix(tableName, str) {
				tableName = strings.Replace(tableName, str, "", 1) //注意，只替换一次
			}
		}
	}
	return tableName
}

// 获取数据库驱动类型
func (s *sToolsGenTable) getDbDriver(group ...string) string {
	if len(group) == 0 {
		group = []string{"default"}
	}
	config := g.DB(group[0]).GetConfig()
	return gstr.ToLower(config.Type)
}

func (s *sToolsGenTable) IsPg(group ...string) bool {
	return s.getDbDriver(group...) == "pgsql"
}

func (s *sToolsGenTable) IsMysql(group ...string) bool {
	return s.getDbDriver(group...) == "mysql"
}

func (s *sToolsGenTable) IsDM(group ...string) bool {
	return s.getDbDriver(group...) == "dm"
}

func (s *sToolsGenTable) IsDefault(group string) bool {
	return group == "default" || group == ""
}

// DeleteTable 删除表信息
func (s *sToolsGenTable) DeleteTable(ctx context.Context, req *system.ToolsGenTableDeleteReq) error {
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err := g.Try(ctx, func(ctx context.Context) {
			_, err := tx.Model(dao.ToolsGenTable.Table()).Where(dao.ToolsGenTable.Columns().TableId+" in(?)", req.Ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除表格数据失败")
			_, err = tx.Model(dao.ToolsGenTableColumn.Table()).Where(dao.ToolsGenTableColumn.Columns().TableId+" in(?)", req.Ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除表格字段数据失败")
		})
		return err
	})
	return err
}

// ColumnList 表格字段列表数据
func (s *sToolsGenTable) ColumnList(ctx context.Context, req *system.ToolsGenTableEditReq) (res *system.ToolsGenTableEditRes, err error) {
	res = new(system.ToolsGenTableEditRes)
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			tableInfo *entity.ToolsGenTable
		)
		res.List, err = service.ToolsGenTableColumn().SelectGenTableColumnListByTableId(ctx, req.TableId)
		liberr.ErrIsNil(ctx, err)
		tableInfo, err = s.GetTableInfoByTableId(ctx, req.TableId)
		liberr.ErrIsNil(ctx, err)
		if tableInfo == nil {
			liberr.ErrIsNil(ctx, gerror.New("表数据不存在"))
		}
		res.Info = &model.ToolsGenTableEditData{
			ToolsGenTable: tableInfo,
		}
		//如果是树表则设置树表配置
		if tableInfo != nil && tableInfo.TplCategory == "tree" {
			options := gjson.New(tableInfo.Options)
			res.Info.TreeCode = options.Get("treeCode")
			res.Info.TreeParentCode = options.Get("treeParentCode")
			res.Info.TreeName = options.Get("treeName")
		}
	})
	return
}

// GetTableInfoByTableId 获取表格数据
func (s *sToolsGenTable) GetTableInfoByTableId(ctx context.Context, tableId int64) (data *entity.ToolsGenTable, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ToolsGenTable.Ctx(ctx).WherePri(tableId).Scan(&data)
		liberr.ErrIsNil(ctx, err, "获取表格信息出错")
	})
	return
}

// GetRelationTable 获取关联表数据
func (s *sToolsGenTable) GetRelationTable(ctx context.Context, req *system.ToolsGenRelationTableReq) (res *system.ToolsGenRelationTableRes, err error) {
	var tableColumnsAll []*entity.ToolsGenTableColumn
	res = new(system.ToolsGenRelationTableRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.ToolsGenTable.Ctx(ctx)
		if req.TableName != "" {
			m = m.WhereLike(dao.ToolsGenTable.Columns().TableName, "%"+req.TableName+"%")
		}
		if req.TableComment != "" {
			m = m.WhereLike(dao.ToolsGenTable.Columns().TableComment, "%"+req.TableComment+"%")
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取表数据总数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		//获取表数据
		err = m.Page(req.PageNum, req.PageSize).Order(dao.ToolsGenTable.Columns().TableId + " ASC ").Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "获取表数据失败")
		//获取表字段数据
		tableColumnsAll, err = service.ToolsGenTableColumn().GetAllTableColumns(ctx)
		liberr.ErrIsNil(ctx, err)
		for k, v := range res.Data {
			for _, cv := range tableColumnsAll {
				if cv.TableId == v.TableId {
					res.Data[k].Columns = append(res.Data[k].Columns, cv)
				}
			}
		}
	})
	return
}

// SaveEdit 更新表及字段生成信息
func (s *sToolsGenTable) SaveEdit(ctx context.Context, req *system.ToolsGenTableColumnsEditReq) (err error) {
	if req == nil {
		err = gerror.New("参数错误")
		return
	}
	var table *entity.ToolsGenTable
	err = dao.ToolsGenTable.Ctx(ctx).Where("table_id=?", req.TableId).Scan(&table)
	if err != nil || table == nil {
		err = gerror.New("数据不存在")
		return
	}
	if req.Group != "" {
		table.Group = req.Group
	}
	if req.TableName != "" {
		table.TableName = req.TableName
	}
	if req.TableComment != "" {
		table.TableComment = req.TableComment
	}
	if req.BusinessName != "" {
		table.BusinessName = req.BusinessName
	}
	if req.ClassName != "" {
		table.ClassName = req.ClassName
	}
	if req.FunctionAuthor != "" {
		table.FunctionAuthor = req.FunctionAuthor
	}
	if req.FunctionName != "" {
		table.FunctionName = req.FunctionName
	}
	if req.ModuleName != "" {
		table.ModuleName = req.ModuleName
	}
	if req.PackageName != "" {
		table.PackageName = req.PackageName
		//从包名中获取模块名
		lastIndex := gstr.PosR(table.PackageName, "/")
		if lastIndex != -1 {
			table.ModuleName = gstr.SubStr(table.PackageName, lastIndex+1)
		}
	}
	if req.Remark != "" {
		table.Remark = req.Remark
	}
	if req.Overwrite != "" {
		table.Overwrite = gconv.Bool(req.Overwrite)
		table.OverwriteInfo = req.OverwriteInfo
	}
	if req.SortColumn != "" {
		table.SortColumn = req.SortColumn
	}
	if req.SortType != "" {
		table.SortType = req.SortType
	}
	if req.ShowDetail != "" {
		table.ShowDetail = gconv.Bool(req.ShowDetail)
	}
	if req.ExcelPort != "" {
		table.ExcelPort = gconv.Bool(req.ExcelPort)
	}
	if req.ExcelImp != "" {
		table.ExcelImp = gconv.Bool(req.ExcelImp)
	}
	if req.UseSnowId != "" {
		table.UseSnowId = gconv.Bool(req.UseSnowId)
	}
	if req.UseVirtual != "" {
		table.UseVirtual = gconv.Bool(req.UseVirtual)
	}
	if req.TplCategory != "" {
		table.TplCategory = req.TplCategory
	}
	if req.Params != "" {
		table.Options = req.Params
	}
	table.UpdateTime = gtime.Now()
	var options g.Map
	if req.TplCategory == "tree" {
		//树表设置options
		options = g.Map{
			"treeCode":       req.TreeCode,
			"treeParentCode": req.TreeParentCode,
			"treeName":       req.TreeName,
		}
		table.Options = gconv.String(options)
	} else {
		table.Options = ""
	}
	table.MenuPid = req.MenuPid

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = tx.Model(dao.ToolsGenTable.Table()).
				WherePri(table.TableId).
				Update(do.ToolsGenTable{
					Group:          table.Group,
					TableName:      table.TableName,
					TableComment:   table.TableComment,
					ClassName:      table.ClassName,
					TplCategory:    table.TplCategory,
					PackageName:    table.PackageName,
					ModuleName:     table.ModuleName,
					BusinessName:   table.BusinessName,
					FunctionName:   table.FunctionName,
					FunctionAuthor: table.FunctionAuthor,
					Options:        table.Options,
					CreateTime:     table.CreateTime,
					UpdateTime:     table.UpdateTime,
					Remark:         table.Remark,
					Overwrite:      table.Overwrite,
					SortColumn:     table.SortColumn,
					SortType:       table.SortType,
					ShowDetail:     table.ShowDetail,
					ExcelPort:      table.ExcelPort,
					ExcelImp:       table.ExcelImp,
					UseSnowId:      table.UseSnowId,
					UseVirtual:     table.UseVirtual,
					OverwriteInfo:  gconv.String(table.OverwriteInfo),
					MenuPid:        table.MenuPid,
				})
			liberr.ErrIsNil(ctx, err, "保存表数据失败")
			//保存列数据
			if req.Columns != nil {
				for _, column := range req.Columns {
					if column.ColumnId > 0 {
						var dbColumn *entity.ToolsGenTableColumn
						err = dao.ToolsGenTableColumn.Ctx(ctx).Where("column_id=?", column.ColumnId).Scan(&dbColumn)
						liberr.ErrIsNil(ctx, err, "列数据信息不存在")
						if dbColumn != nil {
							dbColumn.ColumnComment = column.ColumnComment
							dbColumn.GoType = column.GoType
							dbColumn.HtmlType = column.HtmlType
							dbColumn.HtmlField = column.HtmlField
							dbColumn.QueryType = column.QueryType
							dbColumn.GoField = column.GoField
							dbColumn.DictType = column.DictType
							dbColumn.IsEdit = column.IsEdit
							dbColumn.IsList = column.IsList
							dbColumn.IsDetail = column.IsDetail
							dbColumn.IsQuery = column.IsQuery
							dbColumn.SortOrderEdit = column.SortOrderEdit
							dbColumn.SortOrderList = column.SortOrderList
							dbColumn.SortOrderDetail = column.SortOrderDetail
							dbColumn.SortOrderQuery = column.SortOrderQuery
							dbColumn.IsRequired = column.IsRequired
							dbColumn.ColSpan = column.ColSpan
							dbColumn.RowSpan = column.RowSpan
							dbColumn.IsRowStart = column.IsRowStart
							dbColumn.MinWidth = column.MinWidth
							dbColumn.IsFixed = column.IsFixed
							dbColumn.IsOverflowTooltip = column.IsOverflowTooltip
							dbColumn.IsCascade = column.IsCascade
							dbColumn.ParentColumnName = column.ParentColumnName
							dbColumn.CascadeColumnName = column.CascadeColumnName
							dbColumn.TsType = column.TsType
							if tc, e := options["treeParentCode"]; options != nil && e && tc != "" && tc == dbColumn.HtmlField {
								dbColumn.IsQuery = false
								dbColumn.IsList = false
								dbColumn.HtmlType = "treeSelect"
							}
							//获取字段关联表信息
							if column.LinkLabelName != "" {
								dbColumn.LinkTableName = column.LinkTableName
								dbColumn.LinkLabelId = column.LinkLabelId
								dbColumn.LinkLabelName = column.LinkLabelName
								var linkTable *entity.ToolsGenTable
								err = dao.ToolsGenTable.Ctx(ctx).Where("table_name =?", column.LinkTableName).Scan(&linkTable)
								liberr.ErrIsNil(ctx, err, "获取关联表数据失败")
								dbColumn.LinkTableClass = linkTable.ClassName
								dbColumn.LinkTableModuleName = linkTable.ModuleName
								dbColumn.LinkTableBusinessName = linkTable.BusinessName
								dbColumn.LinkTablePackage = linkTable.PackageName
							} else {
								dbColumn.LinkTableName = ""
								dbColumn.LinkTableClass = ""
								dbColumn.LinkTableModuleName = ""
								dbColumn.LinkTableBusinessName = ""
								dbColumn.LinkTablePackage = ""
								dbColumn.LinkLabelId = ""
								dbColumn.LinkLabelName = ""
							}
							_, err = tx.Model(dao.ToolsGenTableColumn.Table()).
								WherePri(dbColumn.ColumnId).
								FieldsEx(dao.ToolsGenTableColumn.Columns().ColumnId).
								Update(dbColumn)
							liberr.ErrIsNil(ctx, err, "保存列:"+dbColumn.ColumnName+"，数据失败")
						}
					}
				}
			}
		})
		return
	})
	return
}

// GenData 获取生成数据
func (s *sToolsGenTable) GenData(ctx context.Context, tableId int64) (data g.MapStrStr, extendData *model.ToolsGenTableEx, err error) {
	extendData, err = s.SelectRecordById(ctx, tableId)
	if err != nil {
		return
	}
	if extendData == nil {
		err = gerror.New("表格数据不存在")
		return
	}
	service.ToolsGenTableColumn().SetPkColumn(extendData, extendData.Columns)
	view := gview.New()
	err = view.SetConfigWithMap(g.Map{
		"Paths":      g.Cfg().MustGet(ctx, "gen.templatePath").String(),
		"Delimiters": []string{"{{", "}}"},
	})
	if err != nil {
		return
	}
	view.BindFuncMap(g.Map{
		"UcFirst": func(str string) string {
			return gstr.UcFirst(str)
		},
		"Sum": func(a, b int) int {
			return a + b
		},
		"CaseCamelLower": gstr.CaseCamelLower, //首字母小写驼峰
		"CaseCamel":      gstr.CaseCamel,      //首字母大写驼峰
		"HasSuffix":      gstr.HasSuffix,      //是否存在后缀
		"ContainsI":      gstr.ContainsI,      //是否包含子字符串
		"VueTag": func(t string) string {
			return t
		},
		"inArray": func(arr []interface{}, val interface{}) bool {
			for _, v := range arr {
				if reflect.DeepEqual(v, val) {
					return true
				}
			}
			return false
		},
		"append": func(arr []interface{}, val ...interface{}) []interface{} {
			arr = append(arr, val...)
			return arr
		},
		"newArray": func() []interface{} {
			return []interface{}{}
		},
		"strTrim": func(str interface{}, trimStr interface{}) string {
			strCon := gconv.String(str)
			strSub := gconv.String(trimStr)
			if len(strCon) > len(strSub) {
				strCon = gstr.StrEx(strCon, strSub)
			}
			return strCon
		},
	})

	//树形菜单选项
	tplData := g.Map{
		"table":      extendData,
		"goModName":  g.Cfg().MustGet(ctx, "gen.goModName").String(),
		"apiVersion": g.Cfg().MustGet(ctx, "gen.apiName").String(),
		"modulePath": gstr.StrEx(extendData.PackageName, "internal/app/"),
	}
	apiKey := "api"
	apiValue := ""
	var tmpApi string
	if tmpApi, err = view.Parse(ctx, "go/api.template", tplData); err == nil {
		apiValue = tmpApi
		apiValue, err = s.trimBreak(apiValue)
	} else {
		return
	}
	daoKey := "dao"
	daoValue := ""
	var tmpDao string
	if tmpDao, err = view.Parse(ctx, "go/dao.template", tplData); err == nil {
		daoValue = tmpDao
		daoValue, err = s.trimBreak(daoValue)
	} else {
		return
	}
	daoInternalKey := "dao_internal"
	daoInternalValue := ""
	var tmpInternalDao string
	if tmpInternalDao, err = view.Parse(ctx, "go/dao_internal.template", tplData); err == nil {
		daoInternalValue = tmpInternalDao
		daoInternalValue, err = s.trimBreak(daoInternalValue)
	} else {
		return
	}
	modelKey := "model"
	modelValue := ""
	var tmpModel string
	if tmpModel, err = view.Parse(ctx, "go/model.template", tplData); err == nil {
		modelValue = tmpModel
		modelValue, err = s.trimBreak(modelValue)
	} else {
		return
	}
	modelEntityKey := "model_entity"
	modelEntityValue := ""
	var tmpModelEntity string
	if tmpModelEntity, err = view.Parse(ctx, "go/model_entity.template", tplData); err == nil {
		modelEntityValue = tmpModelEntity
		modelEntityValue, err = s.trimBreak(modelEntityValue)
	} else {
		return
	}
	modelDoKey := "model_do"
	modelDoValue := ""
	var tmpModelDo string
	if tmpModelDo, err = view.Parse(ctx, "go/model_do.template", tplData); err == nil {
		modelDoValue = tmpModelDo
		modelDoValue, err = s.trimBreak(modelDoValue)
	} else {
		return
	}
	controllerKey := "controller"
	controllerValue := ""
	var tmpController string
	if tmpController, err = view.Parse(ctx, "go/controller.template", tplData); err == nil {
		controllerValue = tmpController
		controllerValue, err = s.trimBreak(controllerValue)
	} else {
		return
	}

	serviceKey := "service"
	serviceValue := ""
	var tmpService string
	if tmpService, err = view.Parse(ctx, "go/service.template", tplData); err == nil {
		serviceValue = tmpService
		serviceValue, err = s.trimBreak(serviceValue)
	} else {
		return
	}

	logicKey := "logic"
	logicValue := ""
	var tmpLogic string
	if tmpLogic, err = view.Parse(ctx, "go/logic.template", tplData); err == nil {
		logicValue = tmpLogic
		logicValue, err = s.trimBreak(logicValue)
	} else {
		return
	}

	routerKey := "router"
	routerValue := ""
	var tmpRouter string
	if tmpRouter, err = view.Parse(ctx, "go/router.template", tplData); err == nil {
		routerValue = tmpRouter
		routerValue, err = s.trimBreak(routerValue)
	} else {
		return
	}

	routerFuncKey := "router_func"
	routerFuncValue := ""
	var tmpRouterFunc string
	if tmpRouterFunc, err = view.Parse(ctx, "go/router_func.template", tplData); err == nil {
		routerFuncValue = tmpRouterFunc
		routerFuncValue, err = s.trimBreak(routerFuncValue)
	} else {
		return
	}

	sqlKey := "sql"
	sqlValue := ""
	var tmpSql string
	var tmpSqlPath string
	if s.IsPg(extendData.Group) {
		tmpSqlPath = "sql/pgsql.template"
	} else {
		tmpSqlPath = "sql/mysql.template"
	}
	if tmpSql, err = view.Parse(ctx, tmpSqlPath, tplData); err == nil {
		sqlValue = tmpSql
		sqlValue, err = s.trimBreak(sqlValue)
	} else {
		return
	}

	tsApiKey := "tsApi"
	tsApiValue := ""
	var tmpTsApi string
	if tmpTsApi, err = view.Parse(ctx, "ts/api.template", tplData); err == nil {
		tsApiValue = tmpTsApi
		tsApiValue, err = s.trimBreak(tsApiValue)
	} else {
		return
	}

	tsModelKey := "tsModel"
	tsModelValue := ""
	var tmpTsModel string
	if tmpTsModel, err = view.Parse(ctx, "ts/model.template", tplData); err == nil {
		tsModelValue = tmpTsModel
		tsModelValue, err = s.trimBreak(tsModelValue)
	} else {
		return
	}

	vueKey := "vue"
	vueValue := ""
	var tmpVue string
	tmpFile := "vue/list-vue.template"
	if extendData.TplCategory == "tree" {
		if extendData.UseVirtual {
			//使用虚拟表树表
			tmpFile = "vue/tree-virtual-vue.template"
		} else {
			//树表
			tmpFile = "vue/tree-vue.template"
		}
	}
	if tmpVue, err = view.Parse(ctx, tmpFile, tplData); err == nil {
		vueValue = tmpVue
		vueValue, err = s.trimBreak(vueValue)
	} else {
		return
	}

	vueEditKey := "vueEdit"
	vueEditValue := ""
	var tmpVueEditModel string
	if tmpVueEditModel, err = view.Parse(ctx, "vue/edit-vue.template", tplData); err == nil {
		vueEditValue = tmpVueEditModel
		vueEditValue, err = s.trimBreak(vueEditValue)
	} else {
		return
	}

	vueDetailKey := "vueDetail"
	vueDetailValue := ""
	var tmpVueDetailModel string
	if tmpVueDetailModel, err = view.Parse(ctx, "vue/detail-vue.template", tplData); err == nil {
		vueDetailValue = tmpVueDetailModel
		vueDetailValue, err = s.trimBreak(vueDetailValue)
	} else {
		return
	}

	data = g.MapStrStr{
		apiKey:         apiValue,
		daoKey:         daoValue,
		daoInternalKey: daoInternalValue,
		modelKey:       modelValue,
		modelEntityKey: modelEntityValue,
		modelDoKey:     modelDoValue,
		controllerKey:  controllerValue,
		serviceKey:     serviceValue,
		logicKey:       logicValue,
		routerKey:      routerValue,
		routerFuncKey:  routerFuncValue,
		sqlKey:         sqlValue,
		tsApiKey:       tsApiValue,
		tsModelKey:     tsModelValue,
		vueKey:         vueValue,
		vueEditKey:     vueEditValue,
		vueDetailKey:   vueDetailValue,
	}
	return
}

func (s *sToolsGenTable) SelectRecordById(ctx context.Context, tableId int64) (tableEx *model.ToolsGenTableEx, err error) {
	var table *entity.ToolsGenTable
	table, err = s.GetTableInfoByTableId(ctx, tableId)
	if err != nil {
		return
	}
	m := gconv.Map(table)
	err = gconv.Struct(m, &tableEx)
	if err != nil {
		return
	}
	if tableEx.TplCategory == "tree" {
		opt := gjson.New(tableEx.Options)
		tableEx.TreeParentCode = opt.Get("treeParentCode").String()
		tableEx.TreeCode = opt.Get("treeCode").String()
		tableEx.TreeName = opt.Get("treeName").String()
	}
	//表字段数据
	var columns []*entity.ToolsGenTableColumn
	columns, err = service.ToolsGenTableColumn().SelectGenTableColumnListByTableId(ctx, tableId)
	if err != nil {
		return
	}
	for _, column := range columns {
		if column.GoType == "Time" {
			tableEx.HasTimeColumn = true
			continue
		}
	}

	allColumnExs := make([]*model.ToolsGenTableColumnEx, len(columns))

	var (
		editColumns   []*model.ToolsGenTableColumnEx
		listColumns   []*model.ToolsGenTableColumnEx
		detailColumns []*model.ToolsGenTableColumnEx
		queryColumns  []*model.ToolsGenTableColumnEx
	)

	columnMap := make(map[string]*model.ToolsGenTableColumnEx, len(columns))
	for i, column := range columns {
		columnName := column.ColumnName
		columnEx := &model.ToolsGenTableColumnEx{}
		m1 := gconv.Map(column)
		err = gconv.Struct(m1, columnEx)
		columnMap[columnName] = columnEx
		allColumnExs[i] = columnEx
		tableEx.IsPkInsertable = tableEx.IsPkInsertable || column.IsPk && !column.IsIncrement && !tableEx.UseSnowId
		tableEx.IsPkListable = tableEx.IsPkListable || column.IsPk && column.IsList
		if column.IsEdit && !service.ToolsGenTableColumn().IsNotEdit(columnName) && !column.IsPk {
			editColumns = append(editColumns, columnEx)
			columnEx.IsEdit = true
		} else {
			columnEx.IsEdit = false
		}
		if column.IsList && !service.ToolsGenTableColumn().IsNotList(columnName) {
			listColumns = append(listColumns, columnEx)
			columnEx.IsList = true
		} else {
			columnEx.IsList = false
		}
		if column.IsDetail && !service.ToolsGenTableColumn().IsNotDetail(columnName) {
			detailColumns = append(detailColumns, columnEx)
			columnEx.IsDetail = true
		} else {
			columnEx.IsDetail = false
		}
		if column.IsQuery && !service.ToolsGenTableColumn().IsNotQuery(columnName) {
			queryColumns = append(queryColumns, columnEx)
			columnEx.IsQuery = true
		} else {
			columnEx.IsQuery = false
		}
		// validation 规则 和 conversion 方法
		integerValidationRule := "integer"
		floatValidationRule := "float"
		dateValidationRule := "date"
		datetimeValidationRule := "datetime"
		if column.QueryType == "BETWEEN" {
			integerValidationRule += "-array"
			floatValidationRule += "-array"
			dateValidationRule += "-array"
			datetimeValidationRule += "-array"
		}
		if column.IsPk {
			//主键不处理
			continue
		}
		switch column.GoType {
		case "byte":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + integerValidationRule + "|max-length:1#" + column.ColumnComment + "需为整数"
			columnEx.FieldConversion = "gconv.Uint64"
			tableEx.HasConversion = true
			break
		case "int":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + integerValidationRule + "#" + column.ColumnComment + "需为整数"
			columnEx.FieldConversion = "gconv.Int"
			tableEx.HasConversion = true
			break
		case "int64":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + integerValidationRule + "#" + column.ColumnComment + "需为整数"
			columnEx.FieldConversion = "gconv.Int64"
			tableEx.HasConversion = true
			break
		case "uint":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + integerValidationRule + "#" + column.ColumnComment + "需为整数"
			columnEx.FieldConversion = "gconv.Uint"
			tableEx.HasConversion = true
			break
		case "uint64":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + integerValidationRule + "#" + column.ColumnComment + "需为整数"
			columnEx.FieldConversion = "gconv.Uint64"
			tableEx.HasConversion = true
			break
		case "float":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + floatValidationRule + "#" + column.ColumnComment + "需为浮点数"
			columnEx.FieldConversion = "gconv.Float"
			tableEx.HasConversion = true
			break
		case "float64":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + floatValidationRule + "#" + column.ColumnComment + "需为浮点数"
			columnEx.FieldConversion = "gconv.Float64"
			tableEx.HasConversion = true
			break
		case "bool":
			columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@boolean#" + column.ColumnComment + "需为true/false"
			columnEx.FieldConversion = "gconv.Bool"
			tableEx.HasConversion = true
			break
		case "date":
		case "Time":
			if column.ColumnType == "date" {
				columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + dateValidationRule + "#" + column.ColumnComment + "需为YYYY-MM-DD格式"
			} else {
				columnEx.FieldValidation = gstr.CaseCamelLower(columnName) + "@" + datetimeValidationRule + "#" + column.ColumnComment + "需为YYYY-MM-DD hh:mm:ss格式"
			}
			columnEx.FieldConversion = "gconv.Time"
			tableEx.HasConversion = true
			break
		}
	}
	sort.Slice(editColumns, func(i, j int) bool {
		return editColumns[i].SortOrderEdit < editColumns[j].SortOrderEdit
	})
	sort.Slice(listColumns, func(i, j int) bool {
		return listColumns[i].SortOrderList < listColumns[j].SortOrderList
	})
	sort.Slice(detailColumns, func(i, j int) bool {
		return detailColumns[i].SortOrderDetail < detailColumns[j].SortOrderDetail
	})
	sort.Slice(queryColumns, func(i, j int) bool {
		return queryColumns[i].SortOrderQuery < queryColumns[j].SortOrderQuery
	})

	tableEx.Columns = allColumnExs
	tableEx.EditColumns = editColumns
	tableEx.DetailColumns = detailColumns
	tableEx.ListColumns = listColumns
	tableEx.QueryColumns = queryColumns

	_, hasCreatedBy := columnMap["created_by"]
	tableEx.HasCreatedBy = hasCreatedBy
	_, hasDeptId := columnMap["dept_id"]
	tableEx.HasDeptId = hasDeptId
	_, hasUpdateBy := columnMap["updated_by"]
	tableEx.HasUpdatedBy = hasUpdateBy

	for _, column := range allColumnExs {
		if column.IsCascade {
			parentColumnName := column.ParentColumnName
			if g.IsEmpty(parentColumnName) {
				err = gerror.New("级联查询字段\"" + column.ColumnName + "\"并未设置parentColumnName")
				return nil, err
			}
			parent := columnMap[parentColumnName]
			if g.IsEmpty(parent) {
				err = gerror.New("级联查询字段\"" + column.ColumnName + "\"的parentColumnName\"" + parentColumnName + "\"不存在于表中")
				return nil, err
			}
			column.ParentColumnName = parentColumnName
			column.CascadeParent = parent
			parent.IsCascadeParent = true
			parent.CascadeChildrenColumns = gset.NewStrSet()
		}
	}
	for _, column := range allColumnExs {
		if column.IsCascade {
			child := column
			for {
				parentColumnName := child.ParentColumnName
				if g.IsEmpty(parentColumnName) {
					break
				}
				parent := columnMap[parentColumnName]
				parent.CascadeChildrenColumns.Add(column.ColumnName)
				child = parent
			}
		}
	}

	// 提取关联表信息
	var linkedTableMap = gmap.New()
	for _, column := range allColumnExs {
		if len(column.LinkTableName) == 0 {
			continue
		}
		linkedTableName := column.LinkTableName
		// 关联表信息
		linkedTableInfo, err1 := s.getLinkedTableInfo(ctx, linkedTableName)
		if err1 != nil {
			err = err1
			return
		}
		linkedTable := linkedTableMap.GetOrSet(linkedTableName, interface{}(linkedTableInfo)).(*model.ToolsGenTableLinked)

		allColumnMap, err2 := service.ToolsGenTableColumn().SelectDbTableColumnMapByName(ctx, linkedTableName, table.Group) //关联表只能是同数据库group的
		if err2 != nil {
			err = err2
			return
		}
		// 关联字段
		linkedTableIdColumn, foundId := allColumnMap[column.LinkLabelId]
		if !foundId {
			err = gerror.New("无法找到关联表的列" + column.LinkLabelId)
			return
		}
		linkedTableValueColumn, foundValue := allColumnMap[column.LinkLabelName]
		if !foundValue {
			err = gerror.New("无法找到关联表的列" + column.LinkLabelId)
			return
		}
		service.ToolsGenTableColumn().InitColumnField(linkedTableIdColumn, table)
		service.ToolsGenTableColumn().InitColumnField(linkedTableValueColumn, table)
		if linkedTableIdColumn.GoType == "Time" || linkedTableValueColumn.GoType == "Time" {
			tableEx.HasTimeColumn = true
		}
		linkedTable.RefColumns.GetOrSet(linkedTableIdColumn.ColumnName, linkedTableIdColumn)
		linkedTable.RefColumns.GetOrSet(linkedTableValueColumn.ColumnName, &linkedTableValueColumn)
	}

	tableEx.LinkedTables = make([]*model.ToolsGenTableLinked, linkedTableMap.Size())
	for i, v := range linkedTableMap.Values() {
		tableEx.LinkedTables[i] = v.(*model.ToolsGenTableLinked)
	}
	return
}

// GenCode 代码生成操作
func (s *sToolsGenTable) GenCode(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//获取当前运行时目录
		curDir, err := os.Getwd()
		liberr.ErrIsNil(ctx, err, "获取本地路径失败")
		frontDir := g.Cfg().MustGet(ctx, "gen.frontDir").String()
		if !gfile.IsDir(frontDir) {
			panic("项目前端路径不存在，请检查是否已在配置文件中配置！")
		}
		apiName := g.Cfg().MustGet(ctx, "gen.apiName").String()
		goModName := g.Cfg().MustGet(ctx, "gen.goModName").String()
		for _, id := range ids {
			var genData g.MapStrStr
			var extendData *model.ToolsGenTableEx
			genData, extendData, err = s.GenData(ctx, gconv.Int64(id))
			liberr.ErrIsNil(ctx, err)
			packageName := extendData.PackageName
			businessName := gstr.CaseCamelLower(extendData.BusinessName)
			modulePath := gstr.StrEx(extendData.PackageName, "internal/app/")
			overwriteInfo := gmap.NewStrAnyMap()
			for _, ov := range extendData.OverwriteInfo {
				overwriteInfo.Set(ov.Key, ov.Value)
			}
			for key, code := range genData {
				switch key {
				case "api":
					path := strings.Join([]string{curDir, "/", apiName, "/", modulePath, "/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("api")))
					liberr.ErrIsNil(ctx, err)
				case "controller":
					path := strings.Join([]string{curDir, "/", packageName, "/controller/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("controller")))
					liberr.ErrIsNil(ctx, err)
				case "dao":
					path := strings.Join([]string{curDir, "/", packageName, "/dao/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("dao")))
					liberr.ErrIsNil(ctx, err)
				case "dao_internal":
					path := strings.Join([]string{curDir, "/", packageName, "/dao/internal/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("dao_internal")))
					liberr.ErrIsNil(ctx, err)
				case "logic":
					path := strings.Join([]string{curDir, "/", packageName, "/logic/", businessName, "/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("logic")))
					liberr.ErrIsNil(ctx, err)
				case "model":
					path := strings.Join([]string{curDir, "/", packageName, "/model/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("model")))
					liberr.ErrIsNil(ctx, err)
				case "model_do":
					path := strings.Join([]string{curDir, "/", packageName, "/model/do/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("model_do")))
					liberr.ErrIsNil(ctx, err)
				case "model_entity":
					path := strings.Join([]string{curDir, "/", packageName, "/model/entity/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("model_entity")))
					liberr.ErrIsNil(ctx, err)
				case "router":
					if !gstr.ContainsI(packageName, "system") { // system 模块不生成router文件
						path := strings.Join([]string{curDir, "/", packageName, "/router/router", ".go"}, "")
						err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("router")))
						liberr.ErrIsNil(ctx, err)
					}
				case "router_func":
					path := strings.Join([]string{curDir, "/", packageName, "/router/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("router_func")))
					liberr.ErrIsNil(ctx, err)
				case "service":
					path := strings.Join([]string{curDir, "/", packageName, "/service/", extendData.BusinessName, ".go"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("service")))
					liberr.ErrIsNil(ctx, err)
				case "sql":
					path := strings.Join([]string{curDir, "/resource/data/gen_sql/", modulePath, "/", extendData.BusinessName, ".sql"}, "")
					hasSql := gfile.Exists(path)
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("sql")))
					liberr.ErrIsNil(ctx, err)
					if !hasSql || gconv.Bool(overwriteInfo.Get("sql")) {
						//第一次生成则向数据库写入菜单数据
						if s.IsPg(extendData.Group) {
							content := gfile.GetContents(path)
							_, err = g.DB().Exec(ctx, content)
						} else {
							err = s.writeDb(ctx, path)
						}
						liberr.ErrIsNil(ctx, err)
						//清除菜单缓存
						commonService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
					}
				case "tsApi":
					path := strings.Join([]string{frontDir, "/src/api/", modulePath, "/", businessName, ".ts"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("tsApi")))
					liberr.ErrIsNil(ctx, err)
				case "tsModel":
					path := strings.Join([]string{frontDir, "/src/views/", modulePath, "/", businessName + "/list/component/model", ".ts"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("tsModel")))
					liberr.ErrIsNil(ctx, err)
				case "vue":
					path := strings.Join([]string{frontDir, "/src/views/", modulePath, "/", businessName, "/list/index.vue"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("vue")))
					liberr.ErrIsNil(ctx, err)
				case "vueDetail":
					path := strings.Join([]string{frontDir, "/src/views/", modulePath, "/", businessName + "/list/component/detail", ".vue"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("vueDetail")))
					liberr.ErrIsNil(ctx, err)
				case "vueEdit":
					path := strings.Join([]string{frontDir, "/src/views/", modulePath, "/", businessName + "/list/component/edit", ".vue"}, "")
					err = s.createFile(path, code, gconv.Bool(overwriteInfo.Get("vueEdit")))
					liberr.ErrIsNil(ctx, err)
				}
			}
			//生成模块路由
			err = s.genModuleRouter(curDir, goModName, extendData.ModuleName, modulePath)
			liberr.ErrIsNil(ctx, err)
			//生成模块boot logic
			err = s.genModuleBootLogic(curDir, extendData.ModuleName, modulePath)
			liberr.ErrIsNil(ctx, err)
			//生成对应模块的业务logic
			err = s.genModuleLogic(curDir, goModName, extendData.PackageName)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// 写入菜单数据
func (s *sToolsGenTable) writeDb(ctx context.Context, path string) (err error) {
	isAnnotation := false
	var fi *os.File
	fi, err = os.Open(path)
	if err != nil {
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var sqlStr []string
	now := gtime.Now()
	var res sql.Result
	var id int64
	var tx gdb.TX
	tx, err = g.DB().Ctx(ctx).Begin(ctx)
	if err != nil {
		return
	}
	for {
		bytes, e := br.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		str := gstr.Trim(string(bytes))

		if str == "" {
			continue
		}

		if strings.Contains(str, "/*") {
			isAnnotation = true
		}

		if isAnnotation {
			if strings.Contains(str, "*/") {
				isAnnotation = false
			}
			continue
		}

		if str == "" || strings.HasPrefix(str, "--") || strings.HasPrefix(str, "#") {
			continue
		}
		if strings.HasSuffix(str, ";") {
			if gstr.ContainsI(str, "select") {
				if gstr.ContainsI(str, "@now") {
					continue
				}
				if gstr.ContainsI(str, "@parentId") {
					id, err = res.LastInsertId()
				}
			}
			sqlStr = append(sqlStr, str)
			sql := strings.Join(sqlStr, "")
			gstr.ReplaceByArray(sql, []string{"@parentId", gconv.String(id), "@now", now.Format("Y-m-d H:i:s")})
			//插入业务
			res, err = tx.Exec(sql)
			if err != nil {
				tx.Rollback()
				return
			}
			sqlStr = nil
		} else {
			sqlStr = []string{str}
		}
	}
	tx.Commit()
	return
}

// createFile 创建文件
func (s *sToolsGenTable) createFile(fileName, data string, cover bool) (err error) {
	if !gfile.Exists(fileName) || cover {
		var f *os.File
		f, err = gfile.Create(fileName)
		if err == nil {
			f.WriteString(data)
		}
		f.Close()
		err = s.goFmt(fileName)
	}
	return
}

func (s *sToolsGenTable) getLinkedTableInfo(ctx context.Context, tableName string) (linkedTable *model.ToolsGenTableLinked, err error) {
	var table *entity.ToolsGenTable
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ToolsGenTable.Ctx(ctx).Where(dao.ToolsGenTable.Columns().TableName, tableName).Scan(&table)
		liberr.ErrIsNil(ctx, err, "获取关联表格"+tableName+"信息出错")
		m := gconv.Map(table)
		err = gconv.Struct(m, &linkedTable)
		liberr.ErrIsNil(ctx, err)
		if linkedTable.Options != "" {
			err = gconv.Struct(linkedTable.Options, &linkedTable.OptionsStruct)
			liberr.ErrIsNil(ctx, err)
			err = dao.ToolsGenTableColumn.Ctx(ctx).Where(dao.ToolsGenTableColumn.Columns().TableId, table.TableId).
				Where(dao.ToolsGenTableColumn.Columns().ColumnName, gstr.CaseSnake(linkedTable.OptionsStruct.TreeParentCode)).
				Scan(&linkedTable.OptionsStruct.ColumnAttr)
			liberr.ErrIsNil(ctx, err)
		}
		linkedTable.RefColumns = gmap.NewListMap()
	})
	return
}

// 剔除多余的换行
func (s *sToolsGenTable) trimBreak(str string) (rStr string, err error) {
	var b []byte
	if b, err = gregex.Replace("(([\\s\t]*)\r?\n){2,}", []byte("$2\n"), []byte(str)); err != nil {
		return
	}
	if b, err = gregex.Replace("(([\\s\t]*)/{4}\r?\n)", []byte("$2\n\n"), b); err == nil {
		rStr = gconv.String(b)
	}
	return
}

// GenModuleRouter 生成模块logic
func (s *sToolsGenTable) genModuleLogic(curDir, goModName, packageName string) (err error) {
	var (
		packages []string
		code     string
	)
	//logic路径
	path := strings.Join([]string{curDir, "/", packageName, "/logic"}, "")
	//获取该路径下所有需要加载的包
	packages, err = gfile.ScanDirFunc(path, "*", false, func(p string) string {
		if gfile.IsDir(p) {
			return p
		}
		return ""
	})
	if err != nil {
		return
	}
	for _, v := range packages {
		v = gstr.Replace(v, "\\", "/")
		v = gstr.SubStr(v, gstr.PosR(v, "/"))
		code += strings.Join([]string{"import _ ", `"`, goModName, "/", packageName, "/logic", v, `"`, "\n"}, "")
	}
	code = fmt.Sprintf(`package logic%s%s`, "\n\n", code)
	err = s.createFile(path+"/logic.go", code, true)
	return
}

func (s *sToolsGenTable) genModuleRouter(curDir, goModName, moduleName, modulePath string) (err error) {
	modulePathName := gstr.CaseCamelLower(gstr.Replace(modulePath, "/", "_"))
	path := strings.Join([]string{curDir, "/internal/router/" + modulePathName + ".go"}, "")
	if gfile.IsFile(path) || moduleName == "system" {
		return
	}
	moduleNameUpper := gstr.CaseCamel(modulePathName)
	code := fmt.Sprintf(`package router
import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	%sRouter "%s/internal/app/%s/router"
)

func (router *Router) Bind%sModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	%sRouter.R.BindController(ctx, group)
}`, moduleName, goModName, modulePath, moduleNameUpper, moduleName)
	err = s.createFile(path, code, true)
	return
}

func (s *sToolsGenTable) genModuleBootLogic(curDir, moduleName, modulePath string) (err error) {
	modulePathName := gstr.CaseCamelLower(gstr.Replace(modulePath, "/", "_"))
	path := strings.Join([]string{curDir, "/internal/app/boot/" + modulePathName + ".go"}, "")
	if gfile.IsFile(path) || moduleName == "system" {
		return
	}
	code := fmt.Sprintf(`package boot
import (
	_ "github.com/tiger1103/gfast/v3/internal/app/%s/logic"
)
`, modulePath)
	err = s.createFile(path, code, true)
	return
}

func (s *sToolsGenTable) SyncTable(ctx context.Context, tableId int64) (err error) {
	var (
		extendData      *model.ToolsGenTableEx
		table           *entity.ToolsGenTable
		genTableColumns []*entity.ToolsGenTableColumn
	)
	extendData, err = s.SelectRecordById(ctx, tableId)
	if err != nil {
		return
	}
	if extendData == nil {
		err = gerror.New("表格数据不存在")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			table, err = s.GetTableInfoByTableId(ctx, tableId)
			liberr.ErrIsNil(ctx, err)
			if table == nil {
				liberr.ErrIsNil(ctx, gerror.New("表格数据不存在"))
			}
			genTableColumns, err = service.ToolsGenTableColumn().SelectDbTableColumnsByName(ctx, extendData.TableName, extendData.Group)
			liberr.ErrIsNil(ctx, err, "获取列数据失败")
			if len(genTableColumns) <= 0 {
				liberr.ErrIsNil(ctx, gerror.New("获取列数据失败"))
			}
			for _, column := range genTableColumns {
				alreadyExists := false //字段是否存在
				for _, tableColumn := range extendData.Columns {
					if column.ColumnName == tableColumn.ColumnName {
						alreadyExists = true
						break
					}
				}
				//字段不存在则添加
				if !alreadyExists {
					service.ToolsGenTableColumn().InitColumnField(column, table)
					_, err = tx.Model(dao.ToolsGenTableColumn.Table()).
						FieldsEx(dao.ToolsGenTableColumn.Columns().ColumnId).
						Insert(column)
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("保存列`%s`数据失败", column.ColumnName))
				}
			}
			for _, tableColumn := range extendData.Columns {
				alreadyDelete := true //数据表中字段是否已删除
				for _, column := range genTableColumns {
					if column.ColumnName == tableColumn.ColumnName {
						alreadyDelete = false
						break
					}
				}
				if alreadyDelete {
					//删除columns表中字段
					_, err = dao.ToolsGenTableColumn.Ctx(ctx).
						Where(dao.ToolsGenTableColumn.Columns().ColumnId, tableColumn.ColumnId).
						Delete()
					liberr.ErrIsNil(ctx, err)
				}
			}
		})
		return err
	})
	return
}

// goFmt formats the source file and adds or removes import statements as necessary.
func (s *sToolsGenTable) goFmt(path string) (err error) {
	replaceFunc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			g.Log().Printf(context.Background(), `error format "%s" go files: %v`, path, err)
			return content
		}
		return string(res)
	}
	if gfile.IsFile(path) {
		// File format.
		if gfile.ExtName(path) != "go" {
			return
		}
		err = gfile.ReplaceFileFunc(replaceFunc, path)
	} else {
		// Folder format.
		err = gfile.ReplaceDirFunc(replaceFunc, path, "*.go", true)
	}
	if err != nil {
		err = fmt.Errorf(`error format "%s" go files: %v`, path, err)
	}
	return
}
