# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download -x
# Copy the go source
COPY main.go main.go
COPY pkg/ pkg/
COPY web/ web/
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -a -o /workspace/goflat main.go

FROM scratch
WORKDIR /
COPY --from=builder /workspace/goflat /
COPY --from=builder /workspace/web/dist/ /dist/
USER 65532:65532
ENTRYPOINT ["/goflat"]
EXPOSE 3000
