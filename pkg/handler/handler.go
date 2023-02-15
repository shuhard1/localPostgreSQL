package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shuhard1/localPostgreSQL/internal/order"
	"github.com/shuhard1/localPostgreSQL/pkg/cache"
)

type Handler struct {
	Repositry order.Repository
}

func (h *Handler) InitRouter(c *cache.Cache) *gin.Engine {
	var err error

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		item, found := c.Get(id)
		info := fmt.Sprintf("%v", item)
		ord := order.Order{Info: info, ID: id}
		if !found {
			ord, err = h.Repositry.FindOne(context.Background(), id)
			if err != nil {
				fmt.Fprintf(ctx.Writer, "QueryRow failed: %v\n", err)
			} else {
				c.Set(ord.ID, ord.Info+" from cache", 5*time.Minute)
			}
		}
		ctx.JSON(200, ord.Info)
	})
	router.Run(":8080")

	return router
}
