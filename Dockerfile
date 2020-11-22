# Container image that runs your code
FROM golang:latest

WORKDIR $GOPATH/src/github.com/3mard/go-linus

COPY ./ ./

RUN go get -v ./...

RUN go install -v ./...

ENTRYPOINT [ "go-linus" ]
