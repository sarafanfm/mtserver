module github.com/sarafanfm/mtserver

go 1.20

replace (
	github.com/gin-gonic/gin v1.6.3 => github.com/gin-gonic/gin v1.7.4
	nhooyr.io/websocket v1.8.6 => nhooyr.io/websocket v1.8.7
)

require (
	github.com/joho/godotenv v1.4.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.52.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/rs/cors v1.7.0 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	github.com/improbable-eng/grpc-web v0.15.0
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
