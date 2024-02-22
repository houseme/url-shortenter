// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TotalCountTop25Dao is the data access object for table total_count_top25.
type TotalCountTop25Dao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns TotalCountTop25Columns // columns contains all the column names of Table for convenient usage.
}

// TotalCountTop25Columns defines and stores column names for table total_count_top25.
type TotalCountTop25Columns struct {
	ShortUrl        string // 短链内容
	TodayCount      string //
	YesterdayCount  string //
	Last7DaysCount  string //
	MonthlyCount    string //
	TotalCount      string //
	DTodayCount     string //
	DYesterdayCount string //
	DLast7DaysCount string //
	DMonthlyCount   string //
	DTotalCount     string //
	Id              string // ID
	DestUrl         string // 原始链接
	CreateTime      string // 创建时间
	IsValid         string // 是否可用 0 默认 100 正常 200 失效
	Memo            string // 备注
}

// totalCountTop25Columns holds the columns for table total_count_top25.
var totalCountTop25Columns = TotalCountTop25Columns{
	ShortUrl:        "short_url",
	TodayCount:      "today_count",
	YesterdayCount:  "yesterday_count",
	Last7DaysCount:  "last_7_days_count",
	MonthlyCount:    "monthly_count",
	TotalCount:      "total_count",
	DTodayCount:     "d_today_count",
	DYesterdayCount: "d_yesterday_count",
	DLast7DaysCount: "d_last_7_days_count",
	DMonthlyCount:   "d_monthly_count",
	DTotalCount:     "d_total_count",
	Id:              "id",
	DestUrl:         "dest_url",
	CreateTime:      "create_time",
	IsValid:         "is_valid",
	Memo:            "memo",
}

// NewTotalCountTop25Dao creates and returns a new DAO object for table data access.
func NewTotalCountTop25Dao() *TotalCountTop25Dao {
	return &TotalCountTop25Dao{
		group:   "default",
		table:   "total_count_top25",
		columns: totalCountTop25Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TotalCountTop25Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TotalCountTop25Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TotalCountTop25Dao) Columns() TotalCountTop25Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TotalCountTop25Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TotalCountTop25Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TotalCountTop25Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
