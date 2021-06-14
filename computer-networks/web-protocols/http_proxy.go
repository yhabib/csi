package main

import (
	"fmt"
	"http_proxy/socket"
)

func main() {
	data := make([]byte, 1024)

	s := socket.New()
	s.Bind(8080, [4]byte{0, 0, 0, 0})
	s.Listen(2048)
	s.Accept()
	for {
		size := s.Read(data)
		fmt.Printf("Size: %d\nData: %s", size, data)
	}
}
