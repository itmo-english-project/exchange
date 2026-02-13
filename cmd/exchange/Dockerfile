FROM golang:1.24 as builder

WORKDIR /build

COPY . .

RUN make build TARGET=exchange CGO_ENABLED=0

FROM ubuntu:latest

USER 2000

WORKDIR /app

COPY --from=builder /build/bin/exchange ./
COPY cmd/exchange/deploy /configs

ENTRYPOINT ["/app/exchange"]