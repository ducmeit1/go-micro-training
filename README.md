# MICROSERVICES TRAINING, GO, GRPC

## Proto Gen

```shell
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb
```

## Test with grpcui

```
grpcui -plaintext 127.0.0.1:2222
```