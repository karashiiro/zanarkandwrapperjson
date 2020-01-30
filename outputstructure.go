package zanarkandwrapperjson

import "time"

// OutputBase - All packet outputs have these fields
type OutputBase struct {
	Type        string `json:"type"`
	Opcode      uint16 `json:"opcode"`
	Region      string `json:"region"`
	PacketSize  uint32 `json:"packetSize"`
	SegmentType uint16 `json:"segmentType"`
	Body        []byte `json:"body"`
}

// IpcBase - All IPC packet outputs have these fields
type IpcBase struct {
	OutputBase
	SourceActor uint32    `json:"sourceActorSessionID"`
	TargetActor uint32    `json:"targetActorSessionID"`
	ServerID    uint16    `json:"serverID"`
	Timestamp   time.Time `json:"timestamp"`
}

// IpcActorClientControl - ActorControl and ClientTrigger packets have these additional fields
type IpcActorClientControl struct {
	IpcBase
	SuperType string `json:"superType"`
	SubType   string `json:"subType"`
}
