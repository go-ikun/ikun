package ikun

import (
	"net/http"
)

// Router 接口定义了路由的方法
type Router interface {
	Handle(path string, handler http.Handler)
	HandleFunc(path string, handlerFunc http.HandlerFunc)
	GET(path string, handlerFunc http.HandlerFunc)
}

// muxRouter 是 Router 接口的具体实现，使用 http.ServeMux
type muxRouter struct {
	mux *http.ServeMux
}

// NewRouter 创建一个新的 muxRouter 实例
func NewRouter() Router {
	return &muxRouter{mux: http.NewServeMux()}
}

// Handle 注册一个新的路由及其处理器
func (r *muxRouter) Handle(path string, handler http.Handler) {
	r.mux.Handle(path, handler)
}

// HandleFunc 注册一个新的路由及其处理函数
func (r *muxRouter) HandleFunc(path string, handlerFunc http.HandlerFunc) {
	r.mux.HandleFunc(path, handlerFunc)
}

// GET 注册一个新的 GET 路由及其处理函数
func (r *muxRouter) GET(path string, handlerFunc http.HandlerFunc) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlerFunc(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

// GET 添加一个新的 GET 路由到路由器
func (s *srv) GET(path string, handlerFunc http.HandlerFunc) {
	s.router.GET(path, handlerFunc)
}
