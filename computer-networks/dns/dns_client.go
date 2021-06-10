package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

const (
	GOOGLE_DNS    = "8.8.8.8:53"
	CLOUDFARE_DNS = "1.1.1.1"
)

type HEADER struct {
	ID      uint16
	FLAGS   [2]byte
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

type QUESTION struct {
	QNAME  []byte // As long as required
	QTYPE  uint16 // 1 is A
	QCLASS uint16 // 1 is IN
}

type RESOURCE_RECORD struct {
	NAME     uint32
	TYPE     uint16
	CLASS    uint16
	TTL      uint16
	RDLENGTH uint16
	RDATA    uint16
}

type QUERY struct {
	Header     HEADER
	Question   QUESTION
	Answer     RESOURCE_RECORD
	Authority  RESOURCE_RECORD
	Additional RESOURCE_RECORD
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	hostname := "bradfield.com"

	conn, err := net.Dial("udp", GOOGLE_DNS)
	checkErr(err)
	fmt.Println(conn)
	dnsHeader := HEADER{1, [2]byte{0x01, 0x00}, 1, 0, 0, 0}

	parts := strings.Split(hostname, ".")
	msg := []byte{}
	var network bytes.Buffer
	binary.Write(&network, binary.BigEndian, dnsHeader)
	// add message
	for _, v := range parts {
		size := len(v)
		msg = append(msg, byte(size))
		for _, r := range v {
			msg = append(msg, byte(r))
		}
	}
	msg = append(msg, byte(0x00))
	question := QUESTION{msg, 1, 1}
	// Query message delimeter
	binary.Write(&network, binary.BigEndian, question.QNAME)
	binary.Write(&network, binary.BigEndian, question.QCLASS)
	binary.Write(&network, binary.BigEndian, question.QTYPE)
	binary.Write(&network, binary.BigEndian, RESOURCE_RECORD{})
	binary.Write(&network, binary.BigEndian, RESOURCE_RECORD{})
	binary.Write(&network, binary.BigEndian, RESOURCE_RECORD{})

	// binary.Write(conn, binary.BigEndian, network)

	conn.Write(network.Bytes())

	defer conn.Close()
}
