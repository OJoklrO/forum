.PHONY: serve build
serve:
	swag init
	go run main.go

build:
	swag init
	mkdir build -p
	go build .
	rm ./build/forum
	mv ./forum build

test:build
	cd ./build; ./forum


