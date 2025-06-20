FROM golang:1.24.3-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/cart/cart ./cmd/cart/main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /app/cmd/cart/cart ./cmd/cart/cart
COPY ./cmd/cart/config.json ./cmd/cart/

EXPOSE 5002

CMD ["./cmd/cart/cart"]