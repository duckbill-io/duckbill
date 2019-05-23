// 包router实现网络请求的多路复用
package router

import (
	"github.com/duckbill-io/duckbill/controllers"
	"fmt"
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
		fmt.Fprint(w, "hello duckbill")
	}
	// posts index
	if urlpath == "/posts" || urlpath == "/posts/" {
		controllers.Posts(w)
	}
	// show post
	if strings.HasPrefix(urlpath, "/posts/") && urlpath != "/posts/" {
		_, postname := path.Split(urlpath)
		controllers.ShowPost(w, postname)
	}
	// tags index
	if urlpath == "/tags" || urlpath == "/tags/" {
		fmt.Fprint(w, "tags index")
	}
}
