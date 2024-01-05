#!/bin/bash

./include.py ./main.mtl ./nominal.mtl
cp ./nominal.mtl ../../mtl_linux
cd ../../mtl_linux
./mtl_simu --dologtime --logs init,vm,simunet,simuleds,simuaudio