package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/textproto"
)

type Request struct {
	Method     string
	Path       string
	Version    string
	Accept     string
	Host       string
	Connection string
}

func HttpRequest(data []byte) Request {
	lines := bytes.SplitN(data, []byte{'\r', '\n'}, 2)
	requestLine := bytes.Split(lines[0], []byte{' '})
	mimeHeader := parseMimeHeader(lines[1])

	return Request{
		Method:     string(requestLine[0]),
		Path:       string(requestLine[1]),
		Version:    string(requestLine[2]),
		Host:       mimeHeader.Get("host"),
		Accept:     mimeHeader.Get("Accept"),
		Connection: mimeHeader.Get("connect"),
	}
}

func (r Request) String() string {
	return fmt.Sprintf("---------------------\n%s %s %s\nHost: %s\nAccept: %s\nConnection: %s\n---------------------\n", r.Method, r.Path, r.Version, r.Host, r.Accept, r.Connection)
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
	lines := bytes.SplitN(data, []byte{'\r', '\n'}, 2)
	requestLine := bytes.Split(lines[0], []byte{' '})
	mimeHeader := parseMimeHeader(lines[1])
	body := bytes.SplitN(data, []byte{'\r', '\n', '\r', '\n'}, 2)[1]
	return Response{
		Version: string(requestLine[0]),
		Status:  string(requestLine[1]),
		Message: string(requestLine[2]),
		Date:    mimeHeader.Get("date"),
		Length:  mimeHeader.Get("content-length"),
		Type:    mimeHeader.Get("type"),
		Body:    body,
	}
}

func (r Response) String() string {
	return fmt.Sprintf("---------------------\n%s %s %s\nDate: %s\nContent-Length: %s\nContent-Type: %s\nBody: %s\n---------------------\n",
		r.Version, r.Status, r.Message, r.Date, r.Length, r.Type, r.Body)
}

func parseMimeHeader(data []byte) textproto.MIMEHeader {
	tp := textproto.NewReader(bufio.NewReader(bytes.NewReader(data)))
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		log.Fatalf("parseMimeHeader: %s", err)
	}
	return mimeHeader
}
