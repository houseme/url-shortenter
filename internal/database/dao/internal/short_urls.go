// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortUrlsDao is the data access object for table short_urls.
type ShortUrlsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ShortUrlsColumns // columns contains all the column names of Table for convenient usage.
}

// ShortUrlsColumns defines and stores column names for table short_urls.
type ShortUrlsColumns struct {
	Id            string // ID
	AccountNo     string // 所属用户表示
	UserNo        string // 用户表示
	ShortNo       string // 短链的唯一 ID
	ShortUrl      string // 短链内容
	ShortDomain   string // 短链域名
	ShortDomainNo string // 短链域名标识
	DestUrl       string // 原始链接
	DestHash      string // 原始链接 Hash
	Domain        string // 主域名
	IsValid       string // 是否可用 0 默认 100 正常 200 失效
	DisableTime   string // 失效时间
	Memo          string // 备注
	RawState      string // 原始状态 0 默认 100 附带额外参数
	Sort          string // 排序字段
	CollectState  string // 镜像采集状态 0 默认 100 已采集 200 采集失败
	CollectTime   string // 采集时间
	DelState      string // 是否删除 0 默认 200 删除
	DelTime       string // 删除时间
	CreateTime    string // 创建时间
	ModifyTime    string // 修改时间
}

// shortUrlsColumns holds the columns for table short_urls.
var shortUrlsColumns = ShortUrlsColumns{
	Id:            "id",
	AccountNo:     "account_no",
	UserNo:        "user_no",
	ShortNo:       "short_no",
	ShortUrl:      "short_url",
	ShortDomain:   "short_domain",
	ShortDomainNo: "short_domain_no",
	DestUrl:       "dest_url",
	DestHash:      "dest_hash",
	Domain:        "domain",
	IsValid:       "is_valid",
	DisableTime:   "disable_time",
	Memo:          "memo",
	RawState:      "raw_state",
	Sort:          "sort",
	CollectState:  "collect_state",
	CollectTime:   "collect_time",
	DelState:      "del_state",
	DelTime:       "del_time",
	CreateTime:    "create_time",
	ModifyTime:    "modify_time",
}

// NewShortUrlsDao creates and returns a new DAO object for table data access.
func NewShortUrlsDao() *ShortUrlsDao {
	return &ShortUrlsDao{
		group:   "default",
		table:   "short_urls",
		columns: shortUrlsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortUrlsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortUrlsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortUrlsDao) Columns() ShortUrlsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortUrlsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortUrlsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortUrlsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
