package main

import (
	"bytes"
	"log"
	"net"

	"github.com/ayyaruq/zanarkand"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	jsoniter "github.com/json-iterator/go"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

var actorControl uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControl"]
var actorControlSelf uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControlSelf"]
var actorControlTarget uint16 = sapphire.ServerZoneIpcType.ByKeys["ActorControlTarget"]
var clientTrigger uint16 = sapphire.ClientZoneIpcType.ByKeys["ClientTrigger"]

// MessageDirection represents the direction a message's packet was sent in.
type MessageDirection bool

// Packet direction.
const (
	Egress  MessageDirection = true
	Ingress MessageDirection = false
)

// ParseMessage wraps the game event message information into an IpcStructure.
func ParseMessage(message *zanarkand.GameEventMessage, region string, packetDirection MessageDirection, isDev bool) *IpcStructure {
	ipcStructure := NewIpcStructure(message, region, packetDirection)

	if !isDev {
		ipcStructure.Body = nil
	}

	return ipcStructure
}

// SerializePackout - *Serialize* the *pack*et and send it *out* over the network.
func SerializePackout(ipcStructure *IpcStructure, cnctns []net.Conn, isDev bool) {
	stringBytes, err := jsoniter.Marshal(ipcStructure)
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

// SwitchRegion is responsible for getting region-specific opcodes after initialization.
func SwitchRegion(region string, dataPath string) {
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
