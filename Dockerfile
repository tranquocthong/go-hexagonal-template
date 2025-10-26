# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go build -ldflags="-s -w" -o /out/service ./cmd/service

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /out/service /service
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/service"]


