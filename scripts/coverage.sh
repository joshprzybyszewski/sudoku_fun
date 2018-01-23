#!/usr/bin/env bash

set -e

# http://stackoverflow.com/a/21142256/2055281
# I got this directly from https://github.com/resal81/molkit/blob/master/scripts/coverage.sh

echo "mode: atomic" > coverage.txt

for d in $(find ./* -maxdepth 10 -type d -not -path "*vendor*"); do
    if ls $d/*_test.go &> /dev/null; then
        go test -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            echo "FINISHED testing $d"
            cat profile.out | grep -v "mode: " >> coverage.txt
            rm profile.out
        fi
    fi
done
