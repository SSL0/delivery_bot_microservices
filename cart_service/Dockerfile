FROM golang:1.24.3-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/cart_service/cart_service ./cmd/cart_service/cart_service.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /app/cmd/cart_service/cart_service ./cmd/cart_service/cart_service
COPY ./migrations/ ./migrations
COPY ./cmd/cart_service/config.json ./cmd/cart_service/

EXPOSE 5001

CMD ["./cmd/cart_service/cart_service"]