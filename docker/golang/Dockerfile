FROM golang:1.17-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY ./app/src/main.go ${ROOT}



RUN apk update && apk add git

COPY ./go.mod ${ROOT}


WORKDIR ${ROOT}
RUN go mod tidy
RUN go build

CMD ["go", "run", "main.go"]
