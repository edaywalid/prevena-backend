FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp cmd/server/main.go

FROM alpine:latest

COPY --from=builder /app/myapp /myapp
COPY .env /app/.env
EXPOSE 8080

CMD ["/myapp"]
