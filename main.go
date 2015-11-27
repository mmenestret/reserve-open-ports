package main

import (
	"flag"
	"fmt"
	"net"
)

func accept() bool {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error!")
		return false
		// handle error
	}
	_, err2 := ln.Accept()
	if err2 != nil {
		fmt.Println("Error!")
		return false
	}
	return true
}

func bookPort(min int, max int) {
	//currentPort := min
	fmt.Println("Success!")
}

func main() {
	flag.Parse()
	bookPort(8080, 8090)
	fmt.Println("Finished!")
}
