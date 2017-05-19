CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

build:
	@go build -a -tags 'netgo' -ldflags '-w -linkmode external -extldflags -static' .