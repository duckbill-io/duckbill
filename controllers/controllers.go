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
