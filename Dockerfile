# Container image that runs your code
FROM golang:alpine

WORKDIR $GOPATH/src/github.com/3mard/go-linus

COPY ./ ./


RUN go install -v ./...

ENTRYPOINT [ "go-linus" ]
