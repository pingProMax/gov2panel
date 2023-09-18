package model

import "gov2panel/internal/model/entity"

type RechargeRecordsInfo struct {
	V2User            *entity.V2User            `json:"user"`
	V2RechargeRecords *entity.V2RechargeRecords `json:"recharge_records"`
}
