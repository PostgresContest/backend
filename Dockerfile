FROM golang:1.19.7-alpine

RUN mkdir /src
ADD . /src/

WORKDIR /src
RUN go build -o backend

ENTRYPOINT ["/src/backend", "serve"]