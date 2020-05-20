# go-grpc

## Description

go-grpc is a Go example that provides how to use GRpc in golang with some configuration.
It has the following configuration:

* using normal configuration.
* using certificate to connect between GRpc server and client.
* using middleware configuration.


## Requirements

Go 1.12 or above.


### Installation

Run the following command to install the package:
[grpc-protoc](https://grpc.io/docs/quickstart/go/)

to comiple file `*.proto`, run this command:
```
protoc -I pkg/utl/grpc pkg/utl/grpc/helloworld.proto --go_out=plugins=grpc:pkg/utl/grpc 
```


## Getting Started

to running with certificate (server)
```
go run cmd/server_cert/server.go
```
to running with certificate (client)
```
go run cmd/client_cert/client.go
```
