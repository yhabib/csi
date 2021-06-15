package main

import (
	"fmt"
	"http_proxy/parser"
	"http_proxy/socket"
	"log"
)

var (
	PROXY_PORT  = 8080
	SERVER_PORT = 9000
	ADDRESS     = [4]byte{0, 0, 0, 0}
	BUFFER_SIZE = 1024
	PROXY_PATH  = "/proxy"
)

func main() {
	serverSocket := socket.New(PROXY_PORT, ADDRESS)
	cache := make(map[string][]byte)
	size := 0
	isRequestCacheable := false

	defer serverSocket.Close()
	serverSocket.Bind()
	serverSocket.Listen(BUFFER_SIZE)

	for {
		log.Printf("Proxy listening on port: %d...", PROXY_PORT)
		serverBuffer := make([]byte, BUFFER_SIZE)
		clientBuffer := make([]byte, BUFFER_SIZE)

		connectionSocket := serverSocket.Accept()
		size = connectionSocket.Receive(serverBuffer)
		req := parser.HttpRequest(serverBuffer)
		fmt.Println(req.String())
		if req.Path == PROXY_PATH {
			if res, ok := cache[req.Path]; ok {
				log.Println("Serving response from cache")
				connectionSocket.Send(res)
				continue
			} else {
				log.Printf("Request will be cached")
				isRequestCacheable = true
			}
		}

		// HTTP 1.0 -> Connect proxy to Server for every request
		clientSocket := socket.New(SERVER_PORT, ADDRESS)
		clientSocket.Connect()
		clientSocket.Send(serverBuffer[:size])
		size = clientSocket.ReceiveAll(clientBuffer)
		res := parser.HttpResponse(clientBuffer)
		fmt.Println(res.String())

		connectionSocket.Send(clientBuffer[:size])
		if isRequestCacheable {
			cache[req.Path] = clientBuffer[:size]
		}

		// Close connections
		connectionSocket.Close()
		clientSocket.Close()
	}
}
