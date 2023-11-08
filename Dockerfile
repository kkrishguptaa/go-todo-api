FROM golang:1.21

WORKDIR /usr/src/go-todo-api

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/go-todo-api ./...

ENV GIN_MODE=release
EXPOSE $PORT
EXPOSE 10000

CMD ["go-todo-api"]
