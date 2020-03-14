package sapphire

/**
* Structural representation of the packet sent by the server as response
* To a ping packet
 */
type Ping struct {
	/* 0000 */ TimeInMilliseconds uint64
	/* 0008 */ Unknown_8 [0x38]uint8
}

/**
* Structural representation of the packet sent by the server as response
* To a ping packet
 */
type Init struct {
	Unknown  uint64
	CharId   uint32
	Unknown1 uint32
}

/**
* Structural representation of the packet sent by the server
* Carrying chat messages
 */
type Chat struct {
	/* 0000 */ Padding [14]uint8 //Maybe this is SubCode, or some kind of talker ID...
	ChatType           uint16
	Name               [32]byte
	Msg                [1012]byte
}

type ChatBanned struct {
	Padding [4]uint8 // I was not sure reinterpreting ZST is valid behavior in C++.
	// Client doesn't care about the data (zero sized) for this opcode anyway.
}

/**
* Structural representation of the packet sent by the server
* To show a list of worlds for world visit
 */
type WorldVisitList struct {
	World [16]struct {
		Id     uint16 // this is the id of the world from lobby
		Status uint16 // 1 = available (this is what retail sends) | 2+ = unavailable (this will need to be checked with retail if it's exactly 2 or not since it does not actually lock the option)
	}
}

/**
* Structural representation of the packet sent by the server
* Carrying chat messages
 */
type Logout struct {
	Flags1 uint32
	Flags2 uint32
}

/**
* Structural representation of the packet sent by the server
* Sent to show the play time
 */
type PlayTime struct {
	PlayTimeInMinutes uint32
	Padding           uint32
}

/**
* Structural representation of the packet sent by the server
* With a list of players ( party list | friend list | search results )
 */
type PlayerEntry struct {
	ContentId        uint64
	Bytes            [12]uint8
	ZoneId           uint16
	ZoneId1          uint16
	Bytes1           [8]byte
	OnlineStatusMask uint64
	ClassJob         uint8
	Padding          uint8
	Level            uint8
	Padding1         uint8
	Padding2         uint16
	One              uint8
	Name             [0x20]byte
	FcTag            [9]byte
}

type SocialList struct {
	Padding    uint32
	Padding1   uint32
	Padding2   uint32
	SocialType uint8 // Changed name from "type" to "socialType"
	Sequence   uint8
	Padding3   uint16

	Entries [10]PlayerEntry
}

type ExamineSearchInfo struct {
	Unknown       uint32
	Unknown1      uint16
	Unknown2      uint16
	Padding       [16]byte
	Unknown3      uint32
	Unknown4      uint16
	Unknown5      uint16
	Unknown6      uint16
	WorldId       uint8
	SearchMessage [193]byte
	FcName        [24]byte
	Unknown7      uint8
	Padding1      uint16
	LevelEntries  [CLASSJOB_TOTAL]struct {
		Id    uint16
		Level uint16
	}
}

type SetSearchInfo struct {
	OnlineStatusFlags uint64
	Unknown           uint64
	Unknown1          uint32
	Padding           uint8
	SelectRegion      uint8
	SearchMessage     [193]byte
	Padding2          uint8
}

type InitSearchInfo struct {
	OnlineStatusFlags uint64
	Unknown           uint64
	Unknown1          uint8
	SelectRegion      uint8
	SearchMessage     [193]byte
	Padding           [5]byte
}

type ExamineSearchComment struct {
	CharId uint32
	// Packet only has 196 bytes after the charid
	// Likely utf8
	SearchComment [195]byte
	Padding       byte
}

/**
* Structural representation of the packet sent by the server
* To display a server notice message
 */
type ServerNoticeShort struct {
	// These are actually display flags
	/* 0000 */ Padding uint8
	// 0 = Chat log
	// 2 = Nothing
	// 4 = On screen message
	// 5 = On screen message + chat log
	Message [538]byte
}

/**
* Structural representation of the packet sent by the server
* To display a server notice message
 */
type ServerNotice struct {
	// These are actually display flags
	/* 0000 */ Padding uint8
	// 0 = Chat log
	// 2 = Nothing
	// 4 = On screen message
	// 5 = On screen message + chat log
	Message [775]byte
}

type SetOnlineStatus struct {
	OnlineStatusFlags uint64
}

type BlackList struct {
	Entry [20]struct {
		ContentId uint64
		Name      [32]byte
	}
	Padding  uint8
	Padding1 uint8
	Sequence uint16
	Padding2 uint32
}

type LogMessage struct {
	Field_0    uint32
	Field_4    uint32
	Field_8    uint32
	Field_12   uint32
	Category   uint32
	LogMessage uint32
	Field_24   uint8
	Field_25   uint8
	Field_26   [32]uint8
	Field_58   uint32
}

type LinkshellList struct {
	Entry [8]struct {
		LsId      uint64
		UnknownId uint64
		Unknown   uint8
		Rank      uint8
		Padding   uint16
		LsName    [20]uint8
		Unk       [16]uint8
	}
}

/**
* Structural representation of the packet sent by the server
* To send a list of mail the player has
 */
type ReqMoogleMailList struct {
	Letter [5]struct {
		Unk        [0x8]byte
		TimeStamp  uint32     // The time the mail was sent (this also seems to be used as a Id)
		Unk1       [0x30]byte // This should be items, gil, etc for the letter
		Read       bool       // 0 = false | 1 = true
		LetterType uint8      // 0 = Friends | 1 = Rewards | 2 = GM // Changed from "type" to "letterType"
		Unk2       uint8
		SenderName [0x20]byte // The name of the sender
		Summary    [0x3C]byte // The start of the full letter text
		Padding2   [0x5]byte
	}
	Unk3 [0x08]byte
}

/**
* Structural representation of the packet sent by the server
* To show the mail delivery notification
 */
type MailLetterNotification struct {
	SendbackCount uint32    // The amount of letters sent back since you ran out of room (moogle dialog changes based on this)
	FriendLetters uint16    // The amount of letters in the friends section of the letterbox
	UnreadCount   uint16    // The amount of unreads in the letterbox (this is the number that shows up)
	RewardLetters uint16    // The amount of letters in the rewards section of the letterbox
	IsGmLetter    uint8     // Makes the letter notification flash red
	IsSupportDesk uint8     // After setting this to 1 we can no longer update mail notifications (more research needed on the support desk)
	Unk2          [0x4]byte // This has probs something to do with the support desk (inquiry id?)
}

