FROM reg-test.joinpay.com/library/golang:1.17.6-alpine3.14 AS builder
ENV GOPROXY=https://goproxy.cn
workdir /go/src
COPY main.go .
COPY go.mod .
COPY go.sum .
COPY metrics /usr/local/go/src/httpserver/metrics
RUN go build -tags netgo -o /go/bin/httpserver main.go

FROM reg-test.joinpay.com/library/alpine
EXPOSE 8080
USER nobody
ENTRYPOINT ["/opt/httpserver"]

COPY --from=builder /go/bin/httpserver /opt/
