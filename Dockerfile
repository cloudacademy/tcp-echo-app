FROM golang:1.19.0-bullseye as builder

WORKDIR /go/src/tcpapp/
COPY main.go ./
COPY go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o tcpapp .

FROM alpine:3.16.2
RUN apk add ca-certificates

COPY --from=builder /go/src/tcpapp/tcpapp /go/bin/

ENV HOSTPORT=0.0.0.0:9091
EXPOSE 9091

CMD ["/go/bin/tcpapp"]