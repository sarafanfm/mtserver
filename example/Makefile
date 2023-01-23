proto:
	./protoc.sh

server:
# auto preload ENV_FILE when pkg init
# ENV_FILE is a comma separated list of files to load
# @see https://github.com/joho/godotenv for more details
	ENV_FILE=example.env go run cmd/server/main.go

client:
	ENV_FILE=example.env go run cmd/client/main.go

.PHONY: proto server client