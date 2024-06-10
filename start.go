package ikun

import (
	"net/http"
)

// Server 接口定义了服务器的基本行为
type Server interface {
	ReadConfig(path string) (*ServerConfig, error) //读取配置文件自动解析

	StartServer(port ...string) error //服务访问

	GET(path string, handlerFunc http.HandlerFunc) // GET 请求方式
}

// srv 结构体定义了服务器的基本属性，未导出
type srv struct {
	handler http.Handler
	router  Router
}

// NewServer 创建一个新的服务器实例并返回 Server 接口类型
func NewServer() Server {
	r := NewRouter()
	return &srv{router: r, handler: r.(*muxRouter).mux}
}

// StartServer 启动服务器，实现 Server 接口
func (s *srv) StartServer(port ...string) error {
	return s.start(port...)
}

// start 启动服务器，监听指定端口
func (s *srv) start(port ...string) error {
	// 获取配置文件中的端口
	configPort := Config.Server.Port

	// 如果配置文件中端口为空，则使用传入的参数端口，如果参数端口也为空，则使用默认常量端口
	p := DEFAULT_PORT
	if configPort != "" {
		p = configPort
	} else if len(port) > 0 {
		p = port[0]
	}

	if s.handler == nil {
		s.handler = http.NewServeMux()
	}

	// 启动服务器
	err := http.ListenAndServe(":"+p, s.handler)
	if err != nil {
		return err
	}

	return nil
}
