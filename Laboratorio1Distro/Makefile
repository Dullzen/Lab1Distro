# Makefile para cazarrecompenzas
build:
	sudo docker build -t cazarrecompenzas ./cazarrecompenzas

run: build
	-docker rm -f cazarrecompenzas || true
	sudo docker run -p 50053:50053 --name cazarrecompenzas cazarrecompenzas

clean:
	sudo docker rmi -f cazarrecompenzas
