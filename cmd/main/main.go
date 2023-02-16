package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	stan "github.com/nats-io/stan.go"
	"github.com/shuhard1/localPostgreSQL/internal/order/db"

	"github.com/shuhard1/localPostgreSQL/internal/natsStreaming"
	"github.com/shuhard1/localPostgreSQL/pkg/cache"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
	"github.com/shuhard1/localPostgreSQL/pkg/handler"
)

func main() {
	cache := cache.New(10*time.Minute, 20*time.Minute)
	sc, _ := stan.Connect("test-cluster", "1")
	defer sc.Close()

	conn := postgresql.NewClient()
	defer conn.Close(context.Background())

	var h handler.Handler
	h.Repositry = db.NewRepository(conn)

	orders, err := h.Repositry.FindAll(context.Background())
	if err != nil {
		fmt.Printf("FindAll failed: %v\n", err)
	}

	for _, order := range orders {
		cache.Set(order.ID, order.Info, 10*time.Minute)
	}

	sc.Subscribe("foo", func(m *stan.Msg) {
		var dat natsStreaming.Order
		err := json.Unmarshal(m.Data, &dat)
		if err != nil {
			fmt.Printf("unmarshal error: %s\n", err)
		}
		info, _ := json.Marshal(dat)
		h.Repositry.Create(context.Background(), dat.Order_uid, string(info))
	})

	h.InitRouter(cache)
}
