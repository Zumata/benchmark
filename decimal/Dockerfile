FROM golang:1.5

RUN mkdir -p /go/src/benchmark/decimal
WORKDIR /go/src/benchmark/decimal

# this will ideally be built by the ONBUILD below ;)
CMD ["go", "test", "-bench=."]

COPY . /go/src/benchmark/decimal
RUN go-wrapper download
RUN go-wrapper install
