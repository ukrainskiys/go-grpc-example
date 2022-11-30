generate proto
```
protoc --go-grpc_out=. proto/*.proto
```

run server
```
go run server.go
```

run client
```
go run client.go
```