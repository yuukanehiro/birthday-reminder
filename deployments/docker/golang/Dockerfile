FROM golang:1.17-alpine
ENV ROOT=/go/src/app
ENV DIR_CMD=/go/src/app/cmd

RUN apk update && apk add git
RUN apk --no-cache add postgresql-client
COPY ./ ${ROOT}

WORKDIR ${DIR_CMD}
RUN go mod tidy \
  && go build
CMD ["go", "run", "main.go"]
