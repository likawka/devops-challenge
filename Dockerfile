# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Runtime stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
RUN chmod +x ./main
CMD ["./main"]
