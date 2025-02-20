# Go-gRPC-Node-example

A simple example of a Go gRPC server and a Node.js gRPC client.
Will do four types of RPCs:
- Unary
- Client streaming
- Server streaming
- Bidirectional streaming

## Prerequisites

- Go 1.24 or higher (for correct [tool](https://go.dev/doc/modules/gomod-ref#tool)-directive support in go.mod)
- Node.js with npm

## How to run

1. Clone the repository
    ```bash
    git clone git@github.com:oxyii/go-grpc-node-example.git
    ``` 
2. Run the Go server
    ```bash
    go run server.go
    ```
3. Run the Node.js client
    ```bash
    npm install
    npm start
    ```
4. Observe the output in the terminal

## Compiling the proto file

This repository contains a precompiled files in the `proto` directory.
If you want to recompile it, you need to have the `protoc` compiler installed.
You can delete all files and directories in the `proto` directory except for the `test.proto` file, but it is not necessary.
Then you can recompile the proto file by running the following command:
   ```bash
   make
   ```
_see [Makefile](Makefile) for more details_

If you get an error like `google/protobuf/wrappers.proto: File not found`, you can download the missing files from the [protobuf repository](https://github.com/protocolbuffers/protobuf/tree/main/src/google/protobuf) and put them in the `proto` directory.
Contents of the `proto` directory (before compiling) should look like this:
```
proto
├── test.proto
├── google
│   └── protobuf
│       └── wrappers.proto
