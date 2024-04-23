// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccessLogsSummaryDao is the data access object for table access_logs_summary.
type AccessLogsSummaryDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns AccessLogsSummaryColumns // columns contains all the column names of Table for convenient usage.
}

// AccessLogsSummaryColumns defines and stores column names for table access_logs_summary.
type AccessLogsSummaryColumns struct {
	Id         string // ID
	SummaryNo  string // 统计记录唯一标识
	AccountNo  string // 账号标识
	UserNo     string // 用户标识
	ShortNo    string // 短链标识
	ShortUrl   string // 短链内容
	ShortAll   string // 带参数 URL
	YearTime   string // 年份
	MonthTime  string // 月份
	DayTime    string // 日期
	AccessDate string // 访问日期
	UserAgent  string // 访问 user_agent
	Ip         string // 访问 IP
	Summary    string // 访问汇总
	SuccessSum string // 成功
	FailSum    string // 失败
	CreateTime string // 创建时间
	ModifyTime string // 修改时间
}

// accessLogsSummaryColumns holds the columns for table access_logs_summary.
var accessLogsSummaryColumns = AccessLogsSummaryColumns{
	Id:         "id",
	SummaryNo:  "summary_no",
	AccountNo:  "account_no",
	UserNo:     "user_no",
	ShortNo:    "short_no",
	ShortUrl:   "short_url",
	ShortAll:   "short_all",
	YearTime:   "year_time",
	MonthTime:  "month_time",
	DayTime:    "day_time",
	AccessDate: "access_date",
	UserAgent:  "user_agent",
	Ip:         "ip",
	Summary:    "summary",
	SuccessSum: "success_sum",
	FailSum:    "fail_sum",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewAccessLogsSummaryDao creates and returns a new DAO object for table data access.
func NewAccessLogsSummaryDao() *AccessLogsSummaryDao {
	return &AccessLogsSummaryDao{
		group:   "default",
		table:   "access_logs_summary",
		columns: accessLogsSummaryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccessLogsSummaryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccessLogsSummaryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccessLogsSummaryDao) Columns() AccessLogsSummaryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccessLogsSummaryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccessLogsSummaryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccessLogsSummaryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
