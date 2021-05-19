migrate:
	GO_ENV="development" go run migrations/main.go

run:
	GO_ENV="development" go run main.go

test:
	GO_ENV="development" go test ./... -v

test-cov:
	GO_ENV="development" go test ./... -cover -v -coverprogile=c.out && go tool cover -html=c.out -o coverage.html

build:
	go build -ldflags "-s -w" main.go