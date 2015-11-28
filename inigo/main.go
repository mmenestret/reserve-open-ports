package main

import (
	"flag"
	"fmt"
	"net"
)

func port_is_taken(port int) bool {
	fmt.Printf("dialing for port %d\n", port)
	_, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	return err == nil
}

func bookPorts(min, max, ports int) []int {
	currentPort := min - 1
	taken := []int{}
	for len(taken) < ports {
		if currentPort == max {
			currentPort = min
		} else {
			currentPort++
		}
		if !port_is_taken(currentPort) {
			fmt.Printf("blocking for port %d\n", currentPort)
			go net.Listen("tcp", fmt.Sprintf("localhost:%d", currentPort))
			taken = append(taken, currentPort)
		}
	}
	return taken
}

func main() {
	var initialPort = flag.Int("initialPort", 8080, "first port in the range")
	var lastPort = flag.Int("lastPort", 8090, "last port in the range")
	var numberOfPorts = flag.Int("numberOfPorts", 3, "number of ports to book")
	flag.Parse()
	fmt.Println("Port Free: %s!", bookPorts(*initialPort, *lastPort, *numberOfPorts))
}
