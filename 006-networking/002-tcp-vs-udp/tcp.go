package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	// Start listening for TCP connections on the given address
	stream, err := net.Listen("tcp", ":8080")
	check(err)
	defer stream.Close()
	for {
		// Accept new connections
		con, err := stream.Accept()
		check(err)
		// Handle new connections in a Goroutine
		go handle(con)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	for {
		// message := make([]byte, 20)
		// len, err := c.Read(message)
		// check(err)
		// data := strings.TrimSpace(string(message[:len]))
		// fmt.Printf("received: %s from %s\n", data, c.RemoteAddr().String())

		// Read from the connection until a new line is send
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the data read from the connection to the terminal
		fmt.Print("> ", string(data))

		str := fmt.Sprintf("Hello TCP client from %s", c.RemoteAddr().String())

		// Send message to the client
		c.Write([]byte(str))
	}
}
