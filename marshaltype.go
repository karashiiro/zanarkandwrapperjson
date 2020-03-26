package main

import (
	"bytes"
	"encoding/binary"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

// marshalType - Marshal an []byte to a packet structure
func marshalType(packetType *string, data *[]byte, isDirectionEgress *bool) interface{} {
	var generic interface{}
	if *isDirectionEgress {
		generic = getTypeEgress(*packetType)
	} else {
		generic = getTypeIngress(*packetType)
	}

	buf := bytes.NewReader(*data)
	if generic != new(interface{}) {
		binary.Read(buf, binary.LittleEndian, generic)
	}

	return &generic
}

func getTypeIngress(packetType string) interface{} {
	switch packetType {
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

func getTypeEgress(packetType string) interface{} {
	switch packetType {
	case "InventoryModifyHandler":
		return new(sapphire.InventoryModifyHandler)
	}

	return new(interface{})
}
