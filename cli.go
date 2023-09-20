package main

import (
	"flag"
	"fmt"
	"net"
)


type CLIInputs struct {
	IP   net.IP
	Port uint
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
