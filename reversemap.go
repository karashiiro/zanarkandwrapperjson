package main

import "github.com/karashiiro/ZanarkandWrapperJSON/sapphire"

// ActorControlType - Value-first version of sapphire.ActorControlType
var ActorControlType = reverseMap(sapphire.ActorControlType)

// ClientTriggerType - Value-first version of sapphire.ClientTriggerType
var ClientTriggerType = reverseMap(sapphire.ClientTriggerType)

// ServerLobbyIpcType - Value-first version of sapphire.ServerLobbyIpcType
var ServerLobbyIpcType = reverseMap(sapphire.ServerLobbyIpcType)

// ClientLobbyIpcType - Value-first version of sapphire.ClientLobbyIpcType
var ClientLobbyIpcType = reverseMap(sapphire.ClientLobbyIpcType)

// ServerZoneIpcType - Value-first version of sapphire.ServerZoneIpcType
var ServerZoneIpcType = reverseMap(sapphire.ServerZoneIpcType)

// ClientZoneIpcType - Value-first version of sapphire.ClientZoneIpcType
var ClientZoneIpcType = reverseMap(sapphire.ClientZoneIpcType)

// ServerChatIpcType - Value-first version of sapphire.ServerChatIpcType
var ServerChatIpcType = reverseMap(sapphire.ServerChatIpcType)

// ClientChatIpcType - Value-first version of sapphire.ClientChatIpcType
var ClientChatIpcType = reverseMap(sapphire.ClientChatIpcType)

// ServerLobbyIpcTypeCN - Value-first version of sapphire.ServerLobbyIpcTypeCN
var ServerLobbyIpcTypeCN = reverseMap(sapphire.ServerLobbyIpcTypeCN)

// ClientLobbyIpcTypeCN - Value-first version of sapphire.ClientLobbyIpcTypeCN
var ClientLobbyIpcTypeCN = reverseMap(sapphire.ClientLobbyIpcTypeCN)

// ServerZoneIpcTypeCN - Value-first version of sapphire.ServerZoneIpcTypeCN
var ServerZoneIpcTypeCN = reverseMap(sapphire.ServerZoneIpcTypeCN)

// ClientZoneIpcTypeCN - Value-first version of sapphire.ClientZoneIpcTypeCN
var ClientZoneIpcTypeCN = reverseMap(sapphire.ClientZoneIpcTypeCN)

// ServerChatIpcTypeCN - Value-first version of sapphire.ServerChatIpcTypeCN
var ServerChatIpcTypeCN = reverseMap(sapphire.ServerChatIpcTypeCN)

// ClientChatIpcTypeCN - Value-first version of sapphire.ClientChatIpcTypeCN
var ClientChatIpcTypeCN = reverseMap(sapphire.ClientChatIpcTypeCN)

// ServerLobbyIpcTypeKR - Value-first version of sapphire.ServerLobbyIpcTypeKR
var ServerLobbyIpcTypeKR = reverseMap(sapphire.ServerLobbyIpcTypeKR)

// ClientLobbyIpcTypeKR - Value-first version of sapphire.ClientLobbyIpcTypeKR
var ClientLobbyIpcTypeKR = reverseMap(sapphire.ClientLobbyIpcTypeKR)

// ServerZoneIpcTypeKR - Value-first version of sapphire.ServerZoneIpcTypeKR
var ServerZoneIpcTypeKR = reverseMap(sapphire.ServerZoneIpcTypeKR)

// ClientZoneIpcTypeKR - Value-first version of sapphire.ClientZoneIpcTypeKR
var ClientZoneIpcTypeKR = reverseMap(sapphire.ClientZoneIpcTypeKR)

// ServerChatIpcTypeKR - Value-first version of sapphire.ServerChatIpcTypeKR
var ServerChatIpcTypeKR = reverseMap(sapphire.ServerChatIpcTypeKR)

// ClientChatIpcTypeKR - Value-first version of sapphire.ClientChatIpcTypeKR
var ClientChatIpcTypeKR = reverseMap(sapphire.ClientChatIpcTypeKR)

func reverseMap(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}
