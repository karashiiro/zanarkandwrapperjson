package zanarkandwrapperjson

// IpcHead - Distance to IPC header
const IpcHead uint8 = 0x60

// Offsets from the beginning of the IPC header
const (
	PacketSize  uint8 = IpcHead + 0x00
	SourceActor uint8 = IpcHead + 0x04
	TargetActor uint8 = IpcHead + 0x08
	SegmentType uint8 = IpcHead + 0x0C
	IpcType     uint8 = IpcHead + 0x12
	ServerID    uint8 = IpcHead + 0x16
	Timestamp   uint8 = IpcHead + 0x18
	IpcData     uint8 = IpcHead + 0x20
)