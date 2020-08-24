FROM golang:1.15-alpine AS builder

WORKDIR /build
COPY . /build/
RUN ./build.sh


FROM alpine:3
RUN apk add --no-cache \
    dumb-init

WORKDIR /app
COPY --from=builder /build/output /app/

EXPOSE 8080
ENTRYPOINT [ "dumb-init", "--", "/app/webrcon-server" ]