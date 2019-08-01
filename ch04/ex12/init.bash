#!/usr/bin/env bash

for i in $(seq 1 100); do
    url="https://xkcd.com/$i/info.0.json"
    echo $i
    curl --silent -o data/$i.json $url
    sleep 1
done
