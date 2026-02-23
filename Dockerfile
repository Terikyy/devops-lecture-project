# Build stage
FROM golang:1.25.7 AS builder

ARG SERVICE

WORKDIR /app

COPY ${SERVICE}/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /binary ./cmd/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /binary .

CMD ["./binary"]
