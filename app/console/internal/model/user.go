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
