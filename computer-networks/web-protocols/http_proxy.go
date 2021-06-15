package main

import (
	"http_proxy/parser"
	"http_proxy/socket"
	"log"
)

var (
	PROXY_PORT  = 8080
	SERVER_PORT = 9000
	ADDRESS     = [4]byte{0, 0, 0, 0}
	BUFFER_SIZE = 1024
)

func main() {
	serverSocket := socket.New(PROXY_PORT, ADDRESS)
	// size := 0
	// cache := make(map[string][]byte)

	defer serverSocket.Close()
	serverSocket.Bind()
	serverSocket.Listen(BUFFER_SIZE)

	for {
		log.Printf("Proxy listening on port: %d...", PROXY_PORT)
		serverBuffer := make([]byte, BUFFER_SIZE)
		clientBuffer := make([]byte, BUFFER_SIZE)

		connectionSocket := serverSocket.Accept()
		connectionSocket.Receive(serverBuffer)
		req := parser.HttpRequest(serverBuffer)
		log.Println(req.String())

		// HTTP 1.0 -> Connect proxy to Server for every request
		clientSocket := socket.New(SERVER_PORT, ADDRESS)
		clientSocket.Connect()
		clientSocket.Send(serverBuffer)
		clientSocket.Receive(clientBuffer)
		res := parser.HttpResponse(clientBuffer)
		log.Println(res.String())

		connectionSocket.Send(clientBuffer)

		// Close connections
		connectionSocket.Close()
		clientSocket.Close()
	}
}
