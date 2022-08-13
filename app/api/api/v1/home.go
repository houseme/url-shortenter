package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/api/internal/model"
)

// HomeReq is the request struct for api.v1.Home.
type HomeReq struct {
	g.Meta `path:"/:short" tags:"home" method:"get" summary:"You first home api"`
	*model.HomeInput
}

// HomeRes is the response struct.
type HomeRes string
