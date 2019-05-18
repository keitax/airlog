FROM golang:1.12-alpine

WORKDIR /app

RUN apk add --no-cache git

CMD ["go", "run", "./cmd/airlog"]

EXPOSE 8080
