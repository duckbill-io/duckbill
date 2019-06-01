// 包routes 用于设置路由
package routes

import (
	"net/http"

	"github.com/duckbill-io/duckbill/controllers"
)

// DefaultRouter 创建默认路由器
func DefaultRouter() *Router {
	r := Router{}
	r.addDefaultRoutes()
	return &r
}

// addDefaultRoutes 添加默认路由
func (r *Router) addDefaultRoutes() {
	// 添加中间件
	r.Use(logReq)
	// 添加路由
	r.Add("/", controllers.Home)
	r.Add("/about", controllers.About)
	r.Add("/posts", controllers.Posts)
	r.Add("/posts/:name", controllers.Post)
	r.Add("/tags", controllers.Tags)
	r.Add("/tags/:name", controllers.Tag)
	// 添加NotFound路由
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.NotFound(w)
	})
}
