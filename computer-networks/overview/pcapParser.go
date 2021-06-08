package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("Error in %s:\t%v\n", msg, err))
	}
}

// +------------------------------+
// |        Magic number          |
// +--------------+---------------+
// |Major version | Minor version |
// +--------------+---------------+
// |      Time zone offset        |
// +------------------------------+
// |     Time stamp accuracy      |
// +------------------------------+
// |       Snapshot length        |
// +------------------------------+
// |   Link-layer header type     |
// +------------------------------+
// Size: 24B
func parsePcapHeader(f *os.File) binary.ByteOrder {
	MAGIC_NUMBER := "a1b2c3d4"
	fmt.Println("Parsing PCAP Header:")
	fileHeader := make([]byte, 24)
	f.Read(fileHeader)
	magicNumber := fileHeader[0:4]
	fmt.Printf(" Magic Number: % 0x\n", magicNumber)

	if hex.EncodeToString(magicNumber) == MAGIC_NUMBER {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

//+-----------------------------------------------+
// |	 			  Time stamp, seconds value           |
// +----------------------------------------------+
// |Time stamp, microseconds or nanoseconds value |
// +----------------------------------------------+
// |       Length of captured packet data         |
// +----------------------------------------------+
// |   Un-truncated length of the packet data     |
// +----------------------------------------------+
// Size: 16B
func parsePrePacket(f *os.File, order binary.ByteOrder) (int64, error) {
	fmt.Println("Parsing Pre Packet:")
	packetHeader := make([]byte, 16)
	_, err := f.Read(packetHeader)
	truncatedPacketSize := order.Uint32(packetHeader[8:12])
	packetSize := order.Uint32(packetHeader[12:16])
	var isTruncated string
	if packetSize == truncatedPacketSize {
		isTruncated = "NO"
	} else {
		isTruncated = "YES"
	}

	fmt.Printf("	Truncated Packet Size: %4d Bytes\n", truncatedPacketSize)
	fmt.Printf("	Packet Size: %14d Bytes\n", packetSize)
	fmt.Printf("	Truncated data? %11s\n", isTruncated)
	return int64(packetSize), err
}

// +--------------+---------------+--------------+---------------+--------------+---------------+--------------+-----------------+
// | Preamble | Start of Fram | MAC Destination | MAC Source | 802.1Q tag | Ethertype | Payload  | Frame Check | Interpacket Gap |
// | 7B *			| 1B	*					| 6B						  | 6B				 | 4B *				| 2B			  | 46-1500B | 4B	*				 | 12B						 |
// +--------------+---------------+--------------+---------------+--------------+---------------+--------------+-----------------+
// etherTypeIp4 := 0x800
// etherTypeIp6 := 0x86DD
// BIG ENDIAN
// BIG ENDIAN
func parseEthernet(f *os.File) int64 {
	macSize := 6
	etherTypeSize := 2
	headerSize := 2*macSize + etherTypeSize
	fmt.Println("Parsing Ethernet Frame:")
	frameHeader := make([]byte, headerSize)
	f.Read(frameHeader)
	fmt.Printf("	MAC Destination: %10x \n", binary.BigEndian.Uint32(frameHeader[0:macSize]))
	fmt.Printf("	MAC Source: %15x \n", binary.BigEndian.Uint32(frameHeader[macSize:2*macSize]))
	fmt.Printf("	Ethertype: %15x \n", binary.BigEndian.Uint16(frameHeader[2*macSize:]))
	return int64(headerSize)
}

// 32bit Words
// BIG ENDIAN
// First Byte contains version + IHL -> zero out the higher part of the Byte
// Total Lenght: Hedar + Payload
// TCP protocol: 0x06
func parseIp(f *os.File) int64 {
	fmt.Println("Parsing IP Datagram:")
	firstByte := make([]byte, 1)
	_, err := f.Read(firstByte)
	handleError(err, "IP: IHL")
	IHL := (firstByte[0] & 0x0f) << 2 // this is in Words but it has to be in Bytes thus * 4

	f.Seek(1, 1) // skips next byte w/ no relevant information
	totalLength := make([]byte, 2)
	_, err = f.Read(totalLength)
	handleError(err, "IP: Total lenght")

	f.Seek(5, 1)
	protocol := make([]byte, 1)
	_, err = f.Read(protocol)
	handleError(err, "IP: Protocol")

	f.Seek(2, 1)

	addresses := make([]byte, 8)
	_, err = f.Read(addresses)
	handleError(err, "IP: Addresses")

	payloadSize := binary.BigEndian.Uint16(totalLength) - uint16(IHL)

	fmt.Printf("	IHL: %18dB \n", IHL)
	fmt.Printf("	Payload Size: %9dB \n", payloadSize)
	fmt.Printf("	Total length: %9dB \n", binary.BigEndian.Uint16(totalLength))
	fmt.Printf("	Protocol: %12x \n", protocol[0])
	fmt.Printf("	Source Address: %13x \n", addresses[0:4])
	fmt.Printf("	Destination Address: %x \n", addresses[4:])

	return int64(IHL)
}

type TCPSegmentHeader struct {
	SrcPort                   uint16
	DstPort                   uint16
	SeqNum                    uint32
	AckNum                    uint32
	DataOffset_Reserved_Flags [2]byte
	WindowSize                uint16
	Checksum                  uint16
	UrgentPtr                 uint16
}

func parseTcp(f *os.File) int64 {
	fmt.Println("Parsing TCP Segment:")
	offset, _ := f.Seek(0, 1)
	fmt.Println(offset)
	tcpHeader := new(TCPSegmentHeader)
	binary.Read(f, binary.BigEndian, tcpHeader)
	dataOffset := (tcpHeader.DataOffset_Reserved_Flags[0] >> 4) * 4
	offsetSize := dataOffset - 20
	fmt.Printf("	Source Port: %d \n", tcpHeader.SrcPort)
	fmt.Printf("	Destination Port: %d \n", tcpHeader.DstPort)
	fmt.Printf("	Sequenze Number: %d \n", tcpHeader.SeqNum)
	fmt.Printf("	Header Size: %dB \n", dataOffset)
	fmt.Printf("	Offset Size: %dB \n", offsetSize)
	offset, _ = f.Seek(int64(offsetSize), 1)
	fmt.Println(offset)
	return int64(dataOffset)
}

func parseHTTP(f *os.File, payloadSize int64) {
	fmt.Printf("	Payload Size: %dB \n", payloadSize)
	f.Seek(payloadSize, 1)
}

func main() {
	f, err := os.Open("./net.cap")
	handleError(err, "Coulnd't open the file")
	defer f.Close()
	order := parsePcapHeader(f)
	numOfPackets := 0

	for {
		fmt.Println("----------------------------------------------------------------")
		packetSize, err := parsePrePacket(f, order)
		if err == io.EOF {
			break
		}
		handleError(err, "Reading packet")
		ethHeaderSize := parseEthernet(f)
		ipHeaderSize := parseIp(f)
		tcpHeaderSize := parseTcp(f)
		httpBytes := packetSize - ethHeaderSize - ipHeaderSize - tcpHeaderSize
		parseHTTP(f, httpBytes)
		numOfPackets++
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Summary: ")
	fmt.Println("	Number of packets: ", numOfPackets)
}
