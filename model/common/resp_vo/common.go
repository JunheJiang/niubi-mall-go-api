package resp_vo

type PageResult struct {
	List       interface{} `json:"list"`
	TotalCount int64       `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	CurPage    int         `json:"curPage"`
	PageSize   int         `json:"pageSize"`
}
