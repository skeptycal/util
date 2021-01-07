#!/usr/bin/env bash

set -e
echo "" >| coverage.txt
# for d in $(go list ./... | grep -v vendor); do echo $d; done;
for d in $(go list ./... | grep -v vendor); do
	case $1 in
	-bm)
		go test -run=^$ -bench="$d" -benchmem "$d"
		;;
	*)
		go test -race -coverprofile=profile.out -covermode=atomic "$d"
		;;
	esac
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
