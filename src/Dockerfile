FROM golang:1.16.4-alpine

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

# RUN go build -v -o main .

ENV PORT ${PORT}

EXPOSE ${PORT}

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -v -o main ." --command=./main
# ENTRYPOINT CompileDaemon --command="go run ."

# CMD ["/app/main"]
