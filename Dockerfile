# syntax = docker/dockerfile:1-experimental

FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS base
RUN apk add git
ENV CGO_ENABLED=0
WORKDIR /src
COPY go.* .
RUN go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
  --mount=type=cache,target=/root/.cache/go-build \
  GOOS=${TARGETOS} \
  GOARCH=${TARGETARCH} \
  go build -o /out/api .

FROM base AS unit-test
RUN --mount=target=. \
  --mount=type=cache,target=/root/.cache/go-build \
  GOOS=${TARGETOS} \
  GOARCH=${TARGETARCH} \
  go test -v ./...

FROM golangci/golangci-lint:v1.27-alpine AS lint-base

FROM base AS lint
RUN --mount=target=. \
  --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/root/.cache/golangci-lint \
  GOOS=${TARGETOS} \
  GOARCH=${TARGETARCH} \
  golangci-lint run --timeout 10m0s ./...

FROM scratch AS bin-unix
COPY --from=build /out/api /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/api /api.exe

FROM bin-${TARGETOS} AS bin

FROM scratch
COPY --from=build /out/api /

EXPOSE 8080

CMD [ "/api" ]