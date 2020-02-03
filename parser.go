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

// Cast the message data to a packet structure
func parseMessage(message *zanarkand.GameEventMessage, region string, port uint16) {
	ipcStructure := new(IpcStructure)
	ipcStructure.GameEventMessage = *message
	ipcStructure.Region = region

	ipcStructure.Type = getPacketType(message.Opcode, region)

	marshalType(ipcStructure.Type, ipcStructure.Body) // no assignment yet, needs work

	ipcStructure.Type = jsifyString(ipcStructure.Type)

	if message.Opcode == actorControl || message.Opcode == actorControlSelf || message.Opcode == actorControlTarget {
		ipcStructure.SuperType = "actorControl"
		ipcStructure.SubType = jsifyString(ActorControlType[binary.LittleEndian.Uint16(message.Body[0:2])])
	} else if message.Opcode == clientTrigger {
		ipcStructure.SuperType = "clientTrigger"
		ipcStructure.SubType = jsifyString(ClientTriggerType[binary.LittleEndian.Uint16(message.Body[0:2])])
	}

	serializePackout(ipcStructure, port)
}

// *Serialize* the *pack*et and send it *out* over the network
func serializePackout(ipcStructure *IpcStructure, port uint16) {
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

func jsifyString(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}
