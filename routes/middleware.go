package routes

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 中间件
type middleware func(httprouter.Handle) httprouter.Handle

// logReq 中间件用于记录访问请求
func logReq(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("%s %q", r.Method, r.URL.Path)
		h(w, r, ps)
	}
}
