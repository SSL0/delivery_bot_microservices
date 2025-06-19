FROM golang:1.24.3-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/order_service/order_service ./cmd/order_service/order_service.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /app/cmd/order_service/order_service ./cmd/order_service/order_service
COPY ./cmd/order_service/config.json ./cmd/order_service/

EXPOSE 5003

CMD ["./cmd/order_service/order_service"]