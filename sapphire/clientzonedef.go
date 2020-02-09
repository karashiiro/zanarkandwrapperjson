package sapphire

type GmCommand1 struct {
	/* 0000 */ commandId uint32
	/* 0004 */ param1 uint32
	/* 0008 */ param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ param4 uint32
	/* 0014 */ unknown1 uint32
	/* 0018 */ target uint32
}

type GmCommand2 struct {
	/* 0000 */ commandId uint32
	/* 0004 */ param1 uint32
	/* 0008 */ param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ param4 uint32
	/* 0014 */ worldId uint16
	/* 0016 */ target [0x20]byte
	/* 0036 */ unknown1 uint16
}

type ClientTrigger struct {
	/* 0000 */ commandId uint16
	/* 0002 */ unk_2 [2]uint8
	/* 0004 */ param11 uint32
	/* 0008 */ param12 uint32
	/* 000C */ param2 uint32
	/* 0010 */ param4 uint32 // todo: really?
	/* 0014 */ param5 uint32
	/* 0018 */ param3 uint64
}

type UpdatePosition struct {
	/* 0000 */ rotation float32
	/* 0004 */ animationType uint8
	/* 0005 */ animationState uint8
	/* 0006 */ clientAnimationType uint8
	/* 0007 */ headPosition uint8
	/* 0008 */ position FFXIVARR_POSITION3
	/* 000C */ unk [4]uint8 // padding?
}

type SkillHandler struct {
	/* 0000 */ pad_0000 uint8
	/* 0001 */ skillType uint8 // Note: Changed "type" to "skillType"
	/* 0002 */ pad_0002 [2]uint8
	/* 0004 */ actionId uint32
	/* 0008 */ sequence uint16
	/* 000A */ pad_000C [6]uint8
	/* 0010 */ targetId uint64
	/* 0018 */ itemSourceSlot uint16
	/* 001A */ itemSourceContainer uint16
	/* 001C */ unknown uint32
}

type AoESkillHandler struct {
	/* 0000 */ pad_0000 uint8
	/* 0001 */ skillType uint8 // Note: Changed "type" to "skillType"
	/* 0002 */ pad_0002 [2]uint8
	/* 0004 */ actionId uint32
	/* 0008 */ sequence uint16
	/* 000A */ pad_000C [6]uint8
	/* 0010 */ pos FFXIVARR_POSITION3
	/* 001C */ unknown uint32 // could almost be rotation + 16 bits more padding?
}

type ZoneLineHandler struct {
	/* 0000 */ zoneLineId uint32
}

type DiscoveryHandler struct {
	/* 0000 */ positionRef uint32
}

type EventHandlerReturn struct {
	/* 0000 */ eventId uint32
	/* 0004 */ scene uint16
	/* 0006 */ param1 uint16
	/* 0008 */ param2 uint16
	/* 000A */ pad_000A [2]uint8
	/* 000C */ param3 uint16
	/* 000E */ pad_000E [2]uint8
	/* 0010 */ param4 uint16
}

type EnterTerritoryHandler struct {
	/* 0000 */ eventId uint32
	/* 0004 */ param1 uint16
	/* 0006 */ param2 uint16
}

type EventHandlerOutsideRange struct {
	/* 0000 */ param1 uint32
	/* 0004 */ eventId uint32
	/* 0008 */ position FFXIVARR_POSITION3
}

type EventHandlerWithinRange struct {
	/* 0000 */ param1 uint32
	/* 0004 */ eventId uint32
	/* 0008 */ position FFXIVARR_POSITION3
}

type EventHandlerEmote struct {
	/* 0000 */ actorId uint64
	/* 0008 */ eventId uint32
	/* 000C */ emoteId uint16
}

type EventHandlerTalk struct {
	/* 0000 */ actorId uint64
	/* 0008 */ eventId uint32
}

type PingHandler struct {
	/* 0000 */ timestamp uint32 // maybe lol..
}

type SetSearchInfoHandler struct {
	///* 0000 */ status uint64; // Using the more mangled part of the struct
	/* 0000 */ status1 uint32
	/* 0004 */ status2 uint32

	/* 0008 */
	pad_0008 [9]uint8
	/* 0011 */ language uint8
	/* 0012 */ searchComment [193]byte
}

