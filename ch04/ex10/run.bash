#!bin/bash

go build issues.go searchIssues.go
./issues repo:golang/go is:open json decoder
