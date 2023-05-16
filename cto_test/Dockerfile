############################
# Build container
############################
FROM golang:1.14.0-buster AS dep

WORKDIR /ops

ADD go.mod go.sum ./

RUN go get ./...

ADD . .
RUN go build -ldflags="-s -w" -o main && strip -s main && chmod 777 main

############################
# Final container
############################
FROM registry.cto.ai/official_images/base:2-buster-slim

WORKDIR /ops

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists

COPY --chown=ops:9999 --from=dep /ops/main .

##############################################################################
# As a security best practice the container will always run as non-root user.
##############################################################################
USER ops
