package main

import (
	"flag"
	"fmt"
	"log"
	"reliable_transport/socket"
)

var (
	SENDER_ADDRESS        = [4]byte{0, 0, 0, 0}
	RECEIVER_ADDRESS      = [4]byte{127, 0, 0, 1}
	DEFAULT_MODE          = "receiver"
	DEFAULT_RECEIVER_PORT = 8080
	DEFAULT_SENDER_PORT   = 9999
)

func main() {
	senderPort := flag.Int("sPort", DEFAULT_SENDER_PORT, fmt.Sprintf("destination port for sender %d", DEFAULT_SENDER_PORT))
	receiverPort := flag.Int("rPort", DEFAULT_RECEIVER_PORT, fmt.Sprintf("destination port for sender %d", DEFAULT_RECEIVER_PORT))
	mode := flag.String("mode", DEFAULT_MODE, fmt.Sprintf("receiver or sender %s", DEFAULT_MODE))
	flag.Parse()

	if *mode == "sender" {
		Sender(*senderPort)
	} else {
		Receiver(*receiverPort)
	}
}

func Sender(port int) {
	log.Printf("Initiating sender in port: %d\n", port)
	s := socket.New(port, SENDER_ADDRESS)
	buffer := []byte("Hello world!")
	log.Println("Sending message...")
	s.Send(buffer)
}

func Receiver(port int) {
	log.Printf("Initiating receiver in port: %d\n", port)
	s := socket.New(port, RECEIVER_ADDRESS)
	s.Bind()
	buffer := make([]byte, 4096)
	for {
		size := s.Receive(buffer)
		log.Printf("Received message of %d bytes\n", size)
		fmt.Println(string(buffer))
	}
}
