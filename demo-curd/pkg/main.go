package main

import (
	"fmt"
	"golang-code/demo-curd/pkg/controller"
)

func main() {
	manager := controller.NewManager()
	err := manager.StartHttpServer()
	if err == nil {
		fmt.Println(" 启动 HTTP 服务成功...")
	} else {
		fmt.Println(" 启动 HTTP 服务失败...")
	}
}
