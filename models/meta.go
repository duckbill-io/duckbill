package models

import (
	"github.com/duckbill-io/duckbill/orm"
)

// 文章元数据
type Meta struct {
	Name       string   // 标题
	Intro      string   // 简介
	created_at string   // 创建时间
	updated_at string   // 最近更新时间
	tags       []string // 标签列表
}

// 文章元数据列表
type Metas []*Meta

// scan 扫描对应的元信息数据，并更新Meta实例
func (m *Meta) scan() (err error) {
	err = orm.Scan(m)
	return
}

// post 查找文章元数据对应的文章
func (m *Meta) post(hascontent bool) (*Post, error) {
	post := &Post{Meta: *m}
	if hascontent {
		err := post.scan()
		if err != nil {
			return nil, err
		}
	}
	return post, nil
}

// scan 扫描ms对应的元信息数据，并更新ms
func (ms Metas) scan() (err error) {
	err = orm.Scan(ms)
	return
}

// posts 扫描元数据列表对应的文章类表
func (ms Metas) posts(hascontent bool) (posts Posts, err error) {
	posts = make(Posts, len(ms))
	for i, _ := range ms {
		posts[i], err = ms[i].post(hascontent)
		if err != nil {
			return
		}
	}
	return
}
