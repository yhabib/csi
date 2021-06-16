package parser

import (
	"bytes"
	"fmt"
)

type Request struct {
	Method     string
	Path       string
	Version    string
	Host       string
	Connection string
}

func HttpRequest(data []byte) Request {
	lines := bytes.Split(data, []byte{'\r', '\n'})
	requestLine := bytes.Split(lines[0], []byte{' '})
	secondLine := bytes.Split(lines[1], []byte{' '})
	thirdLine := bytes.Split(lines[2], []byte{' '})

	return Request{
		Method:     string(requestLine[0]),
		Path:       string(requestLine[1]),
		Version:    string(requestLine[2]),
		Host:       string(secondLine[1]),
		Connection: string(thirdLine[1]),
	}
}

func (r Request) String() string {
	return fmt.Sprintf("---------------------\n%s %s %s\nHost: %s\nConnection: %s\n---------------------\n", r.Method, r.Path, r.Version, r.Host, r.Connection)
}

type Response struct {
	Version string
	Status  string
	Message string
	Date    string
	Length  string
	Type    string
	Body    []byte
}

func HttpResponse(data []byte) Response {
	parts := bytes.SplitN(data, []byte{'\r', '\n', '\r', '\n'}, 2)
	lines := bytes.Split(parts[0], []byte{'\r', '\n'})
	requestLine := bytes.Split(lines[0], []byte{' '})
	thirdLine := bytes.Split(lines[2], []byte{' '})
	fourthLine := bytes.Split(lines[3], []byte{' '})
	fifthLine := bytes.Split(lines[4], []byte{' '})

	return Response{
		Version: string(requestLine[0]),
		Status:  string(requestLine[1]),
		Message: string(requestLine[2]),
		Date:    string(thirdLine[1]),
		Length:  string(fourthLine[1]),
		Type:    string(fifthLine[1]),
		Body:    parts[1],
	}
}

func (r Response) String() string {
	return fmt.Sprintf("---------------------\n%s %s %s\nDate: %s\nContent-Length: %s\nContent-Type: %s\nBody: %s\n---------------------\n",
		r.Version, r.Status, r.Message, r.Date, r.Length, r.Type, r.Body)
}
