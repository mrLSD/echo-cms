FROM buildpack-deps:jessie-scm
# FROM ubuntu:16.04

COPY . /go
WORKDIR /go
ENTRYPOINT /go/echo-ws

# Open up the port where the app is running.
EXPOSE 3000
