# Container image that runs your code
FROM golang:alpine

COPY ./ ./


CMD go RUN ./...
