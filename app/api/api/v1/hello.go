package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HelloReq is a hello request.
type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}

// HelloRes is a hello response.
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
