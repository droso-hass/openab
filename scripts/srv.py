import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind(("192.168.1.102", 4000))
#sock.bind(("127.0.0.1", 4000))

while True:
    data, addr = sock.recvfrom(1024)
    print(addr)
    print(data)
