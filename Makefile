.PHONY: build

default: build

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a \
		${GO_LDFLAGS} \
		-tags netgo \
		-installsuffix netgo \
		-o app \
		.
