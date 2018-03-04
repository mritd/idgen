#!/bin/bash

ghr -u mritd -t $GITHUB_RELEASE_TOKEN -replace -recreate --prerelease --debug $1 dist/
