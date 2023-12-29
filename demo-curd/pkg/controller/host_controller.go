package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-project/demo-curd/pkg/bean/model"
	"golang-project/demo-curd/pkg/service"
	"golang-project/demo-curd/pkg/service/impl"
	"strconv"
)

var hostController = &HostController{
	hostService: impl.NewHostServiceImpl(),
}

type HostController struct {
	hostService service.HostService
}

func (hostController *HostController) createHost(c *gin.Context) {
	host := model.NewHost()
	if err := c.Bind(host); err != nil {
		_, _ = c.Writer.Write([]byte("解析 Host 参数失败"))
		return
	}

	if _, err := hostController.hostService.InsertHost(c.Request.Context(), host); err != nil {
		_, _ = c.Writer.Write([]byte(fmt.Sprintf("创建 Host 失败, %s", err)))
		return
	}

	_, _ = c.Writer.Write([]byte("创建 Host 成功"))
}

func (hostController *HostController) queryHost(c *gin.Context) {
	hostId, b := c.Params.Get("id")
	if !b {
		_, _ = c.Writer.Write([]byte("不存在 ID 参数"))
	}
	hostList, err := hostController.hostService.SelectOneHost(c.Request.Context(), func() int {
		id, err := strconv.Atoi(hostId)
		if err != nil {
			return 0
		}
		return id
	}())
	if err != nil {
		_, _ = c.Writer.Write([]byte(fmt.Sprintf("查询 Host Id = %s 不存在", hostId)))
		return
	}

	if hostListJson, err := json.Marshal(hostList); err == nil {
		_, _ = c.Writer.Write(hostListJson)
	}
}

func (hostController *HostController) Registry(router gin.IRouter) {
	router.POST("/hosts", hostController.createHost)
	router.GET("/host/:id", hostController.queryHost)
}
