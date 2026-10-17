package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func port_scan(thost *string, port *int) {

	fmt.Printf("Doing a 3 way handshake on traget host = %s & target port = %d\n", *thost, *port)

	_, err := net.Dial("tcp", net.JoinHostPort(*thost, strconv.Itoa(*port)))

	if err == nil {
		fmt.Println("Port is Open")
	} else {
		fmt.Println("Port is closed")
	}
}

func main() {

	target_host_ptr := flag.String("thost", "127.0.0.1", "Set Target Host:: -thost your.target.host.com")
	timeout_ptr := flag.Int("timeout", 3, "Set the timeout in seconds:: -timeout 1")
	port_ptr := flag.Int("tport", 0, "Set Target Port to check:: -tport 80")

	flag.Parse()

	fmt.Println("Target Host = ", *target_host_ptr)
	fmt.Println("Target Port = ", *port_ptr)
	fmt.Println("Timeout set = ", *timeout_ptr)

	port_scan(*&target_host_ptr, *&port_ptr)
}

/*
run this program as follows
go run main.go
go run main.go --help
go run main.go -timeout 5
go run main.go -timeout 5 -tport 90
go run main.go -timeout 5 -tport 90 -thost 192.168.1.1
go run main.go -timeout 5 -tport 90 -thost scanme.nmap.org
*/
