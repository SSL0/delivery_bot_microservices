FROM golang:1.24.3-alpine

WORKDIR /app
COPY ./migrate.sh ./migrate.sh
COPY schemas/ schemas/

ENTRYPOINT [ "migrate.sh" ]