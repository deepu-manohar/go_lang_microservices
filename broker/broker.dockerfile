FROM golang:1.21-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -buildvcs=false -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

#build a tiny docker
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]