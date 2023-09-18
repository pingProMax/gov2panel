package admin

import (
	"context"

	v1 "gov2panel/api/admin/v1"
)

func (c *ControllerV1) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	setTplAdmin(ctx, "index", nil)
	return
}
