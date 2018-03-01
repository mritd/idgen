#!/bin/bash

if [ ! -n "$1" ]; then
    echo "Error:release version is blank!"
else
    rm -rf dist
    gox -cgo=0 -osarch="darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64" -output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
    ghr -u mritd -t $GITHUB_RELEASE_TOKEN -replace -recreate --prerelease --debug $1 dist/
fi
