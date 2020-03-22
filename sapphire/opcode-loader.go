package sapphire

// Bimap is a structure containing two maps, a by-key map and a by-value map
type Bimap struct {
	Keys   map[string]uint16
	Values map[uint16]string
}

// ServerZoneIpcType contains opcode entries for commands executing in the currently-loaded zone, from the server.
var ServerZoneIpcType Bimap

// ClientZoneIpcType contains opcode entries for commands executing in the currently-loaded zone, from the client.
var ClientZoneIpcType Bimap

// ServerLobbyIpcType contains opcode entries for commands executing in the lobby, from the server.
var ServerLobbyIpcType Bimap

// ClientLobbyIpcType contains opcode entries for commands executing in the lobby, from the client.
var ClientLobbyIpcType Bimap

// ServerChatIpcType contains opcode entries for commands executing in chat, from the server.
var ServerChatIpcType Bimap

// ClientChatIpcType contains opcode entries for commands executing in chat, from the client.
var ClientChatIpcType Bimap

// ActorControlTypeReverse contains opcodes for actor controls, by value.
var ActorControlTypeReverse = reverseMap(ActorControlType)

// ClientTriggerTypeReverse contains opcodes for actor controls, by value.
var ClientTriggerTypeReverse = reverseMap(ClientTriggerType)

// LoadOpcodes loads opcodes from the source URL.
func LoadOpcodes(region string) {
	// populate some stuff
}

func reverseMap(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}
