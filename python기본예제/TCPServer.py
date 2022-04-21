#
# SimpleEchoTCPServer.py
#

from socket import *

serverPort = 12000
serverSocket = socket(AF_INET, SOCK_STREAM)
# socket.AF_INET는 IP4인터넷을 사용한다는 뜻이고 데이터를 바이너리(byte 스트림)식으로 사용한다는 뜻입니다.

serverSocket.bind(('', serverPort))
serverSocket.listen(1)

print("Server is ready to receive on port", serverPort)

while True:
    (connectionSocket, clientAddress) = serverSocket.accept()
    print('Connection request from', clientAddress)
    message = connectionSocket.recv(2048)
    modifiedMessage = message.decode().upper()
    connectionSocket.send(modifiedMessage.encode())
    connectionSocket.close()
