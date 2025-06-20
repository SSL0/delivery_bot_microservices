FROM golang:1.24.3-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/catalog/catalog ./cmd/catalog/main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /app/cmd/catalog/catalog ./cmd/catalog/catalog
COPY cmd/catalog/config.json ./cmd/catalog/

EXPOSE 5001

CMD ["./cmd/catalog/catalog"]