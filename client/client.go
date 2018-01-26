package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		fmt.Println("error in dialing to gossip server", err)
		return
	}

	fmt.Println("Start gossiping ;)")


	for {
		reader := bufio.NewReader(os.Stdin)

		gossip, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading gossip:", err)
			return
		}

		fmt.Fprintf(conn, gossip)
	}
}
