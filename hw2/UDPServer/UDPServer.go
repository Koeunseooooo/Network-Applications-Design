/**
 * UDPServer.go
 **/

package main

import (
	"bytes"
	"fmt"
	f "fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	serverPort := "30143"

	pconn, err := net.ListenPacket("udp", ":"+serverPort)
	if err != nil {
		f.Print("Bye bye~")
		os.Exit(0)
		return
	}

	// Measure time for command 4
	startTime := time.Now()
	buffer := make([]byte, 1024)

	for {
		f.Printf("The server is ready to receive on port %s\n", serverPort)

		// when user enters 'Ctrl-C'
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				if sig.String() == "interrupt" {
					f.Printf("Bye bye~")
					os.Exit(0)
				}
			}
		}()

		count := 0

		for {
			n, r_addr, _ := pconn.ReadFrom(buffer)

			fmt.Printf("\nConnection request from %s\n", r_addr.String())

			// Check requests num for command 3
			count += 1

			// Check only the first letter of request: To distinguish it from the text of command1
			msg := string(buffer[:n][0])

			// receive the command and reply with an appropriate response based on command
			if msg == "1" {
				command_1(r_addr, n, pconn, buffer)

			} else if msg == "2" {
				command_2(r_addr, pconn)

			} else if msg == "3" {
				command_3(r_addr, count, pconn)
			} else if msg == "4" {
				command_4(r_addr, startTime, pconn)
			}

		}

	}

}

func command_1(remoteAddr net.Addr, n int, conn net.PacketConn, buffer []byte) {
	f.Printf("Command 1\n\n")
	conn.WriteTo(bytes.ToUpper(buffer[1:n]), remoteAddr)
}

func command_2(remoteAddr net.Addr, conn net.PacketConn) {
	f.Printf("Command 2\n\n")
	conn.WriteTo([]byte(remoteAddr.String()), remoteAddr)
}

func command_3(remoteAddr net.Addr, count int, conn net.PacketConn) {
	f.Printf("Command 3\n\n")
	count_str := strconv.Itoa(count)
	conn.WriteTo([]byte(count_str), remoteAddr)
}

func command_4(remoteAddr net.Addr, startTime time.Time, conn net.PacketConn) {
	f.Printf("Command 4\n\n")
	// check eplasedTime
	elapsedTime := time.Since(startTime)
	conn.WriteTo([]byte(elapsedTime.String()), remoteAddr)
}
