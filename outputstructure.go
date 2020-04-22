package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ayyaruq/zanarkand"
)

// IpcStructure - Struct of the fields that IPC packets can have
type IpcStructure struct {
	zanarkand.GameEventMessage
	Direction       string `json:"direction"`
	IsEgressMessage bool   `json:"-"`
	Region          string `json:"region"`
	SubType         string `json:"subType"`
	SuperType       string `json:"superType"`
	Type            string `json:"type"`
	IpcMessageFields
}

// MarshalJSON overrides all child JSON serialization methods.
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
		Type:      jsifyString(ipc.Type),
		SubType:   jsifyString(ipc.SubType),
		SuperType: jsifyString(ipc.SuperType),
		Direction: ipc.Direction,
		ServerID:  ipc.ServerID,
		Region:    ipc.Region,
		Timestamp: int32(ipc.Timestamp.Unix()),
		Data:      data,
	})
	if err != nil {
		log.Println(err) // shouldn't happen but might
		return nil, err
	}

	b2, err := json.Marshal(ipc.IpcMessageFields)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// This whole string concatenation thing is gross in principle, but it's intuitive and works,
	// as opposed to other things that don't work, like dynamicstruct (can't merge a struct with an interface)
	// or json.Marshal overrides (the fields of the interface become an object under a new key rather than being embedded)
	s1 := string(b1[:len(b1)-1])
	s2 := string(b2[1:])
	compositeJSON := s1 + ", " + s2
	if s2 == "ull" { // "null" with the first rune chopped off
		compositeJSON = string(b1)
	}

	return []byte(compositeJSON), nil
}

func jsifyString(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToLower(str[0:1]) + str[1:]
}

// IpcMessageFields - Holds any IPC struct to be serialized later
type IpcMessageFields interface{}
