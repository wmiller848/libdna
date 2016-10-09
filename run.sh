#!/bin/bash

version=`cat ./VERSION`

docker build --no-cache -t libdna:$version .
#docker run --rm -it libdna:$version sh
docker run --rm libdna:$version /go/bin/flower_genus
