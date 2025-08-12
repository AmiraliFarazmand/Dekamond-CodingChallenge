FROM golang:1.23.12-alpine3.22 AS builder

WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080

CMD ["./main"]