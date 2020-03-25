package sapphire

import (
	"encoding/json"
	"log"
	"net/http"
)

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

// OpcodeEntry has an opcode entry
type OpcodeEntry struct {
	Name   string `json:"name"`
	Opcode uint16 `json:"opcode"`
}

// OpcodeLists has channel types
type OpcodeLists struct {
	ServerZoneIpcType  []OpcodeEntry
	ClientZoneIpcType  []OpcodeEntry
	ServerLobbyIpcType []OpcodeEntry
	ClientLobbyIpcType []OpcodeEntry
	ServerChatIpcType  []OpcodeEntry
	ClientChatIpcType  []OpcodeEntry
}

// OpcodeRegion has all opcodes for a region
type OpcodeRegion struct {
	Region string      `json:"region"`
	Lists  OpcodeLists `json:"lists"`
}

var dataSource = "https://raw.githubusercontent.com/karashiiro/FFXIVOpcodes/master/opcodes.min.json"

// LoadOpcodes loads opcodes from the source URL.
func LoadOpcodes(region string) {
	log.Println("Downloading latest opcodes...")

	// Reset maps
	ServerZoneIpcType.Keys = make(map[string]uint16)
	ClientZoneIpcType.Keys = make(map[string]uint16)
	ServerLobbyIpcType.Keys = make(map[string]uint16)
	ClientLobbyIpcType.Keys = make(map[string]uint16)
	ServerChatIpcType.Keys = make(map[string]uint16)
	ClientChatIpcType.Keys = make(map[string]uint16)

	// Download opcode JSON and marshal it
	res, err := http.Get(dataSource)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Opcode store downloaded. Deserializing...")

	defer res.Body.Close()

	var opcodes []OpcodeRegion

	err = json.NewDecoder(res.Body).Decode(&opcodes)
	if err != nil {
		log.Fatalln(err)
	}

	// Load the opcodes
	for _, val := range opcodes {
		if val.Region == region {
			for _, op := range val.Lists.ServerZoneIpcType {
				ServerZoneIpcType.Keys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ClientZoneIpcType {
				ClientZoneIpcType.Keys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ServerLobbyIpcType {
				ServerLobbyIpcType.Keys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ClientLobbyIpcType {
				ClientLobbyIpcType.Keys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ServerChatIpcType {
				ServerChatIpcType.Keys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ClientChatIpcType {
				ClientChatIpcType.Keys[op.Name] = op.Opcode
			}
		}
	}

	// Make the reversed versions
	ServerZoneIpcType.Values = reverseMap(ServerZoneIpcType.Keys)
	ClientZoneIpcType.Values = reverseMap(ClientZoneIpcType.Keys)
	ServerLobbyIpcType.Values = reverseMap(ServerLobbyIpcType.Keys)
	ClientLobbyIpcType.Values = reverseMap(ClientLobbyIpcType.Keys)
	ServerChatIpcType.Values = reverseMap(ServerChatIpcType.Keys)
	ClientChatIpcType.Values = reverseMap(ClientChatIpcType.Keys)

	log.Println("Done!")
}

func reverseMap(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}
