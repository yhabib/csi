package main

import "http_proxy/socket"

func main() {
	s := socket.New()
	s.Bind(8080, [4]byte{0, 0, 0, 0})
	s.Listen(2048)
	s.Accept()
}
