package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"net"

	"github.com/ayyaruq/zanarkand"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

var actorControl uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControl"]
var actorControlSelf uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControlSelf"]
var actorControlTarget uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControlTarget"]
var clientTrigger uint16 = sapphire.ClientZoneIpcType.ByKeys["ClientTrigger"]

// Cast the message []byte to a packet structure and serialize the whole thing.
func parseMessage(message *zanarkand.GameEventMessage, region string, cnctns []net.Conn, isDirectionEgress bool, isDev bool) {
	ipcStructure := createIpcStructure(message, region, isDirectionEgress)

	//ipcStructure.IpcMessageFields = ipcStructure.UnmarshalType() // TODO: Finish this

	if message.Opcode == actorControl || message.Opcode == actorControlSelf || message.Opcode == actorControlTarget {
		ipcStructure.IdentifyActorControl()
	} else if message.Opcode == clientTrigger {
		ipcStructure.IdentifyClientTrigger()
	}

	// Clear the data array for transport in production.
	if !isDev {
		ipcStructure.Body = make([]byte, 0)
	}

	ipcStructure.SerializePackout(cnctns, isDev)
}

func createIpcStructure(message *zanarkand.GameEventMessage, region string, isDirectionEgress bool) *IpcStructure {
	ipcStructure := new(IpcStructure)
	ipcStructure.GameEventMessage = *message
	ipcStructure.Region = region
	ipcStructure.IsEgressMessage = isDirectionEgress

	ipcStructure.Type = ipcStructure.GetPacketType()
	if isDirectionEgress {
		ipcStructure.Direction = "outbound"
	} else {
		ipcStructure.Direction = "inbound"
	}

	return ipcStructure
}

// GetPacketType - Gets the type of the struct correspnding to the IpcStructure's opcode.
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

// IdentifyActorControl sets the name of the ActorControl category on the packet.
func (ipcStructure *IpcStructure) IdentifyActorControl() {
	ipcStructure.SuperType = "ActorControl"
	ipcStructure.SubType = sapphire.ActorControlTypeReverse[binary.LittleEndian.Uint16(ipcStructure.GameEventMessage.Body[0:2])]
}

// IdentifyClientTrigger sets the name of the ClientTrigger category on the packet.
func (ipcStructure *IpcStructure) IdentifyClientTrigger() {
	ipcStructure.SuperType = "ClientTrigger"
	ipcStructure.SubType = sapphire.ClientTriggerTypeReverse[binary.LittleEndian.Uint16(ipcStructure.GameEventMessage.Body[0:2])]
}

// SerializePackout - *Serialize* the *pack*et and send it *out* over the network.
func (ipcStructure *IpcStructure) SerializePackout(cnctns []net.Conn, isDev bool) {
	stringBytes, err := json.Marshal(ipcStructure)
	if err != nil {
		log.Println(err)
	}
	if len(cnctns) != 0 {
		for _, conn := range cnctns {
			if conn != nil {
				err = wsutil.WriteServerMessage(conn, ws.OpText, stringBytes)
			}
		}
	} else {
		if isDev {
			var buf bytes.Buffer
			buf.Write(stringBytes)
			log.Println(&buf)
		}
	}
	if err != nil {
		log.Println(err)
	}
}

func switchRegion(region string, dataPath string) {
	sapphire.LoadOpcodes(region, dataPath)
	sapphire.LoadConstants(region, dataPath)

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
