FROM --platform=linux/amd64 golang:1.21.3-alpine

WORKDIR /app

COPY . .

RUN go get

RUN go build -tags=jsoniter -o app .

EXPOSE 8080

CMD ["./app"]
