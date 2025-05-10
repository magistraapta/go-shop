run:
	go run cmd/main.go

build:
	go build -o main .

swag:
	swag init -g cmd/main.go