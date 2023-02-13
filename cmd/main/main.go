package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	stan "github.com/nats-io/stan.go"
	"github.com/shuhard1/localPostgreSQL/internal/order/db"

	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
	"github.com/shuhard1/localPostgreSQL/pkg/handler"
	"github.com/shuhard1/localPostgreSQL/pkg/natsStreaming"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "1")
	defer sc.Close()

	conn := postgresql.NewClient()
	defer conn.Close(context.Background())

	var h handler.Handler
	h.Repositry = db.NewRepository(conn)

	sc.Subscribe("foo", func(m *stan.Msg) {
		var dat natsStreaming.Order
		err := json.Unmarshal(m.Data, &dat)
		if err != nil {
			log.Fatalf("unmarshal error: %s\n", err.Error())
		}
		info, _ := json.Marshal(dat)
		fmt.Println(dat.Order_uid)
		fmt.Println(string(info))
		h.Repositry.Create(context.Background(), dat.Order_uid, string(info))
	})

	h.InitRouter()
}
