# localPostgreSQL

## Start instructions

1.
   ``` git clone https://github.com/shuhard1/localPostgreSQL.git```
2. 
   ``` cd $GOPATH/src/github.com/nats-io/nats-streaming-server```<br/>
   ``` go run nats-streaming-server.go```
3.
   ``` cd localPostgreSQL/cmd/main```<br/>
   ``` go run main.go```
4. 
   ``` cd localPostgreSQL/cmd/producer```<br/>
   ``` go run producer.go```
5. http://localhost:8080/?id=f478kekjh9h3d12345678<br/>
### Technologies:
- Golang
- Gin
- PostgreSQL
- Nats streaming
