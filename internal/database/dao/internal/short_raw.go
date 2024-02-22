// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortRawDao is the data access object for table short_raw.
type ShortRawDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ShortRawColumns // columns contains all the column names of Table for convenient usage.
}

// ShortRawColumns defines and stores column names for table short_raw.
type ShortRawColumns struct {
	Id         string // ID
	ShortNo    string // 链接标识
	RawNo      string // 唯一标识
	ShortRaw   string // 额外参数标识
	ShortValue string // 额外参数对应 value
	ShortKey   string // 链接参数 key
	State      string // 0 默认 100 正常 200 失效
	CreateTime string // 创建时间
	ModifyTime string // 更新时间
}

// shortRawColumns holds the columns for table short_raw.
var shortRawColumns = ShortRawColumns{
	Id:         "id",
	ShortNo:    "short_no",
	RawNo:      "raw_no",
	ShortRaw:   "short_raw",
	ShortValue: "short_value",
	ShortKey:   "short_key",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewShortRawDao creates and returns a new DAO object for table data access.
func NewShortRawDao() *ShortRawDao {
	return &ShortRawDao{
		group:   "default",
		table:   "short_raw",
		columns: shortRawColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortRawDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortRawDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortRawDao) Columns() ShortRawColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortRawDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortRawDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortRawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
