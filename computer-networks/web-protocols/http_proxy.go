package main

import (
	"http_proxy/socket"
)

func main() {
	sc := socket.New()
	sp := socket.New()
	defer sc.Close()
	defer sp.Close()

	sc.Bind(8080, [4]byte{0, 0, 0, 0})
	sc.Listen(2048)
	sc.Accept()

	sp.Connect(9000, [4]byte{0, 0, 0, 0})

	for {
		serverBuffer := make([]byte, 1024)
		clientBuffer := make([]byte, 1024)
		size := sc.Receive(serverBuffer)
		if size == 0 {
			break
		}
		sp.Send(serverBuffer)
		sp.Receive(clientBuffer)
		sc.Send(clientBuffer)
	}
}
