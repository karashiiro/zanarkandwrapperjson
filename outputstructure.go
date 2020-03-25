package main

import (
	"encoding/json"

	"github.com/ayyaruq/zanarkand"
)

// IpcStructure - Struct of the fields that IPC packets can have
type IpcStructure struct {
	zanarkand.GameEventMessage
	ParsedIpcHeader
	IpcParameters
}

// ParsedIpcHeader holds parsed header metadata and environment metadata that makes future parsing easier
type ParsedIpcHeader struct {
	Direction string `json:"direction"`
	Region    string `json:"region"`
	SubType   string `json:"subType"`
	SuperType string `json:"superType"`
	Type      string `json:"type"`
}

func (ipc *ParsedIpcHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(&ParsedIpcHeader{
		Direction: ipc.Direction,
		Region:    ipc.Region,
		SubType:   ipc.SubType,
		SuperType: ipc.SuperType,
		Type:      ipc.Type,
	})
}

// IpcParameters - Holds any IPC struct to be serialized later
type IpcParameters interface{}
