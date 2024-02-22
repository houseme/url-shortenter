// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortMirrorDao is the data access object for table short_mirror.
type ShortMirrorDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ShortMirrorColumns // columns contains all the column names of Table for convenient usage.
}

// ShortMirrorColumns defines and stores column names for table short_mirror.
type ShortMirrorColumns struct {
	Id             string // ID
	ShortNo        string // 短链标识
	DestUrl        string // 原始 URL
	FullScreenshot string // 整屏镜像
	ContentPath    string // 内容文件地址
	Content        string // 网页内容
	HashContent    string // hash 值 sha256
	CreateTime     string // 创建时间
	ModifyTime     string // 修改时间
}

// shortMirrorColumns holds the columns for table short_mirror.
var shortMirrorColumns = ShortMirrorColumns{
	Id:             "id",
	ShortNo:        "short_no",
	DestUrl:        "dest_url",
	FullScreenshot: "full_screenshot",
	ContentPath:    "content_path",
	Content:        "content",
	HashContent:    "hash_content",
	CreateTime:     "create_time",
	ModifyTime:     "modify_time",
}

// NewShortMirrorDao creates and returns a new DAO object for table data access.
func NewShortMirrorDao() *ShortMirrorDao {
	return &ShortMirrorDao{
		group:   "default",
		table:   "short_mirror",
		columns: shortMirrorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortMirrorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortMirrorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortMirrorDao) Columns() ShortMirrorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortMirrorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortMirrorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortMirrorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
