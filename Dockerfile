FROM golang
WORKDIR /go/src/github.com/olivere/go-container-debugging
EXPOSE 8080 2345
RUN go get github.com/derekparker/delve/cmd/dlv
ADD . /go/src/github.com/olivere/go-container-debugging

CMD ["dlv", "debug", "github.com/olivere/go-container-debugging", "--headless", "--listen=:2345", "--api-version=2", "--log"]
