package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shuhard1/localPostgreSQL/internal/natsStreaming"
	"github.com/shuhard1/localPostgreSQL/internal/order"
	"github.com/shuhard1/localPostgreSQL/internal/table"
	"github.com/shuhard1/localPostgreSQL/pkg/cache"
)

type Handler struct {
	Repositry order.Repository
}

var index = template.Must(template.New("index").Parse(table.IndexTmpl))

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
				fmt.Fprintf(ctx.Writer, "FindOne failed: %v\n", err)
			} else {
				c.Set(ord.ID, ord.Info, 5*time.Minute)
			}
		}
		var dat natsStreaming.Order
		err := json.Unmarshal([]byte(ord.Info), &dat)
		if err != nil {
			fmt.Fprintf(ctx.Writer, "unmarshal error: %s\n", err)
		}
		if err := index.Execute(ctx.Writer, dat); err != nil {
			fmt.Fprintf(ctx.Writer, "index.Execute error: %s\n", err)
		}
	})
	router.Run(":8080")

	return router
}
