package main

import (
	"fmt"
	"http_proxy/socket"
)

func main() {
	s := socket.New()
	defer s.Close()

	s.Bind(8080, [4]byte{0, 0, 0, 0})
	s.Listen(2048)
	s.Accept()
	for {
		data := make([]byte, 1024)
		size := s.Receive(data)
		fmt.Printf("%s\n", data)
		if size == 0 {
			break
		}
		s.Send(data)
	}
}
