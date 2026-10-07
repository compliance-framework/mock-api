VERSION ?= dev
LDFLAGS := -X main.buildVersion=$(VERSION)

.PHONY: build test docker-build

build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o dist/mock-api .

test:
	go test ./... -coverprofile cover.out

docker-build:
	docker build --build-arg VERSION=$(VERSION) -t mock-api:$(VERSION) .
