# localPostgreSQL

## Start instructions

1.
   ``` create table orders (id text, info json)```
2.
   ``` git clone https://github.com/shuhard1/localPostgreSQL.git```
3. 
   ``` cd $GOPATH/src/github.com/nats-io/nats-streaming-server```<br/>
   ``` go run nats-streaming-server.go```
4.
   ``` cd localPostgreSQL/cmd/main```<br/>
   ``` go run main.go```
5. 
   ``` cd localPostgreSQL/cmd/producer```<br/>
   ``` go run producer.go```
6. http://localhost:8080/?id=f478kekjh9h3d12345678<br/>
### Technologies:
- Golang
- Gin
- PostgreSQL
- Nats streaming
