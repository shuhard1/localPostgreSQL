package main

import (
	stan "github.com/nats-io/stan.go"
)

func main() {
	jsonDataBytes1 := []byte(`{
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

	jsonDataBytes2 := []byte(`{
		"order_uid": "abcd8kekjh123412341234",
		"track_number": "FJILMTESTTRACK",
		"entry": "FGIL",
		"delivery": {
		  "name": "Murza D",
		  "phone": "+9729999999",
		  "zip": "7891789",
		  "city": "Moscow",
		  "address": "Ploshad Dira 90",
		  "region": "Moscow",
		  "email": "uberslavmensch@gmail.com"
		},
		"payment": {
		  "transaction": "g000kekjh9h3d12367843",
		  "request_id": "",
		  "currency": "USA",
		  "provider": "Lpay",
		  "amount": 1300,
		  "payment_dt": 1637909999,
		  "bank": "omega",
		  "delivery_cost": 2599,
		  "goods_total": 999,
		  "custom_fee": 5
		},
		"items": [
		  {
			"chrt_id": 9939999,
			"track_number": "FGILMTESTTRACK",
			"price": 999,
			"rid": "ab4219087a76499999999",
			"name": "Bobris",
			"sale": 3000,
			"size": "25",
			"total_price": 400,
			"nm_id": 2381234,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "usa",
		"internal_signature": "",
		"customer_id": "999test",
		"delivery_service": "pup",
		"shardkey": "99",
		"sm_id": 399,
		"date_created": "2030-11-26T06:22:19Z",
		"oof_shard": "35"
	  }`)

	jsonDataBytes3 := []byte(`{"order_uid": "b563feb7b2b84b6test",
	  "track_number": "WBILMTESTTRACK",
	  "entry": "WBIL",
	  "delivery": {
		"name": "Test Testov",
		"phone": "+9720000000",
		"zip": "2639809",
		"city": "Kiryat Mozkin",
		"address": "Ploshad Mira 15",
		"region": "Kraiot",
		"email": "test@gmail.com"
	  },
	  "payment": {
		"transaction": "b563feb7b2b84b6test",
		"request_id": "",
		"currency": "USD",
		"provider": "wbpay",
		"amount": 1817,
		"payment_dt": 1637907727,
		"bank": "alpha",
		"delivery_cost": 1500,
		"goods_total": 317,
		"custom_fee": 0
	  },
	  "items": [
		{
		  "chrt_id": 9934930,
		  "track_number": "WBILMTESTTRACK",
		  "price": 453,
		  "rid": "ab4219087a764ae0btest",
		  "name": "Mascaras",
		  "sale": 30,
		  "size": "0",
		  "total_price": 317,
		  "nm_id": 2389212,
		  "brand": "Vivienne Sabo",
		  "status": 202
		}
	  ],
	  "locale": "en",
	  "internal_signature": "",
	  "customer_id": "test",
	  "delivery_service": "meest",
	  "shardkey": "9",
	  "sm_id": 99,
	  "date_created": "2021-11-26T06:22:19Z",
	  "oof_shard": "1"}`)

	sc, _ := stan.Connect("test-cluster", "2")
	defer sc.Close()
	sc.Publish("foo", jsonDataBytes1)
	sc.Publish("foo", jsonDataBytes2)
	sc.Publish("foo", jsonDataBytes3)
}
