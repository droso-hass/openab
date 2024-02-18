#!/bin/bash

cp main.mtl _main.mtl
sed -i '1s/^/#define SIMU\n/' _main.mtl
./include.py ./_main.mtl ./nominal.mtl
rm _main.mtl
cp ./nominal.mtl ../mtl_linux
cd ../mtl_linux
./mtl_simu --dologtime --logs init,vm,simunet,simuleds,simuaudio