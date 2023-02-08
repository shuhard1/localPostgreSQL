package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shuhard1/localPostgreSQL/internal/order/db"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
)

func main() {
	conn := postgresql.NewClient()
	defer conn.Close(context.Background())

	repository := db.NewRepository(conn)

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		order, err := repository.FindOne(context.Background(), id)
		if err != nil {
			fmt.Fprintf(ctx.Writer, "QueryRow failed: %v\n", err)
		}
		ctx.JSON(200, order.Info)
	})
	server.Run(":8080")
}
