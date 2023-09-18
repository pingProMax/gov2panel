package v1

//公用 struct

//排序用
type SortOrder struct {
	Sort  string `json:"sort"`
	Order string `json:"order"`
}

//分页用
type OffsetLimit struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
