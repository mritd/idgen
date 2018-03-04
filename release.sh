#!/bin/bash

if [ ! -n "$1" ]; then
    echo "Error:release version is blank!"
else
    ghr -u mritd -t $GITHUB_RELEASE_TOKEN -replace -recreate --prerelease --debug $1 dist/
fi
