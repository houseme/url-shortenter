// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDomainDao is the data access object for table users_domain.
type UsersDomainDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns UsersDomainColumns // columns contains all the column names of Table for convenient usage.
}

// UsersDomainColumns defines and stores column names for table users_domain.
type UsersDomainColumns struct {
	Id          string // ID
	AccountNo   string // 账号标识
	DomainNo    string // 域名标识
	Domain      string // 域名 不需要 http 等协议信息
	Memo        string // 备注信息
	License     string // icp 备案号
	SubLicense  string // icp 备案号 带序号
	State       string // 状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用
	CreateTime  string // 创建时间
	AuditTime   string // 审核时间
	DisableTime string // 禁用时间
	ModifyTime  string // 修改时间
}

// usersDomainColumns holds the columns for table users_domain.
var usersDomainColumns = UsersDomainColumns{
	Id:          "id",
	AccountNo:   "account_no",
	DomainNo:    "domain_no",
	Domain:      "domain",
	Memo:        "memo",
	License:     "license",
	SubLicense:  "sub_license",
	State:       "state",
	CreateTime:  "create_time",
	AuditTime:   "audit_time",
	DisableTime: "disable_time",
	ModifyTime:  "modify_time",
}

// NewUsersDomainDao creates and returns a new DAO object for table data access.
func NewUsersDomainDao() *UsersDomainDao {
	return &UsersDomainDao{
		group:   "default",
		table:   "users_domain",
		columns: usersDomainColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDomainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDomainDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDomainDao) Columns() UsersDomainColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDomainDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDomainDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDomainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
