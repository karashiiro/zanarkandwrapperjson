package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"strings"

	"github.com/ayyaruq/zanarkand"
	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
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

// NewIpcStructure creates a new IpcStructure.
func NewIpcStructure(message *zanarkand.GameEventMessage, region string, isDirectionEgress bool) *IpcStructure {
	ipcStructure := new(IpcStructure)
	ipcStructure.GameEventMessage = *message
	ipcStructure.Region = region
	ipcStructure.IsEgressMessage = isDirectionEgress

	ipcStructure.Type = ipcStructure.GetPacketType()
	if isDirectionEgress {
		ipcStructure.Direction = "send"
	} else {
		ipcStructure.Direction = "receive"
	}

	//ipcStructure.IpcMessageFields = ipcStructure.UnmarshalType() // TODO: Finish this

	if message.Opcode == actorControl || message.Opcode == actorControlSelf || message.Opcode == actorControlTarget {
		ipcStructure.SuperType = "ActorControl"
		ipcStructure.SubType = sapphire.ActorControlTypeReverse[binary.LittleEndian.Uint16(ipcStructure.GameEventMessage.Body[0:2])]
	} else if message.Opcode == clientTrigger {
		ipcStructure.SuperType = "ClientTrigger"
		ipcStructure.SubType = sapphire.ClientTriggerTypeReverse[binary.LittleEndian.Uint16(ipcStructure.GameEventMessage.Body[0:2])]
	}

	return ipcStructure
}

// GetPacketType gets the type of the struct correspnding to the IpcStructure's opcode.
func (ipcStructure *IpcStructure) GetPacketType() string {
	var ipcType string
	var ok bool
	if ipcStructure.IsEgressMessage {
		ipcType, ok = sapphire.ClientZoneIpcType.ByValues[ipcStructure.Opcode]
		if !ok {
			ipcType, ok = sapphire.ClientLobbyIpcType.ByValues[ipcStructure.Opcode]
		}
		if !ok {
			ipcType, ok = sapphire.ClientChatIpcType.ByValues[ipcStructure.Opcode]
		}
	} else {
		ipcType, ok = sapphire.ServerZoneIpcType.ByValues[ipcStructure.Opcode]
		if !ok {
			ipcType, ok = sapphire.ServerLobbyIpcType.ByValues[ipcStructure.Opcode]
		}
		if !ok {
			ipcType, ok = sapphire.ServerChatIpcType.ByValues[ipcStructure.Opcode]
		}
	}
	if !ok {
		ipcType = "unknown"
	}

	return ipcType
}

// MarshalJSON overrides all child JSON serialization methods.
func (ipc *IpcStructure) MarshalJSON() ([]byte, error) {
	data := make([]int, len(ipc.Body))
	for i, b := range ipc.Body {
		data[i] = int(b)
	}

	b1, err := json.Marshal(&struct {
		Opcode        uint16 `json:"opcode"`
		Type          string `json:"type"`
		SubType       string `json:"subType"`
		SuperType     string `json:"superType"`
		Direction     string `json:"direction"`
		ServerID      uint16 `json:"serverID"`
		Region        string `json:"region"`
		Timestamp     int32  `json:"timestamp"`
		SourceActorID uint32 `json:"sourceActorSessionID"`
		TargetActorID uint32 `json:"targetActorSessionID"`
		Data          []int  `json:"data"`
	}{
		Opcode:        ipc.Opcode,
		Type:          jsifyString(ipc.Type),
		SubType:       jsifyString(ipc.SubType),
		SuperType:     jsifyString(ipc.SuperType),
		Direction:     ipc.Direction,
		ServerID:      ipc.ServerID,
		Region:        ipc.Region,
		Timestamp:     int32(ipc.Timestamp.Unix()),
		SourceActorID: ipc.GameEventMessage.SourceActor,
		TargetActorID: ipc.GameEventMessage.TargetActor,
		Data:          data,
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
