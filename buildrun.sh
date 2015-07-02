#!/bin/bash -xe
echo getting dependencies
go get -u github.com/base2services/golang-jenkins
echo running
go run *.go $@
