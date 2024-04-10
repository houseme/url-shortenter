// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortTagDao is the data access object for table short_tag.
type ShortTagDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ShortTagColumns // columns contains all the column names of Table for convenient usage.
}

// ShortTagColumns defines and stores column names for table short_tag.
type ShortTagColumns struct {
	Id         string // ID
	TagNo      string // 唯一标识
	TagName    string // 标签名称
	TagPinyin  string // 名称拼音 中划线链接
	TagEn      string // 名称英文 中划线链接
	State      string // 状态 0 默认 100 正常 200 失效
	CreateTime string // 创建时间
	ModifyTime string // 修改时间
}

// shortTagColumns holds the columns for table short_tag.
var shortTagColumns = ShortTagColumns{
	Id:         "id",
	TagNo:      "tag_no",
	TagName:    "tag_name",
	TagPinyin:  "tag_pinyin",
	TagEn:      "tag_en",
	State:      "state",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// NewShortTagDao creates and returns a new DAO object for table data access.
func NewShortTagDao() *ShortTagDao {
	return &ShortTagDao{
		group:   "default",
		table:   "short_tag",
		columns: shortTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortTagDao) Columns() ShortTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
