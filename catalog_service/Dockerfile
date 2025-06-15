FROM golang:1.24.3-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/catalog_service/catalog_service ./cmd/catalog_service/catalog_service.go

FROM alpine:3.22

COPY --from=build /app/cmd/catalog_service/catalog_service /app/catalog_service
WORKDIR /app
CMD ["/app/catalog_service"]