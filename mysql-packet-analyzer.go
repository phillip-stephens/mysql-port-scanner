package main

import (
	"bytes"
	"encoding/binary"

	"github.com/zhuangsirui/binpacker"
)

type MySQLHandshake struct {
	PacketLength         uint32 `json:"packet_len"`
	Protocol             uint8  `json:"protocol_version"`
	Version              string `json:"server_version"`
	ThreadID             uint32 `json:"thread_id"`
	AuthenticationPlugin string `json:"auth_plugin_name"`
}

func parsePacket(packet []byte) (*MySQLHandshake, error) {
	// var protocol uint8
	handshake := MySQLHandshake{}
	packer := binpacker.NewUnpacker(binary.LittleEndian, bytes.NewReader(packet))

	packer.FetchUint32(&handshake.PacketLength)
	packer.FetchUint8(&handshake.Protocol)

	// next arg. is a null-terminated string. Keep fetching bytes until we find null byte
	handshake.Version = getNullTerminatedString(packer)
	packer.FetchUint32(&handshake.ThreadID)

	// next 27 bytes are various server statuses/states per the MySQL handshake spec, skipping these
	filler := []byte{}
	packer.FetchBytes(27, &filler)

	// auth plugin data - Salt - appears in Wireshark as a null-terminated string
	_ = getNullTerminatedString(packer)

	// Final piece is the auth plugin name
	handshake.AuthenticationPlugin = getNullTerminatedString(packer)

	return &handshake, nil
}

func getNullTerminatedString(packer *binpacker.Unpacker) string {
	chars := []byte{}
	var currentChar byte
	packer.FetchByte(&currentChar)
	for currentChar != byte(0) {
		chars = append(chars, currentChar)
		packer.FetchByte(&currentChar)
	}
	return string(chars)
}
