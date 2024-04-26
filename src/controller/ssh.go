package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/info24/eva/common"
	"github.com/info24/eva/core"
	"github.com/info24/eva/model"
	"github.com/info24/eva/store"
	"net/http"
	"strconv"
)

func RegisterSsh(c *gin.Context) {
	id := c.Param("id")
	token := c.Query("token")
	tokenString, err := core.GetJwtMiddleware().ParseTokenString(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if !tokenString.Valid {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "jwt token invalid",
		})
		return
	}
	row, _ := strconv.Atoi(c.Query("row"))
	col, _ := strconv.Atoi(c.Query("col"))
	var device model.Device
	tx := store.GetDB().Where("id = ?", id).First(&device)
	if tx.Error != nil {

		c.JSON(200, gin.H{"code": common.DeviceNoFound, "msg": "device not found"})
		return
	}
	core.RegisterWsInstance(c.Request, c.Writer, device.ToMap(row, col))

}
