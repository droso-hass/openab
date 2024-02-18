#!/bin/bash

ssh root@192.168.1.34 -oHostKeyAlgorithms=+ssh-rsa -p 22 'tcpdump -U -i wlan1 -w -' | wireshark -k -S -i -
