include .bingo/Variables.mk

.PHONY: drone, test

GRAFANA_TARGET_VERSION ?= main
SWAGGER_SPEC_LOCAL ?= $$(pwd)/../grafana/public/api-merged.json

generate:
	mkdir -p ./goclient
	$(SWAGGER) generate client \
	-f https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json \
	-t ./goclient \
	--skip-validation \
	--with-flatten=verbose \
	--with-flatten=remove-unused \
	--additional-initialism=DTO,API,OK \
	--keep-spec-order

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

