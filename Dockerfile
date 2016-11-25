# FROM buildpack-deps:wheezy-scm
FROM ubuntu:16.04

COPY . /go
WORKDIR /go
ENTRYPOINT /go/echo-cms

# Open up the port where the app is running.
EXPOSE 3000
