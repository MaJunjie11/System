#!/bin/bash
f="./output/"
if [ -e $f ]
then
  go build -o output/main
else
  mkdir output
  go build -o output/main
fi

