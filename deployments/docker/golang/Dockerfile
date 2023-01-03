##
## Build
##
FROM golang:1.18-alpine3.16

ENV APP_NAME=birthday-reminder
ENV ROOT=/go/src/app
# override by CI/CD! ["develop"|"staging"|"production"]
ENV ENVIRONMENT="local"

COPY ./ ${ROOT}
WORKDIR ${ROOT}

ENV GO111MODULE on
RUN apk update && \
  apk --no-cache add git postgresql-client && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go install github.com/go-delve/delve/cmd/dlv@latest

RUN set -eux && \
  go mod tidy && \
  go build
CMD go run $APP_NAME



