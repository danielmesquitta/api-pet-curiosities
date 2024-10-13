.PHONY: default install update run clear gen build lint

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

include .env

default: run

install:
	@go mod download && go install github.com/air-verse/air@latest
update:
	@go mod tidy && go get -u ./...
run:
	@air -c .air.toml
clear:
	@find ./tmp -mindepth 1 ! -name '.gitkeep' -delete
generate:
	@go generate ./...
build:
	@GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./tmp/restapi ./cmd/restapi
lint:
	@golangci-lint run && nilaway ./...
new_entity:
	@go run -mod=mod entgo.io/ent/cmd/ent new --target=internal/provider/db/ent/schema $(ARGS)
%::
	@true

