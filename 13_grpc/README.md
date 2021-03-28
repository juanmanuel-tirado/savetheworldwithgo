# gRPC

These folder contains the examples for the gRPC chapter. In order to generate
the gRPC stubs you need to have the corresponding plugin installed in your
environment and made it available to be executed by the `protoc` tool. Follow
the official documentation [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
to install all the elements you need.

The gRPC stubs are already generated in the folder. If you want to compile
them you can run:

```
protoc -I=. --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src *.proto
```

Observe that depending on the output folders `go_out` and `go-grpc_out` you may
have to change your import statements.

## A minor observation

Some examples have a client and a server. In order to make them run remember
that the server port must be free, and you have to run first the server, 
then the client.
