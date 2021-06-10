package main

import (
	"encoding/binary"
	"fmt"
	"net"
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
	QNAME  [4]byte
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
	// init
	conn, err := net.Dial("udp", GOOGLE_DNS)
	checkErr(err)
	fmt.Println(conn)
	dnsHeader := HEADER{1, [2]byte{0x01, 0x00}, 1, 0, 0, 0}

	// msg, _ := strconv.Atoi("www.bradfield.com")
	// dnsType, _ := strconv.Atoi("A")
	// dnsClass, _ := strconv.Atoi("IN")
	// fmt.Println(msg)
	// fmt.Println(dnsType)
	// fmt.Println(dnsClass)
	// dnsQuestion := QUESTION{uint32(msg), uint16(dnsType), uint16(dnsClass)}
	// msg := []byte("bradfield.com")
	msg := [4]byte{0xFF, 0xFF, 0xFF, 0xFF}
	fmt.Println(msg)

	dnsQuestion := QUESTION{msg, 1, 1}
	record := RESOURCE_RECORD{}
	query := QUERY{dnsHeader, dnsQuestion, record, record, record}
	binary.Write(conn, binary.BigEndian, query)

	defer conn.Close()
}
