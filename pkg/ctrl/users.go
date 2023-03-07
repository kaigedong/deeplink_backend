package ctrl

import (
	"deeplink_backend/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct{}

var Engines = &Controller{}

func (ec *Controller) NotFound(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "NotFound",
		})
}

func (ec *Controller) OK(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		})
}

// 增加被控设备
func (ec *Controller) NewDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}

// 删除被控设备
func (ec *Controller) DeleteDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}

// 修改被控设备?

// 查找被控设备
func (ec *Controller) Devides(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		types.Device{
			ID: uuid.New(),
			IP: "",
		})
}
