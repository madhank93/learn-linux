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
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	})
	check(err)

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 20)
		len, remote, err := conn.ReadFromUDP(message[:])
		check(err)
		data := strings.TrimSpace(string(message[:len]))
		fmt.Printf("received: %s from %s\n", data, remote)
	}
}
