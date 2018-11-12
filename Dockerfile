FROM golang:1.11.2-alpine

ADD . /go/src/github.com/shivakar/docker-hello-world
RUN go install -x github.com/shivakar/docker-hello-world/cmd/docker-hello-world

ENTRYPOINT ["docker-hello-world"]
EXPOSE 8080
