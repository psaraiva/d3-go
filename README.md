# D3

## Deploy
- `docker-compose up -d`

## logs
- `sudo docker logs -f d3-go`
- `sudo docker logs -f d3-mysql`
- `sudo docker logs -f d3-rabbitmq`

## RabbitMQ
-  @see docker-compose.yml
- `docker ps`

## MySql
-  @see docker-compose.yml
- `docker ps`
- `docker exec -it <container_id> mysql -uroot -p`
```Sql
USE orders;
SELECT * FROM orders;
```

## gRPC
- `localhost:40051`
- `grpcurl -plaintext localhost:40051 list`
```sh
evans \
  --path ./internal/infra/grpc/protofiles \
  --proto order.proto \
  --host localhost \
  --port 40051 \
  repl

show service
```

## WebServer API
-  @see plugin REST Client
- `locahost:7000`
- `./api/order.http`

## GraphQL
- `localhost:7080`
```GraphQL
mutation createOrder{
  createOrder(input: {
    id: "aaa-bbb",
    Price: 120.00,
    Tax: 0.5
  }){
    id
    Price
    Tax
  }
}

mutation updateOrder{
  updateOrder(input: {
    id: "aaa-bbb",
    Price: 199.00,
    Tax: 0.75
  }){
    id
    Price
    Tax
  }
}

query ListOrders {
    listOrder {
        id
        Price
        Tax
        FinalPrice
    }
}
```
