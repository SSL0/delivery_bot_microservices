FROM golang:1.24.3-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/catalog_service/catalog_service ./cmd/catalog_service/catalog_service.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /app/cmd/catalog_service/catalog_service ./cmd/catalog_service/catalog_service
COPY ./cmd/catalog_service/config.json ./cmd/catalog_service/

EXPOSE 5001

CMD ["./cmd/catalog_service/catalog_service"]