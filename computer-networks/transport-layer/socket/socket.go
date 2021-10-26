package socket

import (
	"syscall"
)

// fd: File Descriptor
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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	check(err)

	return Socket{fd, &syscall.SockaddrInet4{Port: port, Addr: addr}}
}

func (s *Socket) Bind() {
	err := syscall.Bind(s.Fd, s.Addr)
	check(err)
}

func (s *Socket) Receive(buffer []byte) (size int, addr syscall.Sockaddr) {
	size, addr, err := syscall.Recvfrom(s.Fd, buffer, 0)
	check(err)
	return size, addr
}

func (s *Socket) ReceiveNonBlocking(buffer []byte) (size int) {
	size, _, err := syscall.Recvfrom(s.Fd, buffer, 0^syscall.MSG_DONTWAIT)
	check(err)
	return size
}

func (s *Socket) Send(buffer []byte) {
	err := syscall.Sendto(s.Fd, buffer, 0, s.Addr)
	check(err)
}

func (s *Socket) SendTo(buffer []byte, addr syscall.Sockaddr) {
	err := syscall.Sendto(s.Fd, buffer, 0, addr)
	check(err)
}

func (s *Socket) Close() {
	err := syscall.Close(s.Fd)
	check(err)
}
