package model

// ShortRedis is a struct for redis data
type ShortRedis struct {
	ShortNo   uint64 `json:"shortNo" dc:"shortNo 短链编号"`
	AccountNo uint64 `json:"accountNo" dc:"accountNo 账户编号"`
	DestURL   string `json:"destUrl" dc:"destUrl 目标链接"`
	IsValid   uint   `json:"isValid" dc:"isValid 是否可用 0默认 100正常 200失效"`
}
