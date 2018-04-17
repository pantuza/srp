FROM golang

RUN go get github.com/tools/godep

RUN go get github.com/pantuza/srp

WORKDIR $GOPATH/src/github.com/pantuza/srp

EXPOSE 4242

RUN godep save ./...

ENTRYPOINT make run
