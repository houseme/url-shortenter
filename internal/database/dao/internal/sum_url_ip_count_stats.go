// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SumUrlIpCountStatsDao is the data access object for table sum_url_ip_count_stats.
type SumUrlIpCountStatsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns SumUrlIpCountStatsColumns // columns contains all the column names of Table for convenient usage.
}

// SumUrlIpCountStatsColumns defines and stores column names for table sum_url_ip_count_stats.
type SumUrlIpCountStatsColumns struct {
	TodayCount  string //
	DTodayCount string //
}

// sumUrlIpCountStatsColumns holds the columns for table sum_url_ip_count_stats.
var sumUrlIpCountStatsColumns = SumUrlIpCountStatsColumns{
	TodayCount:  "today_count",
	DTodayCount: "d_today_count",
}

// NewSumUrlIpCountStatsDao creates and returns a new DAO object for table data access.
func NewSumUrlIpCountStatsDao() *SumUrlIpCountStatsDao {
	return &SumUrlIpCountStatsDao{
		group:   "default",
		table:   "sum_url_ip_count_stats",
		columns: sumUrlIpCountStatsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SumUrlIpCountStatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SumUrlIpCountStatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SumUrlIpCountStatsDao) Columns() SumUrlIpCountStatsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SumUrlIpCountStatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SumUrlIpCountStatsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SumUrlIpCountStatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
