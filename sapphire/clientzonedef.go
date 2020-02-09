package sapphire

type GmCommand1 struct {
	/* 0000 */ CommandId uint32
	/* 0004 */ Param1 uint32
	/* 0008 */ Param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ Param4 uint32
	/* 0014 */ Unknown1 uint32
	/* 0018 */ Target uint32
}

type GmCommand2 struct {
	/* 0000 */ CommandId uint32
	/* 0004 */ Param1 uint32
	/* 0008 */ Param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ Param4 uint32
	/* 0014 */ WorldId uint16
	/* 0016 */ Target [0x20]byte
	/* 0036 */ Unknown1 uint16
}

type ClientTrigger struct {
	/* 0000 */ CommandId uint16
	/* 0002 */ Unk_2 [2]uint8
	/* 0004 */ Param11 uint32
	/* 0008 */ Param12 uint32
	/* 000C */ param2 uint32
	/* 0010 */ Param4 uint32 // todo: really?
	/* 0014 */ Param5 uint32
	/* 0018 */ Param3 uint64
}

type UpdatePosition struct {
	/* 0000 */ Rotation float32
	/* 0004 */ AnimationType uint8
	/* 0005 */ AnimationState uint8
	/* 0006 */ ClientAnimationType uint8
	/* 0007 */ HeadPosition uint8
	/* 0008 */ Position FFXIVARR_POSITION3
	/* 000C */ unk [4]uint8 // padding?
}

type SkillHandler struct {
	/* 0000 */ Pad_0000 uint8
	/* 0001 */ SkillType uint8 // Note: Changed "type" to "skillType"
	/* 0002 */ Pad_0002 [2]uint8
	/* 0004 */ ActionId uint32
	/* 0008 */ Sequence uint16
	/* 000A */ pad_000C [6]uint8
	/* 0010 */ TargetId uint64
	/* 0018 */ ItemSourceSlot uint16
	/* 001A */ itemSourceContainer uint16
	/* 001C */ unknown uint32
}

type AoESkillHandler struct {
	/* 0000 */ Pad_0000 uint8
	/* 0001 */ SkillType uint8 // Note: Changed "type" to "skillType"
	/* 0002 */ Pad_0002 [2]uint8
	/* 0004 */ ActionId uint32
	/* 0008 */ Sequence uint16
	/* 000A */ pad_000C [6]uint8
	/* 0010 */ Pos FFXIVARR_POSITION3
	/* 001C */ unknown uint32 // could almost be rotation + 16 bits more padding?
}

type ZoneLineHandler struct {
	/* 0000 */ ZoneLineId uint32
}

type DiscoveryHandler struct {
	/* 0000 */ PositionRef uint32
}

type EventHandlerReturn struct {
	/* 0000 */ EventId uint32
	/* 0004 */ Scene uint16
	/* 0006 */ Param1 uint16
	/* 0008 */ Param2 uint16
	/* 000A */ pad_000A [2]uint8
	/* 000C */ param3 uint16
	/* 000E */ pad_000E [2]uint8
	/* 0010 */ Param4 uint16
}

type EnterTerritoryHandler struct {
	/* 0000 */ EventId uint32
	/* 0004 */ Param1 uint16
	/* 0006 */ Param2 uint16
}

type EventHandlerOutsideRange struct {
	/* 0000 */ Param1 uint32
	/* 0004 */ EventId uint32
	/* 0008 */ Position FFXIVARR_POSITION3
}

type EventHandlerWithinRange struct {
	/* 0000 */ Param1 uint32
	/* 0004 */ EventId uint32
	/* 0008 */ Position FFXIVARR_POSITION3
}

type EventHandlerEmote struct {
	/* 0000 */ ActorId uint64
	/* 0008 */ EventId uint32
	/* 000C */ emoteId uint16
}

type EventHandlerTalk struct {
	/* 0000 */ ActorId uint64
	/* 0008 */ EventId uint32
}

type PingHandler struct {
	/* 0000 */ Timestamp uint32 // maybe lol..
}

type SetSearchInfoHandler struct {
	///* 0000 */ Status uint64; // Using the more mangled part of the struct
	/* 0000 */ Status1 uint32
	/* 0004 */ Status2 uint32

	/* 0008 */
	Pad_0008 [9]uint8
	/* 0011 */ Language uint8
	/* 0012 */ SearchComment [193]byte
}

