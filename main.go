package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"unsafe"
)

const (
	maxPortNumber = 65353
)

type CLIInputs struct {
	IP   net.IP
	Port uint
}

type MySQLHandshake struct {
	PacketLength      uint32 `json:"packet_len"`
	Protocol          uint8  `json:"protocol_version"`
	Version           string `json:"server_version"`
	ThreadID          int
	Salt              string
	ServerCapabilites []byte
	// ServerLanguage
	// ServerStatus
	// ExtendedServerCapabilities
	AuthPluginLength     int
	AuthenticationPlugin string
}

func main() {
	fmt.Println("MySQL scanner invoked")
	inputs, err := collectInputs()
	if err != nil {
		fmt.Printf("unable to parse inputs: %v\nExiting", err)
		return
	}
	fmt.Printf("IP: %v\nPort: %d\n\n", inputs.IP, inputs.Port)

	// connect
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", inputs.IP, inputs.Port))
	if err != nil {
		// TODO add logging about how MySQL isn't running
		fmt.Printf("could not dial MySQL server: %v\n", err)
		return
	}
	fmt.Println("About to read from conn")
	packet := make([]byte, 128)
	bytesRead, err := bufio.NewReader(conn).Read(packet)

	fmt.Printf("Rec. %d bytes\n%x\n", bytesRead, packet)

	fmt.Printf("Packet as string: %s\n", packet)

	// var protocol uint8
	handshake := MySQLHandshake{}
	err = binary.Read(bytes.NewReader(packet), binary.LittleEndian, &handshake.PacketLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Packet Length: %d\n", handshake.PacketLength)

	packet = packet[unsafe.Sizeof(handshake.PacketLength):]
	err = binary.Read(bytes.NewReader(packet), binary.LittleEndian, &handshake.Protocol)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Protocol: %d\n", handshake.Protocol)

	packet = packet[unsafe.Sizeof(handshake.Protocol):]

	// next arg is a null-terminated string, find the null byte
	index := -1
	for i, b := range packet {
		if b == 0x00 {
			index = i
			break
		}
	}
	handshake.Version = string(packet[0:index])
	fmt.Printf("Version: %s\n", handshake.Version)

}

// colectInputs parses the CLI arguements provided
// Returns arguements in a CLIInputs struct and any error that occured
func collectInputs() (*CLIInputs, error) {
	// Parse CLI inputs
	ipPtr := flag.String("ip", "0.0.0.0", "the IP address this scanner should connect to")
	portPtr := flag.Uint("port", 0, "the port this scanner should connect to")

	flag.Parse()

	// Validate inputs
	ip := net.ParseIP(*ipPtr)
	if ip == nil {
		// IP argument provided to CLI was an invalid IP
		return nil, fmt.Errorf("the IP entered %q was not a valid IPv4 IP address", *ipPtr)
	}

	// Ports should be 0 <= x <= 65353
	if *portPtr > maxPortNumber {
		return nil, fmt.Errorf("%d is too large to be a valid port. Ports must be between 0 and %d", *portPtr, maxPortNumber)
	}

	return &CLIInputs{
		IP:   ip,
		Port: *portPtr,
	}, nil
}
