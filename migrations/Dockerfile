FROM golang:1.24.3-alpine

RUN apk add --no-cache bash
WORKDIR /app
COPY ./ /app/
ENTRYPOINT [ "/bin/bash", "/app/migrate.sh" ]