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

test:build
	rm ./build/config/config.yaml
	mv ./build/config/8082.txt ./build/config/config.yaml
	cd ./build; ./forum


