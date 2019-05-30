// 包router	实现多路复用
package router

import (
	"github.com/duckbill-io/duckbill/controllers"

	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type Router struct {
	httprouter.Router
}

func New() (r *Router) {
	r = &Router{*httprouter.New()}
	r.InitRoutes()
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.NotFound(w)
	})
	return
}

// Init 初始化路由
func (r *Router) InitRoutes() {
	r.GET("/", handler(controllers.Home))
	r.GET("/about", handler(controllers.About))
	r.GET("/posts", handler(controllers.Posts))
	r.GET("/posts/:name", handler(controllers.Post))
	r.GET("/tags", handler(controllers.Tags))
	r.GET("/tags/:name", handler(controllers.Tag))
}

func handler(vi interface{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch v := vi.(type) {
		case func(w io.Writer):
			v(w)
		case func(w io.Writer, name string):
			v(w, ps.ByName("name"))
		}
	}
}
