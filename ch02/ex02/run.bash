#!/bin/bash

go build cf.go length.go weight.go
./cf -t temp 20
./cf -t temp <<EOS
20
exit
EOS
./cf -t length 20
./cf -t length <<EOS
20
exit
EOS
./cf -t weight 20
./cf -t weight <<EOS
20
exit
EOS
./cf -t tem  20