FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./service cmd/service/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/service .

EXPOSE 8080

CMD ["./service"]
