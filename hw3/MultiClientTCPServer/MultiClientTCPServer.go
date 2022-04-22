/**
 * koeunseo
 * 20190143
 * TCPServer.go
 **/

package main

// Import Required Modules
import (
	"bytes"
	f "fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var startTime time.Time
var count int

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
	count = 0

	// Measure time for command 4
	startTime = time.Now()
	f.Printf("The server is ready to receive on port %s\n", serverPort)

	for {
		//connect client socket
		conn, err := listener.Accept()
		count += 1
		if nil != err {
			f.Print("Bye bye~")
			os.Exit(0)
			return
		}
		defer conn.Close()

		remoteAddr := conn.RemoteAddr()
		// remoteAddrs := strings.Split(string(remoteAddr.String()), "]")
		// ip := remoteAddrs[0][1:]
		// port := remoteAddrs[1]
		f.Printf("\nConnection request from %s\n", remoteAddr)
		go ConnHandler(conn, remoteAddr)

		// when user enters 'Ctrl-C'
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				if sig.String() == "interrupt" {
					conn.Close()
					f.Printf("Bye bye~")
					os.Exit(0)
				}
			}
		}()

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

func ConnHandler(conn net.Conn, remoteAddr net.Addr) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if nil != err {
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
			command_3(count, conn)
		} else if msg == "4" {
			command_4(startTime, conn)
		}
	}
}
