package model

import "gov2panel/internal/model/entity"

type TicketInfo struct {
	V2Ticket *entity.V2Ticket `json:"ticket"`
	V2User   *entity.V2User   `json:"user"`
}
