#!/usr/bin/env bash 
set -e

args="${@:2}"
version=`cat ./VERSION`

case "$1" in
  run)
    ./scripts/run.sh -v $version $args
    ;;
  test)
    ./scripts/test.sh -v $version $args
    ;;
  *)
    echo "Unknown command $1"
    exit 2
    ;;
esac



