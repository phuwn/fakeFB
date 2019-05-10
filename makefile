#Install missing dependencies
dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

remove:
	docker-compose stop; docker-compose rm -f

init:
	docker-compose up -d
	sql-migrate up -env="local"

#Run app on local
dev:
	@while ! nc -z localhost 5433; do echo "WAITING FOR POSTGRES STARTUP" && sleep 1; done;
	go run cmd/server/*.go

run:
	docker build -t fakeFb .
	docker-compose up -d

build: 
	go build -o server r cmd/server/*.go

build-alpine:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/*.go
