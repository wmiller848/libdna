FROM golang:1.7.1-alpine

ADD ./ /go/src/github.com/wmiller848/libdna

RUN go install github.com/wmiller848/libdna

RUN go install github.com/wmiller848/libdna/examples/flower_genus

WORKDIR /opt
RUN echo "0,1,2,3,4" > tmp.dat
RUN echo "9,8,7,6,5" >> tmp.dat
