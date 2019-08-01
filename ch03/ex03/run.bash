#!bin/bash

go run main.go > surface.svg
open -a Google\ Chrome surface.svg
