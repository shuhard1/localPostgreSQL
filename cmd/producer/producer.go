package main

import (
	stan "github.com/nats-io/stan.go"
)

func main() {
	jsonDataBytes := []byte(`{
		"order_uid": "f478kekjh9h3d12345678",
		"track_number": "DBILMTESTTRACK",
		"entry": "DBIL",
		"delivery": {
		  "name": "Abzal Baida",
		  "phone": "+9720001234",
		  "zip": "2631234",
		  "city": "Bashkortostansk",
		  "address": "Ploshad Mira 25",
		  "region": "Bashkortostan",
		  "email": "uberslav@gmail.com"
		},
		"payment": {
		  "transaction": "f478kekjh9h3d12341234",
		  "request_id": "",
		  "currency": "RU",
		  "provider": "wbpay",
		  "amount": 2000,
		  "payment_dt": 1637901234,
		  "bank": "beta",
		  "delivery_cost": 2500,
		  "goods_total": 123,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 9931234,
			"track_number": "DBILMTESTTRACK",
			"price": 453,
			"rid": "ab4219087a7641234test",
			"name": "Dascaras",
			"sale": 2000,
			"size": "0",
			"total_price": 400,
			"nm_id": 2381234,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "ru",
		"internal_signature": "",
		"customer_id": "testID",
		"delivery_service": "meest",
		"shardkey": "2",
		"sm_id": 100,
		"date_created": "2023-11-26T06:22:19Z",
		"oof_shard": "2"
	  }`)
	sc, _ := stan.Connect("test-cluster", "2")
	defer sc.Close()
	sc.Publish("foo", jsonDataBytes)
}
