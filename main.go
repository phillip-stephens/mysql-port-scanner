package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	maxPortNumber = 65353
)

func main() {
	inputs, err := collectInputs()
	if err != nil {
		fmt.Printf("unable to parse inputs: %v\nExiting", err)
		os.Exit(1)
	}

	// connect to MySQL server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", inputs.IP, inputs.Port))
	if err != nil {
		fmt.Printf("no MySQL instance responded on IP %v and port %d", inputs.IP, inputs.Port)
		return
	}

	packet := make([]byte, 128)
	bytesRead, err := bufio.NewReader(conn).Read(packet)
	if err != nil {
		fmt.Printf("error: couldn't copy server reply into a byte array: %v", err)
		os.Exit(1)
	}

	handshake, err := parsePacket(packet[0:bytesRead])
	if err != nil {
		fmt.Printf("error: could not parse handshake packet: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("MySQL instance is running on:\n"+
		"IP: %s\n"+
		"Port: %d\n\n"+
		"%s", inputs.IP, inputs.Port, handshake)
}
