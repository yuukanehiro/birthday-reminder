FROM golang:1.17-alpine
ENV ROOT=/go/src/app

RUN apk update && apk add git
COPY ./app/src/ ${ROOT}


WORKDIR ${ROOT}
RUN go mod tidy \
  && go build
CMD ["go", "run", "main.go"]
