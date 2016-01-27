FROM golang:1.5.2-alpine
MAINTAINER Byron Mansfield <byron@byronmansfield.com>

WORKDIR /go/src/github.com/byronmansfield/go-server

RUN apk update && \
    apk add curl git && \
    rm -rf /var/cache/apk/*

COPY . /go/src/github.com/byronmansfield/go-server
RUN go get github.com/byronmansfield/go-server

RUN go get
RUN go build

EXPOSE 8080
ENTRYPOINT [ "./go-server" ]
