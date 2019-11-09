FROM golang:1.13-alpine3.10 AS builder

ENV GO111MODULE=on \
    GOOS=linux \
    CGO_ENABLED=0

WORKDIR /go/src/github.com/cyberzeed/scg-go/

COPY . .

RUN go mod download \
    && go build -mod=readonly -v -o server

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/cyberzeed/scg-go/server /usr/local/bin

EXPOSE 8080

CMD ["server"]
