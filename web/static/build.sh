#!/bin/bash -e
# Simple script that I use to build the gh-pages branch content

echo "+ rm -rf ./public/*"
rm -rf ./public/* # if we set -x before that, it displays the fully extended paths instead of the wildcard

set -x
go run generator.go -dest=./public/

cp -r ../assets/dist ./public/assets