type TellHandler struct {
	contentId uint64
	worldId   uint16
	u0A       uint16
	u0C       uint32
	worldId1  uint16
	preName   uint8
	/* 0004 */ targetPCName [32]byte
	/* 0024 */ message [1029]byte
}

type ChatHandler struct {
	/* 0000 */ pad_0000 [4]uint8
	/* 0004 */ sourceId uint32
	/* 0008 */ pad_0008 [16]uint8
	/* 0018 */ chatType uint16
	/* 001A */ message [1012]byte
}

type ShopEventHandler struct {
	/* 0000 */ eventId uint32
	/* 0004 */ param uint32
}

type LinkshellEventHandler struct {
	/* 0000 */ eventId uint32
	/* 0004 */ scene uint16
	/* 0006 */ pad_0006 [1]uint8
	/* 0007 */ lsName [21]byte
}

type InventoryModifyHandler struct {
	/* 0000 */ seq uint32
	/* 0004 */ action uint8
	/* 0005 */ pad_0005 [7]uint8
	/* 000C */ fromContainer uint16
	/* 000E */ pad_000E [2]uint8
	/* 0010 */ fromSlot uint8
	/* 0011 */ pad_0011 [15]uint8
	/* 0020 */ toContainer uint16
	/* 0022 */ pad_0022 [2]uint8
	/* 0024 */ toSlot uint8
	/* 0025 */ pad_0025 [3]uint8
	/* 0028 */ splitCount uint32
}

type RenameLandHandler struct {
	/* 0000 */ ident LandIdent
	/* 0008 */ houseName [20]byte
	/* 0028 */ padding uint32
}

type HousingUpdateHouseGreeting struct {
	/* 0000 */ ident LandIdent
	/* 0008 */ greeting [200]byte
}

type BuildPresetHandler struct {
	/* 0000 */ itemId uint32
	/* 0004 */ plotNum uint8
	/* 0005 */ stateString [27]byte
}

type SetSharedEstateSettings struct {
	/* 0000 */ char1ContentId uint64
	/* 0008 */ char2ContentId uint64
	/* 0010 */ char3ContentId uint64
	/* 0018 */ char1Permissions uint8
	/* 0019 */ padding1 [0x7]uint8
	/* 0020 */ char2Permissions uint8
	/* 0021 */ padding2 [0x7]uint8
	/* 0028 */ char3Permissions uint8
	/* 0029 */ padding3 [0x7]uint8
}

type MarketBoardRequestItemListings struct {
	/* 0000 */ padding1 uint16
	/* 0002 */ itemCatalogId uint16
	/* 0004 */ padding2 uint32
}

type ReqPlaceHousingItem struct {
	/* 0000 */ landId uint16 // 0 when plot 0 or inside an estate
	/* 0002 */ unknown1 uint16
	/* 0004 */ unknown2 uint32
	/* 0008 */ sourceInvContainerId uint16
	/* 000A */ sourceInvSlotId uint16

	/* 000C */
	position FFXIVARR_POSITION3
	/* 0018 */ rotation float32

	/* 001C */
	shouldPlaceItem uint32 // 1 if placing an item, 0 if placing in store
	/* 0020 */ unknown4 [2]uint32 // always 0 it looks like
}

type HousingUpdateObjectPosition struct {
	/* 0000 */ ident LandIdent
	/* 0008 */ slot uint16
	/* 000A */ unk uint16

	/* 000C */
	pos FFXIVARR_POSITION3
	/* 0018 */ rotation float32

	/* 001C */
	padding uint32
}

type MarketBoardSearch struct {
	/* 0000 */ startIdx uint32
	/* 0004 */ requestId uint16
	/* 0006 */ itemSearchCategory uint8
	/* 0007 */ shouldCheckClassJobId uint8 // wat? seems only 1 there at least...
	/* 0008 */ maxEquipLevel uint8
	/* 0009 */ classJobId uint8
	/* 000A */ searchStr [40]byte
	/* 0032 */ unk4 [43]uint16
}

type MarketBoardRequestItemListingInfo struct {
	/* 0000 */ catalogId uint32
	/* 0000 */ requestId uint32
}

type FreeCompanyUpdateShortMessageHandler struct {
	shortMessage [104]byte
	padding      uint8
	unknown      uint8
	unknown1     uint32
	unknown2     uint16
}
