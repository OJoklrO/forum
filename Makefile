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
	cp -r ./storage build

test:build
	cd ./build; ./forum


