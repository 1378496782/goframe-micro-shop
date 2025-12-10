package model

// BaseDao 基础数据访问对象
type BaseDao struct {
	Table string // 表名
}

// NewBaseDao 创建基础数据访问对象
func NewBaseDao(table string) BaseDao {
	return BaseDao{
		Table: table,
	}
}
