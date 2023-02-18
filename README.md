# gRPC with go

## Requirements

- golang v1.19
- protobuf v3.21
- protoc-gen-go v1.28
- protoc-gen-go-grpc v1.2.0


Create pb files
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```
