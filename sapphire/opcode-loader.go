package sapphire

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
)

// ServerZoneIpcType contains opcode entries for commands executing in the currently-loaded zone, from the server.
var ServerZoneIpcType Bimap16

// ClientZoneIpcType contains opcode entries for commands executing in the currently-loaded zone, from the client.
var ClientZoneIpcType Bimap16

// ServerLobbyIpcType contains opcode entries for commands executing in the lobby, from the server.
var ServerLobbyIpcType Bimap16

// ClientLobbyIpcType contains opcode entries for commands executing in the lobby, from the client.
var ClientLobbyIpcType Bimap16

// ServerChatIpcType contains opcode entries for commands executing in chat, from the server.
var ServerChatIpcType Bimap16

// ClientChatIpcType contains opcode entries for commands executing in chat, from the client.
var ClientChatIpcType Bimap16

// ActorControlTypeReverse contains opcodes for actor controls, by value.
var ActorControlTypeReverse = ReverseMap16(ActorControlType)

// ClientTriggerTypeReverse contains opcodes for actor controls, by value.
var ClientTriggerTypeReverse = ReverseMap16(ClientTriggerType)

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

var opcodeSource = "https://raw.githubusercontent.com/karashiiro/FFXIVOpcodes/master/opcodes.min.json"

// LoadOpcodes loads opcodes from the source URL.
func LoadOpcodes(region string, dataPath string) {
	log.Println("Downloading latest opcodes...")

	// Reset maps
	ServerZoneIpcType.ByKeys = make(map[string]uint16)
	ClientZoneIpcType.ByKeys = make(map[string]uint16)
	ServerLobbyIpcType.ByKeys = make(map[string]uint16)
	ClientLobbyIpcType.ByKeys = make(map[string]uint16)
	ServerChatIpcType.ByKeys = make(map[string]uint16)
	ClientChatIpcType.ByKeys = make(map[string]uint16)

	// Download opcode JSON and marshal it
	if !exists(dataPath) {
		os.MkdirAll(dataPath, 0664)
	}

	fileName := path.Join(dataPath, "opcodes.json")
	opcodeFile, err := GetFile(fileName, opcodeSource)
	if err != nil {
		log.Fatalln(err)
	}

	go PollForUpdates(fileName, opcodeSource)
	go WatchFile(fileName, func(newData io.Reader) {
		log.Println("Got new opcodes, reloading...")
		unmarshalOpcodes(newData, region)
		log.Println("Done!")
	})

	unmarshalOpcodes(opcodeFile, region)

	log.Println("Done!")
}

func unmarshalOpcodes(stream io.Reader, region string) {
	var opcodes []OpcodeRegion
	err := json.NewDecoder(stream).Decode(&opcodes)
	if err != nil {
		log.Fatalln(err)
	}

	// Load the opcodes
	for _, val := range opcodes {
		if val.Region == region {
			for _, op := range val.Lists.ServerZoneIpcType {
				ServerZoneIpcType.ByKeys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ClientZoneIpcType {
				ClientZoneIpcType.ByKeys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ServerLobbyIpcType {
				ServerLobbyIpcType.ByKeys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ClientLobbyIpcType {
				ClientLobbyIpcType.ByKeys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ServerChatIpcType {
				ServerChatIpcType.ByKeys[op.Name] = op.Opcode
			}
			for _, op := range val.Lists.ClientChatIpcType {
				ClientChatIpcType.ByKeys[op.Name] = op.Opcode
			}
		}
	}

	// Make the reversed versions
	ServerZoneIpcType.ByValues = ReverseMap16(ServerZoneIpcType.ByKeys)
	ClientZoneIpcType.ByValues = ReverseMap16(ClientZoneIpcType.ByKeys)
	ServerLobbyIpcType.ByValues = ReverseMap16(ServerLobbyIpcType.ByKeys)
	ClientLobbyIpcType.ByValues = ReverseMap16(ClientLobbyIpcType.ByKeys)
	ServerChatIpcType.ByValues = ReverseMap16(ServerChatIpcType.ByKeys)
	ClientChatIpcType.ByValues = ReverseMap16(ClientChatIpcType.ByKeys)
}
