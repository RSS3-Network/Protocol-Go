generate: go-generate gen-openapi

lint: generate
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 run

go-generate:
	go generate ./...


gen-openapi: gen-type gen-metadata gen-api

gen-type:
	scripts/gen-type.sh

gen-metadata:
	scripts/gen-metadata.sh

gen-api:
	docker run --rm -v ${PWD}:/local \
		redocly/openapi-cli@latest \
		bundle /local/base.yaml -o /local/api.yaml
# or just use npx @redocly/openapi-cli@latest bundle base.yaml > api.yaml