type MarketTaxRates struct {
	Unknown1 uint32
	Padding1 uint16
	Unknown2 uint16
	TaxRate  [TOWN_COUNT]uint32 // In the order of Common::Town
	Unknown3 uint64
}

type MarketBoardItemListingCount struct {
	ItemCatalogId uint32
	Unknown1      uint32 // does some shit if nonzero
	RequestId     uint16
	Quantity      uint16 // high/low u8s read separately?
	Unknown3      uint32
}

type MarketBoardItemListing struct {
	Listing [10]struct // 152 bytes each
	{
		ListingId       uint64
		RetainerId      uint64
		RetainerOwnerId uint64
		ArtisanId       uint64
		PricePerUnit    uint32
		TotalTax        uint32
		ItemQuantity    uint32
		ItemId          uint32
		LastReviewTime  uint16
		ContainerId     uint16
		SlotId          uint32
		Durability      uint16
		SpiritBond      uint16
		/**
		* Auto materiaId = (i & 0xFF0) >> 4;
		* Auto index = i & 0xF;
		* Auto leftover = i >> 8;
		 */
		MateriaValue [5]uint16
		Padding1     uint16
		Padding2     uint32
		RetainerName [32]byte
		PlayerName   [32]byte
		Hq           bool
		MateriaCount uint8
		OnMannequin  uint8
		MarketCity   uint8
		DyeId        uint16
		Padding3     uint16
		Padding4     uint32
	} // Multiple packets are sent if there are more than 10 search results.
	ListingIndexEnd   uint8
	ListingIndexStart uint8
	RequestId         uint16
	Padding7          [16]byte
	Unknown13         uint8
	Padding8          uint16
	Unknown14         uint8
	Padding9          uint64
	Unknown15         uint32
	Padding10         uint32
}

type MarketBoardItemListingHistory struct {
	ItemCatalogId  uint32
	ItemCatalogId2 uint32

	Listing [20]struct {
		SalePrice    uint32
		PurchaseTime uint32
		Quantity     uint32
		IsHq         uint8
		Padding      uint8
		OnMannequin  uint8

		BuyerName [33]byte

		ItemCatalogId uint32
	}
}

type MarketBoardSearchResult struct {
	Items [20]struct {
		ItemCatalogId uint32
		Quantity      uint16
		Demand        uint16
	}

	ItemIndexEnd   uint32
	Padding1       uint32
	ItemIndexStart uint32
	RequestId      uint32
}

type ExamineFreeCompanyInfo struct {
	Unknown         [0x20]byte // likely fc allegiance/icon/housing info etc
	CharId          uint32
	FcTimeCreated   uint32
	Unknown2        [0x10]byte
	Unknown3        uint16
	FcName          [0x14]byte // 20 limit
	Padding         uint16
	FcTag           [0x05]byte // 5 tag limit
	Padding2        uint16     // null terminator?
	FcLeader        [0x20]byte // leader name (32 bytes)
	FcSlogan        [192]byte  // source: https://ffxiv.gamerescape.com/wiki/Free_Company (packet cap confirms this size also)
	Padding3        byte       // null terminator?
	FcEstateProfile [20]byte   // todo: size needs confirmation
	Padding4        uint32
}

type FreeCompanyUpdateShortMessage struct {
	Unknown      uint32
	Unknown1     uint16
	Unknown2     uint16
	Unknown3     uint32
	Unknown5     uint32
	ShortMessage [104]byte
}

type StatusEffectList struct {
	ClassId          uint8
	Level1           uint8
	Level            uint16
	Current_hp       uint32
	Max_hp           uint32
	Current_mp       uint16
	Max_mp           uint16
	CurrentTp        uint16
	ShieldPercentage uint8
	Unknown1         uint8
	Effect           [30]StatusEffect
	Padding          uint32
}

type FFXIVGCAffiliation struct {
	GcId   uint8
	GcRank [3]uint8
}

/**
* Structural representation of the packet sent by the server
* Add a status effect
 */
type EffectResult struct {
	GlobalSequence   uint32
	Actor_id         uint32
	Current_hp       uint32
	Max_hp           uint32
	Current_mp       uint16
	Current_tp       uint16
	Max_mp           uint16
	Unknown1         uint8
	ClassId          uint8
	ShieldPercentage uint8
	EntryCount       uint8
	Unknown2         uint16

	StatusEntries [4]struct {
		Index         uint8 // which position do i display this
		Unknown3      uint8
		Id            uint16
		Param         uint16
		Unknown4      uint16 // Sort this out (old right half of power/param property)
		Duration      float32
		SourceActorId uint32
	}

	Unknown5 uint32
}

/**
* Structural representation of the packet sent by the server
* To update certain player details / status
 */
type ActorControl struct {
	/* 0000 */ Category uint16
	/* 0002 */ Padding uint16
	/* 0004 */ Param1 uint32
	/* 0008 */ Param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ Param4 uint32
	/* 0014 */ Padding1 uint32
}

/**
* Structural representation of the packet sent by the server
* To update certain player details / status
 */
type ActorControlSelf struct {
	/* 0000 */ Category uint16
	/* 0002 */ Padding uint16
	/* 0004 */ Param1 uint32
	/* 0008 */ Param2 uint32
	/* 000C */ Param3 uint32
	/* 0010 */ Param4 uint32
	/* 0014 */ Param5 uint32
	/* 0018 */ Param6 uint32
	/* 001C */ Padding1 uint32
}

/**
* Structural representation of the packet sent by the server
* To update certain player details / status
 */
type ActorControlTarget struct {
	/* 0000 */ Category uint16
	/* 0002 */ Padding uint16
	/* 0004 */ Param1 uint32
	/* 0008 */ Param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ Param4 uint32
	/* 0014 */ Padding1 uint32
	/* 0018 */ TargetId uint64
}

/**
* Structural representation of the packet sent by the server
* To update HP / MP / TP
 */
type UpdateHpMpTp struct {
	/* 0000 */ Hp uint32
	/* 0004 */ Mp uint16
	/* 0006 */ Tp uint16
	/* 0008 */ Gp uint16
	/* 0010 */ Unknown_10 uint16
	/* 0012 */ Unknown_12 uint32
}

