# MultiTransport Server example

1. Compile proto (already compiled in repo) - `make proto`
2. Start server from first TTY - `make server`
3. Start client from another TTY - `make client`

also you can open browser with addresses below while server is running

- `http://localhost:8083/hello/v2/notify` - [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) server-side stream request
- `http://localhost:8083/hello/v2/error` - [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) unary request with Forbidden response
- `http://localhost:8083/hello/v1/Alex` - [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) unary request
- `http://localhost:8083/hello/v2/Alex` - [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) unary request with notify to server-side stream
- `http://localhost:8083/static/` - static server (cannot handle `/` because of gRPC Gateway handle it)
- `http://localhost:8084/` - the same static as in previous but in another endpoint and can handle `/`

see `Makefile` for understading how it works
