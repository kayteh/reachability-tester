#!/bin/sh
set -e
build () {
  docker build . -f cmd/app/Dockerfile -t "${1:-katie/reachability-ui}" 
  docker build . -f cmd/node/Dockerfile -t "${2:-katie/reachability-node}" 
}

push () {
  docker push "${1:-katie/reachability-ui}" 
  docker push "${2:-katie/reachability-node}"
}

main () {
  case $1 in
  build) shift; build $@;;
  push) shift; push $@;;
  esac
}

main "$@"