type Effect struct {
	AnimationTargetId uint64 // who the animation targets

	ActionId uint32 // what the casting player casts, shown in battle log/ui
	/*!
	* @Brief Zone sequence for the effect. Used to link effects that are split across multiple packets as one
	 */
	Sequence uint32

	AnimationLockTime float32 // maybe? doesn't seem to do anything
	SomeTargetId      uint32  // always 0x0E000000?

	/*!
	* @Brief The cast sequence from the originating player. Should only be sent to the source, 0 for every other player.
	*
	* This needs to match the sequence sent from the player in the action start packet otherwise you'll cancel the
	* Initial animation and start a new one once the packet arrives.
	 */
	SourceSequence    uint16
	Rotation          uint16
	ActionAnimationId uint16 // the animation that is played by the casting character
	Variation         uint8  // variation in the animation
	EffectDisplayType uint8

	Unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	EffectCount uint8 // ignores effects if 0, otherwise parses all of them
	Padding_21  uint16

	Padding_22 [3]uint16

	Effects [8 * 8]uint8

	Padding_6A [3]uint16

	EffectTargetId uint32 // who the effect targets
	EffectFlags    uint32 // nonzero = effects do nothing, no battle log, no ui text - only shows animations

	Padding_78 uint32
}

type AoeEffect8 struct {
	AnimationTargetId uint64 // who the animation targets

	ActionId       uint32 // what the casting player casts, shown in battle log/ui
	GlobalSequence uint32 // seems to only increment on retail?

	AnimationLockTime float32 // maybe? doesn't seem to do anything
	SomeTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	SourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	Rotation          uint16
	ActionAnimationId uint16 // the animation that is played by the casting character
	Variation         uint8  // variation in the animation
	EffectDisplayType uint8

	Unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	EffectCount uint8 // ignores effects if 0, otherwise parses all of them
	Padding_21  [3]uint16
	Padding     uint16

	Effects [8]struct {
		Entries [8]EffectEntry
	}

	Padding_6A [3]uint16

	EffectTargetId [8]uint64
	UnkFlag        [3]uint16 // all 0x7FFF
	Unk            [3]uint16
}

type AoeEffect16 struct {
	AnimationTargetId uint64 // who the animation targets

	ActionId       uint32 // what the casting player casts, shown in battle log/ui
	GlobalSequence uint32 // seems to only increment on retail?

	AnimationLockTime float32 // maybe? doesn't seem to do anything
	SomeTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	SourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	Rotation          uint16
	ActionAnimationId uint16 // the animation that is played by the casting character
	Variation         uint8  // variation in the animation
	EffectDisplayType uint8

	Unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	EffectCount uint8 // ignores effects if 0, otherwise parses all of them
	Padding_21  [3]uint16
	Padding     uint16

	Effects [16]struct {
		Entries [8]EffectEntry
	}

	Padding_6A [3]uint16

	EffectTargetId [16]uint64
	UnkFlag        [3]uint16 // all 0x7FFF
	Unk            [3]uint16
}

type AoeEffect24 struct {
	AnimationTargetId uint64 // who the animation targets

	ActionId       uint32 // what the casting player casts, shown in battle log/ui
	GlobalSequence uint32 // seems to only increment on retail?

	AnimationLockTime float32 // maybe? doesn't seem to do anything
	SomeTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	SourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	Rotation          uint16
	ActionAnimationId uint16 // the animation that is played by the casting character
	Variation         uint8  // variation in the animation
	EffectDisplayType uint8

	Unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	EffectCount uint8 // ignores effects if 0, otherwise parses all of them
	Padding_21  [3]uint16
	Padding     uint16

	Effects [24]struct {
		Entries [8]EffectEntry
	}

	Padding_6A [3]uint16

	EffectTargetId [24]uint64
	UnkFlag        [3]uint16 // all 0x7FFF
	Unk            [3]uint16
}

type AoeEffect32 struct {
	AnimationTargetId uint64 // who the animation targets

	ActionId       uint32 // what the casting player casts, shown in battle log/ui
	GlobalSequence uint32 // seems to only increment on retail?

	AnimationLockTime float32 // maybe? doesn't seem to do anything
	SomeTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	SourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	Rotation          uint16
	ActionAnimationId uint16 // the animation that is played by the casting character
	Variation         uint8  // variation in the animation
	EffectDisplayType uint8

	Unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	EffectCount uint8 // ignores effects if 0, otherwise parses all of them
	Padding_21  [3]uint16
	Padding     uint16

	Effects [32]struct {
		Entries [8]EffectEntry
	}

	Padding_6A [3]uint16

	EffectTargetId [32]uint64
	UnkFlag        [3]uint16 // all 0x7FFF
	Unk            [3]uint16
}

/**
* Structural representation of the packet sent by the server
* To spawn an actor
 */
type PlayerSpawn struct {
	Title          uint16
	U1b            uint16
	CurrentWorldId uint16
	HomeWorldId    uint16

	GmRank       uint8
	U3c          uint8
	U4           uint8
	OnlineStatus uint8

	Pose uint8
	U5a  uint8
	U5b  uint8
	U5c  uint8

	TargetId        uint64
	U6              uint32
	U7              uint32
	MainWeaponModel uint64
	SecWeaponModel  uint64
	CraftToolModel  uint64

	U14             uint32
	U15             uint32
	BNPCBase        uint32
	BNPCName        uint32
	U18             uint32
	U19             uint32
	DirectorId      uint32
	OwnerId         uint32
	U22             uint32
	HPMax           uint32
	HPCur           uint32
	DisplayFlags    uint32
	FateID          uint16
	MPCurr          uint16
	TPCurr          uint16
	MPMax           uint16
	TPMax           uint16
	ModelChara      uint16
	Rotation        uint16
	ActiveMinion    uint16
	SpawnIndex      uint8
	State           uint8
	PersistentEmote uint8
	ModelType       uint8
	Subtype         uint8
	Voice           uint8
	U25c            uint16
	EnemyType       uint8
	Level           uint8
	ClassJob        uint8
	U26d            uint8
	U27a            uint16
	CurrentMount    uint8
	MountHead       uint8
	MountBody       uint8
	MountFeet       uint8
	MountColor      uint8
	Scale           uint8
	ElementalLevel  uint32
	Element         uint32
	Effect          [30]StatusEffect
	Pos             FFXIVARR_POSITION3
	Models          [10]uint32
	Name            [32]byte
	Look            [26]uint8
	FcTag           [6]byte
	Unk30           uint32
}

/**
* Structural representation of the packet sent by the server
* To spawn an actor
 */
