// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortDomainDao is the data access object for table short_domain.
type ShortDomainDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ShortDomainColumns // columns contains all the column names of Table for convenient usage.
}

// ShortDomainColumns defines and stores column names for table short_domain.
type ShortDomainColumns struct {
	Id         string // ID
	DomainNo   string // 域名编号
	Domain     string // 短链域名
	Memo       string // 备注
	State      string // 状态  0 默认 100 正常 200 失效
	CreateTime string // 创建时间
	ModifyTime string // 更新时间
}

// shortDomainColumns holds the columns for table short_domain.
var shortDomainColumns = ShortDomainColumns{
	Id:         "id",
	DomainNo:   "domain_no",
	Domain:     "domain",
	Memo:       "memo",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewShortDomainDao creates and returns a new DAO object for table data access.
func NewShortDomainDao() *ShortDomainDao {
	return &ShortDomainDao{
		group:   "default",
		table:   "short_domain",
		columns: shortDomainColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortDomainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortDomainDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortDomainDao) Columns() ShortDomainColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortDomainDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortDomainDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortDomainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
