migrate:
	go run migrations/main.go

run:
	go run main.go

test:
	go test ./... -v

test-cov:
	go test ./... -cover -v -coverprogile=c.out && go tool cover -html=c.out -o coverage.html

build:
	go build -ldflags "-s -w" main.go