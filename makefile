migrate:
	SERVICES_PROFILE="development" go run migrations/main.go

run:
	SERVICES_PROFILE="development" go run main.go

test:
	SERVICES_PROFILE="development" go test ./... -v

test-cov:
	SERVICES_PROFILE="development" go test ./... -cover -v -coverprogile=c.out && go tool cover -html=c.out -o coverage.html

build:
	go build -ldflags "-s -w" main.go