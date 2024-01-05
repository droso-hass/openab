#!/bin/bash

ssh root@192.168.1.4 -oHostKeyAlgorithms=+ssh-rsa -p 22 'tcpdump -U -i wlan0 -w -' | sudo wireshark -k -S -i -
