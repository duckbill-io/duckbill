package routes

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 路由器
type Router struct {
	middlewareChain []middleware
	httprouter.Router
}

// Use 添加中间件
func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

// Add 添加路由
func (r *Router) Add(route string, h interface{}) {
	var mergedHandle = r.handler(h)
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandle = r.middlewareChain[i](mergedHandle)
	}
	// 默认是用GET方法
	r.GET(route, mergedHandle)
}

func (r *Router) handler(vi interface{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch v := vi.(type) {
		case func(w io.Writer):
			v(w)
		case func(w io.Writer, name string):
			v(w, ps.ByName("name"))
		}
	}
}
