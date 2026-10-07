# syntax=docker/dockerfile:1

FROM golang:1.26.1 AS builder

ARG VERSION=dev

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . ./

# The Makefile owns the build flags (static binary, version ldflags).
RUN GOOS=linux make build VERSION=${VERSION}

FROM gcr.io/distroless/static-debian12:nonroot

ARG VERSION=dev
LABEL org.opencontainers.image.source="https://github.com/compliance-framework/mock-api" \
	org.opencontainers.image.version="${VERSION}"

COPY --from=builder /src/dist/mock-api /mock-api

EXPOSE 8080

ENTRYPOINT ["/mock-api"]
