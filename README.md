# Grafana HTTP API Client for Go

This library provides a low-level client to access Grafana [HTTP API](https://grafana.com/docs/grafana/latest/http_api/).

:warning: This repository is still active but not under heavy development.
Contributions to this library offering support for the [Terraform provider for Grafana](https://github.com/grafana/terraform-provider-grafana) will be prioritized over generic ones.

## Tests

To lint and run unit tests:

```
make test
```

To run integration tests against a local Docker instance of Grafana:

```
make integ-test-docker
```