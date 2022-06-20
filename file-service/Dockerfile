FROM golang:1.17.11-alpine3.16 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir -p /go/src/file-service

WORKDIR /go/src/file-service
COPY . /go/src/file-service
COPY ./go.mod ./
COPY ./go.sum ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main
WORKDIR /go/src/file-service/preload
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o init
WORKDIR /go/src/file-service

FROM alpine
RUN mkdir -p /app
RUN mkdir -p /app/preload
RUN mkdir -p /app/config
RUN mkdir -p /vfile
COPY    --from=builder /go/src/file-service/main  /app/main
COPY    --from=builder /go/src/file-service/preload/init  /app/preload/init
COPY    --from=builder /go/src/file-service/config/config.yaml  /app/config/config.yaml
WORKDIR /app
RUN chmod 777 /app/main
RUN chmod 777 /app/preload/init
EXPOSE 8200
CMD ["/app/main"]