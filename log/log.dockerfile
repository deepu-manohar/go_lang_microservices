FROM golang:1.21-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o logApp ./cmd/api

RUN chmod +x /app/logApp

#build a tiny docker
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/logApp /app

CMD ["/app/logApp"]