FROM golang:alpine as builder
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"
WORKDIR /src
COPY . /src
RUN go build -o app

From alpine:latest

WORKDIR /root/
COPY --from=builder /src .
ENTRYPOINT ["./app"]
