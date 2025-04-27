# Makefile para cazarrecompenzas
build:
	docker build -t cazarrecompenzas ./cazarrecompenzas

run: build
	docker run -p 50053:50053 --name cazarrecompenzas cazarrecompenzas

clean:
	docker rmi -f cazarrecompenzas