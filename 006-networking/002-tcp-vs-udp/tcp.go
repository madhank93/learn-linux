package main

import (
	"fmt"
	"net"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	stream, err := net.Listen("tcp", ":8080")
	check(err)
	defer stream.Close()
	for {
		con, err := stream.Accept()
		check(err)
		go handle(con)
	}
}

func handle(c net.Conn) {
	for {
		message := make([]byte, 20)
		len, err := c.Read(message)
		check(err)
		data := strings.TrimSpace(string(message[:len]))
		fmt.Printf("received: %s from %s\n", data, c.RemoteAddr().String())
	}
}
