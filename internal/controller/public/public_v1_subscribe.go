package public

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gov2panel/api/public/v1"
)

func (c *ControllerV1) Subscribe(ctx context.Context, req *v1.SubscribeReq) (res *v1.SubscribeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
