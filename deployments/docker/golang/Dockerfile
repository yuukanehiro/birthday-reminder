FROM golang:1.19-alpine as deploy-builder

ENV APP_NAME=birthday-reminder
ENV ROOT=/go/src/app
# override by CI/CD! ["develop"|"staging"|"production"]
ENV ENVIRONMENT="local"

COPY ./ ${ROOT}
WORKDIR ${ROOT}

RUN apk update && \
  apk --no-cache add git postgresql-client

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN set -eux && \
  go mod tidy && \
  go build -ldflags '-s -w'


# ----------------------------------------
FROM debian:bookworm-slim as deploy

ENV APP_NAME=birthday-reminder
ENV ROOT=/go/src/app
# override by CI/CD! ["develop"|"staging"|"production"]
ENV ENVIRONMENT="local"

COPY --from=deploy-builder ./ ${ROOT}
WORKDIR ${ROOT}

RUN apt-get update
COPY --from=deploy-builder ${ROOT}/$APP_NAME .

CMD ./$APP_NAME
