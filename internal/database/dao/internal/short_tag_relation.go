// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortTagRelationDao is the data access object for table short_tag_relation.
type ShortTagRelationDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns ShortTagRelationColumns // columns contains all the column names of Table for convenient usage.
}

// ShortTagRelationColumns defines and stores column names for table short_tag_relation.
type ShortTagRelationColumns struct {
	Id         string // ID
	AccountNo  string // 用户标识
	TagNo      string // 标签标识
	ShortNo    string // 短链的唯一 ID
	State      string // 状态 0 默认 100 正常 200 失效
	CreateTime string // 创建时间
	ModifyTime string // 更新时间
}

// shortTagRelationColumns holds the columns for table short_tag_relation.
var shortTagRelationColumns = ShortTagRelationColumns{
	Id:         "id",
	AccountNo:  "account_no",
	TagNo:      "tag_no",
	ShortNo:    "short_no",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewShortTagRelationDao creates and returns a new DAO object for table data access.
func NewShortTagRelationDao() *ShortTagRelationDao {
	return &ShortTagRelationDao{
		group:   "default",
		table:   "short_tag_relation",
		columns: shortTagRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortTagRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortTagRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortTagRelationDao) Columns() ShortTagRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortTagRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortTagRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortTagRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
