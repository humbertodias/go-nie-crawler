#!/bin/sh

project="go-nie-crawler"

echo "whoami: $(whoami)"
ls -lha $(pwd)

build(){
  echo "Attempting to build $project for OS X"
  make build
}


build

echo 'Logs from build'
cat $(pwd)/unity.log