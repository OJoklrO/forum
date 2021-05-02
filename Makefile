.PHONY: serve build
serve:
	swag init
	go run main.go

build:
	rm ./build -rf
	mkdir build
	go build .
	cp -r ./config ./storage build
	mv ./forum build

