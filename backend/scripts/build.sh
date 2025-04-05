#!/bin/sh

go build -o bin/game_server cmd/game_server/main.go
go build -o bin/auth_server cmd/auth_server/main.go
go build -o bin/game_hub cmd/game_hub/main.go