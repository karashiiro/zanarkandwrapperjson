package sapphire

// Trimmed stuff for testing while there's no opcode json

var ServerLobbyIpcType = map[string]uint16{}

var ClientLobbyIpcType = map[string]uint16{}

var ServerZoneIpcType = map[string]uint16{
	"ActorControl":       0x0246, // updated 5.21
	"ActorControlSelf":   0x02D1, // updated 5.21
	"ActorControlTarget": 0x02FB, // updated 5.21
}

var ClientZoneIpcType = map[string]uint16{
	"ClientTrigger": 0x01A0, // updated 5.21
}

var ServerChatIpcType = map[string]uint16{}

var ClientChatIpcType = map[string]uint16{}
