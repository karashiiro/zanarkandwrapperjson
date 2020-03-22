package main

import (
	"bytes"
	"encoding/binary"
)

// marshalType - Marshal an []byte to a packet structure
func marshalType(packetType string, data []byte) interface{} {
	generic := towerOfBabelSwitchEdition(packetType)

	buf := bytes.NewReader(data)
	binary.Read(buf, binary.LittleEndian, generic)

	return new(interface{})
}

// I call it this, but it's no VVVVV I guess
func towerOfBabelSwitchEdition(packetType string) interface{} {
	switch packetType {
	}

	return new(interface{})
}
