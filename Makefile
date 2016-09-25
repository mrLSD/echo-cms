run:
	@go run server.go

test:
	@./runtest.sh

build:
	@go build

fmt:
	@gofmt -w -l .

docker-build:
	@go build server.go && docker build -t echo-cms . && docker tag echo-cms 192.168.1.107:5000/echo-cms && docker push 192.168.1.107:5000/echo-cms

docker-rmi:
	@docker rmi echo-cms

docker-bld:
	@docker rmi echo-cms && docker build -t echo-cms . 
