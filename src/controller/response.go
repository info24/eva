package controller

import "github.com/gin-gonic/gin"

func response(code int, msg string) gin.H {
	return gin.H{"code": code, "message": msg}
}
