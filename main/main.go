package main

import (
	"encoding/json"
	"fmt"
	"ikun"
	"net/http"
)

// 定义框架实例
var (
	server = ikun.NewServer()
)

func init() {
	InitRouter(server)
}

// main 1.程序入口
func main() {
	data, err := server.ReadConfig("./main/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data.Server.Port)
	//启动服务器并指定端口
	err = server.StartServer(":9001")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// InitRouter 2.初始化路由
func InitRouter(router ikun.Server) {
	router.GET("/ping", Controller)
}

// Controller 3.控制器
func Controller(writer http.ResponseWriter, request *http.Request) {
	// 设置响应头为 JSON
	writer.Header().Set("Content-Type", "application/json")

	// 定义要返回的数据
	response := map[string]string{
		"message": "pong",
	}

	// 将数据编码为 JSON 并写入响应
	json.NewEncoder(writer).Encode(response)
}
