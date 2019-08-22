#!bin/bash

go run $GOPATH/go-lang/practice/ch01/fetch/main.go https://golang.org/ | go run findlinks1.go
