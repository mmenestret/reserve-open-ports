package main

import (
	"flag"
	"fmt"
	"net"
)

func bookPort(min int, max int) {
	//currentPort := min
	_, err := net.Listen("tcp", ":8080")
	for err != nil {
		if err != nil {
			fmt.Println("Error!")
			// handle error
		}

	}
	fmt.Println("Success!")
}

func main() {
	flag.Parse()
	bookPort(8080, 8090)
	fmt.Println("Finished!")
}
