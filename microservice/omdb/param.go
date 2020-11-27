package omdb

type Param struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
}

func (p Param) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"s":    p.Keyword,
		"page": p.Page,
	}
}

func (p Param) Valid() bool {
	return len(p.Keyword) > 0 && p.Page > 0
}
