#!/bin/bash

if [ -z "$1" ]
then
  echo "No repo github supplied"
  exit 1
fi

url=$1

curl \
  -H "Accept: application/vnd.github.v3+json" \
  -s https://api.github.com/repos/$url/releases |jq -r '.[0].tag_name'
