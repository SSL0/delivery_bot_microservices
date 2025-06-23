FROM node:18-alpine as builder
WORKDIR /app

COPY ./ ./ 

RUN npm i

ENTRYPOINT ["node", "main"]
