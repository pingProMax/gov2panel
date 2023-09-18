package public

import (
	"context"

	v1 "gov2panel/api/public/v1"
)

func (c *ControllerV1) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	setTplPublc(ctx, "index", nil)
	// return res, err
	return
}
