package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ayyaruq/zanarkand"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

var actorControl uint16 = sapphire.ServerZoneIpcType["ActorControl"]
var actorControlSelf uint16 = sapphire.ServerZoneIpcType["ActorControlSelf"]
var actorControlTarget uint16 = sapphire.ServerZoneIpcType["ActorControlTarget"]
var clientTrigger uint16 = sapphire.ClientZoneIpcType["ClientTrigger"]

func parseMessage(message *zanarkand.GameEventMessage, region string, port uint16) {
	ipcType := getPacketType(message.Opcode, region)

	ipcType = strings.ToLower(ipcType[0:1]) + ipcType[1:]

	var actorControlCategory string
	var clientTriggerCategory string
	if message.Opcode == actorControl || message.Opcode == actorControlSelf || message.Opcode == actorControlTarget {
		actorControlCategory = ActorControlType[binary.LittleEndian.Uint16(message.Body[0:2])]
		actorControlCategory = strings.ToLower(actorControlCategory[0:1]) + actorControlCategory[1:]
	} else if message.Opcode == clientTrigger {
		clientTriggerCategory = ClientTriggerType[binary.LittleEndian.Uint16(message.Body[0:2])]
		clientTriggerCategory = strings.ToLower(clientTriggerCategory[0:1]) + clientTriggerCategory[1:]
	}

	serializePacket(message, ipcType, actorControlCategory, clientTriggerCategory, region, port)
}

func serializePacket(message *zanarkand.GameEventMessage, ipcType string, actorControlCategory string, clientTriggerCategory string, region string, port uint16) {
	var ipcStructure IpcStructure

	ipcStructure.Type = ipcType
	ipcStructure.Opcode = message.Opcode
	ipcStructure.Region = region
	ipcStructure.PacketSize = message.Length
	ipcStructure.SegmentType = message.Segment

	if message.Segment == 3 {
		ipcStructure.SourceActor = message.SourceActor
		ipcStructure.TargetActor = message.TargetActor
		ipcStructure.ServerID = message.ServerID
		ipcStructure.Timestamp = message.Timestamp.Unix()

		if actorControlCategory != "" {
			ipcStructure.SuperType = "actorControl"
			ipcStructure.SubType = actorControlCategory
		} else if clientTriggerCategory != "" {
			ipcStructure.SuperType = "clientTrigger"
			ipcStructure.SubType = clientTriggerCategory
		}
	}

	// JSON marshalling doesn't work on bytes, byte arrays are converted into strings and need to be bitshifted out later, presumably?
	// At any rate, it's not compatible with how .NET serializes byte arrays.
	serializedData := make([]int, len(message.Body))
	for i, b := range message.Body {
		serializedData[i] = int(b)
	}
	ipcStructure.Body = serializedData

	var buf bytes.Buffer
	var bytes []byte
	bytes, _ = json.Marshal(ipcStructure)
	buf.Write(bytes)
	_, err := http.Post("http://localhost:"+fmt.Sprint(port), "application/json", &buf)
	if err != nil {
		log.Println(err)
	}
}

func getPacketType(opcode uint16, region string) string {
	var ipcType string
	var ok bool
	if region == "Global" {
		ipcType, ok = ServerLobbyIpcType[opcode]
		if !ok {
			ipcType, ok = ClientLobbyIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ServerZoneIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ClientZoneIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ServerChatIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ClientChatIpcType[opcode]
		}
	} else if region == "CN" {
		ipcType, ok = ServerLobbyIpcTypeCN[opcode]
		if !ok {
			ipcType, ok = ClientLobbyIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ServerZoneIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ClientZoneIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ServerChatIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ClientChatIpcTypeCN[opcode]
		}
	} else if region == "KR" {
		ipcType, ok = ServerLobbyIpcTypeKR[opcode]
		if !ok {
			ipcType, ok = ClientLobbyIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ServerZoneIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ClientZoneIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ServerChatIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ClientChatIpcTypeKR[opcode]
		}
	}

	if !ok {
		ipcType = "unknown"
	}

	return ipcType
}
