NAME=httpserver
IMAGE_NAME=metazone/$(NAME)
CGO_ENABLED=0
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
BUILD_DATE=$(shell date +%s)
VERSION_PACKAGE=main
LDFLAGS=-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) -X $(VERSION_PACKAGE).GitTag=$(GIT_TAG) -X $(VERSION_PACKAGE).BuildDate=$(BUILD_DATE)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-s -w ${LDFLAGS}" -o ../../bin/$(NAME)
.PHONY: build

docker:
	sed "s/ENV VERSION=\"\"/ENV VERSION=${GIT_TAG}/" Dockerfile > Dockerfile.tmp
	docker build -f Dockerfile.tmp -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	rm -rf Dockerfile.tmp
.PHONY: docker

docker_push: docker
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):latest
.PHONY: docker_push

clean:
	go clean
	rm -rf ../../bin/$(NAME)
