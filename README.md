# GRPC

## 1. Install dependency
### Protoc: to define and gen a sample of code API
- go get -u github.com/golang/protobuf/protoc-gen-go
### GRPC
- go get -u google.golang.org/grpc

## 2. Sample structure of proto file 
```
  syntax = "proto3";

  package <package_name>;

  option go_package = <dir to go package>;

  message RequestMessage {
    <type> name = 1;
  }

  message ResponseMessage {
    <type> name = 1;
  }

  service GetMessage {
    rpc GetName(RequestMessage) returns (ResponseMessage);
  }
```

## 3. Instruction:
### In a server side:
- Declare new server struct
  ```
  type server struct {
    pb.SomeNameSever
  }
  ```
- We want to listening at some address, to do this, follow this pattern:
  ```
  lis, err := net.Listen("tcp", <address>)
  ```
- Create a new server:
  ```
  s := grpc.NewServer()
  ```
- Register new server:
  ```
  pb.RegisterSomeServer(s, &server{})
  ```
### In a client side:
- Dial (connect) to address
  ```
  c, err := grpc.Dial("address", <secure>)
  ```
- Init new server:
  ```
  client := c.NewSomeServer()

## 4. 4 Type of gRPC APIS
- Unary: client send one request and server response one request.
  ```
  rpc Service(Request) returns (Response);
  ```
- Server streaming: client send one request and server sent multiple responses.
  ```
  rpc Service(Request) returns (stream Response);
  ```
- Client streaming: client sent multiple requests and server sent one response.
  ```
  rpc Service(stream Request) returns (Response);
  ```
- Bi-direction: client sent multiple requests and server sent multiple responses.
  ```
  rpc Service(stream Request) returns (stream Response);
  ```
