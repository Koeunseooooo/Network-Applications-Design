#
# SimpleEchoTCPClient.py
#

from socket import *

# 'nsl2.cau.ac.kr'
serverName = 'localhost'
serverPort = 12000

clientSocket = socket(AF_INET, SOCK_STREAM)
clientSocket.connect((serverName, serverPort))

print("Client is running on port", clientSocket.getsockname()[1])

message = input('Input lowercase sentence: ')

clientSocket.send(message.encode())

modifiedMessage = clientSocket.recv(2048)
# 보통 클라에서의 수신은 별도의 thread를 구성하여 서버측으로부터 받은 메시지를 수신하게 된다.

print('Reply from server:', modifiedMessage.decode())

clientSocket.close()
