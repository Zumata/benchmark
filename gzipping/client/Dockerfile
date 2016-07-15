FROM golang:1.5

RUN go get github.com/tools/godep

ADD . /go/src/github.com/Zumata/benchmark/gzipping/client

WORKDIR /go/src/github.com/Zumata/benchmark/gzipping/client

RUN godep go install

ENTRYPOINT /go/bin/client

EXPOSE 3334
