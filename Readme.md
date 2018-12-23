# grpc-todo

A simple todo app using gRPC on both the server and client side

Code for the article [A TODO app using grpc-web and Vue.js](https://medium.com/@aravindhanjay/a-todo-app-using-grpc-web-and-vue-js-4e0c18461a3e)

## Build process

To start the gRPC server

```go run server.go```

To start the Envoy proxy 

```sudo -E docker build -t envoy:v1 .```

```sudo docker run  -p 8080:8080 --net=host  envoy:v1```

To start the client side frontend app

``` cd todo-client/ && yarn serve```