LDFLAGS="-X main.GitCommit=$(shell git rev-list -1 HEAD)"

.PHONY: build clean run
build: project-mash
clean:
	rm project-mash
run: project-mash
	./project-mash
	
project-mash:
	go build -ldflags $(LDFLAGS)