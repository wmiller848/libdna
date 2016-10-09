FROM golang:1.7.1-alpine

ADD ./ /go/src/github.com/wmiller848/libdna

RUN go install github.com/wmiller848/libdna

RUN go install github.com/wmiller848/libdna/examples/flower_genus
