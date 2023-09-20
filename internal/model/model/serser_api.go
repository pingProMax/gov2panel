package model

//后端 上报用户流量用
type UserTraffic struct {
	UID      int
	Email    string
	Upload   int64
	Download int64
}
