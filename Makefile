.PHONY: drone, test

GRAFANA_TARGET_VERSION ?= main
SWAGGER_CODEGEN_CLI_TAG ?= latest
SWAGGER_SPEC_LOCAL ?= $$(pwd)/../grafana/public/api-merged.json

generate:
	docker run --rm \
	--user $$(id -u):$$(id -g) \
	-v $$(pwd):$$(pwd) \
	swaggerapi/swagger-codegen-cli:${SWAGGER_CODEGEN_CLI_TAG} generate \
	-i https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json \
	-l go \
	-o $$(pwd)/goclient \
	--model-name-suffix=Model \
	--additional-properties packageName=goclient \
	-t $$(pwd)/codegen/templates
	goimports -w -v $$(pwd)/goclient

clean:
	rm -rf $$(pwd)/goclient

drone:
	drone jsonnet --stream --format --source .drone/drone.jsonnet --target .drone/drone.yml
	drone lint .drone/drone.yml
	drone sign --save grafana/grafana-api-golang-client .drone/drone.yml

test:
	go version
	golangci-lint --version
	golangci-lint run ./...
	go test -cover -race -vet all -mod readonly ./...

