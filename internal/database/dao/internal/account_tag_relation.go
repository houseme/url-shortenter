// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountTagRelationDao is the data access object for table account_tag_relation.
type AccountTagRelationDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns AccountTagRelationColumns // columns contains all the column names of Table for convenient usage.
}

// AccountTagRelationColumns defines and stores column names for table account_tag_relation.
type AccountTagRelationColumns struct {
	Id         string // ID
	AccountNo  string // 用户标识
	TagNo      string // 标签标识
	State      string // 状态 0 默认 100 正常 200 失效
	CreateTime string // 创建时间
	ModifyTime string // 更新时间
}

// accountTagRelationColumns holds the columns for table account_tag_relation.
var accountTagRelationColumns = AccountTagRelationColumns{
	Id:         "id",
	AccountNo:  "account_no",
	TagNo:      "tag_no",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewAccountTagRelationDao creates and returns a new DAO object for table data access.
func NewAccountTagRelationDao() *AccountTagRelationDao {
	return &AccountTagRelationDao{
		group:   "default",
		table:   "account_tag_relation",
		columns: accountTagRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccountTagRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountTagRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccountTagRelationDao) Columns() AccountTagRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountTagRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountTagRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountTagRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
