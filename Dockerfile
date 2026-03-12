FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o httpecho .

FROM alpine:3.19
COPY --from=builder /app/httpecho /usr/local/bin/
EXPOSE 8080
CMD ["httpecho"]
