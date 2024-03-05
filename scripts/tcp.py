import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.bind(("", 9000))
sock.listen(1)

client, addr = sock.accept()

while True:
    data, addr = client.recv(1024)
    print(data)
