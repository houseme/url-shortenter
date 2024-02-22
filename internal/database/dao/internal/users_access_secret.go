// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersAccessSecretDao is the data access object for table users_access_secret.
type UsersAccessSecretDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns UsersAccessSecretColumns // columns contains all the column names of Table for convenient usage.
}

// UsersAccessSecretColumns defines and stores column names for table users_access_secret.
type UsersAccessSecretColumns struct {
	Id         string // ID
	AccountNo  string // 企业管理员 ID 一致
	SecretId   string // 授权应用 ID
	SecretKey  string // 授权应用 key
	Salt       string // 盐值
	SaltKey    string // 盐值 key
	GrantType  string // 授权类型：默认空，API token 授权:client_credentials
	State      string // 状态 0 默认 100 正常 200 禁用
	CreateTime string // 创建时间
	ModifyTime string // 修改时间
}

// usersAccessSecretColumns holds the columns for table users_access_secret.
var usersAccessSecretColumns = UsersAccessSecretColumns{
	Id:         "id",
	AccountNo:  "account_no",
	SecretId:   "secret_id",
	SecretKey:  "secret_key",
	Salt:       "salt",
	SaltKey:    "salt_key",
	GrantType:  "grant_type",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewUsersAccessSecretDao creates and returns a new DAO object for table data access.
func NewUsersAccessSecretDao() *UsersAccessSecretDao {
	return &UsersAccessSecretDao{
		group:   "default",
		table:   "users_access_secret",
		columns: usersAccessSecretColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersAccessSecretDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersAccessSecretDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersAccessSecretDao) Columns() UsersAccessSecretColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersAccessSecretDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersAccessSecretDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersAccessSecretDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
