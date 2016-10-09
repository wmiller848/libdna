#!/bin/bash
set -e

version=`cat ./VERSION`

docker build --no-cache -t libdna:$version .
#docker run --rm -it libdna:$version sh
docker run --rm libdna:$version sh -c "/go/bin/flower_genus < tmp.dat"
