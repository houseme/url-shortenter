// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersMerchantDao is the data access object for table users_merchant.
type UsersMerchantDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns UsersMerchantColumns // columns contains all the column names of Table for convenient usage.
}

// UsersMerchantColumns defines and stores column names for table users_merchant.
type UsersMerchantColumns struct {
	Id                  string // ID
	AccountNo           string // 合作伙伴 用户 ID
	AppKey              string // 渠道 Key 合作来源 ID
	AppSecret           string // 渠道合作密钥
	MerName             string // 昵称
	MerAvatar           string // 企业头像
	MerMobile           string // 登陆账号
	MerType             string // 用户类型  1 企业、2 个人 (自然人)、3 个体工商户
	CreateDate          string // 开户日期格式:YYYYMMDD
	CorpName            string // 企业的公司全称
	SocialCreditCode    string // 企业的统一社会信用代码
	LicenseStartDate    string // 企业营业执照起始日期
	LicenseEndDate      string // 证件有效期为长期填写:99991231
	LicenseImg          string // 证件照片
	CorpBusinessAddress string // 企业的营业地址
	CorpRegAddress      string // 企业的注册地址
	CorpFixedTelephone  string // 企业固定电话
	BusinessScope       string // 企业经营范围
	LegalName           string // 企业法人姓名
	LegalCertType       string // 参考证件类型
	LegalCertId         string // 与证件类型对应
	LegalCertStartDate  string // 证件起始日期
	LegalCertEndDate    string // 证件有效期为长期填写:99991231
	LegalCertFront      string // 身份证正面
	LegalCertBack       string // 身份证反面
	LegalMobile         string // 法人手机号码
	ContactName         string // 企业联系姓名
	ContactMobile       string // 企业联系人手机号
	ContractEmail       string // 企业联系人邮箱地址
	AgentNo             string // 代理商标识
	StartUsing          string // 是否开始使用 0 默认，100 开始，110 暂停，120 禁用
	StartUseTime        string // 开始使用时间
	State               string // 用户状态 60 待审核，80 审核中，100 正常，210 禁用，110 审核拒绝，120 补充资料中
	CreateTime          string // 创建时间
	ModifyTime          string // 更新时间
}

// usersMerchantColumns holds the columns for table users_merchant.
var usersMerchantColumns = UsersMerchantColumns{
	Id:                  "id",
	AccountNo:           "account_no",
	AppKey:              "app_key",
	AppSecret:           "app_secret",
	MerName:             "mer_name",
	MerAvatar:           "mer_avatar",
	MerMobile:           "mer_mobile",
	MerType:             "mer_type",
	CreateDate:          "create_date",
	CorpName:            "corp_name",
	SocialCreditCode:    "social_credit_code",
	LicenseStartDate:    "license_start_date",
	LicenseEndDate:      "license_end_date",
	LicenseImg:          "license_img",
	CorpBusinessAddress: "corp_business_address",
	CorpRegAddress:      "corp_reg_address",
	CorpFixedTelephone:  "corp_fixed_telephone",
	BusinessScope:       "business_scope",
	LegalName:           "legal_name",
	LegalCertType:       "legal_cert_type",
	LegalCertId:         "legal_cert_id",
	LegalCertStartDate:  "legal_cert_start_date",
	LegalCertEndDate:    "legal_cert_end_date",
	LegalCertFront:      "legal_cert_front",
	LegalCertBack:       "legal_cert_back",
	LegalMobile:         "legal_mobile",
	ContactName:         "contact_name",
	ContactMobile:       "contact_mobile",
	ContractEmail:       "contract_email",
	AgentNo:             "agent_no",
	StartUsing:          "start_using",
	StartUseTime:        "start_use_time",
	State:               "state",
	CreateTime:          "create_time",
	ModifyTime:          "modify_time",
}

// NewUsersMerchantDao creates and returns a new DAO object for table data access.
func NewUsersMerchantDao() *UsersMerchantDao {
	return &UsersMerchantDao{
		group:   "default",
		table:   "users_merchant",
		columns: usersMerchantColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersMerchantDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersMerchantDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersMerchantDao) Columns() UsersMerchantColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersMerchantDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersMerchantDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersMerchantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
