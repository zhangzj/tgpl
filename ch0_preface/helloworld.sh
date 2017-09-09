#!/usr/bin/env bash

go get gopl.io/ch1/helloworld
#fetch build install
## will fail because of network
$GOPATH/bin/helloworld
# run

go version
# will print the go version now
