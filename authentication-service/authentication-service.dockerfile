FROM golang:1.21-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -buildvcs=false -o authApp ./cmd/api

RUN chmod +x /app/authApp

#build a tiny docker
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authApp /app

CMD ["/app/authApp"]