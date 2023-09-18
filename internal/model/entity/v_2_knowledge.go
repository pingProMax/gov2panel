// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Knowledge is the golang structure for table v2_knowledge.
type V2Knowledge struct {
	Id        int         `json:"id"         ` //
	Category  string      `json:"category"   ` // 分類名
	Title     string      `json:"title"      ` // 標題
	Body      string      `json:"body"       ` // 內容
	OrderId   int         `json:"order_id"   ` // 排序
	Show      int         `json:"show"       ` // 顯示
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}
