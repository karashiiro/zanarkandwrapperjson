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

var actorControl uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControl"]
var actorControlSelf uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControlSelf"]
var actorControlTarget uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControlTarget"]
var clientTrigger uint16 = sapphire.ClientZoneIpcType.ByKeys["ClientTrigger"]

// Cast the message data to a packet structure
func parseMessage(message *zanarkand.GameEventMessage, region *string, port *uint16, isDirectionEgress bool, isDev *bool) {
	ipcStructure := new(IpcStructure)
	ipcStructure.GameEventMessage = *message
	ipcStructure.Region = *region

	ipcStructure.Type = getPacketType(&message.Opcode, region, &isDirectionEgress)
	if isDirectionEgress {
		ipcStructure.Direction = "outbound"
	} else {
		ipcStructure.Direction = "inbound"
	}

	ipcStructure.IpcParameters = marshalType(&ipcStructure.Type, &ipcStructure.Body, &isDirectionEgress)

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

	serializePackout(ipcStructure, port, isDev)
}

// *Serialize* the *pack*et and send it *out* over the network
func serializePackout(ipcStructure *IpcStructure, port *uint16, isDev *bool) {
	var buf bytes.Buffer
	bytes, _ := json.Marshal(ipcStructure)
	buf.Write(bytes)
	_, err := http.Post("http://localhost:"+fmt.Sprint(*port), "application/json", &buf)
	if err != nil {
		if *isDev {
			log.Println(&buf)
		} else {
			log.Println(err)
		}
	}
}

func getPacketType(opcode *uint16, region *string, isDirectionEgress *bool) string {
	var ipcType string
	var ok bool
	if *isDirectionEgress {
		ipcType, ok = sapphire.ServerZoneIpcType.ByValues[*opcode]
		if !ok {
			ipcType, ok = sapphire.ServerLobbyIpcType.ByValues[*opcode]
		}
		if !ok {
			ipcType, ok = sapphire.ServerChatIpcType.ByValues[*opcode]
		}
	} else {
		ipcType, ok = sapphire.ClientZoneIpcType.ByValues[*opcode]
		if !ok {
			ipcType, ok = sapphire.ClientLobbyIpcType.ByValues[*opcode]
		}
		if !ok {
			ipcType, ok = sapphire.ClientChatIpcType.ByValues[*opcode]
		}
	}
	if !ok {
		ipcType = "unknown"
	}

	return ipcType
}

func switchRegion(region string) {
	sapphire.LoadOpcodes(region)

	if queryActorControl, ok := sapphire.ServerZoneIpcType.ByKeys["ActorControl"]; ok {
		actorControl = queryActorControl
	} else {
		actorControl = 0xFFFF
	}

	if queryActorControlSelf, ok := sapphire.ServerZoneIpcType.ByKeys["ActorControlSelf"]; ok {
		actorControlSelf = queryActorControlSelf
	} else {
		actorControlSelf = 0xFFFF
	}

	if queryActorControlTarget, ok := sapphire.ServerZoneIpcType.ByKeys["ActorControlTarget"]; ok {
		actorControlTarget = queryActorControlTarget
	} else {
		actorControlTarget = 0xFFFF
	}

	if queryClientTrigger, ok := sapphire.ClientZoneIpcType.ByKeys["ClientTrigger"]; ok {
		clientTrigger = queryClientTrigger
	} else {
		clientTrigger = 0xFFFF
	}
}

func jsifyString(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}
