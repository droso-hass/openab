
all: v2-firmware server

server: internal/* main.go go.mod go.sum
	go build

v2-metalc: firmware/v2/mtl_linux/*
	cd firmware/v2/mtl_linux && make comp

v2-firmware: v2-metalc firmware/v2/src/*
	cd firmware/v2/mtl_linux && ./mtl_compiler -s ../src/nominal.mtl ../nominal.bin
