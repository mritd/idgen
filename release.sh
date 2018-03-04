#!/bin/bash

if [ ! -n "$1" ]; then
    echo "Error:release version is blank!"
else
    (cd gopath/src/github.com/mritd/idgen && git pull)
    xgo -targets="darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64" -dest dist -go latest github.com/mritd/idgen
    ghr -u mritd -t $GITHUB_RELEASE_TOKEN -replace -recreate --prerelease --debug $1 dist/
fi
