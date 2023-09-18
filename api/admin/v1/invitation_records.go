package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type InvitationRecordsReq struct {
	g.Meta `path:"/invitation_records" tags:"InvitationRecords" method:"get,post" summary:"邀请记录"`
	SortOrder
	OffsetLimit
	entity.V2InvitationRecords
	UserName     string `json:"user_name"`
	FromUserName string `json:"from_user_name"`
}
type InvitationRecordsRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.InvitationRecordsInfo `json:"data"`
	Totle  int                            `json:"totle"`
}

type InvitationRecordsUpStateReq struct {
	g.Meta `path:"/invitation_records/state" tags:"InvitationRecords" method:"post" summary:"审核状态api"`
	Id     int `json:"id"`
	State  int `json:"state"`
}
type InvitationRecordsUpStateRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
