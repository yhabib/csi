package main

import (
	"bytes"
	"dns_client/dns"
	"fmt"
	"net"
	"os"
)

const (
	GOOGLE_DNS    = "8.8.8.8:53"
	CLOUDFARE_DNS = "1.1.1.1"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// TODO: Read from command line
	hostname := os.Args[1]
	dnsType := os.Args[2]
	fmt.Printf("Making query: %s for hostname: %s\n", dnsType, hostname)

	conn, err := net.Dial("udp", GOOGLE_DNS)
	checkErr(err)

	var network bytes.Buffer
	dns.BuildQuery(hostname, dnsType, &network)
	conn.Write(network.Bytes())

	resp := make([]byte, 2048)
	conn.Read(resp)
	size := network.Len()
	answer := dns.ParseResponse(resp, size)
	fmt.Println(answer.String())
	defer conn.Close()
}
