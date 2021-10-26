package main

import (
	"flag"
	"fmt"
	"log"
	"reliable_transport/socket"
	"time"
)

var (
	ADDRESS               = [4]byte{0, 0, 0, 0}
	DEFAULT_MODE          = "receiver"
	DEFAULT_RECEIVER_PORT = 8080
	DEFAULT_PROXY_PORT    = 9988
	SENDER_PORT           = 9999
	ACK                   = "ACK"
)

func main() {
	senderPort := flag.Int("sPort", DEFAULT_PROXY_PORT, fmt.Sprintf("proxy port %d", DEFAULT_PROXY_PORT))
	receiverPort := flag.Int("rPort", DEFAULT_RECEIVER_PORT, fmt.Sprintf("receiver port %d", DEFAULT_RECEIVER_PORT))
	mode := flag.String("mode", DEFAULT_MODE, fmt.Sprintf("receiver or sender %s", DEFAULT_MODE))
	flag.Parse()

	if *mode == "sender" {
		Sender(*senderPort)
	} else {
		Receiver(*receiverPort)
	}
}

func Sender(proxyPort int) {
	log.Printf("Initiating sender in port: %d\n", proxyPort)
	proxy := socket.New(proxyPort, ADDRESS)

	s := socket.New(SENDER_PORT, ADDRESS)
	defer s.Close()
	s.Bind()
	buffer := []byte("Hello world!")
	ackBuffer := make([]byte, 128)
	for {
		log.Println("Sending message...")
		s.SendTo(buffer, proxy.Addr)
		log.Println("Waiting for ACK...")
		time.Sleep(1 * time.Second)
		size := s.ReceiveNonBlocking(ackBuffer)
		log.Println("TEST")
		if size == 0 {
			log.Println("ACK not received, retrying ...")
		} else {
			log.Printf("ACK received %s\n", ackBuffer)
			break
		}
	}
}

func Receiver(port int) {
	log.Printf("Initiating receiver in port: %d\n", port)
	s := socket.New(port, ADDRESS)
	defer s.Close()
	s.Bind()
	buffer := make([]byte, 4096)
	for {
		size, addr := s.Receive(buffer)
		log.Printf("Received message of %d bytes\n", size)
		fmt.Println(string(buffer))
		if size > 0 {
			log.Printf("Sending ACK")
			s.SendTo([]byte(ACK), addr)
		}
	}
}

// HOW the hell PROXY always sends to RECEIVER
