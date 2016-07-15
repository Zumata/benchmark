FROM golang:1.5

RUN go get github.com/tools/godep

ADD . /go/src/github.com/Zumata/benchmark/gzipping/server

WORKDIR /go/src/github.com/Zumata/benchmark/gzipping/server

RUN godep go install

ENTRYPOINT /go/bin/server

EXPOSE 3333