type NpcSpawn struct {
	GimmickId uint32 // needs to be existing in the map, mob will snap to it
	U2b       uint8
	U2ab      uint8
	GmRank    uint8
	U3b       uint8

	AggressionMode uint8 // 1 passive, 2 aggressive
	OnlineStatus   uint8
	U3c            uint8
	Pose           uint8

	U4 uint32

	TargetId        uint64
	U6              uint32
	U7              uint32
	MainWeaponModel uint64
	SecWeaponModel  uint64
	CraftToolModel  uint64

	U14             uint32
	U15             uint32
	BNPCBase        uint32
	BNPCName        uint32
	LevelId         uint32
	U19             uint32
	DirectorId      uint32
	SpawnerId       uint32
	ParentActorId   uint32
	HPMax           uint32
	HPCurr          uint32
	DisplayFlags    uint32
	FateID          uint16
	MPCurr          uint16
	TPCurr          uint16
	MPMax           uint16
	TPMax           uint16
	ModelChara      uint16
	Rotation        uint16
	ActiveMinion    uint16
	SpawnIndex      uint8
	State           uint8
	PersistantEmote uint8
	ModelType       uint8
	Subtype         uint8
	Voice           uint8
	U25c            uint16
	EnemyType       uint8
	Level           uint8
	ClassJob        uint8
	U26d            uint8
	U27a            uint8
	CurrentMount    uint8
	MountHead       uint8
	MountBody       uint8
	MountFeet       uint8
	MountColor      uint8
	Scale           uint8
	ElementalLevel  uint16 // Eureka
	Element         uint16 // Eureka
	U30b            uint32
	Effect          [30]StatusEffect
	Pos             FFXIVARR_POSITION3
	Models          [10]uint32
	Name            [32]byte
	Look            [26]uint8
	FcTag           [6]byte
	Unk30           uint32
	Unk31           uint32
	BNPCPartSlot    uint8
	Unk32           uint8
	Unk33           uint16
	Unk34           uint32
}

/**
* Structural representation of the packet sent by the server
* To show player movement
 */
type ActorFreeSpawn struct {
	SpawnId uint32
	ActorId uint32
}

/**
* Structural representation of the packet sent by the server
* To show player movement
 */
type ActorMove struct {
	/* 0000 */ HeadRotation uint8
	/* 0001 */ Rotation uint8
	/* 0002 */ AnimationType uint8
	/* 0003 */ AnimationState uint8
	/* 0004 */ AnimationSpeed uint8
	/* 0005 */ UnknownRotation uint8
	/* 0006 */ PosX uint16
	/* 0008 */ PosY uint16
	/* 000A */ posZ uint16
	/* 000C */ unknown_12 uint32
}

/**
* Structural representation of the packet sent by the server
* To set an actors position
 */
type ActorSetPos struct {
	R16         uint16
	WaitForLoad uint8
	Unknown1    uint8
	Unknown2    uint32
	X           float32
	Y           float32
	Z           float32
	Unknown3    uint32
}

/**
* Structural representation of the packet sent by the server
* To start an actors casting
 */
type ActorCast struct {
	Action_id uint16
	SkillType uint8
	Unknown   uint8
	Unknown_1 uint32 // action id or mount id
	Cast_time float32
	Target_id uint32
	Rotation  uint16
	Flag      uint16 // 1 = interruptible blinking cast bar
	Unknown_2 uint32
	PosX      uint16
	PosY      uint16
	PosZ      uint16
	Unknown_3 uint16
}

type HateList struct {
	NumEntries uint32
	Entry      [32]struct {
		ActorId     uint32
		HatePercent uint8
		Unknown     uint8
		Padding     uint16
	}
	Padding uint32
}

type HateRank struct {
	NumEntries uint32
	Entry      [32]struct {
		ActorId    uint32
		HateAmount uint32
	}
	Padding uint32
}

type UpdateClassInfo struct {
	ClassId        uint8
	Level1         uint8
	Level          uint16
	NextLevelIndex uint32
	CurrentExp     uint32
	RestedExp      uint32
}

/**
* Structural representation of the packet sent by the server
* To send the titles available to the player
 */
type PlayerTitleList struct {
	TitleList [48]uint8
}

/**
* Structural representation of the packet sent by the server
* To initialize a zone for the player
 */
type InitZone struct {
	ServerId                 uint16
	ZoneId                   uint16
	Unknown1                 uint16
	ContentfinderConditionId uint16
	Unknown3                 uint32
	Unknown4                 uint32
	WeatherId                uint8
	Bitmask                  uint8
	Bitmask1                 uint8
	// Bitmask1 findings
	//0 = Unknown ( 7B F8 69 )
	//1 = Show playguide window ( 7B 69 )
	//2 = Unknown ( 7B 69 )
	//4 = Disables record ready check ( 7B DF DF F8 F0 E4 110 (all sorts of social packets) )
	//8 = Hide server icon ( 7B 69 )
	//16 = Enable flight ( 7B F8 69 )
	//32 = Unknown ( 7B F8 69 )
	//64 = Unknown ( 7B F8 69 )
	//128 = Shows message "You are now in the instanced area XX A.
	//Current instance can be confirmed at any time using the /instance text command." ( 7B F8 69 )

	Unknown5             uint8
	Unknown8             uint32
	FestivalId           uint16
	AdditionalFestivalId uint16
	Unknown9             uint32
	Unknown10            uint32
	Unknown11            uint32
	Unknown12            [4]uint32
	Unknown13            [3]uint32
	Pos                  FFXIVARR_POSITION3
	Unknown14            [3]uint32
	Unknown15            uint32
}

/**
* Structural representation of the packet sent by the server to initialize
* The client UI upon initial connection.
 */
