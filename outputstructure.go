package main

import (
	"encoding/json"
	"log"

	"github.com/ayyaruq/zanarkand"
)

// IpcStructure - Struct of the fields that IPC packets can have
type IpcStructure struct {
	zanarkand.GameEventMessage
	Direction string `json:"direction"`
	Region    string `json:"region"`
	SubType   string `json:"subType"`
	SuperType string `json:"superType"`
	Type      string `json:"type"`
	IpcParameters
}

// MarshalJSON overrides all child JSON serialization methods. Apparently there's a less redundant way to do this, but at the moment I don't care because it works.
// It also lets me explicitly define field order.
func (ipc *IpcStructure) MarshalJSON() ([]byte, error) {
	data := make([]int, len(ipc.Body))
	for i, b := range ipc.Body {
		data[i] = int(b)
	}
	b1, err := json.Marshal(&struct {
		Opcode    uint16 `json:"opcode"`
		Type      string `json:"type"`
		SubType   string `json:"subType"`
		SuperType string `json:"superType"`
		Direction string `json:"direction"`
		ServerID  uint16 `json:"serverID"`
		Region    string `json:"region"`
		Timestamp int32  `json:"timestamp"`
		Data      []int  `json:"data"`
	}{
		Opcode:    ipc.Opcode,
		Type:      ipc.Type,
		SubType:   ipc.SubType,
		SuperType: ipc.SuperType,
		Direction: ipc.Direction,
		ServerID:  ipc.ServerID,
		Region:    ipc.Region,
		Timestamp: int32(ipc.Timestamp.Unix()),
		Data:      data,
	})
	if err != nil {
		log.Println(err)
	}
	b2, err := json.Marshal(ipc.IpcParameters)
	if err != nil {
		log.Println(err)
	}

	// I give it a solid 3/10, with a TODO: make this not a string manip hack
	s1 := string(b1[:len(b1)-1])
	s2 := string(b2[1:])

	return []byte(s1 + ", " + s2), nil
}

// IpcParameters - Holds any IPC struct to be serialized later
type IpcParameters interface{}
