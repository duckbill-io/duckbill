// 控制器
package controllers

import (
	"github.com/duckbill-io/duckbill/models"
	"github.com/duckbill-io/duckbill/render"
	"io"
	"log"
)

// 显示指定名称的文章
func ShowPost(w io.Writer, name string) {
	post, err := models.FindPost(name)
	if err != nil {
		log.Fatal(err)
	}
	showTemplate := "views/posts/show.html"
	render.Execute(w, showTemplate, post)
}

// 罗列所有的文章
func Posts(w io.Writer) {
	posts, err := models.FindAllPosts()
	if err != nil {
		log.Fatal(err)
	}

	indexTemplate := "views/posts/index.html"
	err = render.Execute(w, indexTemplate, posts)
	if err != nil {
		log.Fatal(err)
	}
}

// 显示制定名称的标签
func ShowTag(w io.Writer, name string) {
	tag, err := models.FindTag(name)
	if err != nil {
		log.Fatal(err)
	}
	showTemplate := "views/tags/show.html"
	render.Execute(w, showTemplate, tag)
}

// 罗列所有的标签
func Tags(w io.Writer) {
	tags, err := models.FindAllTags()
	if err != nil {
		log.Fatal(err)
	}

	indexTemplate := "views/tags/index.html"
	render.Execute(w, indexTemplate, tags)
}
