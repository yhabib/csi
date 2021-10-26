package main

import (
	"fmt"
	"http_proxy/parser"
	"http_proxy/socket"
	"log"
	"strings"
)

var (
	PROXY_PORT  = 8080
	SERVER_PORT = 9000
	ADDRESS     = [4]byte{0, 0, 0, 0}
	BUFFER_SIZE = 1024
	PROXY_PATH  = "/proxy"
)

func main() {
	cache := make(map[string][]byte)
	size := 0

	serverSocket := socket.New(PROXY_PORT, ADDRESS)
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
		doesPathContainCachePath := strings.HasPrefix(req.Path, PROXY_PATH)
		if doesPathContainCachePath {
			if res, ok := cache[req.Path]; ok {
				log.Println("Serving response from cache")
				connectionSocket.Send(res)
				connectionSocket.Close()
				continue
			}
		}

		// HTTP 1.0 -> Connect proxy to Server for every request
		clientSocket := socket.New(SERVER_PORT, ADDRESS)
		clientSocket.Connect()
		clientSocket.Send(serverBuffer[:size])
		size = clientSocket.Receive(clientBuffer)
		res := parser.HttpResponse(clientBuffer)
		fmt.Println(res.String())
		fmt.Println(size)
		fmt.Println(res.Length)
		connectionSocket.Send(clientBuffer[:size])
		if doesPathContainCachePath {
			log.Printf("Caching request")
			cache[req.Path] = clientBuffer[:size]
		}

		// Close connections
		connectionSocket.Close()
		clientSocket.Close()
	}
}
