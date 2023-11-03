package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `path:"/user" tags:"User" method:"get,post" summary:"用户管理"`
	SortOrder
	OffsetLimit

	*entity.V2User

	BalanceS           string `json:"balanceS"            ` // 账户余额
	DiscountS          string `json:"discountS"           ` // 专享折扣
	CommissionRateS    string `json:"commission_rateS"    ` // 返利比例
	CommissionBalanceS string `json:"commission_balanceS" ` // aff余额
	TS                 string `json:"tS"                  ` // 最后在线时间戳
	US                 string `json:"uS"                  ` // 已使用流量
	DS                 string `json:"dS"                  ` // 已使用流量
	TransferEnableS    string `json:"transfer_enableS"    ` // 流量
	ExpiredAtS         string `json:"expired_atS"         ` // 到期时间
	CreatedAtS         string `json:"created_atS"         ` // 创建时间
	UpdatedAtS         string `json:"updated_atS"         ` // 更新时间
}
type UserRes struct {
	g.Meta   `mime:"text/html" example:"string"`
	UserList []*model.UserInfo `json:"user_list"`
	Totle    int               `json:"totle"`
}

type UserAEReq struct {
	g.Meta `path:"/user/ae" tags:"User" method:"post" summary:"用户管理AE"`
	entity.V2User
}
type UserAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UserDelReq struct {
	g.Meta `path:"/user/del" tags:"User" method:"post" summary:"用户管理删除api"`
	Ids    []int `json:"ids"`
}
type UserDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UserUpBanned1Req struct {
	g.Meta `path:"/user/banned1" tags:"User" method:"post" summary:"用户管理冻结api"`
	Ids    []int `json:"ids"`
}
type UserUpBanned1Res struct {
	g.Meta `mime:"text/html" example:"string"`
}

type LogoutReq struct {
	g.Meta  `path:"/logout" tags:"User" method:"get" summary:"退出登录"`
	TUserID int
}
type LogoutRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type RefreshReq struct {
	g.Meta  `path:"/refresh" tags:"User" method:"post" summary:"刷新token"`
	TUserID int
}
type RefreshRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type ResetTokenAndUuidReq struct {
	g.Meta `path:"/user/reset_token_uuid" tags:"User" method:"post" summary:"重置用户uuid和token"`
	UserId int `json:"user_id"`
}
type ResetTokenAndUuidRes struct {
}
