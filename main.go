package main

import (
	"net"
	"log"
	"fmt"
	"syscall"
)
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

	_, errAccepting := listener.Accept()
	if errAccepting != nil {
		log.Println("Couldn't accept on", fullAdd)
	}
}

func PingSocket(add string, port int){
	fullAdd := fmt.Sprintf("%s:%d", add, port)
	conn, err := net.Dial("tcp", fullAdd)
	if err != nil {
		log.Println("Port 9000 is NOT avaiable")
	} else {
		log.Println("Port 9000 is avaiable")
		defer conn.Close()
	}
}
func main() {
	isSocketOpened := make(chan bool)
	port := 9000
	go listenAndAccept(port, isSocketOpened)
	<- isSocketOpened
	PingSocket("localhost", port)
}
