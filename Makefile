BUILD_VERSION   := $(shell cat version)
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

all:
	bash .cross_compile.sh

release: all
	ghr -u mritd -t $(GITHUB_TOKEN) -replace -recreate -name "Bump v${BUILD_VERSION}" --debug ${BUILD_VERSION} dist

pre-release: all
	ghr -u mritd -t $(GITHUB_TOKEN) -replace -recreate -prerelease -name "Bump v${BUILD_VERSION}" --debug ${BUILD_VERSION} dist

clean:
	rm -rf dist

install:
	go install -ldflags "-X 'github.com/mritd/idgen/cmd.Version=${BUILD_VERSION}' \
                         -X 'github.com/mritd/idgen/cmd.BuildDate=${BUILD_DATE}' \
                         -X 'github.com/mritd/idgen/cmd.CommitID=${COMMIT_SHA1}'"

docker:
	docker build -t mritd/idgen:${BUILD_VERSION} .

.PHONY: all release clean install

.EXPORT_ALL_VARIABLES:
