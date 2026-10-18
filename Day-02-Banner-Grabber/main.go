//only works for talkative protocols as silent protocols such as HTTP requrie their corresponding data to be sent first.

package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func port_scan(thost *string, port *int) {

	fmt.Printf("Doing a 3 way handshake on traget host = %s & target port = %d\n", *thost, *port)

	conn, err := net.Dial("tcp", net.JoinHostPort(*thost, strconv.Itoa(*port)))

	buffer := make([]byte, 1024)

	if err == nil {
		fmt.Println("Port is Open")
		//perform banner grabbing now.
		number_of_bytes_read, reading_error := conn.Read(buffer)
		if reading_error == nil {
			//print the banner
			fmt.Printf("Number of bytes read by .Read(): %d\nThis is the buffer read aka Banner: %s\n", number_of_bytes_read, buffer)
			//uptill this point it works for Talkative Protocols but fails for Quiet Protocols like http where it is expected from the client to send some data
		} else {
			// an error occured while reading and print the error
			fmt.Println("An error occured with .Read(): ", reading_error)
		}
	} else {
		fmt.Println("Port is closed")
	}
}

func main() {

	target_host_ptr := flag.String("thost", "127.0.0.1", "Set Target Host:: -thost your.target.host.com")
	port_ptr := flag.Int("tport", 0, "Set Target Port to check:: -tport 80")

	flag.Parse()

	fmt.Println("Target Host = ", *target_host_ptr)
	fmt.Println("Target Port = ", *port_ptr)

	port_scan(*&target_host_ptr, *&port_ptr)
}

/*
run this program as follows
go run main.go
go run main.go -help
go run main.go -tport 90
go run main.go -tport 90 -thost 192.168.1.1
go run main.go -tport 22 -thost scanme.nmap.org
*/
