// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersMerchant is the golang structure of table users_merchant for DAO operations like Where/Data.
type UsersMerchant struct {
	g.Meta              `orm:"table:users_merchant, do:true"`
	Id                  interface{} // ID
	AccountNo           interface{} // 合作伙伴 用户 ID
	AppKey              interface{} // 渠道 Key 合作来源 ID
	AppSecret           interface{} // 渠道合作密钥
	MerName             interface{} // 昵称
	MerAvatar           interface{} // 企业头像
	MerMobile           interface{} // 登陆账号
	MerType             interface{} // 用户类型  1 企业、2 个人 (自然人)、3 个体工商户
	CreateDate          interface{} // 开户日期格式:YYYYMMDD
	CorpName            interface{} // 企业的公司全称
	SocialCreditCode    interface{} // 企业的统一社会信用代码
	LicenseStartDate    interface{} // 企业营业执照起始日期
	LicenseEndDate      interface{} // 证件有效期为长期填写:99991231
	LicenseImg          interface{} // 证件照片
	CorpBusinessAddress interface{} // 企业的营业地址
	CorpRegAddress      interface{} // 企业的注册地址
	CorpFixedTelephone  interface{} // 企业固定电话
	BusinessScope       interface{} // 企业经营范围
	LegalName           interface{} // 企业法人姓名
	LegalCertType       interface{} // 参考证件类型
	LegalCertId         interface{} // 与证件类型对应
	LegalCertStartDate  interface{} // 证件起始日期
	LegalCertEndDate    interface{} // 证件有效期为长期填写:99991231
	LegalCertFront      interface{} // 身份证正面
	LegalCertBack       interface{} // 身份证反面
	LegalMobile         interface{} // 法人手机号码
	ContactName         interface{} // 企业联系姓名
	ContactMobile       interface{} // 企业联系人手机号
	ContractEmail       interface{} // 企业联系人邮箱地址
	AgentNo             interface{} // 代理商标识
	StartUsing          interface{} // 是否开始使用 0 默认，100 开始，110 暂停，120 禁用
	StartUseTime        *gtime.Time // 开始使用时间
	State               interface{} // 用户状态 60 待审核，80 审核中，100 正常，210 禁用，110 审核拒绝，120 补充资料中
	CreateTime          *gtime.Time // 创建时间
	ModifyTime          *gtime.Time // 更新时间
}
