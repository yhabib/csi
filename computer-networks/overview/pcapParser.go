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
	fileHeaderSize := 24
	MAGIC_NUMBER := "a1b2c3d4"
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Parsing PCAP Header:")
	fileHeader := make([]byte, fileHeaderSize)
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
func parsePrePacket(f *os.File, order binary.ByteOrder) (uint32, error) {
	prePacketSize := 16
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Parsing Pre Packet:")
	packetHeader := make([]byte, prePacketSize)
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
	return packetSize, err
}

// +--------------+---------------+--------------+---------------+--------------+---------------+--------------+-----------------+
// | Preamble | Start of Fram | MAC Destination | MAC Source | 802.1Q tag | Ethertype | Payload  | Frame Check | Interpacket Gap |
// | 7B *			| 1B	*					| 6B						  | 6B				 | 4B *				| 2B			  | 46-1500B | 4B	*				 | 12B						 |
// +--------------+---------------+--------------+---------------+--------------+---------------+--------------+-----------------+
// etherTypeIp4 := 0x800
// etherTypeIp6 := 0x86DD
// BIG ENDIAN
func parseEthernet(f *os.File) error {
	macSize := 6
	etherTypeSize := 2
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Parsing Ethernet Frame:")
	frameHeader := make([]byte, 2*macSize+etherTypeSize)
	_, err := f.Read(frameHeader)
	fmt.Println(frameHeader)
	fmt.Printf("	MAC Destination: %10x \n", binary.BigEndian.Uint32(frameHeader[0:macSize]))
	fmt.Printf("	MAC Source: %15x \n", binary.BigEndian.Uint32(frameHeader[macSize:2*macSize]))
	fmt.Printf("	Ethertype: %15x \n", binary.BigEndian.Uint16(frameHeader[2*macSize:]))
	return err
}

// 32bit Words
// BIG ENDIAN
// First Byte contains version + IHL -> zero out the higher part of the Byte
// Total Lenght: Hedar + Payload
// TCP protocol: 0x06
func parseIp(f *os.File) error {
	fmt.Println("----------------------------------------------------------------")
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

	addresses := make([]byte, 8)
	_, err = f.Read(addresses)
	handleError(err, "IP: Addresses")

	fmt.Printf("	IHL: %18dB \n", IHL)
	fmt.Printf("	Payload lenght: %7dB \n", binary.BigEndian.Uint16(totalLength)-uint16(IHL))
	fmt.Printf("	Total length: %9dB \n", binary.BigEndian.Uint16(totalLength))
	fmt.Printf("	Protocol: %12x \n", protocol[0])
	fmt.Printf("	Source Address: %13x \n", addresses[0:4])
	fmt.Printf("	Destination Address: %x \n", addresses[4:])

	return err
}

func main() {
	f, err := os.Open("./net.cap")
	handleError(err, "Coulnd't open the file")
	defer f.Close()

	order := parsePcapHeader(f)
	numOfPackets := 0

	for i := 0; i < 1; i++ {
		_, err := parsePrePacket(f, order)
		if err == io.EOF {
			break
		}
		handleError(err, "Reading packet")
		parseEthernet(f)
		parseIp(f)
		numOfPackets++
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Summary: ")
	fmt.Println("	Number of packets: ", numOfPackets)
}
