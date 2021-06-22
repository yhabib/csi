package parser

import (
	"fmt"
	"net"
	"syscall"
)

func Icmp(data []byte) (icmpType uint8, code uint8) {
	// What are those initial 20B?????? -> IP header -> do it programmatically
	icmpStart := 20
	icmpB := data[icmpStart]
	codeB := data[icmpStart+1]
	return uint8(icmpB), uint8(codeB)
}

type IpAddr [4]byte

func (ip *IpAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func Addr(sAddres syscall.Sockaddr) (ip IpAddr, name string) {
	addr := sAddres.(*syscall.SockaddrInet4)
	ip = IpAddr(addr.Addr)
	// reverse lookupAddr
	names, err := net.LookupAddr(ip.String())
	if err != nil {
		// it means no successfull reverse looup -> use IP
		return ip, ip.String()
	}
	return ip, names[0]
}
