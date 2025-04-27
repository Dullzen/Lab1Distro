# Makefile para cazarrecompenzas

build:
    sudo docker build -t cazarrecompenzas .

run: build
    sudo docker run -p 50053:50053 --name cazarrecompenzas cazarrecompenzas

clean:
    sudo docker rmi -f cazarrecompenzas