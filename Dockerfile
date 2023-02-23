# Use base golang image from Docker Hub
FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN apk add curl

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

RUN ls -a

ENTRYPOINT ["./binary"]
