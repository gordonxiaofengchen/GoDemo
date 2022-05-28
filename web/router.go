package web

import (
	"demo/web/handler"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	handler.InitMiddleWare(Router)
	handler.InitStaffHandler(Router)
}