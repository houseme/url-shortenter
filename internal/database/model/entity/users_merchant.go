// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersMerchant is the golang structure for table users_merchant.
type UsersMerchant struct {
	Id                  uint64      `json:"id"                  description:"ID"`
	AccountNo           uint64      `json:"accountNo"           description:"合作伙伴 用户 ID"`
	AppKey              uint64      `json:"appKey"              description:"渠道 Key 合作来源 ID"`
	AppSecret           string      `json:"appSecret"           description:"渠道合作密钥"`
	MerName             string      `json:"merName"             description:"昵称"`
	MerAvatar           string      `json:"merAvatar"           description:"企业头像"`
	MerMobile           string      `json:"merMobile"           description:"登陆账号"`
	MerType             uint        `json:"merType"             description:"用户类型  1 企业、2 个人 (自然人)、3 个体工商户"`
	CreateDate          string      `json:"createDate"          description:"开户日期格式:YYYYMMDD"`
	CorpName            string      `json:"corpName"            description:"企业的公司全称"`
	SocialCreditCode    string      `json:"socialCreditCode"    description:"企业的统一社会信用代码"`
	LicenseStartDate    string      `json:"licenseStartDate"    description:"企业营业执照起始日期"`
	LicenseEndDate      string      `json:"licenseEndDate"      description:"证件有效期为长期填写:99991231"`
	LicenseImg          string      `json:"licenseImg"          description:"证件照片"`
	CorpBusinessAddress string      `json:"corpBusinessAddress" description:"企业的营业地址"`
	CorpRegAddress      string      `json:"corpRegAddress"      description:"企业的注册地址"`
	CorpFixedTelephone  string      `json:"corpFixedTelephone"  description:"企业固定电话"`
	BusinessScope       string      `json:"businessScope"       description:"企业经营范围"`
	LegalName           string      `json:"legalName"           description:"企业法人姓名"`
	LegalCertType       string      `json:"legalCertType"       description:"参考证件类型"`
	LegalCertId         string      `json:"legalCertId"         description:"与证件类型对应"`
	LegalCertStartDate  string      `json:"legalCertStartDate"  description:"证件起始日期"`
	LegalCertEndDate    string      `json:"legalCertEndDate"    description:"证件有效期为长期填写:99991231"`
	LegalCertFront      string      `json:"legalCertFront"      description:"身份证正面"`
	LegalCertBack       string      `json:"legalCertBack"       description:"身份证反面"`
	LegalMobile         string      `json:"legalMobile"         description:"法人手机号码"`
	ContactName         string      `json:"contactName"         description:"企业联系姓名"`
	ContactMobile       string      `json:"contactMobile"       description:"企业联系人手机号"`
	ContractEmail       string      `json:"contractEmail"       description:"企业联系人邮箱地址"`
	AgentNo             uint64      `json:"agentNo"             description:"代理商标识"`
	StartUsing          uint        `json:"startUsing"          description:"是否开始使用 0 默认，100 开始，110 暂停，120 禁用"`
	StartUseTime        *gtime.Time `json:"startUseTime"        description:"开始使用时间"`
	State               uint        `json:"state"               description:"用户状态 60 待审核，80 审核中，100 正常，210 禁用，110 审核拒绝，120 补充资料中"`
	CreateTime          *gtime.Time `json:"createTime"          description:"创建时间"`
	ModifyTime          *gtime.Time `json:"modifyTime"          description:"更新时间"`
}
