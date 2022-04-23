/**
 * koeunseo
 * 20190143
 * MultiClientTCPServer.go
 **/

package main

// Import Required Modules
import (
	"bytes"
	// _ "fmt"
	f "fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var startTime time.Time

var id int

var numOfReq int
var numOfCon int

func main() {
	// we use designated personal port number as serverPort
	serverPort := "30143"

	listener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		f.Print("Bye bye~")
		os.Exit(0)
		return
	}
	defer listener.Close()

	numOfReq = 0
	numOfCon = 0

	// Measure time for command 4 & to print out the number of clients every 1 minute.
	startTime = time.Now()
	ticker := time.NewTicker(time.Second * 60)
	go func() {
		for i := range ticker.C {
			f.Printf("Number of connected client = %d\n\n", numOfCon)
			_ = i

		}
	}()

	// time.Sleep(time.Second * 3)
	// ticker.Stop()
	f.Printf("The server is ready to receive on port %s\n\n", serverPort)

	// to do : 클라이언트 배열을 생성해서 관리

	for {

		//connect client socket
		conn, err := listener.Accept()
		if nil != err {
			conn.Close()
			ticker.Stop()
			f.Print("Bye bye~")
			os.Exit(0)
		}

		// when user enters 'Ctrl-C'
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				if sig.String() == "interrupt" {
					conn.Close()
					ticker.Stop()
					f.Printf("Bye bye~")
					os.Exit(0)
				}
			}
		}()

		numOfReq += 1
		numOfCon += 1 // +1보다는 배열의 길이로 계산하는게 나을 것 같음
		id += 1

		remoteAddr := conn.RemoteAddr()
		// remoteAddrs := strings.Split(string(remoteAddr.String()), "]")
		// ip := remoteAddrs[0][1:]
		// port := remoteAddrs[1]
		f.Printf("Connection request from %s\n\n", remoteAddr)
		f.Printf("Client %d connected. Number of connected client = %d\n\n", id, numOfCon)
		go ConnHandler(conn, remoteAddr, id)

	}
}

func command_1(n int, conn net.Conn, buffer []byte) {
	f.Printf("Command 1\n\n")
	conn.Write(bytes.ToUpper(buffer[1:n]))
}

func command_2(remoteAddr net.Addr, conn net.Conn) {
	f.Printf("Command 2\n\n")
	conn.Write([]byte(remoteAddr.String()))
}

func command_3(count int, conn net.Conn) {
	f.Printf("Command 3\n\n")
	count_str := strconv.Itoa(count)
	conn.Write([]byte(count_str))
}

func command_4(startTime time.Time, conn net.Conn) {
	f.Printf("Command 4\n\n")
	// check eplasedTime
	elapsedTime := time.Since(startTime)
	conn.Write([]byte(elapsedTime.String()))
}

func ConnHandler(conn net.Conn, remoteAddr net.Addr, id int) {

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if nil != err {
			numOfCon -= 1
			f.Printf("Client %d disconnected. Number of connected clients = %d\n\n", id, numOfCon)
			conn.Close()
			break
		}

		// Check only the first letter of request: To distinguish it from the text of command1
		msg := string(buffer[:n][0])

		// receive the command and reply with an appropriate response based on command
		if msg == "1" {
			command_1(n, conn, buffer)

		} else if msg == "2" {
			command_2(remoteAddr, conn)

		} else if msg == "3" {
			command_3(numOfReq, conn)
		} else if msg == "4" {
			command_4(startTime, conn)
		}
	}
}
