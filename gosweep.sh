#!/bin/bash
# From: https://github.com/h12w/gosweep/blob/master/gosweep.sh
# The script does automatic checking on a Go package and its sub-packages, including:
# 1. gofmt         (http://golang.org/cmd/gofmt/)
# 2. goimports     (https://github.com/bradfitz/goimports)
# 3. golint        (https://github.com/golang/lint)
# 4. go vet        (http://golang.org/cmd/vet)
# 5. race detector (http://blog.golang.org/race-detector)
# 6. test coverage (http://blog.golang.org/cover)

set -e

# Automatic checks
echo "+ gofmt"
test -z "$(gofmt -l -w .     | tee /dev/stderr)"
echo "+ goimports"
test -z "$(goimports -l -w . | tee /dev/stderr)"
echo "+ golint"
test -z "$(golint .          | tee /dev/stderr)"
echo "+ go vet"
go vet ./...
echo "+ test race"
env GORACE="halt_on_error=1" go test -short -race ./...

# Run test coverage on each subdirectories and merge the coverage profile.

echo "mode: count" > profile.cov

# Standard go tooling behavior is to ignore dirs with leading underscores
echo "+ coverage"
for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_*' -type d);
do
if ls $dir/*.go &> /dev/null; then
    go test -short -covermode=count -coverprofile=$dir/profile.tmp $dir
    if [ -f $dir/profile.tmp ]
    then
        cat $dir/profile.tmp | tail -n +2 >> profile.cov
        rm $dir/profile.tmp
    fi
fi
done

go tool cover -func profile.cov

# To submit the test coverage result to coveralls.io,
# use goveralls (https://github.com/mattn/goveralls)
# goveralls -coverprofile=profile.cov -service=travis-ci