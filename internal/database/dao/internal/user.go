// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id         string // ID
	UserNo     string // 账号唯一标识
	Username   string // 用户名称
	Account    string // 账号
	Password   string // 密码
	State      string // 状态 0 默认 100 正常 200 失效
	Avatar     string // 头像
	GroupLevel string // 用户等级 0 默认超级 1000 企业管理员，10000 普通管理员
	AccountNo  string // 用户关联企业 ID 同企业管理员 ID 一致
	Deadline   string // 截止时间
	Salt       string // 盐值
	CreateTime string // 创建时间
	ModifyTime string // 更新时间
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:         "id",
	UserNo:     "user_no",
	Username:   "username",
	Account:    "account",
	Password:   "password",
	State:      "state",
	Avatar:     "avatar",
	GroupLevel: "group_level",
	AccountNo:  "account_no",
	Deadline:   "deadline",
	Salt:       "salt",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
