package model

// Base is the base model
type Base struct {
	AuthAccountNo    uint64 `json:"authAccountNo" dc:"认证账号"`
	AuthAccountLevel uint   `json:"authAccountLevel" dc:"账号级别"`
}

// DefaultHandlerResponse .
type DefaultHandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Time    int64       `json:"time"`    // 返回当前响应时间
	TraceID string      `json:"traceID"` // 请求唯一标识
}
