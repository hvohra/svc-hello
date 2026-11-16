FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/svc-hello .

FROM gcr.io/distroless/static-debian12:nonroot
# Links the GHCR package to this repo.
LABEL org.opencontainers.image.source="https://github.com/hvohra/svc-hello"
COPY --from=build /out/svc-hello /svc-hello
# Numeric, so Kubernetes can verify runAsNonRoot.
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/svc-hello"]
