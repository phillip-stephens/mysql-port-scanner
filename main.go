package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"
)

const (
	maxPortNumber = 65353
)

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
	if err != nil {
		fmt.Printf("couldn't read server reply into a byte array: %v", err)
		return
	}

	fmt.Printf("Rec. %d bytes\n%x\n", bytesRead, packet)

	fmt.Printf("Packet as string: %s\n", packet)

	handshake, err := parsePacket(packet[0:bytesRead])
	if err != nil {
		fmt.Printf("could not parse handshake packet: %v\n", err)
		return
	}
	fmt.Printf("Handshake: %v", spew.Sdump(*handshake))

}
