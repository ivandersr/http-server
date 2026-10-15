package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ivandersr/http-server/internal/request"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Request Line:\n")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", r.RequestLine.HttpVersion)
	}
}
