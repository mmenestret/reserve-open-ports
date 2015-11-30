package main

import (
	"net"
	"log"
	"fmt"
	"syscall"
)

type addr string
func (l addr) Network() string{
	return "tcp"
}
func (l addr) String() string{
	return string(l)
}
func listenAndAccept(port int, isSocketOpened chan bool) {
	fullAdd := fmt.Sprintf("%s:%d", "localhost", port)

	listener, errListening := net.Listen("tcp", fullAdd)
	if errListening != nil {
		log.Println("Couldn't listen on", fullAdd)
		syscall.Exit(1)
	}
	log.Printf("Listening & accepting connections on %s...\n", fullAdd)

	// Raising the caller that it can now Dial on that address
	isSocketOpened <- true

	conn, errAccepting := listener.Accept()
	if errAccepting != nil {
		log.Println("Couldn't accept on", fullAdd)
	}
	log.Println("Closing after having received a connection from", conn.RemoteAddr().String())
	conn.Close()
}

func PingSocket(add string, port int){
	fullAdd := fmt.Sprintf("%s:%d", add, port)
	d := net.Dialer{LocalAddr: addr(fullAdd)}
	conn, err := d.Dial("tcp", fullAdd)
	fmt.Println(d.LocalAddr.String())
	if err != nil {
		log.Printf("Port %d is NOT avaiable\n", port)
	} else {
		log.Printf("Port %d is avaiable\n", port)
		defer conn.Close()
	}
}
func main() {
	isSocketOpened := make(chan bool)
	port := 9001
	listenAndAccept(port, isSocketOpened)
	<- isSocketOpened
	PingSocket("127.0.0.1", port)
}
