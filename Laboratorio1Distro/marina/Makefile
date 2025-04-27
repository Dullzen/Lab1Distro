# Makefile para gobierno

all: run

build:
	sudo docker build -t gobierno .

run: build
	-sudo docker rm -f gobierno || true
	sudo docker run --rm -it -p 50051:50051 --name gobierno gobierno

clean:
	sudo docker rmi -f gobierno