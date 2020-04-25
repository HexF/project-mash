GITCOMMIT=$(shell git rev-list -1 HEAD)
LDFLAGS="-X main.GitCommit=$(GITCOMMIT)"

.PHONY: default build image

default: test image

test:
	go test -v -cover ./...

build:
	go get -v -t -d ./...
	go build -ldflags $(LDFLAGS)

image:
	docker build -t hexf/project-mash:$(GITCOMMIT) -t hexf/project-mash:latest .