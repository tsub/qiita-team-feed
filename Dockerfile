FROM golang:1.9-alpine AS builder

ENV DEP_VERSION="0.3.2"

COPY . /go/src/github.com/tsub/qiita-team-feed
WORKDIR /go/src/github.com/tsub/qiita-team-feed
RUN apk add --update --no-cache \
        git \
        curl && \
    curl -fSL -o /usr/local/bin/dep "https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64" && \
    chmod +x /usr/local/bin/dep && \
    dep ensure && \
    go build -o /qiita-team-feed

FROM alpine:3.7

RUN apk add --update --no-cache \
        ca-certificates

COPY --from=builder /qiita-team-feed /
CMD /qiita-team-feed
