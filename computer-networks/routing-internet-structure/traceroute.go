package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"time"
	"traceroute/listener"
	"traceroute/parser"
	"traceroute/pinger"
)

type IPv4 [4]byte

var (
	PORT                        = 8080
	ADDR                        = [4]byte{0, 0, 0, 0}
	ICMP_TYPE_TTL_EXPIRED uint8 = 11
	ICMP_CODE_TTL_EXPIRED uint8 = 0
	TTL                         = 10
)

// Traceroute:
//  - starts timer
//  - sends UDP packets w/ a predefined TTL
//  - receives IMCP messages, gets name & ip and stops timer
// 3 packets per TTL

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide a destination to traceroute, eg: go run traceroute www.google.com")
	}

	// port should be quite random as it is what we use to know when to stop
	port := flag.Int("p", 45123, "Set port to be used")
	pings := flag.Int("pings", 1, "Number of pings to send per TTL")
	flag.Parse()

	addr := getDestAddr(os.Args[len(os.Args)-1], *port)

	pinger := pinger.New(PORT, ADDR)
	listener := listener.New(PORT, ADDR)
	listener.Bind()

	for ttl := 1; ttl <= TTL; ttl++ {
		fmt.Printf("%d. ", ttl)
		var prevIp parser.IpAddr

		for ping := 0; ping < *pings; ping++ {
			start := time.Now()
			pinger.Ping(&addr, ttl)

			for {
				rec := make([]byte, 512)
				addr, err := listener.Receive(rec)
				duration := time.Since(start)
				icmpType, icmpCode := parser.Icmp(rec)
				ipAddr, ipName := parser.Addr(addr)
				if icmpType == ICMP_TYPE_TTL_EXPIRED || icmpCode == ICMP_CODE_TTL_EXPIRED {
					if prevIp.String() != ipAddr.String() {
						fmt.Printf("%s (%s)", ipName, ipAddr.String())
					}
					fmt.Printf(" %.2fms", float32(duration)/1000000)
					prevIp = ipAddr
					break
				}

				if err != nil {
					fmt.Println(err)
					break
				}
			}
		}
		fmt.Println()
	}
	defer pinger.Close()
	defer listener.Close()
}

func getDestAddr(arg string, port int) syscall.SockaddrInet4 {
	ips, _ := net.LookupIP(arg)
	addr := [4]byte{}

	// len(ips[0]) = 16 -> because of ipv6 8 groups of 2B
	addr[0] = ips[0][12]
	addr[1] = ips[0][13]
	addr[2] = ips[0][14]
	addr[3] = ips[0][15]
	return syscall.SockaddrInet4{Port: port, Addr: addr}
}