type PlayerSetup struct {
	// Plain C types for a bit until the packet is actually fixed.
	// Makes conversion between different editors easier.
	ContentId                       uint64
	Unknown8                        uint32
	UnknownC                        uint32
	CharId                          uint32
	RestedExp                       uint32
	CompanionCurrentExp             uint32
	Unknown1C                       uint32
	FishCaught                      uint32
	UseBaitCatalogId                uint32
	Unknown28                       uint32
	UnknownPvp2C                    uint16
	Unknown3                        uint16
	PvpFrontlineOverallCampaigns    uint32
	UnknownTimestamp34              uint32
	UnknownTimestamp38              uint32
	Unknown3C                       uint32
	Unknown40                       uint32
	Unknown44                       uint32
	CompanionTimePassed             float32
	Unknown4C                       uint32
	Unknown50                       uint16
	UnknownPvp52                    [4]uint16
	PlayerCommendations             uint16
	Unknown5C                       uint16
	Unknown5E                       uint16
	PvpFrontlineWeeklyCampaigns     uint16
	EnhancedAnimaGlassProgress      uint16
	Unknown64                       [4]uint16
	PvpRivalWingsTotalMatches       uint16
	PvpRivalWingsTotalVictories     uint16
	PvpRivalWingsWeeklyMatches      uint16
	PvpRivalWingsWeeklyVictories    uint16
	MaxLevel                        uint8
	Expansion                       uint8
	Unknown76                       uint8
	Unknown77                       uint8
	Race                            uint8
	Tribe                           uint8
	Gender                          uint8
	CurrentJob                      uint8
	CurrentClass                    uint8
	Deity                           uint8
	NamedayMonth                    uint8
	NamedayDay                      uint8
	CityState                       uint8
	Homepoint                       uint8
	Unknown82                       uint8
	PetHotBar                       uint8
	CompanionRank                   uint8
	CompanionStars                  uint8
	CompanionSp                     uint8
	CompanionUnk86                  uint8
	CompanionColor                  uint8
	CompanionFavoFeed               uint8
	Unknown89                       uint8
	Unknown8A                       [4]uint8
	HasRelicBook                    uint8
	RelicBookId                     uint8
	Unknown90                       [4]uint8
	CraftingMasterMask              uint8
	Unknown95                       [9]uint8
	Unknown9F                       [2]uint8
	UnknownA1                       [3]uint8
	Exp                             [CLASSJOB_SLOTS]uint32
	Unknown108                      uint32
	PvpTotalExp                     uint32
	UnknownPvp110                   uint32
	PvpExp                          uint32
	PvpFrontlineOverallRanks        [3]uint32
	Levels                          [CLASSJOB_SLOTS]uint16
	Unknown15C                      [9]uint16
	U1                              uint16
	U2                              uint16
	Unknown112                      [23]uint16
	FishingRecordsFish              [26]uint16
	BeastExp                        [11]uint16
	Unknown1EA                      [5]uint16
	PvpFrontlineWeeklyRanks         [3]uint16
	UnknownMask1FA                  [4]uint16
	CompanionName                   [21]uint8
	CompanionDefRank                uint8
	CompanionAttRank                uint8
	CompanionHealRank               uint8
	U19                             [8]uint8
	MountGuideMask                  [22]uint8
	Name                            [32]byte
	UnknownOword                    [16]uint8
	UnknownOw                       uint8
	UnlockBitmask                   [64]uint8
	Aetheryte                       [21]uint8
	Discovery                       [445]uint8
	Howto                           [34]uint8
	Minions                         [45]uint8
	ChocoboTaxiMask                 [10]uint8
	WatchedCutscenes                [124]uint8
	CompanionBardingMask            [10]uint8
	CompanionEquippedHead           uint8
	CompanionEquippedBody           uint8
	CompanionEquippedLegs           uint8
	Unknown52A                      [4]uint8
	UnknownMask52E                  [11]uint8
	FishingGuideMask                [105]uint8
	FishingSpotVisited              [31]uint8
	Unknown59A                      [27]uint8
	Unknown5A9                      [7]uint8
	BeastRank                       [11]uint8
	UnknownPvp5AB                   [11]uint8
	Unknown5B9                      [5]uint8
	Pose                            uint8
	Unknown5B91                     uint8
	ChallengeLogComplete            [9]uint8
	WeaponPose                      uint8
	UnknownMask673                  [10]uint8
	UnknownMask5DD                  [28]uint8
	RelicCompletion                 [12]uint8
	SightseeingMask                 [26]uint8
	HuntingMarkMask                 [55]uint8
	TripleTriadCards                [32]uint8
	U12                             [11]uint8
	U13                             uint8
	AetherCurrentMask               [22]uint8
	U10                             [3]uint8
	OrchestrionMask                 [40]uint8
	HallOfNoviceCompletion          [3]uint8
	AnimaCompletion                 [11]uint8
	U14                             [16]uint8
	U15                             [13]uint8
	UnlockedRaids                   [28]uint8
	UnlockedDungeons                [18]uint8
	UnlockedGuildhests              [10]uint8
	UnlockedTrials                  [8]uint8
	UnlockedPvp                     [5]uint8
	ClearedRaids                    [28]uint8
	ClearedDungeons                 [18]uint8
	ClearedGuildhests               [10]uint8
	ClearedTrials                   [8]uint8
	ClearedPvp                      [5]uint8
	FishingRecordsFishWeight        [26]uint16
	ExploratoryMissionNextTimestamp uint32
	PvpLevel                        uint8
}

/**
* Structural representation of the packet sent by the server
* To set a players stats
 */
type PlayerStats struct {
	// Order comes from baseparam order column
	Strength            uint32
	Dexterity           uint32
	Vitality            uint32
	Intelligence        uint32
	Mind                uint32
	Piety               uint32
	Hp                  uint32
	Mp                  uint32
	Tp                  uint32
	Gp                  uint32
	Cp                  uint32
	Delay               uint32
	Tenacity            uint32
	AttackPower         uint32
	Defense             uint32
	DirectHitRate       uint32
	Evasion             uint32
	MagicDefense        uint32
	CriticalHit         uint32
	AttackMagicPotency  uint32
	HealingMagicPotency uint32
	ElementalBonus      uint32
	Determination       uint32
	SkillSpeed          uint32
	SpellSpeed          uint32
	Haste               uint32
	Craftsmanship       uint32
	Control             uint32
	Gathering           uint32
	Perception          uint32

	// Todo: what is here?
	Unknown [26]uint32
}

/**
* Structural representation of the packet sent by the server
* To set an actors current owner
 */
type ActorOwner struct {
	ActorType uint8 // Note: Changed "type" to "actorType"
	Padding   [7]uint8
	ActorId   uint32
	ActorId2  uint32
}

/**
* Structural representation of the packet sent by the server
* To set a players state
 */
type PlayerStateFlags struct {
	Flags   [12]uint8
	Padding uint32
}

/**
* Structural representation of the packet sent by the server
* Containing current class information
 */
type PlayerClassInfo struct {
	ClassId      uint32
	Unknown      uint8
	IsSpecialist uint8
	SyncedLevel  uint16 // Locks actions, equipment, prob more. Player's current level (synced).
	ClassLevel   uint16 // Locks roles, prob more. Player's actual unsynced level.
	RoleActions  [10]uint32
}

/**
* Structural representation of the packet sent by the server
* To update a players appearance
 */
type ModelEquip struct {
	/* 0000 */ MainWeapon uint64
	/* 0008 */ OffWeapon uint64
	/* 0010 */ Unk1 uint8
	/* 0011 */ ClassJobId uint8
	/* 0012 */ Level uint8
	/* 0013 */ Unk2 uint8
	/* 0014 */ Models [10]uint32
	/* 003C */ padding2 uint32
}

