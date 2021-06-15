package main

import (
	"fmt"
	"http_proxy/socket"
)

const PORT = 1234

func main() {
	sc := socket.New()
	sp := socket.New()
	defer sc.Close()

	sc.Bind(8080, [4]byte{0, 0, 0, 0})
	sc.Listen(2048)
	sc.Accept()

	sp.Connect(8081, [4]byte{0, 0, 0, 0})

	for {
		data := make([]byte, 1024)
		size := sc.Receive(data)
		fmt.Printf("%s\n", data)
		if size == 0 {
			break
		}
		sp.Send(data)
	}
}
