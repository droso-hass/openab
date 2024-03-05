
all: v2-fw server

server: server/*
	cd server && go build

v2-metalc: tagtag/mtl_linux/*
	cd tagtag/mtl_linux && make all

v2-fw: v2-metalc tagtag/src*
	cd tagtag/src && ./include.py ./main.mtl ./nominal.mtl
	cd tagtag/mtl_linux && ./mtl_compiler -s ../src/nominal.mtl ../nominal.bin
	cp tagtag/nominal.bin server/static

v2-simu: v2-metalc tagtag/src*
	cd tagtag/src && cp main.mtl _main.mtl
	sed -i '1s/^/#define SIMU\n/' tagtag/src/_main.mtl
	cd tagtag/src && ./include.py ./_main.mtl ./nominal.mtl
	rm tagtag/src/_main.mtl
	cp tagtag/src/nominal.mtl tagtag/mtl_linux
	cd tagtag/mtl_linux && ./mtl_simu --dologtime --logs init,vm,simunet,simuleds,simuaudio

run: all
	sudo ./server/openab --log-level=debug --nats-server="nats://localhost:4222" --nats-user="openab" --nats-password="1234"

