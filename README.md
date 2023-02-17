# localPostgreSQL

## Start instructions

1.
   ``` create table orders (id text, info json)```<br/>
2.
   ``` git clone https://github.com/shuhard1/localPostgreSQL.git```<br/>
3. 
   ``` cd $GOPATH/src/github.com/nats-io/nats-streaming-server```<br/>
   ``` go run nats-streaming-server.go```<br/>
4.
   ``` cd localPostgreSQL/cmd/main```<br/>
   ``` go run main.go```<br/>
5. 
   ``` cd localPostgreSQL/cmd/producer```<br/>
   ``` go run producer.go```<br/>
6. http://localhost:8080/?id=f478kekjh9h3d12345678<br/><br/>
### Technologies:
- Golang
- Gin
- PostgreSQL
- Nats streaming