type Examine struct {
	UnkFlag1         uint8
	UnkFlag2         uint8
	ClassJob         byte
	Level            byte
	Padding          uint16
	TitleId          uint16
	GrandCompany     byte
	GrandCompanyRank byte

	Unknown         [6]byte
	U6_fromPSpawn   uint32
	U7_fromPSpawn   uint32
	Padding1        [8]byte
	MainWeaponModel uint64
	SecWeaponModel  uint64
	Unknown2        uint8
	WorldId         uint16
	Unknown3        [12]byte
	Entries         [14]struct {
		CatalogId           uint32
		AppearanceCatalogId uint32
		CrafterId           uint64
		Quality             uint8
		Unknown             [3]uint8
		Materia             [5]struct {
			MateriaId uint16
			Tier      uint16
		}
	}
	Name     [32]byte
	Padding2 byte
	Unk3     [16]byte
	Look     [26]byte
	Padding3 [5]byte
	Models   [10]uint32
	Unknown4 [200]byte
}

type CharaNameReq struct {
	ContentId uint64
	Name      [32]byte
}

/**
* Structural representation of the packet sent by the server
* To update a players appearance
 */
type ItemInfo struct {
	ContainerSequence uint32
	Unknown           uint32
	ContainerId       uint16
	Slot              uint16
	Quantity          uint32
	CatalogId         uint32
	ReservedFlag      uint32
	SignatureId       uint64
	HqFlag            uint8
	Unknown2          uint8
	Condition         uint16
	SpiritBond        uint16
	Stain             uint16
	GlamourCatalogId  uint32
	Materia1          uint16
	Materia2          uint16
	Materia3          uint16
	Materia4          uint16
	Materia5          uint16
	Tier1             uint8
	Tier2             uint8
	Tier3             uint8
	Tier4             uint8
	Tier5             uint8
	Padding           uint8
	Unknown10         uint32
}

/**
* Structural representation of the packet sent by the server
* To update a players appearance
 */
type ContainerInfo struct {
	ContainerSequence uint32
	NumItems          uint32
	ContainerId       uint32
	Unknown           uint32
}

/**
* Structural representation of the packet sent by the server
* To update a players appearance
 */
type CurrencyCrystalInfo struct {
	ContainerSequence uint32
	ContainerId       uint16
	Slot              uint16
	Quantity          uint32
	Unknown           uint32
	CatalogId         uint32
	Unknown1          uint32
	Unknown2          uint32
	Unknown3          uint32
}

type InventoryTransactionFinish struct {
	SequenceId  uint32
	SequenceId1 uint32
	Padding     uint64
}

type InventoryTransaction struct {
	Sequence        uint32
	TransactionType uint8 // Note: Changed "type" to "transactionType"
	Padding         uint8
	Padding1        uint16
	OwnerId         uint32
	StorageId       uint32
	SlotId          uint16
	Padding2        uint16
	StackSize       uint32
	CatalogId       uint32
	SomeActorId     uint32
	TargetStorageId int32
	Padding3        [3]uint32
}

type InventoryActionAck struct {
	Sequence   uint32
	ActionType uint16 // Note: Changed "type" to "actionType"
	Padding    uint16
	Padding1   uint32
	Padding2   uint32
}

/**
* Structural representation of the packet sent by the server
* To update a slot in the inventory
 */
type UpdateInventorySlot struct {
	Sequence         uint32
	Unknown          uint32
	ContainerId      uint16
	Slot             uint16
	Quantity         uint32
	CatalogId        uint32
	ReservedFlag     uint32
	SignatureId      uint64
	HqFlag           uint16
	Condition        uint16
	SpiritBond       uint16
	Color            uint16
	GlamourCatalogId uint32
	Materia1         uint16
	Materia2         uint16
	Materia3         uint16
	Materia4         uint16
	Materia5         uint16
	Tier1            uint8
	Tier2            uint8
	Tier3            uint8
	Tier4            uint8
	Tier5            uint8
	Padding          uint8
	Unknown10        uint32
}

/**
* Structural representation of the packet sent by the server
* To start an event, not actually playing it, but registering
 */
type EventStart struct {
	/* 0000 */ ActorId uint64
	/* 0008 */ EventId uint32
	/* 000C */ param1 uint8
	/* 000D */ param2 uint8
	/* 000E */ padding uint16
	/* 0010 */ Param3 uint32
	/* 0014 */ Padding1 uint32
}

/**
* Structural representation of the packet sent by the server
* To fill a huntin log entry
 */
type HuntingLogEntry struct {
	U0            int32 // -1 for all normal classes
	Rank          uint8 // starting from 0
	Index         uint8 // classes and gcs
	Entries       [10][4]uint8
	Pad           uint16
	CompleteFlags uint64 // 4 bit for each potential entry and the 5th bit for completion of the section
	Pad1          uint64
}

/**
* Structural representation of the packet sent by the server
* To play an event
 */
type EventPlay struct {
	ActorId  uint64
	EventId  uint32
	Scene    uint16
	Padding  uint16
	Flags    uint32
	Param3   uint32
	Param4   uint8
	Padding1 [3]uint8
	Param5   uint32
	Unknown  [8]uint8
}

/**
* Structural representation of the packet sent by the server
* To play an event
 */
type DirectorPlayScene struct {
	ActorId  uint64
	EventId  uint32
	Scene    uint16
	Padding  uint16
	Flags    uint32
	Param3   uint32
	Param4   uint8
	Padding1 [3]uint8
	Param5   uint32
	Unknown8 [0x08]uint8
	Unknown  [0x38]uint8
}

/**
* Structural representation of the packet sent by the server
* To finish an event
 */
type EventFinish struct {
	/* 0000 */ EventId uint32
	/* 0004 */ Param1 uint8
	/* 0005 */ Param2 uint8
	/* 0006 */ Padding uint16
	/* 0008 */ Param3 uint32
	/* 000C */ padding1 uint32
}

type EventPlay4 struct {
	ActorId    uint64
	EventId    uint32
	Scene      uint16
	Padding    uint16
	SceneFlags uint32
	ParamCount uint8
	Padding2   [3]uint8
	Params     [4]uint32
}

type EventPlay8 struct {
	ActorId    uint64
	EventId    uint32
	Scene      uint16
	Padding    uint16
	SceneFlags uint32
	ParamCount uint8
	Padding2   [3]uint8
	Params     [8]uint32
}

