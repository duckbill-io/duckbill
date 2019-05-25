// 博客引擎的model模块
package models

import (
	"github.com/duckbill-io/duckbill/orm"
)

//	=====================
//	=== Post 相关函数 ===
//	=====================

// FindPost 根据name查找对应的文章
func FindPost(name string) (*Post, error) {
	post, err := newPost(name)
	if err != nil {
		return nil, err
	}

	err = post.scan()
	if err != nil {
		return nil, err
	}
	return post, nil
}

// newPost 根据name初始化一个Post实例
func newPost(name string) (*Post, error) {
	meta, err := findMeta(name)
	if err != nil {
		return nil, err
	}
	post, err := meta.post(false)
	if err != nil {
		return nil, err
	}
	return post, nil
}

//  FindAllPosts 查找所有的文章
func FindAllPosts() (Posts, error) {
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

//	=====================
//	=== Meta 相关函数 ===
//	=====================

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

// findAllMetas 找找所有的元数据
func findAllMetas() (Metas, error) {
	metas := Metas{&Meta{}}
	err := metas.scan()
	if err != nil {
		return nil, err
	}
	return metas, nil
}

//	=====================
//	=== Tag 相关函数 ====
//	=====================

// FindTag 根据name查找对应的标签
func FindTag(name string) (*Tag, error) {
	tag := newTag(name)
	err := tag.scan()
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// newTag 根据name初始化一个Tag实例
func newTag(name string) *Tag {
	return &Tag{Name: name}
}

// FindAllTags 查找所有的标签
func FindAllTags() (Tags, error) {
	tags := Tags{&Tag{}}
	err := tags.scan()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

//	=====================
//	===   orm.Scan    ===
//	=====================

func scan(v interface{}) error {
	err := orm.Scan(v)
	return err
}
