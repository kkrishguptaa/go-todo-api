FROM --platform=linux/amd64 golang:1.26.2-alpine@sha256:f85330846cde1e57ca9ec309382da3b8e6ae3ab943d2739500e08c86393a21b1

WORKDIR /app

COPY . .

RUN go get

RUN go build -tags=jsoniter -o app .

EXPOSE 8080

CMD ["./app"]
