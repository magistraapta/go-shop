# Stage 1: Build
FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .
COPY . .      

EXPOSE 8080
CMD ["./main"]
