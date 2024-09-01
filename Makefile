run:
	go run cmd/app/main.go

build:
	go build -o bin/bot cmd/bot/main.go

regcmd:
	go run cmd/register_commands/main.go

rmcmd:
	go run cmd/remove_commands/main.go