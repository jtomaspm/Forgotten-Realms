FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/game_server

FROM gcr.io/distroless/base-debian12
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/migrations/game_server ./migrations/game_server
ARG REALM_SETTINGS
COPY --from=builder /app/realm_settings/$REALM_SETTINGS/* ./realm_settings/

CMD ["./main"]