type EventPlay16 struct {
	ActorId    uint64
	EventId    uint32
	Scene      uint16
	Padding    uint16
	SceneFlags uint32
	ParamCount uint8
	Padding2   [3]uint8
	Params     [16]uint32
}

type EventPlay64 struct {
	ActorId    uint64
	EventId    uint32
	Scene      uint16
	Padding    uint16
	SceneFlags uint32
	ParamCount uint8
	Padding2   [3]uint8
	Params     [64]uint32
}

type EventPlay128 struct {
	ActorId    uint64
	EventId    uint32
	Scene      uint16
	Padding    uint16
	SceneFlags uint32
	ParamCount uint8
	Padding2   [3]uint8
	Params     [128]uint32
}

type EventPlay255 struct {
	ActorId    uint64
	EventId    uint32
	Scene      uint16
	Padding    uint16
	SceneFlags uint32
	ParamCount uint8
	Padding2   [3]uint8
	Params     [255]uint32
}

/**
* Structural representation of the packet sent by the server
* To respond to a linkshell creation event
 */
type EventLinkshell struct {
	EventId  uint32
	Scene    uint8
	Param1   uint8
	Param2   uint8
	Param3   uint8
	Unknown1 uint32
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
}

/**
* Structural representation of the packet sent by the server
* To send the active quests
 */
type QuestActiveList struct {
	ActiveQuests [30]QuestActive
}

/**
* Structural representation of the packet sent by the server
* To send update a quest slot
 */
type QuestUpdate struct {
	Slot      uint16
	Padding   uint16
	QuestInfo QuestActive
}

/**
* Structural representation of the packet sent by the server
* To send the completed quests mask
 */
type QuestCompleteList struct {
	QuestCompleteMask   [480]uint8
	UnknownCompleteMask [80]uint8
}

/**
* Structural representation of the packet sent by the server
* To finish a quest
 */
type QuestFinish struct {
	QuestId uint16
	Flag1   uint8
	Flag2   uint8
	Padding uint32
}

/**
* Structural representation of the packet sent by the server
* To send a quest message
* Type 0 default
* Type 1 icon
* Type 5 status
 */
type QuestMessage struct {
	/* 0000 */ QuestId uint32
	/* 0000 */ MsgId uint8
	/* 0000 */ QuestType uint8 // Note: Changed "type" to "questType"
	/* 0000 */ Padding1 uint16
	/* 0000 */ Var1 uint32
	/* 0000 */ Var2 uint32
}

type QuestTracker struct {
	Entry [5]struct {
		Active     uint8
		QuestIndex uint8
	}
	Padding [3]uint16
}

type WeatherChange struct {
	WeatherId uint32
	Delay     float32
}

/**
* Structural representation of the packet sent by the server
* To send a unviel a map
 */
type Discovery struct {
	/* 0000 */ MapPartId uint32
	/* 0004 */ MapId uint32
}

/**
* UNKOWN TYPE
 */
type FFXIVARR_IPC_UNK322 struct {
	/* 0000 */ Unk [8]uint8
}

/**
* UNKOWN TYPE
 */
type FFXIVARR_IPC_UNK320 struct {
	/* 0000 */ Unk [0x38]uint8
}

/**
* Structural representation of the packet sent by the server
* Prepare zoning, showing screenmessage
 */
type PrepareZoning struct {
	LogMessage  uint32
	TargetZone  uint16
	Animation   uint16
	Param4      uint8
	Hide        uint8
	FadeOut     uint8
	Param7      uint8
	FadeOutTime uint8
	Unknown     uint8 // this changes whether or not the destination zone's name displays during the loading screen. Seems to always be 9 (=hidden) when going to an instance and certain zones, 0 otherwise.
	Padding     uint16
}

/**
* Structural representation of the packet sent by the server
* To trigger content finder events
*
* See https://gist.github.com/Minoost/c35843c4c8a7a931f31fdaac9bce64c2
 */
type CFNotify struct {
	State1 uint32 // 3 = cancelled, 4 = duty ready
	State2 uint32 // if state1 == 3, state2 is cancelled reason

	Param1 uint32 // usually classJobId
	Param2 uint32 // usually flag
	Param3 uint32 // usually languages, sometimes join in progress timestamp

	Param4   uint16 // usually roulette id
	Contents [5]uint16
}

/**
* Structural representation of the packet sent by the server
* To update contents available in duty finder or raid finder
*
* Do note that this packet has to come early in login phase (around initui packet)
* Or it won't be applied until you reconnect
 */
type CFAvailableContents struct {
	Contents [0x48]uint8
}

/**
* Structural representation of the packet sent by the server
* To update adventure in needs in duty roulette
 */
type CFPlayerInNeed struct {
	// Ordered by roulette id
	InNeeds [0x10]uint8
}

/**
* Structural representation of the packet sent by the server
* To update duty info in general
 */
type CFDutyInfo struct {
	PenaltyTime uint8
	Unknown     [7]uint8
}

type CFRegisterDuty struct {
	Unknown0   uint32 // 0x301
	RouletteId uint8  // if it's a daily roulette
	Unknown1   uint8  // 0xDB
	ContentId  uint16
}

type CFMemberStatus struct {
	ContentId     uint16
	Unknown1      uint16
	Status        uint8
	CurrentTank   uint8
	CurrentDps    uint8
	CurrentHealer uint8
	EstimatedTime uint8
	Unknown2      [3]uint8
	Unknown3      uint32
}

type EorzeaTimeOffset struct {
	Timestamp uint64
}

/**
* Structural representation of the packet sent by the server
* To set the gear show/hide status of a character
 */
type EquipDisplayFlags struct {
	Bitmask uint8
}

/**
* Structural representation of the packet sent by the server
* To mount a player
 */
type Mount struct {
	Id uint32
}

/**
* Structural representation of the packet sent by the server
* To mount a player
 */
type DirectorVars struct {
	/*! DirectorType | ContentId */
	M_directorId uint32
	/*! Currect sequence */
	M_sequence uint8
	/*! Current branch */
	M_branch uint8
	/*! Raw storage for flags/vars */
	M_unionData [10]uint8
	/*! Unknown */
	U20 uint16
	U22 uint16
	U24 uint16
	U28 uint16
}

type DirectorPopUp struct {
	DirectorId    uint32
	Pad1          [2]uint16
	SourceActorId uint64
	/*!
	* 2 = Green text in log
	 */
	Flags       uint8
	Pad2        [3]uint8
	BNPCName    uint32
	TextId      uint32
	PopupTimeMs uint32
	Pad3        [4]uint32
}

type ActorGauge struct {
	ClassJobId uint8
	Data       [15]uint8 // depends on classJobId
}

