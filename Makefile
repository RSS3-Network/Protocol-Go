generate: go-generate gen-type

lint: generate
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 run

go-generate:
	go generate ./...

gen-type:
	scripts/gen-type.sh