package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" tags:"Index" method:"get" summary:"首页"`
}
type IndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
