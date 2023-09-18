package model

import "gov2panel/internal/model/entity"

type TicketMessageInfo struct {
	V2TicketMessage *entity.V2TicketMessage `json:"ticket"`
	V2User          *entity.V2User          `json:"user"`
}