type TellHandler struct {
	ContentId uint64
	WorldId   uint16
	U0A       uint16
	U0C       uint32
	WorldId1  uint16
	PreName   uint8
	/* 0004 */ TargetPCName [32]byte
	/* 0024 */ Message [1029]byte
}

type ChatHandler struct {
	/* 0000 */ Pad_0000 [4]uint8
	/* 0004 */ SourceId uint32
	/* 0008 */ Pad_0008 [16]uint8
	/* 0018 */ ChatType uint16
	/* 001A */ message [1012]byte
}

type ShopEventHandler struct {
	/* 0000 */ EventId uint32
	/* 0004 */ Param uint32
}

type LinkshellEventHandler struct {
	/* 0000 */ EventId uint32
	/* 0004 */ Scene uint16
	/* 0006 */ Pad_0006 [1]uint8
	/* 0007 */ LsName [21]byte
}

type InventoryModifyHandler struct {
	/* 0000 */ Seq uint32
	/* 0004 */ Action uint8
	/* 0005 */ Pad_0005 [7]uint8
	/* 000C */ fromContainer uint16
	/* 000E */ pad_000E [2]uint8
	/* 0010 */ FromSlot uint8
	/* 0011 */ Pad_0011 [15]uint8
	/* 0020 */ ToContainer uint16
	/* 0022 */ Pad_0022 [2]uint8
	/* 0024 */ ToSlot uint8
	/* 0025 */ Pad_0025 [3]uint8
	/* 0028 */ SplitCount uint32
}

type RenameLandHandler struct {
	/* 0000 */ Ident LandIdent
	/* 0008 */ HouseName [20]byte
	/* 0028 */ Padding uint32
}

type HousingUpdateHouseGreeting struct {
	/* 0000 */ Ident LandIdent
	/* 0008 */ Greeting [200]byte
}

type BuildPresetHandler struct {
	/* 0000 */ ItemId uint32
	/* 0004 */ PlotNum uint8
	/* 0005 */ StateString [27]byte
}

type SetSharedEstateSettings struct {
	/* 0000 */ Char1ContentId uint64
	/* 0008 */ Char2ContentId uint64
	/* 0010 */ Char3ContentId uint64
	/* 0018 */ Char1Permissions uint8
	/* 0019 */ Padding1 [0x7]uint8
	/* 0020 */ Char2Permissions uint8
	/* 0021 */ Padding2 [0x7]uint8
	/* 0028 */ Char3Permissions uint8
	/* 0029 */ Padding3 [0x7]uint8
}

type MarketBoardRequestItemListings struct {
	/* 0000 */ Padding1 uint16
	/* 0002 */ ItemCatalogId uint16
	/* 0004 */ Padding2 uint32
}

type ReqPlaceHousingItem struct {
	/* 0000 */ LandId uint16 // 0 when plot 0 or inside an estate
	/* 0002 */ Unknown1 uint16
	/* 0004 */ Unknown2 uint32
	/* 0008 */ SourceInvContainerId uint16
	/* 000A */ sourceInvSlotId uint16

	/* 000C */
	Position FFXIVARR_POSITION3
	/* 0018 */ Rotation float32

	/* 001C */
	ShouldPlaceItem uint32 // 1 if placing an item, 0 if placing in store
	/* 0020 */ Unknown4 [2]uint32 // always 0 it looks like
}

type HousingUpdateObjectPosition struct {
	/* 0000 */ Ident LandIdent
	/* 0008 */ Slot uint16
	/* 000A */ unk uint16

	/* 000C */
	Pos FFXIVARR_POSITION3
	/* 0018 */ Rotation float32

	/* 001C */
	Padding uint32
}

type MarketBoardSearch struct {
	/* 0000 */ StartIdx uint32
	/* 0004 */ RequestId uint16
	/* 0006 */ ItemSearchCategory uint8
	/* 0007 */ ShouldCheckClassJobId uint8 // wat? seems only 1 there at least...
	/* 0008 */ MaxEquipLevel uint8
	/* 0009 */ ClassJobId uint8
	/* 000A */ searchStr [40]byte
	/* 0032 */ Unk4 [43]uint16
}

type MarketBoardRequestItemListingInfo struct {
	/* 0000 */ CatalogId uint32
	/* 0000 */ RequestId uint32
}

type FreeCompanyUpdateShortMessageHandler struct {
	ShortMessage [104]byte
	Padding      uint8
	Unknown      uint8
	Unknown1     uint32
	Unknown2     uint16
}
