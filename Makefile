# use linter for formatted code
lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.1.6 golangci-lint run

build:
	go mod vendor
	go build -o ivan cmd/ivan/main.go