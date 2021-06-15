package socket

import (
	"syscall"
)

// fd: File Descriptor
// nfd: New File Descriptor

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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	check(err)

	return Socket{fd, &syscall.SockaddrInet4{Port: port, Addr: addr}}
}

func (s *Socket) Bind() {
	err := syscall.Bind(s.Fd, s.Addr)
	check(err)
}

func (s *Socket) Connect() {
	err := syscall.Connect(s.Fd, s.Addr)
	check(err)
}

func (s *Socket) Listen(backlog int) {
	err := syscall.Listen(s.Fd, backlog)
	check(err)
}

func (s *Socket) Accept() *Socket {
	fd, addr, err := syscall.Accept(s.Fd)
	check(err)
	return &Socket{Fd: fd, Addr: addr}
}

func (s *Socket) Receive(data []byte) (size int) {
	size, _, err := syscall.Recvfrom(s.Fd, data, 0)
	check(err)
	return size
}

func (s *Socket) Send(buffer []byte) {
	err := syscall.Sendto(s.Fd, buffer, 0, s.Addr)
	check(err)
}

func (s *Socket) Close() {
	err := syscall.Close(s.Fd)
	check(err)
}
