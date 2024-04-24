// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMerchant is the golang structure for table user_merchant.
type UserMerchant struct {
	Id                  uint64      `json:"id"                  orm:"id"                    description:"ID"`
	AccountNo           uint64      `json:"accountNo"           orm:"account_no"            description:"合作伙伴 用户 ID"`
	UserNo              uint64      `json:"userNo"              orm:"user_no"               description:"用户标识"`
	AppKey              uint64      `json:"appKey"              orm:"app_key"               description:"渠道 Key 合作来源 ID"`
	AppSecret           string      `json:"appSecret"           orm:"app_secret"            description:"渠道合作密钥"`
	MerName             string      `json:"merName"             orm:"mer_name"              description:"昵称"`
	MerAvatar           string      `json:"merAvatar"           orm:"mer_avatar"            description:"企业头像"`
	MerMobile           string      `json:"merMobile"           orm:"mer_mobile"            description:"登陆账号"`
	MerType             uint        `json:"merType"             orm:"mer_type"              description:"用户类型  1 企业、2 个人 (自然人)、3 个体工商户"`
	CreateDate          string      `json:"createDate"          orm:"create_date"           description:"开户日期格式:YYYYMMDD"`
	CorpName            string      `json:"corpName"            orm:"corp_name"             description:"企业的公司全称"`
	SocialCreditCode    string      `json:"socialCreditCode"    orm:"social_credit_code"    description:"企业的统一社会信用代码"`
	LicenseStartDate    string      `json:"licenseStartDate"    orm:"license_start_date"    description:"企业营业执照起始日期"`
	LicenseEndDate      string      `json:"licenseEndDate"      orm:"license_end_date"      description:"证件有效期为长期填写:99991231"`
	LicenseImg          string      `json:"licenseImg"          orm:"license_img"           description:"证件照片"`
	CorpBusinessAddress string      `json:"corpBusinessAddress" orm:"corp_business_address" description:"企业的营业地址"`
	CorpRegAddress      string      `json:"corpRegAddress"      orm:"corp_reg_address"      description:"企业的注册地址"`
	CorpFixedTelephone  string      `json:"corpFixedTelephone"  orm:"corp_fixed_telephone"  description:"企业固定电话"`
	BusinessScope       string      `json:"businessScope"       orm:"business_scope"        description:"企业经营范围"`
	LegalName           string      `json:"legalName"           orm:"legal_name"            description:"企业法人姓名"`
	LegalCertType       string      `json:"legalCertType"       orm:"legal_cert_type"       description:"参考证件类型"`
	LegalCertId         string      `json:"legalCertId"         orm:"legal_cert_id"         description:"与证件类型对应"`
	LegalCertStartDate  string      `json:"legalCertStartDate"  orm:"legal_cert_start_date" description:"证件起始日期"`
	LegalCertEndDate    string      `json:"legalCertEndDate"    orm:"legal_cert_end_date"   description:"证件有效期为长期填写:99991231"`
	LegalCertFront      string      `json:"legalCertFront"      orm:"legal_cert_front"      description:"身份证正面"`
	LegalCertBack       string      `json:"legalCertBack"       orm:"legal_cert_back"       description:"身份证反面"`
	LegalMobile         string      `json:"legalMobile"         orm:"legal_mobile"          description:"法人手机号码"`
	ContactName         string      `json:"contactName"         orm:"contact_name"          description:"企业联系姓名"`
	ContactMobile       string      `json:"contactMobile"       orm:"contact_mobile"        description:"企业联系人手机号"`
	ContractEmail       string      `json:"contractEmail"       orm:"contract_email"        description:"企业联系人邮箱地址"`
	AgentNo             uint64      `json:"agentNo"             orm:"agent_no"              description:"代理商标识"`
	StartUsing          uint        `json:"startUsing"          orm:"start_using"           description:"是否开始使用 0 默认，100 开始，110 暂停，120 禁用"`
	StartUseTime        *gtime.Time `json:"startUseTime"        orm:"start_use_time"        description:"开始使用时间"`
	State               uint        `json:"state"               orm:"state"                 description:"用户状态 60 待审核，80 审核中，100 正常，210 禁用，110 审核拒绝，120 补充资料中"`
	CreateTime          *gtime.Time `json:"createTime"          orm:"create_time"           description:"创建时间"`
	ModifyTime          *gtime.Time `json:"modifyTime"          orm:"modify_time"           description:"更新时间"`
}
