package user

import (
	"context"
	"fmt"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Knowledge(ctx context.Context, req *v1.KnowledgeReq) (res *v1.KnowledgeRes, err error) {
	res = &v1.KnowledgeRes{}
	res.Data, err = service.Knowledge().GetKnowledgeShowList(req.V2Knowledge)
	fmt.Println(req.Id)
	if req.Id == 0 {
		setTplUser(ctx, "knowledge", g.Map{"data": res.Data})
		return
	} else {
		setTplUser(ctx, "knowledge_id", g.Map{"data": res.Data})
		return
	}
}
