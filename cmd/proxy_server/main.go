package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		origin, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer origin.Close()

		go reverseProxy(origin)
	}
}

func reverseProxy(origin net.Conn) {
	destiny, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destiny.Close()

	go io.Copy(destiny, origin)
	io.Copy(origin, destiny)
}
