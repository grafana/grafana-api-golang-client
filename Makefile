GRAFANA_VERSION ?= 9.0.2

.PHONY: drone, test, acc-test-docker, acc-test

drone:
	drone jsonnet --stream --format --source .drone/drone.jsonnet --target .drone/drone.yml
	drone lint .drone/drone.yml
	drone sign --save grafana/grafana-api-golang-client .drone/drone.yml

test:
	go version
	golangci-lint --version
	golangci-lint run ./...
	go test -cover -race -vet all -mod readonly ./...

acc-test-docker:
	GRAFANA_VERSION=$(GRAFANA_VERSION) \
		docker-compose \
		-f ./docker-compose.yml \
		run --rm -e TESTARGS="$(TESTARGS)" \
		grafana-api-golang-client \
		make acc-test

acc-test:
	go test --tags=acceptance ./... -v $(TESTARGS) -timeout 120m
