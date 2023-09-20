package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
)

type MySQLHandshake struct {
	PacketLength uint32 `json:"packet_len"`
	Protocol     uint8  `json:"protocol_version"`
	// Version      string `json:"server_version"`
	// ThreadID     int
	// Salt              string
	// ServerCapabilites []byte
	// // ServerLanguage
	// // ServerStatus
	// // ExtendedServerCapabilities
	// AuthPluginLength     int
	// AuthenticationPlugin string
}

func parsePacket(packet []byte) (*MySQLHandshake, error) {
	// var protocol uint8
	handshake := MySQLHandshake{}
	err := binary.Read(bytes.NewReader(packet), binary.LittleEndian, &handshake)
	if err != nil {
		return nil, errors.Wrap(err, "could not read from packet")
	}
	fmt.Printf("Handshake: %v\n", handshake)

	return &handshake, nil
}
