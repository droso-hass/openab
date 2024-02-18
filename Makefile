
all: v2-firmware copy server

server: server/internal/* server/main.go server/go.mod server/go.sum
	cd server && go build

v2-metalc: tagtag_fw/mtl_linux/*
	cd tagtag_fw/mtl_linux && make comp

v2-firmware: v2-metalc tagtag_fw/src*
	cd tagtag_fw/src && ./include.py ./main.mtl ./nominal.mtl
	cd tagtag_fw/mtl_linux && ./mtl_compiler -s ../src/nominal.mtl ../nominal.bin

copy:
	cp tagtag_fw/nominal.bin server/static
