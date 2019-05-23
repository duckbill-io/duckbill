// 博客引擎的model模块
package models

// FindPost 根据name查找对应的文章
func FindPost(name string) (*Post, error) {
	post, err := newPost(name)
	if err != nil {
		return nil, err
	}

	err := post.scan()
	if err != nil {
		return nil, err
	}
	return post, nil
}

//  FindAllPosts 查找所有的文章
func FindAllPosts() ([]*Post, error) {
	metas, err := findAllMetas()
	if err != nil {
		return nil, err
	}

	posts, err := metas.posts(false)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// findAllMetas 找找所有的元数据
func findAllMetas() ([]*Meta, error) {
	metas := make([]*Meta)
	metas
}

// newPost 根据name初始化一个Post实例
func newPost(name string) (*Post, error) {
	meta := *findMeta(name)
	post, err := meta.post(false)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// findMeta 根据name查找对应的文章元数据
func findMeta(name string) (*Meta, error) {
	meta := newMeta(name)
	err := meta.scan()
	if err != nil {
		return nil, err
	}

	return meta, err
}

// newMeat 根据name初始化一个Meta实例
func newMeta(name string) *Meta {
	return &Meta{Name: name}
}
