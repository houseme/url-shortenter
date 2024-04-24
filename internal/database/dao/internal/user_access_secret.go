// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserAccessSecretDao is the data access object for table user_access_secret.
type UserAccessSecretDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserAccessSecretColumns // columns contains all the column names of Table for convenient usage.
}

// UserAccessSecretColumns defines and stores column names for table user_access_secret.
type UserAccessSecretColumns struct {
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

// userAccessSecretColumns holds the columns for table user_access_secret.
var userAccessSecretColumns = UserAccessSecretColumns{
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

// NewUserAccessSecretDao creates and returns a new DAO object for table data access.
func NewUserAccessSecretDao() *UserAccessSecretDao {
	return &UserAccessSecretDao{
		group:   "default",
		table:   "user_access_secret",
		columns: userAccessSecretColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserAccessSecretDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserAccessSecretDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserAccessSecretDao) Columns() UserAccessSecretColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserAccessSecretDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserAccessSecretDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserAccessSecretDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
