# Network-Applications-and-Design_2022
> 중앙대학교 소프트웨어학부 4-1학기 네트워크응용설계 과목   
> 2022.03 ~ 06

## Table of Contents

- [HW2](##HW2)


## HW2
1. **Topic** | Socket Prgramming with Header Information  
2. **Submit** | class용 서버에 softcopy 명령어로 과제 제출  
3. **Files** | (in hw2 folder) TCPClient.go TCPServer.go UDPClient.go UDPServer.go  
4. **Language** | go
5. **Main Module** | net / bytes / signal / time 등
6. **Feature**
  - **Client (5 option)**
    - option 1) convert text to UPPER-case letters
    - option 2) ask the server what the IP address and port number of the client is
    - option 3) ask the server how many client requests(commands) it has served so far
    - option 4) ask the server program how long it has been running for since it started
    - option 5) exit client program  
    
  - **Server (reply with an response based on the option)**
    - option 1) convert text to UPPER-case letters
    - option 2) tell the client what the IP address and port number of the client is
    - option 3) tell the client how many client requests it has served so far
    - option 4) tell the client how long it (server program) has been running for (unit:seconds)
  - Etc.
    - Client and Server should exit the program gracefullly 
      - This means the program should not show ANY error messages!!
      - we use os/signal modules for implementation
    - Echo Client & Echo Server
      - Not use multi-thread (goroutine..)
      - we will use gorountine in hw3!


