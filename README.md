# Network-Applications-and-Design_2022

> 중앙대학교 소프트웨어학부 4-1학기 네트워크응용설계   
> 2022.03 ~ 06

## Table of Contents

-   [HW2](##HW2)
-   [HW3](##HW3)
-   [HW4](##HW4)
-   [HW5](##HW5)

## HW2

1. **Topic** | Socket Prgramming with Header Information
2. **duration** | ~2022.04.10
3. **Submit** | class용 서버에 softcopy 명령어로 과제 제출
4. **Files** | (in hw2 folder) TCPClient.go TCPServer.go UDPClient.go UDPServer.go
5. **Language** | go
6. **Main Module** | net / bytes / signal / time 등
7. **Features I Implemented 😀**

-   **Client (5 option)**
    -   option 1) convert text to UPPER-case letters
    -   option 2) ask the server what the IP address and port number of the client is
    -   option 3) ask the server how many client requests(commands) it has served so far
    -   option 4) ask the server program how long it has been running for since it started
    -   option 5) exit client program
-   **Server (reply with an response based on the option)**
    -   option 1) convert text to UPPER-case letters
    -   option 2) tell the client what the IP address and port number of the client is
    -   option 3) tell the client how many client requests it has served so far
    -   option 4) tell the client how long it (server program) has been running for (unit:seconds)
-   **Etc.**
    -   Client and Server should exit the program gracefullly
        -   This means the program should not show ANY error messages!!
        -   I used os/signal modules for implementation
    -   Echo Client & Echo Server
        -   Not use multi-thread (goroutine..)
        -   I will use gorountine in hw3 to support multi-client

## HW3

1. **Topic** | Socket Prgramming with Multiple Clients
2. **duration** | ~2022.05.03
3. **Submit** | class용 서버에 softcopy 명령어로 과제 제출
4. **Files** | (in hw3 folder) TCPClient.go MultiClientTCPServer.go
5. **Language** | go, C
6. **Main Func** | goroutine, NewTicker
7. **Features I Implemented 🥰**

-   **Server (Modified only server!)**
    -   support multiple clients using Go routine
    -   assign each client a unique ID such as {client 1, client 2, client 3 , ...} when they connect
        -   This ID should not change even if a client disconnects
    -   Whenever a new client connects, or an existing client disconnects, print out the client ID and the number of clients on the server as well as ALL connected clients
    -   print out the number of clients ever 1 MINUTE (I use NewTicker to output repeatedly)


## HW4
(작성 예정)

## HW5
(작성 예정)
