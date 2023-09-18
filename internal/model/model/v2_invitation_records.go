package model

import "gov2panel/internal/model/entity"

type InvitationRecordsInfo struct {
	User              *entity.V2User              `json:"user"`
	FromUser          *entity.V2User              `json:"from_user"`
	InvitationRecords *entity.V2InvitationRecords `json:"invitation_records"`
}
