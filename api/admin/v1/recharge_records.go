package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RechargeRecordsReq struct {
	g.Meta `path:"/recharge_records" tags:"RechargeRecords" method:"get,post" summary:"充值消费记录"`
	SortOrder
	OffsetLimit
	entity.V2RechargeRecords
	UserName string `json:"user_name"`
}
type RechargeRecordsRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.RechargeRecordsInfo `json:"data"`
	Totle  int                          `json:"totle"`
}

type RechargeRecordsAddReq struct {
	g.Meta `path:"/recharge_records/add" tags:"RechargeRecords" method:"post" summary:"充值消费记录添加"`
	entity.V2RechargeRecords
}
type RechargeRecordsAddRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type RechargeRecordsUpRemarksReq struct {
	g.Meta  `path:"/recharge_records/up_remarks" tags:"RechargeRecords" method:"post" summary:"更新备注"`
	Id      int    `json:"id"`
	Remarks string `json:"remarks"`
}
type RechargeRecordsUpRemarksRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
