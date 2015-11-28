package main

import (
	"flag"
	"net"
	"strconv"
	"strings"
)

func open_port(port int) bool {
	address := strings.Join([]string{":", strconv.Itoa(port)}, "")
	list, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}

	con, acceptErr := list.Accept()
	buff := make([]byte, 1)
	if acceptErr != nil {
		return false
	}

	nb, readErr := con.Read(buff)
	if readErr != nil {
		return false
	}
	return (nb > 0)
}
func check_port_open(port int, done chan bool) {
	done <- open_port(port)
}
func main() {
	flag.Parse()
	first_port := 9000
	port_range := 10
	done := make(chan bool)
	for i := 0; i <= port_range; i++ {
		go check_port_open(first_port+i, done)
	}
	c := <-done
	print(c)
}
