# gRPC with go

## Requirements

- golang v1.19
- protobuf v3.21
- protoc-gen-go v1.28 ([link](https://grpc.io/docs/languages/go/quickstart/))
- protoc-gen-go-grpc v1.2.0 ([link](https://grpc.io/docs/languages/go/quickstart/))
- evans ([link](https://github.com/ktr0731/evans))

Create pb files
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

Run the server on port 50051
```
make server
```

Run the client
```
make client
```
