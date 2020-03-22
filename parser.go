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

var actorControl uint16 = sapphire.ServerZoneIpcType.Keys["ActorControl"]
var actorControlSelf uint16 = sapphire.ServerZoneIpcType.Keys["ActorControlSelf"]
var actorControlTarget uint16 = sapphire.ServerZoneIpcType.Keys["ActorControlTarget"]
var clientTrigger uint16 = sapphire.ClientZoneIpcType.Keys["ClientTrigger"]

// Cast the message data to a packet structure
func parseMessage(message *zanarkand.GameEventMessage, region *string, port *uint16, isDirectionEgress bool, isDev *bool) {
	ipcStructure := new(IpcStructure)
	ipcStructure.GameEventMessage = *message
	ipcStructure.Region = *region

	ipcStructure.Type = getPacketType(message.Opcode, *region, isDirectionEgress)
	if isDirectionEgress {
		ipcStructure.Direction = "outbound"
	} else {
		ipcStructure.Direction = "inbound"
	}

	ipcStructure.IpcParameters = marshalType(ipcStructure.Type, ipcStructure.Body)

	ipcStructure.Type = jsifyString(ipcStructure.Type)

	if message.Opcode == actorControl || message.Opcode == actorControlSelf || message.Opcode == actorControlTarget {
		ipcStructure.SuperType = "actorControl"
		ipcStructure.SubType = jsifyString(sapphire.ActorControlTypeReverse[binary.LittleEndian.Uint16(message.Body[0:2])])
	} else if message.Opcode == clientTrigger {
		ipcStructure.SuperType = "clientTrigger"
		ipcStructure.SubType = jsifyString(sapphire.ClientTriggerTypeReverse[binary.LittleEndian.Uint16(message.Body[0:2])])
	}

	if !*isDev {
		ipcStructure.Body = make([]byte, 0)
	}

	serializePackout(ipcStructure, port)
}

// *Serialize* the *pack*et and send it *out* over the network
func serializePackout(ipcStructure *IpcStructure, port *uint16) {
	var buf bytes.Buffer
	var bytes []byte
	bytes, _ = json.Marshal(ipcStructure)
	buf.Write(bytes)
	_, err := http.Post("http://localhost:"+fmt.Sprint(*port), "application/json", &buf)
	if err != nil && ipcStructure.Opcode == actorControl || ipcStructure.Opcode == actorControlSelf || ipcStructure.Opcode == actorControlTarget || ipcStructure.Opcode == clientTrigger {
		log.Println(&buf)
	}
}

func getPacketType(opcode uint16, region string, isDirectionEgress bool) string {
	var ipcType string
	var ok bool
	if isDirectionEgress {
		ipcType, ok = sapphire.ServerZoneIpcType.Values[opcode]
		if !ok {
			ipcType, ok = sapphire.ServerLobbyIpcType.Values[opcode]
		}
		if !ok {
			ipcType, ok = sapphire.ServerChatIpcType.Values[opcode]
		}
	} else {
		ipcType, ok = sapphire.ClientZoneIpcType.Values[opcode]
		if !ok {
			ipcType, ok = sapphire.ClientLobbyIpcType.Values[opcode]
		}
		if !ok {
			ipcType, ok = sapphire.ClientChatIpcType.Values[opcode]
		}
	}
	if !ok {
		ipcType = "unknown"
	}

	return ipcType
}

func switchRegion(region string) {
	sapphire.LoadOpcodes(region)

	if queryActorControl, ok := sapphire.ServerZoneIpcType.Keys["ActorControl"]; ok {
		actorControl = queryActorControl
	} else {
		actorControl = 0xFFFF
	}

	if queryActorControlSelf, ok := sapphire.ServerZoneIpcType.Keys["ActorControlSelf"]; ok {
		actorControlSelf = queryActorControlSelf
	} else {
		actorControlSelf = 0xFFFF
	}

	if queryActorControlTarget, ok := sapphire.ServerZoneIpcType.Keys["ActorControlTarget"]; ok {
		actorControlTarget = queryActorControlTarget
	} else {
		actorControlTarget = 0xFFFF
	}

	if queryClientTrigger, ok := sapphire.ClientZoneIpcType.Keys["ClientTrigger"]; ok {
		clientTrigger = queryClientTrigger
	} else {
		clientTrigger = 0xFFFF
	}
}

func jsifyString(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}
