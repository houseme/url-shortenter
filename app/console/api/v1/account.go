package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// CreateAccountReq is the request struct for the CreateAccount endpoint.
type CreateAccountReq struct {
	g.Meta `path:"/account" tags:"Account Service" method:"Post" summary:"create an account"`
	*model.CreateAccountInput
}

// CreateAccountRes is the response struct for the CreateAccount endpoint.
type CreateAccountRes struct {
	*model.CreateAccountOutput
}

// ModifyAccountReq is the request struct for the ModifyAccount endpoint.
type ModifyAccountReq struct {
	g.Meta `path:"/account/:account/update" tags:"Account Service" method:"Post" summary:"You say, I modify an account"`
	*model.ModifyAccountInput
}

// ModifyAccountRes is the response struct for the ModifyAccount endpoint.
type ModifyAccountRes struct {
	*model.ModifyAccountOutput
}
