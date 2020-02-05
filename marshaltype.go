package main

import (
	"bytes"
	"encoding/binary"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

// marshalType - Marshal an []byte to a packet structure
func marshalType(packetType string, data []byte) interface{} {
	generic := towerOfBabelSwitchEdition(packetType)

	buf := bytes.NewReader(data)
	binary.Read(buf, binary.LittleEndian, generic)

	return generic
}

// I call it this, but it's no VVVVV I guess
func towerOfBabelSwitchEdition(packetType string) interface{} {
	switch packetType {
	// ServerZoneDef.h
	case "Ping":
		return new(sapphire.Ping)
	case "Init":
		return new(sapphire.Init)
	// ClientZoneDef.h
	case "GmCommand1":
		return new(sapphire.GmCommand1)
	case "GmCommand2":
		return new(sapphire.GmCommand2)
	case "ClientTrigger":
		return new(sapphire.ClientTrigger)
	case "UpdatePosition":
		return new(sapphire.UpdatePosition)
	case "SkillHandler":
		return new(sapphire.SkillHandler)
	case "AoESkillHandler":
		return new(sapphire.AoESkillHandler)
	case "ZoneLineHandler":
		return new(sapphire.ZoneLineHandler)
	case "DiscoveryHandler":
		return new(sapphire.DiscoveryHandler)
	case "EventHandlerReturn":
		return new(sapphire.EventHandlerReturn)
	case "EnterTerritoryHandler":
		return new(sapphire.EnterTerritoryHandler)
	case "EventHandlerOutsideRange":
		return new(sapphire.EventHandlerOutsideRange)
	case "EventHandlerWithinRange":
		return new(sapphire.EventHandlerWithinRange)
	case "EventHandlerEmote":
		return new(sapphire.EventHandlerEmote)
	case "EventHandlerTalk":
		return new(sapphire.EventHandlerTalk)
	case "PingHandler":
		return new(sapphire.PingHandler)
	case "SetSearchInfo":
		return new(sapphire.SetSearchInfo)
	case "TellHandler":
		return new(sapphire.TellHandler)
	case "ChatHandler":
		return new(sapphire.ChatHandler)
	case "ShopEventHandler":
		return new(sapphire.ShopEventHandler)
	case "LinkshellEventHandler":
		return new(sapphire.LinkshellEventHandler)
	case "InventoryModifyHandler":
		return new(sapphire.InventoryModifyHandler)
	case "RenameLandHandler":
		return new(sapphire.RenameLandHandler)
	case "HousingUpdateHouseGreeting":
		return new(sapphire.HousingUpdateHouseGreeting)
	case "BuildPresetHandler":
		return new(sapphire.BuildPresetHandler)
	case "SetSharedEstateSettings":
		return new(sapphire.SetSharedEstateSettings)
	case "MarketBoardRequestItemListings":
		return new(sapphire.MarketBoardRequestItemListings)
	case "ReqPlaceHousingItem":
		return new(sapphire.ReqPlaceHousingItem)
	case "HousingUpdateObjectPosition":
		return new(sapphire.HousingUpdateObjectPosition)
	case "MarketBoardSearch":
		return new(sapphire.MarketBoardSearch)
	case "MarketBoardRequestItemListingInfo":
		return new(sapphire.MarketBoardRequestItemListingInfo)
	case "FreeCompanyUpdateShortMessageHandler":
		return new(sapphire.FreeCompanyUpdateShortMessageHandler)
	// ServerLobbyDef.h
	case "RetainerList":
		return new(sapphire.RetainerList)
	case "ServiceIdInfo":
		return new(sapphire.ServiceIdInfo)
	case "ServerList":
		return new(sapphire.ServerList)
	case "CharList":
		return new(sapphire.CharList)
	case "EnterWorld":
		return new(sapphire.EnterWorld)
	case "CharCreate":
		return new(sapphire.CharCreate)
	case "LobbyError":
		return new(sapphire.LobbyError)
	// ClientChatDef.h
	case "Tell":
		return new(sapphire.Tell)
	case "TellErrNotFound":
		return new(sapphire.TellErrNotFound)
	case "FreeCompanyEvent":
		return new(sapphire.FreeCompanyEvent)
	}

	return new(interface{})
}
