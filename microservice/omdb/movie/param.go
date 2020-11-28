package movie

import (
	"omdb/proto"
	"strconv"
)

type Param struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
}

func (p Param) ToMap() map[string]string {
	page := strconv.Itoa(p.Page)
	return map[string]string{
		"s":    p.Keyword,
		"page": page,
	}
}

func (p Param) Valid() bool {
	return len(p.Keyword) > 0 && p.Page > 0
}

func (p *Param) FromProto(param *proto.Param) {
	p.Keyword = param.Keyword
	p.Page = int(param.Page)
}
