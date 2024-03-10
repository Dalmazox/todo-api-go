export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

PKG_NAME=./dist/server
PKG_PATH=./cmd/main.go

all: build
PHONY: all

run: build
	$(PKG_NAME)
PHONY: run

build: clean
	go build -o $(PKG_NAME) $(PKG_PATH)
PHONY: build

clean:
	rm -rf $(PKG_NAME)
PHONY: clean

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseDependency  --parseInternal -d cmd,internal/controllers -o api --ot go,json
PHONY: swagger