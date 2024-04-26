export PATH := $(PATH):`go env GOPATH`/bin
export GO111MODULE=on
LDFLAGS := -s -w

VERSION ?= nil

all: env build

env:
	@go version
	@node --version

build: web eva

eva: copy-web
	cd src
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -ldflags "-X 'main.Version=$(VERSION)'" -o bin/eva ./cmd/server
	cd ..

web: npm-install
	npm run build --prefix font

npm-install:
	npm install --prefix font

copy-web:
	rm -rf ./src/serve/static
	cp -rf ./font/dist/* ./src/serve/static

clean:
	rm -rf ./bin
	rm -rf ./font/dist