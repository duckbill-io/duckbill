package models

// 标签
type Tag struct {
	Name      string   `json:"name"`  // 标签名称
	Count     string   `json:"count"` // 打有该标签的文章数量
	Articles  Posts    // 打有该标签的文章列表
	PostNames []string `json:"articles"` // 打有该标签的文章名列表
}

// 标签列表
type Tags []*Tag

// scan 扫描并更新标签
func (t *Tag) scan() error {
	err := scan(t)
	if err != nil {
		return err
	}
	//　查找标签的所有文章
	t.Articles = make(Posts, len(t.PostNames))
	for i := range t.PostNames {
		t.Articles[i], err = newPost(t.PostNames[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// scan 扫描并更新标签列表
func (ts *Tags) scan() error {
	err := scan(ts)
	return err
}

