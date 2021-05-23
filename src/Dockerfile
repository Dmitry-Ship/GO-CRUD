# stage 1: install dependencies
FROM golang:1.16.4-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get github.com/codegangsta/gin

# stage 2: build binary for production
FROM base as build

COPY ./ ./app

RUN go build -v -o main .

# stage 3: generate minimal image for production
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app .

CMD ["./main"]  



