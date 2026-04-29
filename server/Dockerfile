FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download
COPY ./server/ ./
RUN CGO_ENABLED=0 go build -o /server ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /server .
EXPOSE 9090
CMD ["./server"]
