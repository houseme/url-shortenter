package model

// CreateAccountInput is the input for CreateAccount
type CreateAccountInput struct {
	*Base      `json:"-"`
	Account    string `json:"account" dc:"账号" v:"required|passport#账号唯一标识|账号包含字母、数字和下划线，长度在6~18之间"`
	Password   string `json:"password" dc:"密码" v:"required|password2#请填写密码|密码需要6-18位,必须包含大小写字母和数字"`
	GroupLevel uint   `json:"groupLevel" dc:"用户登记" v:"required|integer|in,0,1000,10000#用户等级|必须是整数|用户等级只能是0,1000,10000"`
}

// CreateAccountOutput is the output for CreateAccount
type CreateAccountOutput bool

// ModifyAccountInput is the input for ModifyAccount
type ModifyAccountInput struct {
	*Base    `json:"-"`
	Account  string `json:"account" dc:"账号" v:"required|passport#账号唯一标识|账号包含字母、数字和下划线，长度在6~18之间"`
	Password string `json:"password" dc:"密码" v:"required|password2#请填写密码|密码需要6-18位,必须包含大小写字母和数字"`
}

// ModifyAccountOutput is the output for ModifyAccount
type ModifyAccountOutput bool

// ModifyPasswordInput is the input for ModifyPassword
type ModifyPasswordInput struct {
	*Base    `json:"-"`
	Password string `json:"password" dc:"密码" v:"required|password2#请填写密码|密码需要6-18位,必须包含大小写字母和数字"`
	Code     string `json:"code" dc:"验证码" v:"required|between:4,6#请填写验证码|验证码长度为4-6位"`
}

// ModifyPasswordOutput is the output for ModifyPassword
type ModifyPasswordOutput bool
