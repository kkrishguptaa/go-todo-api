FROM --platform=linux/amd64 golang:1.21.5-alpine@sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b

WORKDIR /app

COPY . .

RUN go get

RUN go build -tags=jsoniter -o app .

EXPOSE 8080

CMD ["./app"]
