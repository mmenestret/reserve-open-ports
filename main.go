package main

import (
	"flag"
	"fmt"
	"net"
)

func open_port() bool {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error!")
		return false
		// handle error
	}
	_, err1 := net.ListenTCP("tcp", tcpAddr)
	if err1 != nil {
		fmt.Println("Error!")
		return false
		// handle error
	}
	return true
}

func bookPort(min int, max int) bool {
	//currentPort := min
	//has_port := false
	//for !has_port {
	return open_port()
	//}
}

func main() {
	flag.Parse()

	port_is_open := bookPort(8080, 8090)
	fmt.Println("Finished!")
	fmt.Println(port_is_open)
}
