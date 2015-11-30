package main

import (
	"flag"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"syscall"
	"os"
)

func port_is_taken(port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err == nil {
		conn.Close()
	}
	return err == nil
}

func bookPorts(min, max, ports int) []int {
	currentPort := min - 1
	freePorts := []int{}
	for len(freePorts) < ports {
		if currentPort == max {
			currentPort = min
		} else {
			currentPort++
		}
		if !port_is_taken(currentPort) {
			freePorts = append(freePorts, currentPort)
		}
	}
	return freePorts
}

func main() {
	var initialPort = flag.Int("initialPort", 8080, "first port in the range")
	var lastPort = flag.Int("lastPort", 8090, "last port in the range")
	var numberOfPorts = flag.Int("numberOfPorts", 3, "number of ports to book")
	var scriptToExec = flag.String("scriptToExec", "", "path to the script to exec")
	flag.Parse()
	freePorts := bookPorts(*initialPort, *lastPort, *numberOfPorts)
	freePortsAsStrings := make([]string, 3)
	for _,p := range(freePorts){
		freePortsAsStrings = append(freePortsAsStrings, strconv.Itoa(p))
	}
	cmd := exec.Command(*scriptToExec, freePortsAsStrings...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}
	syscall.Exit(0)
}
