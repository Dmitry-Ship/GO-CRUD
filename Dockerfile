FROM golang:1.16.4-alpine

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -v -o main .

ENV PORT ${PORT}

EXPOSE ${PORT}

# CMD ["/app/main"]
