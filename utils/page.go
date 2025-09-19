package utils

// 分页查找的实现，对应之后argodesion的前端UI
type PageRequest struct {
	// 分页总数
	PageSize uint `json:"page_size"`
	// 当前分页
	PageNum uint `json:"page_num"`
}

func NewPageRequest() *PageRequest {
	return &PageRequest{
		// 默认20页
		PageSize: 20,
		// 默认第一页
		PageNum: 1,
	}
}

// 分页处理
func (p *PageRequest) Offset() int64 {
	return (int64(p.PageNum) - 1) * int64(p.PageSize)
}
