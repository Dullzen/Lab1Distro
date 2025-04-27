# Makefile para marina

all: run

build:
	sudo docker build -t marina .

run: build
	-sudo docker rm -f marina || true
	sudo docker run --rm -it -p 50052:50052 --name marina marina

clean:
	sudo docker rmi -f marina