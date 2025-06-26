package router

import (
	"net/http"
	"timenote/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/activity_logs", handler.CreateActivityLogHandler)

	return r
}
