FROM golang:alpine

RUN apk add --no-cache git make && rm -rf /var/cache/apk/*
WORKDIR $GOPATH/src
COPY . .
RUN make build

ENTRYPOINT ["./project-mash"]