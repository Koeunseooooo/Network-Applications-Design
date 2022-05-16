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
	"strconv"
	"strings"
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

const serverName string = "localhost"
const serverPort string = "30143"

const errHead string = "0"
const msgHead string = "1"
const listHead string = "2"
const dmHead string = "3"
const verHead string = "4"
const rttHead string = "5"

func main() {
	// buffer := make([]byte, 1024)

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
		f.Printf("Bye bye~")
		os.Exit(0)
	}
	defer conn.Close()

	// Check the address and port number of the client(self)
	// localAddr := conn.LocalAddr().(*net.TCPAddr)
	// f.Printf("The client is running on port %d\n", localAddr.Port)

	for {
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

		//read - goroutine
		go ReadHandler(conn, nickName)

		//write
		num_reader := bufio.NewReader(os.Stdin)
		text, err := num_reader.ReadString('\n')
		if nil != err {
			f.Printf("Bye bye~")
			os.Exit(0)
		}
		// conn.Write([]byte(msgHead + msg))

		if strings.HasPrefix(text, "\"list") && len(text) == 5 { // \list command
			f.Print("list를 보여줘")
		} else if strings.HasPrefix(text, "\"dm ") && len(strings.Split(text, " ")) == 3 { // \dm command
			f.Print("dm을 보낼거야")
		} else if strings.HasPrefix(text, "\"exit") && len(text) == 5 { // \exit command
			conn.Close()
			f.Printf("Bye bye~")
			os.Exit(0)
		} else if strings.HasPrefix(text, "\"ver") && len(text) == 4 { // \ver command
			f.Printf("ver을 보여줘")
		} else if strings.HasPrefix(text, "\"rtt") && len(text) == 4 { // \rtt command
			f.Printf("rtt를 보여줘")
		} else if strings.HasPrefix(text, "\"") {
			f.Print("invalid command")
		} else { // user can type a text message
			conn.Write([]byte(msgHead + text))
		}

	}

}

func ReadHandler(conn net.Conn, nickName string) {
	buffer := make([]byte, 1024)

	//read
	for {
		n, err := conn.Read(buffer)
		if nil != err {
			conn.Close()
			f.Print("Bye bye~")
			os.Exit(0)
		}

		header := string(buffer[:n][0])
		body := string(buffer[:n][1:])

		if header == "0" {
			f.Print(body)
			f.Print("오류래!")
			conn.Close()
			os.Exit(0)
		} else if header == "1" {
			welcome_msg := body[:7]
			clients_num := body[7:]
			f.Printf("[%s %s to CAU network class chat room at %s:%s.]\n[There are %s users connected.]\n", welcome_msg, nickName, serverName, serverPort, clients_num)

		} else if header == "2" {
			body := string(buffer[:n][1:])
			f.Print(body)
		} else {
			f.Printf("신종오류다..")
		}
	}

}

// command_1 ) print out the reply text string
func command_1(num string, conn net.Conn) {
	f.Print("Input sentence: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// measure RTT time
	startTime := time.Now()
	conn.Write([]byte(num + input))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if nil != err {
		f.Printf("Bye bye~")
		os.Exit(0)
	}
	elapsedTime := time.Since(startTime)

	response := string(buffer[:n])
	f.Printf("\nReply from server: %s", strings.Trim(response, "\n"))

	// convert rttTime format from seconds to milliseconds
	rttTime := elapsedTime.Seconds() * 1000
	f.Printf("\nRTT = %.4f\n\n", rttTime)
}

// command_2 ) print out the reply client's ip & port
func command_2(num string, conn net.Conn) {
	startTime := time.Now()
	conn.Write([]byte(num))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if nil != err {
		f.Printf("Bye bye~")
		os.Exit(0)
	}
	// f.Print(string(buffer[:n]))
	elapsedTime := time.Since(startTime)

	responses := strings.Split(string(buffer[:n]), "]")
	ip := responses[0][1:]
	port := responses[1][1:]

	f.Printf("\nReply from server: client IP = %s, port = %s \n", ip, port)
	rttTime := elapsedTime.Seconds() * 1000
	f.Printf("RTT = %.4f\n\n", rttTime)
}

// command_3 ) print out the reply server's request total count
func command_3(num string, conn net.Conn) {
	startTime := time.Now()
	conn.Write([]byte(num))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if nil != err {
		f.Printf("Bye bye~")
		os.Exit(0)
	}
	elapsedTime := time.Since(startTime)

	response := string(buffer[:n])
	f.Printf("\nReply from server: requests serverd = %s \n", response)
	rttTime := elapsedTime.Seconds() * 1000
	f.Printf("RTT = %.4f\n", rttTime)

}

// command_4 ) print out the reply server's runtime
func command_4(num string, conn net.Conn) {
	startTime := time.Now()
	conn.Write([]byte(num))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if nil != err {
		f.Printf("Bye bye~")
		os.Exit(0)
	}
	elapsedTime := time.Since(startTime)

	// Convert runtime in second format to HH:MM:SS format
	response := string(buffer[:n])
	response = strings.Split(response, ".")[0]
	ss := ""
	mm := ""
	hh := ""

	isminute := strings.Split(response, "m")
	if len(isminute) > 1 {
		mm = isminute[0]
		ss = isminute[1]
		ishour := strings.Split(response, "h")
		if len(ishour) > 2 {
			hh = ishour[0]
			mm = ishour[1]
			ss = ishour[2]
		}
	} else {
		ss = response
	}
	ss_int, _ := strconv.Atoi(ss)
	mm_int, _ := strconv.Atoi(mm)
	hh_int, _ := strconv.Atoi(hh)
	f.Printf("\nReply from server: runtime = %02d:%02d:%02d \n", hh_int, mm_int, ss_int)
	rttTime := elapsedTime.Seconds() * 1000
	f.Printf("RTT = %.4f\n\n", rttTime)

}
