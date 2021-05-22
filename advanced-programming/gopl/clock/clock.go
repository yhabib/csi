package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("Starting tcp server ...")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println(" -> accepting connection")
		if err != nil {
			log.Print(err) // connectino aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnects
		}
		time.Sleep(1 * time.Second)
	}
}
