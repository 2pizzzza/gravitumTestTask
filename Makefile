.PHONY: generate
generate:
	sqlc generate

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fmt
fmt:
	go fmt ./...