# Use base golang image from Docker Hub
FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

# Install dependencies in go.mod and go.sum
# COPY go.mod go.sum ./
COPY . .
# RUN go mod download

# Copy rest of the application source code
RUN go mod tidy
# Compile the application to /fiber-app
RUN go build -o binary

RUN ls -a

# Expose port and start the application
ENTRYPOINT ["/app/binary"]
