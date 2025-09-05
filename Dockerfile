# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cgemini .

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cgemini .

CMD ["./cgemini"]
