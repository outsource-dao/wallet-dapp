# wallet-service
## Generate protobuf
```
protoc --proto_path=./proto --go_out=plugins=grpc:./proto ./proto/dApp.proto
```

# Docker Build Instruction

```
docker build . -t dapp-backend:0.1
docker run -p 4551:4551 --env-file=./env dapp-backend:0.1
```65