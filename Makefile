
all: v2-fw server

server: server/internal/* server/main.go server/go.mod server/go.sum
	cd server && go build

v2-metalc: tagtag_fw/mtl_linux/*
	cd tagtag_fw/mtl_linux && make all

v2-fw: v2-metalc tagtag_fw/src*
	cd tagtag_fw/src && ./include.py ./main.mtl ./nominal.mtl
	cd tagtag_fw/mtl_linux && ./mtl_compiler -s ../src/nominal.mtl ../nominal.bin
	cp tagtag_fw/nominal.bin server/static

v2-simu: v2-metalc tagtag_fw/src*
	cd tagtag_fw/src && cp main.mtl _main.mtl
	sed -i '1s/^/#define SIMU\n/' tagtag_fw/src/_main.mtl
	cd tagtag_fw/src && ./include.py ./_main.mtl ./nominal.mtl
	rm tagtag_fw/src/_main.mtl
	cp tagtag_fw/src/nominal.mtl tagtag_fw/mtl_linux
	cd tagtag_fw/mtl_linux && ./mtl_simu --dologtime --logs init,vm,simunet,simuleds,simuaudio

run: all
	./server/openab
