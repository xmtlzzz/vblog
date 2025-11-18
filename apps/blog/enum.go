package blog

import (
	"errors"
	"strings"
)

// 枚举实现
type STAGE int

func (s STAGE) String() string {
	return StageMap[s]
}

// 这里传入指针是因为要修改实际的内容，同时因为STAGE_DIFT都是常量，常量无法取内存地址
// 通过Trim去除前段传过来文章状态中的2个"
// 如果MarshalJSON方法不处理那么这里就不需要这样写，这样写是为了保证前端不需要动内容直接对接后端即可
func (s *STAGE) UnmarshalJSON(data []byte) error {
	res := strings.Trim(string(data), `"`)
	switch res {
	case "草稿":
		*s = STAGE_DRAFT
		return nil
	case "已发布":
		*s = STAGE_PUBLISHED
		return nil
	default:
		return errors.New("请传入正确的状态码")
	}
	return nil
}

// 实现Marshaller接口自定义json序列化，转换为例如"草稿的内容返回给前端"
func (s STAGE) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// 后端处理发布逻辑让前段直接调用即可
var StageMap = map[STAGE]string{
	STAGE_DRAFT:     "草稿",
	STAGE_PUBLISHED: "已发布",
}

const (
	// 未发布
	STAGE_DRAFT STAGE = iota
	// 以发布
	STAGE_PUBLISHED
)
