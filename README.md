# grpc-tutorial

## Creating invoicer

```sh
protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto
```
