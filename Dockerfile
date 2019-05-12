FROM golang:1.12.5-alpine3.9 AS builder

COPY . /go/src/github.com/mritd/idgen

WORKDIR /go/src/github.com/mritd/idgen

ENV GO111MODULE on
ENV GOPROXY https://athens.azurefd.net

RUN apk upgrade \
    && apk add git \
    && BUILD_VERSION=$(cat version) \
    && BUILD_DATE=$(date "+%F %T") \
    && COMMIT_SHA1=$(git rev-parse HEAD) \
    && go install github.com/gobuffalo/packr/v2/packr2 \
    && packr2 clean && packr2 \
    && go install -ldflags  "-X 'github.com/mritd/idgen/cmd.Version=${BUILD_VERSION}' \
                            -X 'github.com/mritd/idgen/cmd.BuildDate=${BUILD_DATE}' \
                            -X 'github.com/mritd/idgen/cmd.CommitID=${COMMIT_SHA1}'"

FROM alpine:3.9 AS dist

LABEL maintainer="mritd <mritd@linux.com>"

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

RUN apk upgrade \
    && apk add bash tzdata \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/idgen /usr/bin/idgen

CMD ["idgen","server"]
