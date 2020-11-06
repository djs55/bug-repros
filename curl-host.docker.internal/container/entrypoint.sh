#!/bin/sh

set -e

i=0
while [ $i -lt 10000 ]; do
  curl http://host.docker.internal:8080/
  i=$(( i + 1 ))
  echo $i
done