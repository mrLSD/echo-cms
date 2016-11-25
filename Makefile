#
# Makefile
# @author Evgeny Ukhanov <mrlsd@ya.ru>
#

.PHONY: run, test, build, fmt, docker-build, docker-rmi, docker-bld, install-go
default: run

GOVERSION = 1.7.3

run:
	@go run server.go

test:
	@./runtest.sh

build:
	@go build

fmt:
	@gofmt -w -l .

docker-build-registry:
	@go build server.go
	@docker build -t echo-cms .
	@docker tag echo-cms 192.168.1.107:5000/echo-cms
	@docker push 192.168.1.107:5000/echo-cms

docker-rmi:
	@docker rmi echo-cms

docker-build:
	@go build
	@docker rmi echo-cms
	@docker build -t echo-cms .

docker-run:
	@docker run -p 3000:3000 --rm -it echo-cms

install-go:
	@echo Installing Golang v$(GOVERSION)...
	@echo [================================================================]
	@wget -O /tmp/go.tar.gz https://storage.googleapis.com/golang/go$(GOVERSION).linux-amd64.tar.gz
	@sudo tar -C /usr/local -xzf /tmp/go.tar.gz
	@sudo ln -s /usr/local/go/bin/go /usr/local/bin/go
	@sudo ln -s /usr/local/go/bin/godoc /usr/local/bin/godoc
	@sudo ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
	@echo [================================================================]
	@echo Done.

vendoring:
	@echo Start vendoring...
	govendor init
	for package in $(VENDORPACKAGES) ; do \
		gvt fetch $$package ; \
	done