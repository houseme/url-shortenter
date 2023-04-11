// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// CreateMerchantInput is the input of CreateMerchant.
type CreateMerchantInput struct {
	*Base               `json:"-"`
	MerName             string `json:"merName" desc:"商户名称" v:"required#商户名称不能为空"`
	MerAvatar           string `json:"merAvatar" desc:"商户头像" v:"required#商户头像不能为空"`
	MerMobile           string `json:"merMobile" desc:"商户手机号" v:"required#商户手机号不能为空"`
	MerType             uint   `json:"merType" desc:"用户类型  1企业、 2个人(自然人)、3个体工商户" v:"required|in:1,2,3#用户类型不能为空|用户类型只能是1,2,3"`
	CreateDate          string `json:"createDate" desc:"创建时间" v:"required#创建时间不能为空"`
	CorpName            string `json:"corpName" desc:"企业名称" v:"required#企业名称不能为空"`
	SocialCreditCode    string `json:"socialCreditCode" desc:"统一社会信用代码" v:"required#统一社会信用代码不能为空"`
	LicenseStartDate    string `json:"licenseStartDate" desc:"营业执照有效期开始时间" v:"required#营业执照有效期开始时间不能为空"`
	LicenseEndDate      string `json:"licenseEndDate" desc:"营业执照有效期结束时间" v:"required#营业执照有效期结束时间不能为空"`
	LicenseImg          string `json:"licenseImg" desc:"营业执照图片" v:"required#营业执照图片不能为空"`
	CorpBusinessAddress string `json:"corpBusinessAddress" desc:"企业经营地址" v:"required#企业经营地址不能为空"`
	CorpRegAddress      string `json:"corpRegAddress" desc:"企业注册地址" v:"required#企业注册地址不能为空"`
	CorpFixedTelephone  string `json:"corpFixedTelephone" desc:"企业固定电话" v:"required#企业固定电话不能为空"`
	BusinessScope       string `json:"businessScope" desc:"经营范围" v:"required#经营范围不能为空"`
	LegalName           string `json:"legalName" desc:"法人姓名" v:"required#法人姓名不能为空"`
	LegalCertType       uint   `json:"legalCertType" desc:"法人证件类型 1身份证、2护照、3军官证、4港澳通行证、5台湾通行证" v:"required|in:1,2,3,4,5#法人证件类型不能为空|法人证件类型只能是1,2,3,4,5"`
	LegalCertID         string `json:"legalCertID" desc:"法人证件号码" v:"required#法人证件号码不能为空"`
	LegalCertStartDate  string `json:"legalCertStartDate" desc:"法人证件有效期开始时间" v:"required#法人证件有效期开始时间不能为空"`
	LegalCertEndDate    string `json:"legalCertEndDate" desc:"法人证件有效期结束时间" v:"required#法人证件有效期结束时间不能为空"`
	LegalCertFront      string `json:"legalCertFront" desc:"法人证件正面照" v:"required#法人证件正面照不能为空"`
	LegalCertBack       string `json:"legalCertBack" desc:"法人证件反面照" v:"required#法人证件反面照不能为空"`
	LegalMobile         string `json:"legalMobile" desc:"法人手机号" v:"required#法人手机号不能为空"`
	ContactName         string `json:"contactName" desc:"联系人姓名" v:"required#联系人姓名不能为空"`
	ContactMobile       string `json:"contactMobile" desc:"联系人手机号" v:"required#联系人手机号不能为空"`
	ContactEmail        string `json:"contactEmail" desc:"联系人邮箱" v:"required#联系人邮箱不能为空"`
	AgentNo             uint64 `json:"agentNo" desc:"代理商编号" v:"required#代理商编号不能为空"`
}

// CreateMerchantOutput is the output of CreateMerchant.
type CreateMerchantOutput struct {
}

// CreateKeySecretInput is the input of CreateKeySecret.
type CreateKeySecretInput struct {
	*Base `json:"-"`
}

// CreateKeySecretOutput is the output of CreateKeySecret.
type CreateKeySecretOutput struct {
	AppKey    string `json:"appKey" description:"渠道合作标识"`
	AppSecret string `json:"appSecret" description:"渠道合作密钥"`
}

// QueryMerchantInput is the input of QueryMerchant.
type QueryMerchantInput struct {
	*Base `json:"-"`
}

// QueryMerchantOutput is the output of QueryMerchant.
type QueryMerchantOutput struct {
	MerNo               uint64 `json:"merNo" description:"商户编号"`
	MerName             string `json:"merName" description:"商户名称"`
	MerAvatar           string `json:"merAvatar" description:"商户头像"`
	MerMobile           string `json:"merMobile" description:"商户手机号"`
	MerType             uint   `json:"merType" description:"用户类型  1企业、 2个人(自然人)、3个体工商户"`
	CreateDate          string `json:"createDate" description:"创建时间"`
	CorpName            string `json:"corpName" description:"企业名称"`
	SocialCreditCode    string `json:"socialCreditCode" description:"统一社会信用代码"`
	LicenseStartDate    string `json:"licenseStartDate" description:"营业执照有效期开始时间"`
	LicenseEndDate      string `json:"licenseEndDate" description:"营业执照有效期结束时间"`
	LicenseImg          string `json:"licenseImg" description:"营业执照图片"`
	CorpBusinessAddress string `json:"corpBusinessAddress" description:"企业经营地址"`
	CorpRegAddress      string `json:"corpRegAddress" description:"企业注册地址"`
	CorpFixedTelephone  string `json:"corpFixedTelephone" description:"企业固定电话"`
	BusinessScope       string `json:"businessScope" description:"经营范围"`
	LegalName           string `json:"legalName" description:"法人姓名"`
	LegalCertType       uint   `json:"legalCertType" description:"法人证件类型 1身份证、2护照、3军官证、4港澳通行证、5台湾通行证"`
	LegalCertID         string `json:"legalCertID" description:"法人证件号码"`
	LegalCertStartDate  string `json:"legalCertStartDate" description:"法人证件有效期开始时间"`
	LegalCertEndDate    string `json:"legalCertEndDate" description:"法人证件有效期结束时间"`
	LegalCertFront      string `json:"legalCertFront" description:"法人证件正面照"`
	LegalCertBack       string `json:"legalCertBack" description:"法人证件反面照"`
	LegalMobile         string `json:"legalMobile" description:"法人手机号"`
	ContactName         string `json:"contactName" description:"联系人姓名"`
	ContactMobile       string `json:"contactMobile" description:"联系人手机号"`
	ContactEmail        string `json:"contactEmail" description:"联系人邮箱"`
	AgentNo             uint64 `json:"agentNo" description:"代理商编号"`
	CreateTime          string `json:"createTime" description:"创建时间"`
}

// UserDetailInput is the input of UserDetail.
type UserDetailInput struct {
	*Base `json:"-"`
}

// UserDetailOutput is the output of UserDetail.
type UserDetailOutput struct {
	Username string `json:"username" description:"用户名"`
	Avatar   string `json:"avatar" description:"头像"`
}
