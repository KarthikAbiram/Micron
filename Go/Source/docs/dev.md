# Commands used to generate grpc go Template

## Generate Protobuf Messages
```
protoc `
  --proto_path=proto `
  --go_out=gen `
  --go_opt=paths=source_relative `
  proto/micron.proto
```
This generates: gen/micron.pb.go

## Generate grpc Client/Server Interfaces

```
protoc `
  --proto_path=proto `
  --go_out=gen `
  --go_opt=paths=source_relative `
  --go-grpc_out=gen `
  --go-grpc_opt=paths=source_relative `
  proto/micron.proto
```