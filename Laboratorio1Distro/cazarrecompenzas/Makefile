# Makefile para cazarrecompenzas

all: run

build:
	sudo docker build -t cazarrecompenzas .

run: build
	-sudo docker rm -f cazarrecompenzas || true
	sudo docker run --rm -it -p 50053:50053 --name cazarrecompenzas cazarrecompenzas

clean:
	sudo docker rmi -f cazarrecompenzas
