// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDomainAuditDao is the data access object for table users_domain_audit.
type UsersDomainAuditDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UsersDomainAuditColumns // columns contains all the column names of Table for convenient usage.
}

// UsersDomainAuditColumns defines and stores column names for table users_domain_audit.
type UsersDomainAuditColumns struct {
	Id             string // ID
	AccountNo      string // 用户标识
	DomainNo       string // 用户认证域名 ID
	Icp            string // icp 备案号
	QueryResult    string // 查询结果
	AuditAccountNo string // 审核用户 ID
	AuditTime      string // 审核时间
	AuditState     string // 审核状态 0 默认 20:审核通过 30:审核失败
	CreateTime     string // 创建时间
	ModifyTime     string // 修改时间
}

// usersDomainAuditColumns holds the columns for table users_domain_audit.
var usersDomainAuditColumns = UsersDomainAuditColumns{
	Id:             "id",
	AccountNo:      "account_no",
	DomainNo:       "domain_no",
	Icp:            "icp",
	QueryResult:    "query_result",
	AuditAccountNo: "audit_account_no",
	AuditTime:      "audit_time",
	AuditState:     "audit_state",
	CreateTime:     "create_time",
	ModifyTime:     "modify_time",
}

// NewUsersDomainAuditDao creates and returns a new DAO object for table data access.
func NewUsersDomainAuditDao() *UsersDomainAuditDao {
	return &UsersDomainAuditDao{
		group:   "default",
		table:   "users_domain_audit",
		columns: usersDomainAuditColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDomainAuditDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDomainAuditDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDomainAuditDao) Columns() UsersDomainAuditColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDomainAuditDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDomainAuditDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDomainAuditDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
