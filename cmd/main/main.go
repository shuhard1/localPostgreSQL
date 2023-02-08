package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shuhard1/localPostgreSQL/internal/order/db"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
)

func main() {
	conn := postgresql.NewClient()
	defer conn.Close(context.Background())

	repository := db.NewRepository(conn)

	all, err := repository.FindOne(context.Background(), "b563feb7b2b84b6test")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	server := gin.Default()
	server.GET("/id", func(ctx *gin.Context) {
		ctx.JSON(200, all.Info)
	})
	server.Run(":8080")
}
