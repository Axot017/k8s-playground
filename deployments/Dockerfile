FROM golang:1.21-alpine3.18 as builder

ARG MODULE_NAME

WORKDIR /app

RUN mkdir ${MODULE_NAME}

RUN mkdir common

COPY ${MODULE_NAME}/go.mod ${MODULE_NAME}/go.sum ./${MODULE_NAME}/

COPY common/ ./common/

WORKDIR /app/${MODULE_NAME}

RUN go mod download

COPY ${MODULE_NAME}/ ./

RUN go build -o ./service cmd/service/main.go

FROM alpine:3.18

ARG MODULE_NAME

WORKDIR /app

COPY --from=builder /app/${MODULE_NAME}/service .

EXPOSE 8080

CMD ["./service"]
