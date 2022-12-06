FROM golang:1.18-alpine AS builder
WORKDIR /workspace

COPY go.mod go.mod
COPY main.go main.go

RUN CGO_ENABLED=0 go build -a -o testA ./main.go

FROM ubuntu:22.10
WORKDIR /
COPY --from=builder /workspace/testA .
EXPOSE 8000

ENTRYPOINT ["/testA"]
