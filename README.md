# localPostgreSQL

## Start instructions

1.
   ``` create table orders (id text NOT NULL UNIQUE, info json NOT NULL)```
   
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
   http://localhost:8080/?id=abcd8kekjh123412341234<br/>
   http://localhost:8080/?id=b563feb7b2b84b6test<br/>
   
### Technologies:
- Golang
- Gin
- PostgreSQL
- Nats streaming
