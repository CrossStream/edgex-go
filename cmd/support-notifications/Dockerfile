#
# Copyright (c) 2018
# Cavium
#
# SPDX-License-Identifier: Apache-2.0
#

FROM golang:1.12-alpine AS builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/edgexfoundry/edgex-go

# The main mirrors are giving us timeout issues on builds periodically.
# So we can try these.

RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories

RUN apk update && apk add make && apk add bash git

COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . .
RUN make cmd/support-notifications/support-notifications

FROM scratch

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2018: Cavium'

COPY --from=builder /bin/bash /bin/bash
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/support-notifications/Attribution.txt /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/support-notifications/support-notifications /
COPY --from=builder /go/src/github.com/edgexfoundry/edgex-go/cmd/support-notifications/res/docker/configuration.toml /res/docker/configuration.toml
ENTRYPOINT ["/support-notifications","--registry","--profile=docker","--confdir=/res"]

