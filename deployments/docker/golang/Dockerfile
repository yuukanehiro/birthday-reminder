FROM golang:1.17-alpine
ENV ROOT=/go/src/app

RUN apk update && apk add git
RUN apk --no-cache add postgresql-client

COPY ./ ${ROOT}

WORKDIR ${ROOT}
RUN go mod tidy \
  && go build
CMD ["go", "run", "main.go"]
