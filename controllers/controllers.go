// 控制器
package controllers

import (
	"github.com/duckbill-io/duckbill/models"
	"github.com/duckbill-io/duckbill/render"
	"io"
	"log"
)

/*
TODO 根据不同的错误选择相应的错误处理方式
例如:
	1.访问的文章不存在时应该返回什么样的页面(一般的会重定向至404页面，提示用户访问的文章不存在)
	2.服务器内部有错误应该返回什么页面(一般会重定向至500错误页面)
但这些错误处理应该怎么组织才合适?重定向的话似乎要重新组织controller模块与router模块
*/

// 显示指定名称的文章
func Post(w io.Writer, name string) {
	showTemplate := "views/posts/show.html"
	post, err := models.FindPost(name)
	if err != nil {
		logerror(err)
		showTemplate = "views/500.html"
	}
	err = render.Execute(w, showTemplate, post)
	logerror(err)
}

// 罗列所有的文章
func Posts(w io.Writer) {
	posts, err := models.FindAllPosts()
	logerror(err)

	indexTemplate := "views/posts/index.html"
	err = render.Execute(w, indexTemplate, posts)
	logerror(err)
}

// 显示制定名称的标签
func Tag(w io.Writer, name string) {
	tag, err := models.FindTag(name)
	logerror(err)

	showTemplate := "views/tags/show.html"
	err = render.Execute(w, showTemplate, tag)
	logerror(err)
}

// 罗列所有的标签
func Tags(w io.Writer) {
	tags, err := models.FindAllTags()
	logerror(err)

	indexTemplate := "views/tags/index.html"
	err = render.Execute(w, indexTemplate, tags)
	logerror(err)
}

// 主页
func Home(w io.Writer) {
	homeTemplate := "views/home.html"
	err := render.Execute(w, homeTemplate, nil)
	logerror(err)
}

// 关于我
func About(w io.Writer) {
	aboutTemplate := "views/about.html"
	err := render.Execute(w, aboutTemplate, nil)
	logerror(err)
}

// 404 Not Found
func NotFound(w io.Writer) {
	notfoundTemplate := "views/404.html"
	err := render.Execute(w, notfoundTemplate, nil)
	logerror(err)
}

func logerror(err error) {
	if err != nil {
		log.Printf("err: %s", err)
	}
}
