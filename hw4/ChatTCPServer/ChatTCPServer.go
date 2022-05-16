/**
 * koeunseo
 * 20190143
 * ChatTCPServer.go
 **/

package main

// Import Required Modules
import (
	f "fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var startTime time.Time

const errHead string = "0"
const welHead string = "1"
const msgHead string = "2"
const listHead string = "3"
const dmHead string = "4"
const verHead string = "5"
const rttHead string = "6"

var clients []client
var id int

type client struct {
	id       int
	nickName string
	conn     net.Conn
}

func removeClient(client client) {
	for i := 0; i < len(clients); i++ {
		if client.id == clients[i].id {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
		}
	}
}

func removeAllClients() {
	for i := 0; i < len(clients); i++ {
		clients[i].conn.Close()
		clients = nil
	}
}

func main() {
	buffer := make([]byte, 1024)
	// we use designated personal port number as serverPort
	serverPort := "30143"

	listener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		f.Print("Bye bye~")
		os.Exit(0)
		return
	}
	defer listener.Close()

	id = 0

	startTime = time.Now()

	f.Printf("The server is ready to receive on port %s\n\n", serverPort)

	for {
		// when user enters 'Ctrl-C'
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				if sig.String() == "interrupt" {
					removeAllClients()
					f.Printf("Bye bye~")
					os.Exit(0)
				}
			}
		}()
		//connect client socket
		conn, err := listener.Accept()
		if nil != err {
			conn.Close()
			f.Print("Bye bye~1")
			os.Exit(0)
		}

		if len(clients) >= 8 { // 어짜피 바로 disconnect할 꺼니까 스레드로 넘길필요 없지 않을까?
			msg := "chatting room full. cannot connect"
			conn.Write([]byte(errHead + msg))
			conn.Close()
		} else {
			n, err := conn.Read(buffer)
			if nil != err {
				conn.Close()
				f.Print("Bye bye~2")
				os.Exit(0)
			}

			nickName := string(buffer[:n])
			isDuplicate := false
			for i := 0; i < len(clients); i++ {
				if nickName == clients[i].nickName {
					msg := "that nickname is already used by another user.cannot connect."
					f.Print("1")
					conn.Write([]byte(errHead + msg))
					conn.Close()
					isDuplicate = true
				}
			}

			if isDuplicate {
				continue
			}

			id += 1
			var client = client{id, nickName, conn}
			clients = append(clients, client)

			remoteAddr := conn.RemoteAddr()
			go ConnHandler(client, remoteAddr, id)
		}

	}
}

func ConnHandler(client client, remoteAddr net.Addr, id int) {
	buffer := make([]byte, 1024)

	welcome_msg := "welcome" + strconv.Itoa(len(clients))
	client.conn.Write([]byte(welHead + welcome_msg))

	f.Printf("%s joined from %s. There are %d users connected.\n", client.nickName, remoteAddr.String(), len(clients))

	for {

		n, err := client.conn.Read(buffer)
		if nil != err {
			f.Printf("Client %d disconnected.\n\n", client.id)
			removeClient(client)
			client.conn.Close()
			break
		}
		msg := string(buffer[:n][0])
		f.Print(msg)
		if msg == "1" {
			f.Print("...먀!")
		}
	}
}

// func command_1(n int, conn net.Conn, buffer []byte) {
// 	numOfReq += 1
// 	f.Printf("Command 1\n\n")
// 	conn.Write(bytes.ToUpper(buffer[1:n]))
// }

// func command_2(remoteAddr net.Addr, conn net.Conn) {
// 	numOfReq += 1
// 	f.Printf("Command 2\n\n")
// 	conn.Write([]byte(remoteAddr.String()))
// }

// func command_3(conn net.Conn) {
// 	numOfReq += 1
// 	f.Printf("Command 3\n\n")
// 	count_str := strconv.Itoa(numOfReq)
// 	conn.Write([]byte(count_str))
// }

// func command_4(startTime time.Time, conn net.Conn) {
// 	numOfReq += 1
// 	f.Printf("Command 4\n\n")
// 	// check eplasedTime
// 	elapsedTime := time.Since(startTime)
// 	conn.Write([]byte(elapsedTime.String()))
// }

// func main() {
//     var Slice1 = []int{1, 2, 3, 4, 5}
//     fmt.Printf("slice1: %v\n", Slice1)

//     Slice2 := remove(Slice1, 2)
//     fmt.Printf("slice2: %v\n", Slice2)
// }

// 출력
// slice1: [1 2 3 4 5]
// slice2: [1 2 5 4]

// 스레드에 넣었던 코드

// buffer := make([]byte, 1024)
// for {
// n, err := conn.Read(buffer)
// if nil != err {
// 	numOfCon -= 1
// 	f.Printf("Client %d disconnected. Number of connected clients = %d\n\n", id, numOfCon)
// 	conn.Close()
// 	break
// }

// Check only the first letter of request: To distinguish it from the text of command1
// 이 0번째가 이제
// msg := string(buffer[:n][0])

// receive the command and reply with an appropriate response based on command
// if msg == "1" {
// 	command_1(n, conn, buffer)

// } else if msg == "2" {
// 	command_2(remoteAddr, conn)

// } else if msg == "3" {
// 	command_3(conn)
// } else if msg == "4" {
// 	command_4(startTime, conn)
// }
// }
