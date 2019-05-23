package models

import (
	"github.com/duckbill-io/duckbill/orm"
)

// 文章
type Post struct {
	Content string `json:"content"`
	Meta
}

// 文章类表
type Posts []*Post

// scan 扫描对应的文章信息，并更新Post实例
func (p *Post) scan() (err error) {
	err := orm.Scan(p)
	return
}
