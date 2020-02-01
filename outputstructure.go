package main

// IpcStructure - Struct of fields IPC packets can have
type IpcStructure struct {
	Type        string `json:"type"`
	Opcode      uint16 `json:"opcode"`
	Region      string `json:"region"`
	PacketSize  uint32 `json:"packetSize"`
	SegmentType uint16 `json:"segmentType"`
	Body        []int  `json:"data"`
	SourceActor uint32 `json:"sourceActorSessionID"`
	TargetActor uint32 `json:"targetActorSessionID"`
	ServerID    uint16 `json:"serverID"`
	Timestamp   int64  `json:"timestamp"`
	SuperType   string `json:"superType"`
	SubType     string `json:"subType"`
}
