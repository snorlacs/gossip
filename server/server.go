package main

import (
	"net"
	"fmt"
	"bufio"
	"io"
)

const PORT = ":8080"

func main() {

	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println("Error occured during creating server", err)
	}

	fmt.Println("starting gossip server in port", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting connection", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Gossip is starting to handle connections")

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		gossip, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client went offline")
			} else {
				fmt.Println("Error reading from connection", err)
			}
			return
		}

		fmt.Println("client:", string(gossip))
	}
}
