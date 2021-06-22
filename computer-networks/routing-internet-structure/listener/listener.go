package listener

import (
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Socket struct {
	Fd   int
	Addr syscall.Sockaddr
}

func New(port int, addr [4]byte) Socket {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	check(err)

	return Socket{fd, &syscall.SockaddrInet4{Port: port, Addr: addr}}
}

func (s *Socket) Bind() {
	err := syscall.Bind(s.Fd, s.Addr)
	check(err)
}

func (s *Socket) Receive(buffer []byte) (addr syscall.Sockaddr, err error) {
	_, addr, err = syscall.Recvfrom(s.Fd, buffer, 0)
	return
}

func (s *Socket) Close() {
	err := syscall.Close(s.Fd)
	check(err)
}
