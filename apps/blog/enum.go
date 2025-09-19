package blog

// 枚举实现
type STAGE int

const (
	// 未发布
	STAGE_DRAFT STAGE = iota
	// 以发布
	STAGE_PUBLISHED
)
