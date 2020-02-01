package main

import (
	"github.com/ayyaruq/zanarkand"
)

// IpcPacket - Representation of an IPC packet
type IpcPacket struct {
	Metadata              *zanarkand.Frame
	PacketSize            uint32
	SourceActor           uint32
	TargetActor           uint32
	SegmentType           uint16
	Type                  string
	Opcode                uint16
	ServerID              uint16
	Timestamp             uint32
	ActorControlCategory  string
	ClientTriggerCategory string
}
