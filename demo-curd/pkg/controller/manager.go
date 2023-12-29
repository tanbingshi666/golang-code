package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-project/demo-curd/pkg/conf"
	"net/http"
	"strconv"
	"time"
)

type Manager struct {
	server *http.Server
	router gin.IRouter
}

func NewManager() *Manager {
	ginEngine := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              conf.GetConfig().Http.Host + ":" + strconv.Itoa(conf.GetConfig().Http.Port),
		Handler:           ginEngine,
	}

	return &Manager{
		server: server,
		router: ginEngine,
	}
}

func (manager *Manager) StartHttpServer() error {
	hostController.Registry(manager.router)
	if err := manager.server.ListenAndServe(); err != nil {
		return fmt.Errorf("启动 HTTP 服务失败, %s", err.Error())
	}
	return nil
}
