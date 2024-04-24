// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserTagRelationDao is the data access object for table user_tag_relation.
type UserTagRelationDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns UserTagRelationColumns // columns contains all the column names of Table for convenient usage.
}

// UserTagRelationColumns defines and stores column names for table user_tag_relation.
type UserTagRelationColumns struct {
	Id         string // ID
	UserNo     string // 用户标识
	TagNo      string // 标签标识
	State      string // 状态 0 默认 100 正常 200 失效
	CreateTime string // 创建时间
	ModifyTime string // 更新时间
}

// userTagRelationColumns holds the columns for table user_tag_relation.
var userTagRelationColumns = UserTagRelationColumns{
	Id:         "id",
	UserNo:     "user_no",
	TagNo:      "tag_no",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewUserTagRelationDao creates and returns a new DAO object for table data access.
func NewUserTagRelationDao() *UserTagRelationDao {
	return &UserTagRelationDao{
		group:   "default",
		table:   "user_tag_relation",
		columns: userTagRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserTagRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserTagRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserTagRelationDao) Columns() UserTagRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserTagRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserTagRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserTagRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
