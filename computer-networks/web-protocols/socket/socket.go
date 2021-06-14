package socket

import (
	"fmt"
	"syscall"
)

// fd: File Descriptor

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Socket struct {
	fd  int
	nfd int
}

func New() Socket {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	check(err)

	return Socket{fd, 0}
}

func (s *Socket) Bind(port int, addr [4]byte) {
	address := syscall.SockaddrInet4{Port: port, Addr: addr}
	err := syscall.Bind(s.fd, &address)
	check(err)
}

func (s *Socket) Listen(backlog int) {
	err := syscall.Listen(s.fd, backlog)
	check(err)
}

func (s *Socket) Accept() {
	nfd, _, err := syscall.Accept(s.fd)
	check(err)
	s.nfd = nfd
}

func (s *Socket) Read(buffer []byte) (size int) {
	fmt.Println(s)
	size, _, err := syscall.Recvfrom(s.nfd, buffer, 0)
	check(err)
	return size
}
