.PHONY: serve build
init:
	swag init

# with config.yaml
serve:init
	go run main.go

test:init
	HttpPort=8888 go run main.go

build:init
	CGO_ENABLED=0 go build .

deploy:build
	scp -r ./forum tcloud:~/mg-forum

test: init
	go test ./...
