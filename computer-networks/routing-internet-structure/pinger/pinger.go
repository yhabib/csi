package pinger

import (
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Socket struct {
	Fd int
}

func New(port int, addr [4]byte) Socket {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	check(err)

	return Socket{fd}
}

func (s *Socket) Ping(addr syscall.Sockaddr, ttl int) {
	// empty UDP package
	data := make([]byte, 20)
	// https://docs.microsoft.com/en-us/windows/win32/winsock/ipproto-ip-socket-options
	syscall.SetsockoptInt(s.Fd, syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
	err := syscall.Sendto(s.Fd, data, 0, addr)
	check(err)
}

func (s *Socket) Close() {
	err := syscall.Close(s.Fd)
	check(err)
}
