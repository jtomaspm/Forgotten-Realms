FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/game_hub

FROM gcr.io/distroless/base-debian12
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 4000

CMD ["./main"]