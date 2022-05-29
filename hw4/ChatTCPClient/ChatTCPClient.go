/**
 * koeunseo
 * 20190143
 * ChatTCPClient.go
 **/

package main

// Import Required Modules
import (
	"bufio"
	f "fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

func checkStringAlphabet(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}

const serverName string = "nsl2.cau.ac.kr"
const serverPort string = "30143"

const errHead string = "0"
const msgHead string = "1"
const listHead string = "2"
const dmHead string = "3"
const verHead string = "4"
const rttHead string = "5"

var startTime time.Time

func printExitMessage() {
	f.Println("gg~")
}
func main() {

	if len(os.Args) < 2 {
		f.Printf("You must enter a Nickname")
		os.Exit(0)
	} else if len(os.Args) > 2 {
		f.Printf("Nickname must be no spaces")
		os.Exit(0)
	}

	nickName := os.Args[1:2][0]

	if len(nickName) >= 32 {
		f.Printf("Nickname must be within 32 characters.")
		os.Exit(0)
	}

	if !checkStringAlphabet(nickName) {
		f.Printf("Nickname must be in English.")
		os.Exit(0)
	}

	conn, err := net.Dial("tcp", serverName+":"+serverPort)
	conn.Write([]byte(nickName))

	// if server socket not open, Client exits Program
	if err != nil {
		conn.Close()
		os.Exit(0)
	}
	once := new(sync.Once) // create once

	go func() {
		// when user enters 'Ctrl-C'
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			if sig.String() == "interrupt" {
				once.Do(printExitMessage) // gg~
				conn.Write([]byte(errHead))
				conn.Close()
				os.Exit(0)

			}
		}
	}()
	for {

		//read - goroutine
		go ReadHandler(conn, nickName)

		num_reader := bufio.NewReader(os.Stdin)
		text, err := num_reader.ReadString('\n')
		if nil != err {
			os.Exit(0)
		}

		text = strings.TrimSpace(text)

		n_text := strings.Split(text, " ")

		//write
		if strings.HasPrefix(n_text[0], `\`+"list") && len(n_text) == 1 { // \list command
			conn.Write([]byte(listHead))
		} else if strings.HasPrefix(n_text[0], `\`+"dm") && len(n_text) >= 3 {
			text = text[4:]
			conn.Write([]byte(dmHead + text))
		} else if strings.HasPrefix(n_text[0], `\`+"exit") && len(n_text) == 1 { // \exit command
			f.Printf("gg~")
			conn.Write([]byte(errHead))
			conn.Close()
			os.Exit(0)
		} else if strings.HasPrefix(n_text[0], `\`+"ver") && len(n_text) == 1 { // \ver command
			conn.Write([]byte(verHead))
		} else if strings.HasPrefix(n_text[0], `\`+"rtt") && len(n_text) == 1 { // \rtt command
			startTime = time.Now()
			conn.Write([]byte(rttHead))
		} else if strings.HasPrefix(n_text[0], `\`) {
			f.Print("invalid command\n")
		} else { // user can type a text message
			conn.Write([]byte(msgHead + text + "\n"))
		}
		f.Println()

	}

}

func ReadHandler(conn net.Conn, nickName string) {
	defer f.Println("world")
	buffer := make([]byte, 1024)

	//read
	for {

		n, err := conn.Read(buffer)
		if nil != err {
			conn.Close()
			os.Exit(0)
		}
		elapsedTime := time.Since(startTime)

		header := string(buffer[:n][0])
		body := string(buffer[:n][1:])

		if header == "0" {
			f.Print(body)
			conn.Close()
			os.Exit(0)
		} else if header == "1" { // welcome message
			body := string(buffer[:n][1:])
			f.Println(body)

		} else if header == "2" { // plain text message
			body := string(buffer[:n][1:])
			f.Println(body)
		} else if header == "3" { // response "\list"
			body := string(buffer[:n][1:])
			slice := strings.Split(body, "@")
			for i := 0; i < len(slice)-1; i++ {
				list := strings.TrimRight(slice[i], "@")
				f.Printf("[" + list + "]\n\n")
			}
		} else if header == "4" { //response "\dm"
			body := string(buffer[:n][1:])
			f.Printf(body + "\n\n")
		} else if header == "5" { //response "\ver"
			body := string(buffer[:n][1:])
			f.Printf("[version=" + body + "]\n\n")
		} else if header == "6" { //response "\rtt"
			rttTime := elapsedTime.Seconds() * 1000
			f.Printf("[RTT = %.4f]\n\n", rttTime)
		} else if header == "7" { //left_msg
			body := string(buffer[:n][1:])
			f.Println(body)
		}
	}

}
