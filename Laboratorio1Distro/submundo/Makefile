# Makefile para submundo

all: run

build:
	sudo docker build -t submundo .

run: build
	-sudo docker rm -f submundo || true
	sudo docker run --rm -it -p 50052:50052 --name submundo submundo

clean:
	sudo docker rmi -f submundo