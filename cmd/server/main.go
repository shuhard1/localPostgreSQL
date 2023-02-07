package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shuhard1/localPostgreSQL/internal/customer/db"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
)

func main() {
	conn := postgresql.NewClient()
	defer conn.Close(context.Background())

	repository := db.NewRepository(conn)

	all, err := repository.FindAll(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	for _, ath := range all {
		fmt.Println(ath)
	}
}
