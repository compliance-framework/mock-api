# syntax=docker/dockerfile:1

FROM golang:1.26.1 AS builder

ARG VERSION=dev

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath \
	-ldflags "-s -w -X main.buildVersion=${VERSION}" \
	-o /out/mock-api .

FROM gcr.io/distroless/static-debian12:nonroot

ARG VERSION=dev
LABEL org.opencontainers.image.source="https://github.com/compliance-framework/mock-api" \
	org.opencontainers.image.version="${VERSION}"

COPY --from=builder /out/mock-api /mock-api

EXPOSE 8080

ENTRYPOINT ["/mock-api"]
