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
	"strings"
)

const errHead string = "0"
const welHead string = "1"
const msgHead string = "2"
const listHead string = "3"
const dmHead string = "4"
const verHead string = "5"
const rttHead string = "6"
const leftHead string = "7"

const ignore_case string = "i hate professor"

const version string = "2.3.0"

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

	for {
		// when user enters 'Ctrl-C'
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				if sig.String() == "interrupt" {
					removeAllClients()
					os.Exit(0)
				}
			}
		}()

		//connect client socket
		conn, err := listener.Accept()
		if nil != err {
			conn.Close()
		}

		buffer := make([]byte, 1024)
		if len(clients) >= 8 {
			msg := "chatting room full. cannot connect"
			conn.Write([]byte(errHead + msg))
			conn.Close()
		} else {
			n, err := conn.Read(buffer)
			if nil != err {
				conn.Close()
			}

			nickName := string(buffer[:n])
			isDuplicate := false
			for i := 0; i < len(clients); i++ {
				if nickName == clients[i].nickName {
					msg := "that nickname is already used by another user.cannot connect."
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
	localAddr := client.conn.LocalAddr().(*net.TCPAddr)
	welcome_msg := "welcome " + client.nickName + " to CAU network class chat room at " + localAddr.String() + ".]\n[There are " + strconv.Itoa(len(clients)) + "users connected.]\n\n"
	// f.Printf("[%s %s to CAU network class chat room at %s:%s.]\n[There are %s users connected.]\n\n", welcome_msg, nickName, serverName, serverPort, clients_num)
	client.conn.Write([]byte(welHead + welcome_msg))

	f.Printf("[%s joined from %s. There are %d users connected.]\n\n", client.nickName, remoteAddr.String(), len(clients))

	for {

		n, err := client.conn.Read(buffer)
		if nil != err {
			removeClient(client)
			client.conn.Close()
			return
		}

		header := string(buffer[:n][0])

		if header == "0" { //errHead
			client.conn.Close()
			removeClient(client)
			f.Printf("[%s left. There are %d users now.]\n\n", client.nickName, len(clients))
			return

		}

		if header == "1" { //plain msgHead
			body := string(buffer[:n][1:])
			user := client.nickName + "> "
			msg := user + body

			// check ignore case
			checkmsg := strings.ToLower(msg)
			if strings.Contains(checkmsg, ignore_case) {
				exit_msg := "[You left this chat room.]"
				removeClient(client)
				client.conn.Write([]byte(leftHead + exit_msg))
				client.conn.Close()

				ignore_msg := "[" + client.nickName + " is disconnected. There are " + strconv.Itoa(len(clients)) + " users in the chat room]\n\n"
				for i := 0; i < len(clients); i++ {
					clients[i].conn.Write([]byte(msgHead + msg + "\n" + ignore_msg))
				}
				f.Printf(ignore_msg)
			} else {
				for i := 0; i < len(clients); i++ {

					if client.id != clients[i].id {
						clients[i].conn.Write([]byte(msgHead + msg))
					}
				}
			}

		} else if header == "2" { //listHead
			msg := ""
			for i := 0; i < len(clients); i++ {
				slice := strings.Split(clients[i].conn.RemoteAddr().String(), ":")
				ip := slice[0]
				port := slice[1]
				msg += clients[i].nickName + "," + ip + "," + port
				msg += "@"
			}
			client.conn.Write([]byte(listHead + msg))

		} else if header == "3" { //dmHead
			body := string(buffer[:n][1:])
			idx := strings.IndexAny(body, " ")
			nickName := body[:idx]
			msg := body[idx:]
			for i := 0; i < len(clients); i++ {
				if clients[i].nickName == nickName {
					to_msg := "from: " + client.nickName + "> " + msg
					clients[i].conn.Write([]byte(dmHead + to_msg))
				}
			}

		} else if header == "4" { //verHead
			client.conn.Write([]byte(verHead + version))
		} else if header == "5" { //rttHead
			client.conn.Write([]byte(rttHead))
		}
	}
}
