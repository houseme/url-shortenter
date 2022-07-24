package consts

const (
	Logger = "schedule"
	// ShortValid 有效状态 0默认 100正常 200失效
	ShortValid = 100
	// ShortInvalid 失效状态
	ShortInvalid = 200

	// ShortCollectStateSuccess 收集状态 0默认 100成功 200失败
	ShortCollectStateSuccess = 100
	// ShortCollectStateFailed 失败状态
	ShortCollectStateFailed = 200
	// ShortCollectStateProcessing 正在收集状态
	ShortCollectStateProcessing = 0

	// VisitState 访问状态 0默认 100正常 200失效
	VisitState = 0
	// VisitStateNormal 访问状态 0默认 100正常 200失效
	VisitStateNormal = 100
	// VisitStateInvalid 失效状态
	VisitStateInvalid = 200

	// ShortAccountNo 帐号
	ShortAccountNo = "short_url_account_no_"
	// ShortShortNo 短网址
	ShortShortNo = "short_url_short_no_"

	// ContentTypeMirror 内容类型 0默认 100 镜像 ，200审核
	ContentTypeMirror = 100
	// ContentTypeAudit 内容类型 0默认 100 镜像 ，200审核
	ContentTypeAudit = 200
	// ContentTypeDefault 内容类型 0默认 100 镜像 ，200审核
	ContentTypeDefault = 0
)
