FROM golang:1.24.3-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/order/order ./cmd/order/main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /app/cmd/order/order ./cmd/order/order
COPY cmd/order/config.json ./cmd/order/

EXPOSE 5003

CMD ["./cmd/order/order"]