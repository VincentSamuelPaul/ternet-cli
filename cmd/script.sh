#!/bin/bash

# go test -v

rm ./main

go build main.go

./main -help
./main -login
./main -logout
./main -signup
./main -browse
./main -newpost