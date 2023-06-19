# Grafana HTTP API Client for Go

This library provides a low-level client to access Grafana [HTTP API](https://grafana.com/docs/grafana/latest/http_api/).

:warning: This repository is still active but not under heavy development.
Contributions to this library offering support for the [Terraform provider for Grafana](https://github.com/grafana/terraform-provider-grafana) will be prioritized over generic ones.

## Tests

To run the tests:

```
make test
```

## Generate the client

First, install bingo locally. Bingo helps with having consistent tool versioning across dev environments.
```bash
go install github.com/bwplotka/bingo@latest
```

Finally, generate the client using the `make` step. This step will also install swagger through bingo.
```bash
make generate
```

### Linter

THe linter (speficially, `depguard`) complains if an import has not been added to the list of allowed imports. Ensure to add any new imports to `.golangci.toml` under `[linters-settings.depguard.rules.main]`.
