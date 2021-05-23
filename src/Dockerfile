FROM golang:1.16.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o main .

# ENTRYPOINT CompileDaemon --build="go build -v -o main ." --command=./main
# ENTRYPOINT CompileDaemon  --command="go run main.go"

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app .

EXPOSE ${PORT}

CMD ["./main"]  



