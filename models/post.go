package models

// 文章
type Post struct {
	Content string `json:"content"`
	Meta
}

// 文章类表
type Posts []*Post

// scan 扫描对应的文章信息，并更新Post实例
func (p *Post) scan() (err error) {
	err = scan(p)
	return
}

// sortByCreatedAt 将文章列表按照创建时间到需排序
func (ps *Posts) sortByCreatedAt() {
	byCreatedAt(*ps).sort()
}
