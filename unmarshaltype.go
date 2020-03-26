package main

import (
	"bytes"
	"encoding/binary"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

// UnmarshalType - Unmarshal an []byte to a packet structure
func (ipcStructure *IpcStructure) UnmarshalType() interface{} {
	var generic interface{}
	if ipcStructure.IsEgressMessage {
		generic = ipcStructure.GetTypeEgress()
	} else {
		generic = ipcStructure.GetTypeIngress()
	}

	buf := bytes.NewReader(ipcStructure.Body)
	if generic != new(interface{}) {
		binary.Read(buf, binary.LittleEndian, generic)
	}

	return &generic
}

// GetTypeIngress returns an instance of the struct in the []byte of this package for inbound packets.
func (ipcStructure *IpcStructure) GetTypeIngress() interface{} {
	switch ipcStructure.Type {
	// ServerZoneDef
	case "ActorControl":
		return new(sapphire.ActorControl)
	case "ActorControlSelf":
		return new(sapphire.ActorControlSelf)
	case "CurrencyCrystalInfo":
		return new(sapphire.CurrencyCrystalInfo)
	case "EffectResult":
		return new(sapphire.EffectResult)
	case "EventFinish":
		return new(sapphire.EventFinish)
	case "EventPlay":
		return new(sapphire.EventPlay)
	case "EventPlay4":
		return new(sapphire.EventPlay4)
	case "EventStart":
		return new(sapphire.EventStart)
	case "InitZone":
		return new(sapphire.InitZone)
	case "InventoryTransaction":
		return new(sapphire.InventoryTransaction)
	case "ItemInfo":
		return new(sapphire.ItemInfo)
	case "MarketBoardItemListing":
		return new(sapphire.MarketBoardItemListing)
	case "MarketBoardItemListingCount":
		return new(sapphire.MarketBoardItemListingCount)
	case "MarketBoardItemListingHistory":
		return new(sapphire.MarketBoardItemListingHistory)
	case "MarketBoardSearchResult":
		return new(sapphire.MarketBoardSearchResult)
	case "MarketTaxRates":
		return new(sapphire.MarketTaxRates)
	case "NpcSpawn":
		return new(sapphire.NpcSpawn)
	case "PlayerSetup":
		return new(sapphire.PlayerSetup)
	case "PlayerSpawn":
		return new(sapphire.PlayerSpawn)
	case "PlayerStats":
		return new(sapphire.PlayerStats)
	case "RetainerInformation":
		return new(sapphire.RetainerInformation)
	case "SomeDirectorUnk4":
		return new(sapphire.SomeDirectorUnk4)
	case "UpdateClassInfo":
		return new(sapphire.UpdateClassInfo)
	case "UpdateInventorySlot":
		return new(sapphire.UpdateInventorySlot)
	case "WeatherChange":
		return new(sapphire.WeatherChange)
	// ServerLobbyDef
	case "LobbyRetainerList":
		return new(sapphire.LobbyRetainerList)
	case "LobbyServiceAccountList":
		return new(sapphire.LobbyServiceAccountList)
	case "LobbyServerList":
		return new(sapphire.LobbyServerList)
	}

	return new(interface{})
}

// GetTypeEgress returns an instance of the struct in the []byte of this package for outbound packets.
func (ipcStructure *IpcStructure) GetTypeEgress() interface{} {
	switch ipcStructure.Type {
	case "InventoryModifyHandler":
		return new(sapphire.InventoryModifyHandler)
	}

	return new(interface{})
}
