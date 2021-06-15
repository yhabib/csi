package main

import (
	"http_proxy/socket"
)

func main() {
	serverSocket := socket.New(8080, [4]byte{0, 0, 0, 0})
	clientSocket := socket.New(9000, [4]byte{0, 0, 0, 0})

	defer serverSocket.Close()
	defer clientSocket.Close()

	serverSocket.Bind()
	serverSocket.Listen(2048)
	clientSocket.Connect()

	for {
		connectionSocket := serverSocket.Accept()
		serverBuffer := make([]byte, 1024)
		clientBuffer := make([]byte, 1024)

		connectionSocket.Receive(serverBuffer)
		// if size == 0 {
		// 	break
		// }
		clientSocket.Send(serverBuffer)
		clientSocket.Receive(clientBuffer)
		connectionSocket.Send(clientBuffer)
	}
}
