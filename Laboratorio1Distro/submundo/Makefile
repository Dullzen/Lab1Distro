# Makefile para submundo

all: run

build:
	sudo docker build -t submundo .

run: build
	-sudo docker rm -f submundo || true
	sudo docker run --rm -it -p 50053:50053 --name submundo submundo

clean:
	sudo docker rmi -f submundo