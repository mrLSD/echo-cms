FROM buildpack-deps:wheezy-scm

COPY . /go
WORKDIR /go
ENTRYPOINT /go/server

# Open up the port where the app is running.
EXPOSE 3000
