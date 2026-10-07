# mock-api

Mock repo for developing CCF release automation. Not a product.

It mirrors a CCF Go service (like `api`) in miniature:

- `main.go`: HTTP server on `:8080`; `GET /healthz` returns `200 ok`.
- `pkg/version`: exported `Version` const and `Hello()`, imported by `mock-agent`.
- `Dockerfile`: multi-stage build of a static binary; `--build-arg VERSION=...` sets the
  binary's build version.

```sh
make build                      # dist/mock-api (VERSION=dev by default)
make test                       # go test ./... with coverage
make docker-build VERSION=1.2.3 # docker build --build-arg VERSION=1.2.3
```
