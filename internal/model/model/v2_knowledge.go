package model

import "gov2panel/internal/model/entity"

type KnowledgeInfo struct {
	Category string                `json:"category"`
	Data     []*entity.V2Knowledge `json:"data"`
}
