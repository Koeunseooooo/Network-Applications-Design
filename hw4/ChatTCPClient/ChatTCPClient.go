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

func main() {
	serverName := "localhost"
	serverPort := "30143"

	// connect server socket
	conn, err := net.Dial("tcp", serverName+":"+serverPort)

	// if server socket not open, Client exits Program
	if err != nil {
		f.Printf("Bye bye~")
		os.Exit(0)
	}
	defer conn.Close()

	// Check the address and port number of the client(self)
	localAddr := conn.LocalAddr().(*net.TCPAddr)
	f.Printf("The client is running on port %d\n", localAddr.Port)

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

		// user take input for the selection
		var num string
		menu()
		f.Print("Input option :")

		num_reader := bufio.NewReader(os.Stdin)
		num, err = num_reader.ReadString('\n')
		if nil != err {
			f.Printf("Bye bye~")
			os.Exit(0)
		}

		// print out the returened response based on the command
		if num == "1\n" {
			command_1(num, conn)
		} else if num == "2\n" {
			command_2(num, conn)
		} else if num == "3\n" {
			command_3(num, conn)
		} else if num == "4\n" {
			command_4(num, conn)
		} else if num == "5\n" {
			conn.Close()
			f.Print("Bye bye~")
			os.Exit(0)
		} else {
		}

		f.Print("\n")
	}

}

func menu() {
	f.Printf("<Menu>\n")
	f.Printf("1) convert text to UPPER-case\n")
	f.Printf("2) get my IP address and port number\n")
	f.Printf("3) get server request count\n")
	f.Printf("4) get server running time\n")
	f.Printf("5) exit\n")
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
