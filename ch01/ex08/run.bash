#!/bin/bash

go run fetch.go http://gopl.io > url1

go run fetch.go gopl.io > url2

VALUE=$(diff url1 url2)

if [ -z "$VALUE" ]; then
  echo "正常終了"
  exit 0
fi
  echo "異常終了"
  exit 1