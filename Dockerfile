FROM golang:1.20.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD ["/app/main"]
