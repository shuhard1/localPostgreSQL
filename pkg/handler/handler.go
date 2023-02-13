package handler

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shuhard1/localPostgreSQL/internal/order"
)

type Handler struct {
	Repositry order.Repository
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		order, err := h.Repositry.FindOne(context.Background(), id)
		if err != nil {
			fmt.Fprintf(ctx.Writer, "QueryRow failed: %v\n", err)
		}
		ctx.JSON(200, order.Info)
	})
	router.Run(":8080")

	return router
}
