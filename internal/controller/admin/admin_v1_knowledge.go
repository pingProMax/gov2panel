package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"
)

func (c *ControllerV1) Knowledge(ctx context.Context, req *v1.KnowledgeReq) (res *v1.KnowledgeRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "knowledge", nil)
	case "POST":
		res = &v1.KnowledgeRes{}
		res.Data, err = service.Knowledge().GetKnowledgeAllList(req.V2Knowledge)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) KnowledgeAE(ctx context.Context, req *v1.KnowledgeAEReq) (res *v1.KnowledgeAERes, err error) {
	err = service.Knowledge().AEKnowledge(&req.V2Knowledge)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) KnowledgeDel(ctx context.Context, req *v1.KnowledgeDelReq) (res *v1.KnowledgeDelRes, err error) {
	err = service.Knowledge().DelKnowledge(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}
