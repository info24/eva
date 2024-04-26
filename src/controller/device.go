package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/info24/eva/model"
	"github.com/info24/eva/store"
	"strconv"
)

func GetAllDevice(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("current", "1"))
	name := c.Query("name")
	ip := c.Query("ip")

	var devices []model.Device
	var total int64
	tx := store.GetDB().Model(devices)
	if name != "" {
		tx = tx.Where("name LIKE ?", "%"+name+"%")
	}
	if ip != "" {
		tx = tx.Where("ip = ?", ip)
	}
	tx.Count(&total).Limit(size).Offset(size * (page - 1)).Find(&devices)
	c.JSON(200, gin.H{
		"total": total,
		"data":  devices,
		"msg":   "success",
	})
}

func AddDevice(c *gin.Context) {
	var device model.Device
	err := c.ShouldBind(&device)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	tx := store.GetDB().Create(&device)
	if tx.Error != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  tx.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func UpdateDevice(c *gin.Context) {
	var device model.Device
	id := c.Param("id")
	err := c.ShouldBind(&device)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	tx := store.GetDB().Where("id = ?", id).Updates(&device)
	if tx.Error != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  tx.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func DeleteDevice(c *gin.Context) {
	id := c.Param("id")
	var device model.Device
	tx := store.GetDB().Where("id = ?", id).Delete(&device)
	if tx.Error != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  tx.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "delete success",
	})
}