type PerformNote struct {
	Data [32]uint8
}

type HousingUpdateLandFlagsSlot struct {
	SlotType uint32 // Note: Changed "type" to "slotType"
	Unknown  uint32
	FlagSet  LandFlagSet
}

type HousingLandFlags struct {
	FreeCompanyHouse LandFlagSet // 00
	Unkown1          uint64
	PrivateHouse     LandFlagSet // 24
	Unkown2          uint64
	Apartment        LandFlagSet // 48
	Unkown3          uint64
	SharedHouse      [2]LandFlagSet //72
	Unkown4          uint64
	UnkownHouse      LandFlagSet
	Unkown5          uint64
}

//Structs
type LandStruct struct {
	PlotSize    uint8     //0
	HouseState  uint8     // 2
	Flags       uint8     // bit1 -> hasPublicAccess; bit2 -> isPersonalHouse
	IconAddIcon uint8     // 6
	FcId        uint32    //8
	FcIcon      uint32    // 12
	FcIconColor uint32    // 16
	HousePart   [8]uint16 // 34
	HouseColour [8]uint8  // 36
}

type LandUpdate struct {
	LandIdent LandIdent
	Land      LandStruct
}

type LandPriceUpdate struct {
	Price    uint32
	TimeLeft uint32
}

type LandInfoSign struct {
	LandIdent      LandIdent
	ownerId        uint64 // ither contentId or fcId
	Unknow1        uint32
	HouseIconAdd   uint8
	HouseSize      uint8
	HouseType      uint8
	EstateName     [23]byte
	EstateGreeting [193]byte
	OwnerName      [31]byte
	FcTag          [7]byte
	Tag            [3]uint8
}

type LandRename struct {
	LandIdent LandIdent
	HouseName [20]byte
	Padding   uint32
}

type LandUpdateHouseName struct {
	Unknown   [3]uint32
	HouseName [20]byte
	Unknown2  [2]uint32
}

type LandSetMap struct {
	U1          uint8
	Subdivision uint8
	U3          uint8
	LandInfo    [30]struct {
		Status    uint8
		Size      uint8
		IsPrivate uint8
	}
	Padding [3]uint8
}

type LandSetInitialize struct {
	LandIdent   LandIdent
	Unknown1    uint8
	SubInstance uint8 //  (default
	Unknown3    uint8
	Unknown4    uint8
	Unknown5    uint8
	Unknown6    uint8
	Unknown7    uint8
	Unknown8    uint8
	Land        [30]LandStruct
}

type YardObjectSpawn struct {
	LandId      uint8
	ObjectArray uint8
	Unknown1    uint16
	Object      HousingObject
}

type HousingObjectMove struct {
	ItemRotation uint16
	ObjectArray  uint8
	LandId       uint8
	Pos          FFXIVARR_POSITION3
	Unknown1     uint16
	Unknown2     uint16
	Unknown3     uint16
}

type HousingObjectInitialize struct {
	LandIdent LandIdent
	/*!
	* When this is 2, actrl 0x400 will hide the additional quarters door
	* If it's any other value, it will stay there regardless
	 */
	U1          int8 //Outdoor -1 / Indoor 0 - probably indicator
	PacketNum   uint8
	PacketTotal uint8
	U2          uint8 //Outdoor 0 / Indoor 100(?)
	Object      [100]HousingObject
	Unknown4    uint32 //unused
}

type HousingInternalObjectSpawn struct {
	ContainerId     uint16
	ContainerOffset uint8
	Pad1            uint8

	Object HousingObject
}

type HousingIndoorInitialize struct {
	U1          uint16
	U2          uint16
	U3          uint16
	U4          uint16
	IndoorItems [10]uint32
}

type HousingWardInfo struct {
	LandIdent LandIdent

	HouseInfoEntry [60]struct {
		HousePrice      uint32
		InfoFlags       uint8
		HouseAppeal     [3]uint8
		EstateOwnerName [30]byte
	}
}

type HousingEstateGreeting struct {
	LandIdent LandIdent
	Message   [200]byte
}

type HousingShowEstateGuestAccess struct {
	Unknown [2]uint32
	Ident   LandIdent
}

/**
* Structural representation of the packet sent by the server
* To show the current shared estate settings
 */
type SharedEstateSettingsResponse struct {
	Entry [3]struct {
		ContentId   uint64
		Permissions uint8
		Name        [0x20]byte
		Padding     [0x7]byte
	}
}

type MSQTrackerProgress struct {
	Id      uint32
	Padding uint32
}

type MSQTrackerComplete struct {
	Id       uint32
	Padding1 uint32
	Padding2 uint64
	Padding3 uint64
	Padding4 uint64 // last 4 bytes is but who cares
}

type ObjectSpawn struct {
	SpawnIndex    uint8
	ObjKind       uint8
	State         uint8
	Unknown3      uint8
	ObjId         uint32
	ActorId       uint32
	LevelId       uint32
	Unknown10     uint32
	SomeActorId14 uint32
	GimmickId     uint32
	Scale         float32
	Unknown20a    int16
	Rotation      uint16
	Unknown24a    int16
	Unknown24b    int16
	Unknown28a    uint16
	Unknown28c    int16
	HousingLink   uint32
	Position      FFXIVARR_POSITION3
	Unknown3C     int16
	Unknown3E     int16
}

type ObjectDespawn struct {
	SpawnIndex uint8
	Padding    [7]uint8
}

type DuelChallenge struct {
	OtherClassJobId uint8
	OtherLevel      uint8 // class job level
	ChallengeByYou  uint8 // 0 if the other challenges you, 1 if you challenges the other.
	OtherItemLevel  uint8

	OtherActorId uint32

	OtherName [32]byte
}

type RetainerInformation struct {
	Unknown1         uint64   /* 0x00 */
	RetainerId       uint32   /* 0x08 */
	Unknown2         uint32   /* 0x0C */
	HireOrder        byte     /* 0x10 */
	ItemCount        byte     /* 0x11 */
	Unknown3         uint16   /* 0x12 */
	Gil              uint32   /* 0x14 */
	ItemSellingCount byte     /* 0x18 */
	CityId           byte     /* 0x19 */
	ClassJobId       byte     /* 0x1A */
	Level            byte     /* 0x1B */
	Unknown4         uint32   /* 0x1C */
	VentureId        uint32   /* 0x20 */
	VentureComplete  uint32   /* 0x24 */
	Padding1         byte     /* 0x28 */
	Name             [32]byte /* 0x29 */
}
