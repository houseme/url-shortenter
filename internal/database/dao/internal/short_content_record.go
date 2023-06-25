// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortContentRecordDao is the data access object for table short_content_record.
type ShortContentRecordDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ShortContentRecordColumns // columns contains all the column names of Table for convenient usage.
}

// ShortContentRecordColumns defines and stores column names for table short_content_record.
type ShortContentRecordColumns struct {
	Id          string // ID
	ShortNo     string // 短链标识
	TrxId       string // 唯一标识
	ContentType string // 内容类型 0 默认，100 镜像内容 200 审核内容
	Content     string // 网页内容
	HashContent string // hash 值 sha256
	CreateTime  string // 创建时间
	ModifyTime  string // 修改时间
}

// shortContentRecordColumns holds the columns for table short_content_record.
var shortContentRecordColumns = ShortContentRecordColumns{
	Id:          "id",
	ShortNo:     "short_no",
	TrxId:       "trx_id",
	ContentType: "content_type",
	Content:     "content",
	HashContent: "hash_content",
	CreateTime:  "create_time",
	ModifyTime:  "modify_time",
}

// NewShortContentRecordDao creates and returns a new DAO object for table data access.
func NewShortContentRecordDao() *ShortContentRecordDao {
	return &ShortContentRecordDao{
		group:   "default",
		table:   "short_content_record",
		columns: shortContentRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortContentRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortContentRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortContentRecordDao) Columns() ShortContentRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortContentRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortContentRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortContentRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
