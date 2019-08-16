FROM golang:1.12-alpine

WORKDIR /app

RUN apk add --no-cache git build-base

CMD ["go", "run", "."]

EXPOSE 8080
