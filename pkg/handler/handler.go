package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!!",
		})
	})

	return router
}
