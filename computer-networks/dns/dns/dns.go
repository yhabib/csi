package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type HEADER struct {
	ID      uint16
	FLAGS   [2]byte
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

type QUESTION struct {
	QNAME  []byte // As long as required
	QTYPE  uint16 // 1 is A
	QCLASS uint16 // 1 is IN
}

type RESOURCE_RECORD struct {
	NAME     uint16
	TYPE     uint16
	CLASS    uint16
	TTL      uint32
	RDLENGTH uint16
	RDATA    []byte
}

type QUERY struct {
	Header     HEADER
	Question   QUESTION
	Answer     RESOURCE_RECORD
	Authority  RESOURCE_RECORD
	Additional RESOURCE_RECORD
}

const CLASS_IN uint16 = 1

var fromRecordType = map[string]uint16{
	"A":     1,
	"NS":    2,
	"CNAME": 5,
	"MX":    15,
	"AAAA":  28,
}

var toRecordTyoe = map[uint16]string{
	1:  "A",
	2:  "NS",
	5:  "CNAME",
	15: "MX",
	28: "AAAA",
}

func BuildQuery(hostname string, qType string, buffer *bytes.Buffer) {
	dnsHeader := HEADER{1, [2]byte{0x01, 0x00}, 1, 0, 0, 0}

	parts := strings.Split(hostname, ".")
	msg := []byte{}
	binary.Write(buffer, binary.BigEndian, dnsHeader)
	// add message
	for _, v := range parts {
		msg = append(msg, byte(len(v)))
		msg = append(msg, []byte(v)...)
	}
	msg = append(msg, byte(0x00))
	question := QUESTION{msg, CLASS_IN, fromRecordType[qType]}
	// Query message delimeter
	binary.Write(buffer, binary.BigEndian, question.QNAME)
	binary.Write(buffer, binary.BigEndian, question.QCLASS)
	binary.Write(buffer, binary.BigEndian, question.QTYPE)
	// Handle error
}

func ParseResponse(data []byte, size int) RESOURCE_RECORD {
	rdData := data[size:]
	i := uint16(0)
	answer := RESOURCE_RECORD{
		NAME:     read16(rdData, &i),
		TYPE:     read16(rdData, &i),
		CLASS:    read16(rdData, &i),
		TTL:      read32(rdData, &i),
		RDLENGTH: read16(rdData, &i),
	}
	answer.RDATA = rdData[i : i+answer.RDLENGTH]

	return answer
}

func read32(data []byte, pointer *uint16) uint32 {
	x := binary.BigEndian.Uint32(data[*pointer:])
	*pointer += 4
	return x
}

func read16(data []byte, pointer *uint16) uint16 {
	x := binary.BigEndian.Uint16(data[*pointer:])
	*pointer += 2
	return x
}

func renderAddress(addr []byte) string {
	nums := make([]string, len(addr))
	for i, v := range addr {
		nums[i] = string(uint16(v))
	}
	return strings.Join(nums, ".")
}

func (rr *RESOURCE_RECORD) String() string {
	return fmt.Sprintf("Class: %s\nTTL: %d\nRDLENGHT: %d\nRDATA: %s\n", toRecordTyoe[rr.CLASS], rr.TTL, rr.RDLENGTH, renderAddress(rr.RDATA))
}
