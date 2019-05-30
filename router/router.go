// 包router实现网络请求的多路复用
package router

import (
	"github.com/duckbill-io/duckbill/controllers"
	"net/http"
	"path"
	"strings"
)

type Router struct{}

func New() *Router {
	return &Router{}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlpath := r.URL.Path
	// home page
	if urlpath == "/" {
		controllers.Home(w)
	}
	// about page
	if urlpath == "/about" || urlpath == "/about/" {
		controllers.About(w)
	}
	// posts index
	if urlpath == "/posts" || urlpath == "/posts/" {
		controllers.Posts(w)
	}
	// show post
	if strings.HasPrefix(urlpath, "/posts/") && urlpath != "/posts/" {
		_, postname := path.Split(urlpath)
		controllers.Post(w, postname)
	}
	// tags index
	if urlpath == "/tags" || urlpath == "/tags/" {
		controllers.Tags(w)
	}
	// show tag
	if strings.HasPrefix(urlpath, "/tags/") && urlpath != "/tags/" {
		_, tagname := path.Split(urlpath)
		controllers.Tag(w, tagname)
	}
}
