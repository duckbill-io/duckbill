package models

// 标签
type Tag struct {
	Name      string   // 标签名称
	Count     int      // 打有该标签的文章数量
	Articles  Posts    // 打有该标签的文章列表
	PostNames []string // 打有该标签的文章名列表
}

// 标签列表
type Tags []*Tag

// scan 扫描并更新标签
func (t *Tag) scan() error {
	err := scan(t)
	if err != nil {
		return err
	}
	//　查找标签的所有实例
	arts, err := t.articles()
	if err != nil {
		return err
	}
	*t.Articles = arts
	return nil
}

// scan 扫描并更新标签列表
func (ts *Tags) scan() error {
	err := scan(ts)
	return err
}

// posts 查找打有该标签的文章
func (t *Tag) articles() (Posts, error) {
	postnames := *t.PostNames
	posts := make(Posts, len(postnames))
	for i := range postnames {
		*posts[i].Name = postnames[i]
		err := posts[i].scan()
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}
