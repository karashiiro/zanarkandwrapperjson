package sapphire

/**
* Structural representation of the packet sent by the server as response
* to a ping packet
 */
type Ping struct {
	/* 0000 */ timeInMilliseconds uint64
	/* 0008 */ unknown_8 [0x38]uint8
}

/**
* Structural representation of the packet sent by the server as response
* to a ping packet
 */
type Init struct {
	unknown  uint64
	charId   uint32
	unknown1 uint32
}

/**
* Structural representation of the packet sent by the server
* carrying chat messages
 */
type Chat struct {
	/* 0000 */ padding [14]uint8 //Maybe this is SubCode, or some kind of talker ID...
	chatType           uint16
	name               [32]byte
	msg                [1012]byte
}

type ChatBanned struct {
	padding [4]uint8 // I was not sure reinterpreting ZST is valid behavior in C++.
	// client doesn't care about the data (zero sized) for this opcode anyway.
}

/**
* Structural representation of the packet sent by the server
* to show a list of worlds for world visit
 */
type WorldVisitList struct {
	world [16]struct {
		id     uint16 // this is the id of the world from lobby
		status uint16 // 1 = available (this is what retail sends) | 2+ = unavailable (this will need to be checked with retail if it's exactly 2 or not since it does not actually lock the option)
	}
}

/**
* Structural representation of the packet sent by the server
* carrying chat messages
 */
type Logout struct {
	flags1 uint32
	flags2 uint32
}

/**
* Structural representation of the packet sent by the server
* sent to show the play time
 */
type PlayTime struct {
	playTimeInMinutes uint32
	padding           uint32
}

/**
* Structural representation of the packet sent by the server
* with a list of players ( party list | friend list | search results )
 */
type PlayerEntry struct {
	contentId        uint64
	bytes            [12]uint8
	zoneId           uint16
	zoneId1          uint16
	bytes1           [8]byte
	onlineStatusMask uint64
	classJob         uint8
	padding          uint8
	level            uint8
	padding1         uint8
	padding2         uint16
	one              uint8
	name             [0x20]byte
	fcTag            [9]byte
}

type SocialList struct {
	padding    uint32
	padding1   uint32
	padding2   uint32
	socialType uint8 // Changed name from "type" to "socialType"
	sequence   uint8
	padding3   uint16

	entries [10]PlayerEntry
}

type ExamineSearchInfo struct {
	unknown       uint32
	unknown1      uint16
	unknown2      uint16
	padding       [16]byte
	unknown3      uint32
	unknown4      uint16
	unknown5      uint16
	unknown6      uint16
	worldId       uint8
	searchMessage [193]byte
	fcName        [24]byte
	unknown7      uint8
	padding1      uint16
	levelEntries  [CLASSJOB_TOTAL]struct {
		id    uint16
		level uint16
	}
}

type SetSearchInfo struct {
	onlineStatusFlags uint64
	unknown           uint64
	unknown1          uint32
	padding           uint8
	selectRegion      uint8
	searchMessage     [193]byte
	padding2          uint8
}

type InitSearchInfo struct {
	onlineStatusFlags uint64
	unknown           uint64
	unknown1          uint8
	selectRegion      uint8
	searchMessage     [193]byte
	padding           [5]byte
}

type ExamineSearchComment struct {
	charId uint32
	// packet only has 196 bytes after the charid
	// likely utf8
	searchComment [195]byte
	padding       byte
}

/**
* Structural representation of the packet sent by the server
* to display a server notice message
 */
type ServerNoticeShort struct {
	// these are actually display flags
	/* 0000 */ padding uint8
	// 0 = chat log
	// 2 = nothing
	// 4 = on screen message
	// 5 = on screen message + chat log
	message [538]byte
}

/**
* Structural representation of the packet sent by the server
* to display a server notice message
 */
type ServerNotice struct {
	// these are actually display flags
	/* 0000 */ padding uint8
	// 0 = chat log
	// 2 = nothing
	// 4 = on screen message
	// 5 = on screen message + chat log
	message [775]byte
}

type SetOnlineStatus struct {
	onlineStatusFlags uint64
}

type BlackList struct {
	entry [20]struct {
		contentId uint64
		name      [32]byte
	}
	padding  uint8
	padding1 uint8
	sequence uint16
	padding2 uint32
}

type LogMessage struct {
	field_0    uint32
	field_4    uint32
	field_8    uint32
	field_12   uint32
	category   uint32
	logMessage uint32
	field_24   uint8
	field_25   uint8
	field_26   [32]uint8
	field_58   uint32
}

type LinkshellList struct {
	entry [8]struct {
		lsId      uint64
		unknownId uint64
		unknown   uint8
		rank      uint8
		padding   uint16
		lsName    [20]uint8
		unk       [16]uint8
	}
}

/**
* Structural representation of the packet sent by the server
* to send a list of mail the player has
 */
type ReqMoogleMailList struct {
	letter [5]struct {
		unk        [0x8]byte
		timeStamp  uint32     // The time the mail was sent (this also seems to be used as a Id)
		unk1       [0x30]byte // This should be items, gil, etc for the letter
		read       bool       // 0 = false | 1 = true
		letterType uint8      // 0 = Friends | 1 = Rewards | 2 = GM // Changed from "type" to "letterType"
		unk2       uint8
		senderName [0x20]byte // The name of the sender
		summary    [0x3C]byte // The start of the full letter text
		padding2   [0x5]byte
	}
	unk3 [0x08]byte
}

/**
* Structural representation of the packet sent by the server
* to show the mail delivery notification
 */
type MailLetterNotification struct {
	sendbackCount uint32    // The amount of letters sent back since you ran out of room (moogle dialog changes based on this)
	friendLetters uint16    // The amount of letters in the friends section of the letterbox
	unreadCount   uint16    // The amount of unreads in the letterbox (this is the number that shows up)
	rewardLetters uint16    // The amount of letters in the rewards section of the letterbox
	isGmLetter    uint8     // Makes the letter notification flash red
	isSupportDesk uint8     // After setting this to 1 we can no longer update mail notifications (more research needed on the support desk)
	unk2          [0x4]byte // This has probs something to do with the support desk (inquiry id?)
}

type FMarketTaxRates struct {
	unknown1 uint32
	padding1 uint16
	unknown2 uint16
	taxRate  [TOWN_COUNT]uint32 // In the order of Common::Town
	unknown3 uint64
}

type FMarketBoardItemListingCount struct {
	itemCatalogId uint32
	unknown1      uint32 // does some shit if nonzero
	requestId     uint16
	quantity      uint16 // high/low u8s read separately?
	unknown3      uint32
}

type MarketBoardItemListing struct {
	listing [10]struct // 152 bytes each
	{
		listingId       uint64
		retainerId      uint64
		retainerOwnerId uint64
		artisanId       uint64
		pricePerUnit    uint32
		totalTax        uint32
		itemQuantity    uint32
		itemId          uint32
		lastReviewTime  uint16
		containerId     uint16
		slotId          uint32
		durability      uint16
		spiritBond      uint16
		/**
		* auto materiaId = (i & 0xFF0) >> 4;
		* auto index = i & 0xF;
		* auto leftover = i >> 8;
		 */
		materiaValue [5]uint16
		padding1     uint16
		padding2     uint32
		retainerName [32]byte
		playerName   [32]byte
		hq           bool
		materiaCount uint8
		onMannequin  uint8
		marketCity   uint8
		dyeId        uint16
		padding3     uint16
		padding4     uint32
	} // Multiple packets are sent if there are more than 10 search results.
	listingIndexEnd   uint8
	listingIndexStart uint8
	requestId         uint16
	padding7          [16]byte
	unknown13         uint8
	padding8          uint16
	unknown14         uint8
	padding9          uint64
	unknown15         uint32
	padding10         uint32
}

type MarketBoardItemListingHistory struct {
	itemCatalogId  uint32
	itemCatalogId2 uint32

	listing [20]struct {
		salePrice    uint32
		purchaseTime uint32
		quantity     uint32
		isHq         uint8
		padding      uint8
		onMannequin  uint8

		buyerName [33]byte

		itemCatalogId uint32
	}
}

type MarketBoardSearchResult struct {
	items [20]struct {
		itemCatalogId uint32
		quantity      uint16
		demand        uint16
	}

	itemIndexEnd   uint32
	padding1       uint32
	itemIndexStart uint32
	requestId      uint32
}

type ExamineFreeCompanyInfo struct {
	unknown         [0x20]byte // likely fc allegiance/icon/housing info etc
	charId          uint32
	fcTimeCreated   uint32
	unknown2        [0x10]byte
	unknown3        uint16
	fcName          [0x14]byte // 20 limit
	padding         uint16
	fcTag           [0x05]byte // 5 tag limit
	padding2        uint16     // null terminator?
	fcLeader        [0x20]byte // leader name (32 bytes)
	fcSlogan        [192]byte  // source: https://ffxiv.gamerescape.com/wiki/Free_Company (packet cap confirms this size also)
	padding3        byte       // null terminator?
	fcEstateProfile [20]byte   // todo: size needs confirmation
	padding4        uint32
}

type FreeCompanyUpdateShortMessage struct {
	unknown      uint32
	unknown1     uint16
	unknown2     uint16
	unknown3     uint32
	unknown5     uint32
	shortMessage [104]byte
}

type StatusEffectList struct {
	classId          uint8
	level1           uint8
	level            uint16
	current_hp       uint32
	max_hp           uint32
	current_mp       uint16
	max_mp           uint16
	currentTp        uint16
	shieldPercentage uint8
	unknown1         uint8
	effect           [30]StatusEffect
	padding          uint32
}

type FFXIVGCAffiliation struct {
	gcId   uint8
	gcRank [3]uint8
}

/**
* Structural representation of the packet sent by the server
* add a status effect
 */
type EffectResult struct {
	globalSequence   uint32
	actor_id         uint32
	current_hp       uint32
	max_hp           uint32
	current_mp       uint16
	current_tp       uint16
	max_mp           uint16
	unknown1         uint8
	classId          uint8
	shieldPercentage uint8
	entryCount       uint8
	unknown2         uint16

	statusEntries [4]struct {
		index         uint8 // which position do i display this
		unknown3      uint8
		id            uint16
		param         uint16
		unknown4      uint16 // Sort this out (old right half of power/param property)
		duration      float32
		sourceActorId uint32
	}

	unknown5 uint32
}

/**
* Structural representation of the packet sent by the server
* to update certain player details / status
 */
type ActorControl struct {
	/* 0000 */ category uint16
	/* 0002 */ padding uint16
	/* 0004 */ param1 uint32
	/* 0008 */ param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ param4 uint32
	/* 0014 */ padding1 uint32
}

/**
* Structural representation of the packet sent by the server
* to update certain player details / status
 */
type ActorControlSelf struct {
	/* 0000 */ category uint16
	/* 0002 */ padding uint16
	/* 0004 */ param1 uint32
	/* 0008 */ param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ param4 uint32
	/* 0014 */ param5 uint32
	/* 0018 */ param6 uint32
	/* 0018 */ padding1 uint32
}

/**
* Structural representation of the packet sent by the server
* to update certain player details / status
 */
type ActorControlTarget struct {
	/* 0000 */ category uint16
	/* 0002 */ padding uint16
	/* 0004 */ param1 uint32
	/* 0008 */ param2 uint32
	/* 000C */ param3 uint32
	/* 0010 */ param4 uint32
	/* 0014 */ padding1 uint32
	/* 0018 */ targetId uint64
}

/**
* Structural representation of the packet sent by the server
* to update HP / MP / TP
 */
type UpdateHpMpTp struct {
	/* 0000 */ hp uint32
	/* 0004 */ mp uint16
	/* 0006 */ tp uint16
	/* 0008 */ gp uint16
	/* 0010 */ unknown_10 uint16
	/* 0012 */ unknown_12 uint32
}

type Effect struct {
	animationTargetId uint64 // who the animation targets

	actionId uint32 // what the casting player casts, shown in battle log/ui
	/*!
	* @brief Zone sequence for the effect. Used to link effects that are split across multiple packets as one
	 */
	sequence uint32

	animationLockTime float32 // maybe? doesn't seem to do anything
	someTargetId      uint32  // always 0x0E000000?

	/*!
	* @brief The cast sequence from the originating player. Should only be sent to the source, 0 for every other player.
	*
	* This needs to match the sequence sent from the player in the action start packet otherwise you'll cancel the
	* initial animation and start a new one once the packet arrives.
	 */
	sourceSequence    uint16
	rotation          uint16
	actionAnimationId uint16 // the animation that is played by the casting character
	variation         uint8  // variation in the animation
	effectDisplayType uint8

	unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	effectCount uint8 // ignores effects if 0, otherwise parses all of them
	padding_21  uint16

	padding_22 [3]uint16

	effects [8 * 8]uint8

	padding_6A [3]uint16

	effectTargetId uint32 // who the effect targets
	effectFlags    uint32 // nonzero = effects do nothing, no battle log, no ui text - only shows animations

	padding_78 uint32
}

type AoeEffect8 struct {
	animationTargetId uint64 // who the animation targets

	actionId       uint32 // what the casting player casts, shown in battle log/ui
	globalSequence uint32 // seems to only increment on retail?

	animationLockTime float32 // maybe? doesn't seem to do anything
	someTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	sourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	rotation          uint16
	actionAnimationId uint16 // the animation that is played by the casting character
	variation         uint8  // variation in the animation
	effectDisplayType uint8

	unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	effectCount uint8 // ignores effects if 0, otherwise parses all of them
	padding_21  [3]uint16
	padding     uint16

	effects [8]struct {
		entries [8]EffectEntry
	}

	padding_6A [3]uint16

	effectTargetId [8]uint64
	unkFlag        [3]uint16 // all 0x7FFF
	unk            [3]uint16
}

type AoeEffect16 struct {
	animationTargetId uint64 // who the animation targets

	actionId       uint32 // what the casting player casts, shown in battle log/ui
	globalSequence uint32 // seems to only increment on retail?

	animationLockTime float32 // maybe? doesn't seem to do anything
	someTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	sourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	rotation          uint16
	actionAnimationId uint16 // the animation that is played by the casting character
	variation         uint8  // variation in the animation
	effectDisplayType uint8

	unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	effectCount uint8 // ignores effects if 0, otherwise parses all of them
	padding_21  [3]uint16
	padding     uint16

	effects [16]struct {
		entries [8]EffectEntry
	}

	padding_6A [3]uint16

	effectTargetId [16]uint64
	unkFlag        [3]uint16 // all 0x7FFF
	unk            [3]uint16
}

type AoeEffect24 struct {
	animationTargetId uint64 // who the animation targets

	actionId       uint32 // what the casting player casts, shown in battle log/ui
	globalSequence uint32 // seems to only increment on retail?

	animationLockTime float32 // maybe? doesn't seem to do anything
	someTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	sourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	rotation          uint16
	actionAnimationId uint16 // the animation that is played by the casting character
	variation         uint8  // variation in the animation
	effectDisplayType uint8

	unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	effectCount uint8 // ignores effects if 0, otherwise parses all of them
	padding_21  [3]uint16
	padding     uint16

	effects [24]struct {
		entries [8]EffectEntry
	}

	padding_6A [3]uint16

	effectTargetId [24]uint64
	unkFlag        [3]uint16 // all 0x7FFF
	unk            [3]uint16
}

type AoeEffect32 struct {
	animationTargetId uint64 // who the animation targets

	actionId       uint32 // what the casting player casts, shown in battle log/ui
	globalSequence uint32 // seems to only increment on retail?

	animationLockTime float32 // maybe? doesn't seem to do anything
	someTargetId      uint32  // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

	sourceSequence    uint16 // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
	rotation          uint16
	actionAnimationId uint16 // the animation that is played by the casting character
	variation         uint8  // variation in the animation
	effectDisplayType uint8

	unknown20   uint8 // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
	effectCount uint8 // ignores effects if 0, otherwise parses all of them
	padding_21  [3]uint16
	padding     uint16

	effects [32]struct {
		entries [8]EffectEntry
	}

	padding_6A [3]uint16

	effectTargetId [32]uint64
	unkFlag        [3]uint16 // all 0x7FFF
	unk            [3]uint16
}

/**
* Structural representation of the packet sent by the server
* to spawn an actor
 */
type PlayerSpawn struct {
	title          uint16
	u1b            uint16
	currentWorldId uint16
	homeWorldId    uint16

	gmRank       uint8
	u3c          uint8
	u4           uint8
	onlineStatus uint8

	pose uint8
	u5a  uint8
	u5b  uint8
	u5c  uint8

	targetId        uint64
	u6              uint32
	u7              uint32
	mainWeaponModel uint64
	secWeaponModel  uint64
	craftToolModel  uint64

	u14             uint32
	u15             uint32
	bNPCBase        uint32
	bNPCName        uint32
	u18             uint32
	u19             uint32
	directorId      uint32
	ownerId         uint32
	u22             uint32
	hPMax           uint32
	hPCur           uint32
	displayFlags    uint32
	fateID          uint16
	mPCurr          uint16
	tPCurr          uint16
	mPMax           uint16
	tPMax           uint16
	modelChara      uint16
	rotation        uint16
	activeMinion    uint16
	spawnIndex      uint8
	state           uint8
	persistentEmote uint8
	modelType       uint8
	subtype         uint8
	voice           uint8
	u25c            uint16
	enemyType       uint8
	level           uint8
	classJob        uint8
	u26d            uint8
	u27a            uint16
	currentMount    uint8
	mountHead       uint8
	mountBody       uint8
	mountFeet       uint8
	mountColor      uint8
	scale           uint8
	elementalLevel  uint32
	element         uint32
	effect          [30]StatusEffect
	pos             FFXIVARR_POSITION3
	models          [10]uint32
	name            [32]byte
	look            [26]uint8
	fcTag           [6]byte
	unk30           uint32
}

/**
* Structural representation of the packet sent by the server
* to spawn an actor
 */
type NpcSpawn struct {
	gimmickId uint32 // needs to be existing in the map, mob will snap to it
	u2b       uint8
	u2ab      uint8
	gmRank    uint8
	u3b       uint8

	aggressionMode uint8 // 1 passive, 2 aggressive
	onlineStatus   uint8
	u3c            uint8
	pose           uint8

	u4 uint32

	targetId        uint64
	u6              uint32
	u7              uint32
	mainWeaponModel uint64
	secWeaponModel  uint64
	craftToolModel  uint64

	u14             uint32
	u15             uint32
	bNPCBase        uint32
	bNPCName        uint32
	levelId         uint32
	u19             uint32
	directorId      uint32
	spawnerId       uint32
	parentActorId   uint32
	hPMax           uint32
	hPCurr          uint32
	displayFlags    uint32
	fateID          uint16
	mPCurr          uint16
	tPCurr          uint16
	mPMax           uint16
	tPMax           uint16
	modelChara      uint16
	rotation        uint16
	activeMinion    uint16
	spawnIndex      uint8
	state           uint8
	persistantEmote uint8
	modelType       uint8
	subtype         uint8
	voice           uint8
	u25c            uint16
	enemyType       uint8
	level           uint8
	classJob        uint8
	u26d            uint8
	u27a            uint8
	currentMount    uint8
	mountHead       uint8
	mountBody       uint8
	mountFeet       uint8
	mountColor      uint8
	scale           uint8
	elementalLevel  uint16 // Eureka
	element         uint16 // Eureka
	u30b            uint32
	effect          [30]StatusEffect
	pos             FFXIVARR_POSITION3
	models          [10]uint32
	name            [32]byte
	look            [26]uint8
	fcTag           [6]byte
	unk30           uint32
	unk31           uint32
	bNPCPartSlot    uint8
	unk32           uint8
	unk33           uint16
	unk34           uint32
}

/**
* Structural representation of the packet sent by the server
* to show player movement
 */
type ActorFreeSpawn struct {
	spawnId uint32
	actorId uint32
}

/**
* Structural representation of the packet sent by the server
* to show player movement
 */
type ActorMove struct {
	/* 0000 */ headRotation uint8
	/* 0001 */ rotation uint8
	/* 0002 */ animationType uint8
	/* 0003 */ animationState uint8
	/* 0004 */ animationSpeed uint8
	/* 0005 */ unknownRotation uint8
	/* 0006 */ posX uint16
	/* 0008 */ posY uint16
	/* 000a */ posZ uint16
	/* 000C */ unknown_12 uint32
}

/**
* Structural representation of the packet sent by the server
* to set an actors position
 */
type ActorSetPos struct {
	r16         uint16
	waitForLoad uint8
	unknown1    uint8
	unknown2    uint32
	x           float32
	y           float32
	z           float32
	unknown3    uint32
}

/**
* Structural representation of the packet sent by the server
* to start an actors casting
 */
type ActorCast struct {
	action_id uint16
	skillType uint8
	unknown   uint8
	unknown_1 uint32 // action id or mount id
	cast_time float32
	target_id uint32
	rotation  uint16
	flag      uint16 // 1 = interruptible blinking cast bar
	unknown_2 uint32
	posX      uint16
	posY      uint16
	posZ      uint16
	unknown_3 uint16
}

type HateList struct {
	numEntries uint32
	entry      [32]struct {
		actorId     uint32
		hatePercent uint8
		unknown     uint8
		padding     uint16
	}
	padding uint32
}

type HateRank struct {
	numEntries uint32
	entry      [32]struct {
		actorId    uint32
		hateAmount uint32
	}
	padding uint32
}

type UpdateClassInfo struct {
	classId        uint8
	level1         uint8
	level          uint16
	nextLevelIndex uint32
	currentExp     uint32
	restedExp      uint32
}

/**
* Structural representation of the packet sent by the server
* to send the titles available to the player
 */
type PlayerTitleList struct {
	titleList [48]uint8
}

/**
* Structural representation of the packet sent by the server
* to initialize a zone for the player
 */
type InitZone struct {
	serverId                 uint16
	zoneId                   uint16
	unknown1                 uint16
	contentfinderConditionId uint16
	unknown3                 uint32
	unknown4                 uint32
	weatherId                uint8
	bitmask                  uint8
	bitmask1                 uint8
	// bitmask1 findings
	//0 = unknown ( 7B F8 69 )
	//1 = show playguide window ( 7B 69 )
	//2 = unknown ( 7B 69 )
	//4 = disables record ready check ( 7B DF DF F8 F0 E4 110 (all sorts of social packets) )
	//8 = hide server icon ( 7B 69 )
	//16 = enable flight ( 7B F8 69 )
	//32 = unknown ( 7B F8 69 )
	//64 = unknown ( 7B F8 69 )
	//128 = shows message "You are now in the instanced area XX A.
	//Current instance can be confirmed at any time using the /instance text command." ( 7B F8 69 )

	unknown5             uint8
	unknown8             uint32
	festivalId           uint16
	additionalFestivalId uint16
	unknown9             uint32
	unknown10            uint32
	unknown11            uint32
	unknown12            [4]uint32
	unknown13            [3]uint32
	pos                  FFXIVARR_POSITION3
	unknown14            [3]uint32
	unknown15            uint32
}

/**
* Structural representation of the packet sent by the server to initialize
* the client UI upon initial connection.
 */
type PlayerSetup struct {
	// plain C types for a bit until the packet is actually fixed.
	// makes conversion between different editors easier.
	contentId                       uint64
	unknown8                        uint32
	unknownC                        uint32
	charId                          uint32
	restedExp                       uint32
	companionCurrentExp             uint32
	unknown1C                       uint32
	fishCaught                      uint32
	useBaitCatalogId                uint32
	unknown28                       uint32
	unknownPvp2C                    uint16
	unknown3                        uint16
	pvpFrontlineOverallCampaigns    uint32
	unknownTimestamp34              uint32
	unknownTimestamp38              uint32
	unknown3C                       uint32
	unknown40                       uint32
	unknown44                       uint32
	companionTimePassed             float32
	unknown4C                       uint32
	unknown50                       uint16
	unknownPvp52                    [4]uint16
	playerCommendations             uint16
	unknown5C                       uint16
	unknown5E                       uint16
	pvpFrontlineWeeklyCampaigns     uint16
	enhancedAnimaGlassProgress      uint16
	unknown64                       [4]uint16
	pvpRivalWingsTotalMatches       uint16
	pvpRivalWingsTotalVictories     uint16
	pvpRivalWingsWeeklyMatches      uint16
	pvpRivalWingsWeeklyVictories    uint16
	maxLevel                        uint8
	expansion                       uint8
	unknown76                       uint8
	unknown77                       uint8
	race                            uint8
	tribe                           uint8
	gender                          uint8
	currentJob                      uint8
	currentClass                    uint8
	deity                           uint8
	namedayMonth                    uint8
	namedayDay                      uint8
	cityState                       uint8
	homepoint                       uint8
	unknown82                       uint8
	petHotBar                       uint8
	companionRank                   uint8
	companionStars                  uint8
	companionSp                     uint8
	companionUnk86                  uint8
	companionColor                  uint8
	companionFavoFeed               uint8
	unknown89                       uint8
	unknown8A                       [4]uint8
	hasRelicBook                    uint8
	relicBookId                     uint8
	unknown90                       [4]uint8
	craftingMasterMask              uint8
	unknown95                       [9]uint8
	unknown9F                       [2]uint8
	unknownA1                       [3]uint8
	exp                             [CLASSJOB_SLOTS]uint32
	unknown108                      uint32
	pvpTotalExp                     uint32
	unknownPvp110                   uint32
	pvpExp                          uint32
	pvpFrontlineOverallRanks        [3]uint32
	levels                          [CLASSJOB_SLOTS]uint16
	unknown15C                      [9]uint16
	u1                              uint16
	u2                              uint16
	unknown112                      [23]uint16
	fishingRecordsFish              [26]uint16
	beastExp                        [11]uint16
	unknown1EA                      [5]uint16
	pvpFrontlineWeeklyRanks         [3]uint16
	unknownMask1FA                  [4]uint16
	companionName                   [21]uint8
	companionDefRank                uint8
	companionAttRank                uint8
	companionHealRank               uint8
	u19                             [8]uint8
	mountGuideMask                  [22]uint8
	name                            [32]byte
	unknownOword                    [16]uint8
	unknownOw                       uint8
	unlockBitmask                   [64]uint8
	aetheryte                       [21]uint8
	discovery                       [445]uint8
	howto                           [34]uint8
	minions                         [45]uint8
	chocoboTaxiMask                 [10]uint8
	watchedCutscenes                [124]uint8
	companionBardingMask            [10]uint8
	companionEquippedHead           uint8
	companionEquippedBody           uint8
	companionEquippedLegs           uint8
	unknown52A                      [4]uint8
	unknownMask52E                  [11]uint8
	fishingGuideMask                [105]uint8
	fishingSpotVisited              [31]uint8
	unknown59A                      [27]uint8
	unknown5A9                      [7]uint8
	beastRank                       [11]uint8
	unknownPvp5AB                   [11]uint8
	unknown5B9                      [5]uint8
	pose                            uint8
	unknown5B91                     uint8
	challengeLogComplete            [9]uint8
	weaponPose                      uint8
	unknownMask673                  [10]uint8
	unknownMask5DD                  [28]uint8
	relicCompletion                 [12]uint8
	sightseeingMask                 [26]uint8
	huntingMarkMask                 [55]uint8
	tripleTriadCards                [32]uint8
	u12                             [11]uint8
	u13                             uint8
	aetherCurrentMask               [22]uint8
	u10                             [3]uint8
	orchestrionMask                 [40]uint8
	hallOfNoviceCompletion          [3]uint8
	animaCompletion                 [11]uint8
	u14                             [16]uint8
	u15                             [13]uint8
	unlockedRaids                   [28]uint8
	unlockedDungeons                [18]uint8
	unlockedGuildhests              [10]uint8
	unlockedTrials                  [8]uint8
	unlockedPvp                     [5]uint8
	clearedRaids                    [28]uint8
	clearedDungeons                 [18]uint8
	clearedGuildhests               [10]uint8
	clearedTrials                   [8]uint8
	clearedPvp                      [5]uint8
	fishingRecordsFishWeight        [26]uint16
	exploratoryMissionNextTimestamp uint32
	pvpLevel                        uint8
}

/**
* Structural representation of the packet sent by the server
* to set a players stats
 */
type PlayerStats struct {
	// order comes from baseparam order column
	strength            uint32
	dexterity           uint32
	vitality            uint32
	intelligence        uint32
	mind                uint32
	piety               uint32
	hp                  uint32
	mp                  uint32
	tp                  uint32
	gp                  uint32
	cp                  uint32
	delay               uint32
	tenacity            uint32
	attackPower         uint32
	defense             uint32
	directHitRate       uint32
	evasion             uint32
	magicDefense        uint32
	criticalHit         uint32
	attackMagicPotency  uint32
	healingMagicPotency uint32
	elementalBonus      uint32
	determination       uint32
	skillSpeed          uint32
	spellSpeed          uint32
	haste               uint32
	craftsmanship       uint32
	control             uint32
	gathering           uint32
	perception          uint32

	// todo: what is here?
	unknown [26]uint32
}

/**
* Structural representation of the packet sent by the server
* to set an actors current owner
 */
type ActorOwner struct {
	actorType uint8 // Note: Changed "type" to "actorType"
	padding   [7]uint8
	actorId   uint32
	actorId2  uint32
}

/**
* Structural representation of the packet sent by the server
* to set a players state
 */
type PlayerStateFlags struct {
	flags   [12]uint8
	padding uint32
}

/**
* Structural representation of the packet sent by the server
* containing current class information
 */
type PlayerClassInfo struct {
	classId      uint32
	unknown      uint8
	isSpecialist uint8
	syncedLevel  uint16 // Locks actions, equipment, prob more. Player's current level (synced).
	classLevel   uint16 // Locks roles, prob more. Player's actual unsynced level.
	roleActions  [10]uint32
}

/**
* Structural representation of the packet sent by the server
* to update a players appearance
 */
type ModelEquip struct {
	/* 0000 */ mainWeapon uint64
	/* 0008 */ offWeapon uint64
	/* 0010 */ unk1 uint8
	/* 0011 */ classJobId uint8
	/* 0012 */ level uint8
	/* 0013 */ unk2 uint8
	/* 0014 */ models [10]uint32
	/* 003C */ padding2 uint32
}

type Examine struct {
	unkFlag1         uint8
	unkFlag2         uint8
	classJob         byte
	level            byte
	padding          uint16
	titleId          uint16
	grandCompany     byte
	grandCompanyRank byte

	unknown         [6]byte
	u6_fromPSpawn   uint32
	u7_fromPSpawn   uint32
	padding1        [8]byte
	mainWeaponModel uint64
	secWeaponModel  uint64
	unknown2        uint8
	worldId         uint16
	unknown3        [12]byte
	entries         [14]struct {
		catalogId           uint32
		appearanceCatalogId uint32
		crafterId           uint64
		quality             uint8
		unknown             [3]uint8
		materia             [5]struct {
			materiaId uint16
			tier      uint16
		}
	}
	name     [32]byte
	padding2 byte
	unk3     [16]byte
	look     [26]byte
	padding3 [5]byte
	models   [10]uint32
	unknown4 [200]byte
}

type CharaNameReq struct {
	contentId uint64
	name      [32]byte
}

/**
* Structural representation of the packet sent by the server
* to update a players appearance
 */
type ItemInfo struct {
	containerSequence uint32
	unknown           uint32
	containerId       uint16
	slot              uint16
	quantity          uint32
	catalogId         uint32
	reservedFlag      uint32
	signatureId       uint64
	hqFlag            uint8
	unknown2          uint8
	condition         uint16
	spiritBond        uint16
	stain             uint16
	glamourCatalogId  uint32
	materia1          uint16
	materia2          uint16
	materia3          uint16
	materia4          uint16
	materia5          uint16
	tier1             uint8
	tier2             uint8
	tier3             uint8
	tier4             uint8
	tier5             uint8
	padding           uint8
	unknown10         uint32
}

/**
* Structural representation of the packet sent by the server
* to update a players appearance
 */
type ContainerInfo struct {
	containerSequence uint32
	numItems          uint32
	containerId       uint32
	unknown           uint32
}

/**
* Structural representation of the packet sent by the server
* to update a players appearance
 */
type CurrencyCrystalInfo struct {
	containerSequence uint32
	containerId       uint16
	slot              uint16
	quantity          uint32
	unknown           uint32
	catalogId         uint32
	unknown1          uint32
	unknown2          uint32
	unknown3          uint32
}

type InventoryTransactionFinish struct {
	sequenceId  uint32
	sequenceId1 uint32
	padding     uint64
}

type InventoryTransaction struct {
	sequence        uint32
	transactionType uint8 // Note: Changed "type" to "transactionType"
	padding         uint8
	padding1        uint16
	ownerId         uint32
	storageId       uint32
	slotId          uint16
	padding2        uint16
	stackSize       uint32
	catalogId       uint32
	someActorId     uint32
	targetStorageId int32
	padding3        [3]uint32
}

type InventoryActionAck struct {
	sequence   uint32
	actionType uint16 // Note: Changed "type" to "actionType"
	padding    uint16
	padding1   uint32
	padding2   uint32
}

/**
* Structural representation of the packet sent by the server
* to update a slot in the inventory
 */
type UpdateInventorySlot struct {
	sequence         uint32
	unknown          uint32
	containerId      uint16
	slot             uint16
	quantity         uint32
	catalogId        uint32
	reservedFlag     uint32
	signatureId      uint64
	hqFlag           uint16
	condition        uint16
	spiritBond       uint16
	color            uint16
	glamourCatalogId uint32
	materia1         uint16
	materia2         uint16
	materia3         uint16
	materia4         uint16
	materia5         uint16
	tier1            uint8
	tier2            uint8
	tier3            uint8
	tier4            uint8
	tier5            uint8
	padding          uint8
	unknown10        uint32
}

/**
* Structural representation of the packet sent by the server
* to start an event, not actually playing it, but registering
 */
type EventStart struct {
	/* 0000 */ actorId uint64
	/* 0008 */ eventId uint32
	/* 000C */ param1 uint8
	/* 000D */ param2 uint8
	/* 000E */ padding uint16
	/* 0010 */ param3 uint32
	/* 0014 */ padding1 uint32
}

/**
* Structural representation of the packet sent by the server
* to fill a huntin log entry
 */
type HuntingLogEntry struct {
	u0            int32 // -1 for all normal classes
	rank          uint8 // starting from 0
	index         uint8 // classes and gcs
	entries       [10][4]uint8
	pad           uint16
	completeFlags uint64 // 4 bit for each potential entry and the 5th bit for completion of the section
	pad1          uint64
}

/**
* Structural representation of the packet sent by the server
* to play an event
 */
type EventPlay struct {
	actorId  uint64
	eventId  uint32
	scene    uint16
	padding  uint16
	flags    uint32
	param3   uint32
	param4   uint8
	padding1 [3]uint8
	param5   uint32
	unknown  [8]uint8
}

/**
* Structural representation of the packet sent by the server
* to play an event
 */
type DirectorPlayScene struct {
	actorId  uint64
	eventId  uint32
	scene    uint16
	padding  uint16
	flags    uint32
	param3   uint32
	param4   uint8
	padding1 [3]uint8
	param5   uint32
	unknown8 [0x08]uint8
	unknown  [0x38]uint8
}

/**
* Structural representation of the packet sent by the server
* to finish an event
 */
type EventFinish struct {
	/* 0000 */ eventId uint32
	/* 0004 */ param1 uint8
	/* 0005 */ param2 uint8
	/* 0006 */ padding uint16
	/* 0008 */ param3 uint32
	/* 000C */ padding1 uint32
}

type EventPlayN struct {
	actorId    uint64
	eventId    uint32
	scene      uint16
	padding    uint16
	sceneFlags uint32
	paramCount uint8
	padding2   [3]uint8
	params     [1]uint32
}

type EventPlay255 struct {
	actorId    uint64
	eventId    uint32
	scene      uint16
	padding    uint16
	sceneFlags uint32
	paramCount uint8
	padding2   [3]uint8
	params     [255]uint32
}

/**
* Structural representation of the packet sent by the server
* to respond to a linkshell creation event
 */
type EventLinkshell struct {
	eventId  uint32
	scene    uint8
	param1   uint8
	param2   uint8
	param3   uint8
	unknown1 uint32
	unknown2 uint32
	unknown3 uint32
	unknown4 uint32
}

/**
* Structural representation of the packet sent by the server
* to send the active quests
 */
type QuestActiveList struct {
	activeQuests [30]QuestActive
}

/**
* Structural representation of the packet sent by the server
* to send update a quest slot
 */
type QuestUpdate struct {
	slot      uint16
	padding   uint16
	questInfo QuestActive
}

/**
* Structural representation of the packet sent by the server
* to send the completed quests mask
 */
type QuestCompleteList struct {
	questCompleteMask   [480]uint8
	unknownCompleteMask [80]uint8
}

/**
* Structural representation of the packet sent by the server
* to finish a quest
 */
type QuestFinish struct {
	questId uint16
	flag1   uint8
	flag2   uint8
	padding uint32
}

/**
* Structural representation of the packet sent by the server
* to send a quest message
* type 0 default
* type 1 icon
* type 5 status
 */
type QuestMessage struct {
	/* 0000 */ questId uint32
	/* 0000 */ msgId uint8
	/* 0000 */ questType uint8 // Note: Changed "type" to "questType"
	/* 0000 */ padding1 uint16
	/* 0000 */ var1 uint32
	/* 0000 */ var2 uint32
}

type QuestTracker struct {
	entry [5]struct {
		active     uint8
		questIndex uint8
	}
	padding [3]uint16
}

type WeatherChange struct {
	weatherId uint32
	delay     float32
}

/**
* Structural representation of the packet sent by the server
* to send a unviel a map
 */
type Discovery struct {
	/* 0000 */ mapPartId uint32
	/* 0004 */ mapId uint32
}

/**
* UNKOWN TYPE
 */
type FFXIVARR_IPC_UNK322 struct {
	/* 0000 */ unk [8]uint8
}

/**
* UNKOWN TYPE
 */
type FFXIVARR_IPC_UNK320 struct {
	/* 0000 */ unk [0x38]uint8
}

/**
* Structural representation of the packet sent by the server
* prepare zoning, showing screenmessage
 */
type PrepareZoning struct {
	logMessage  uint32
	targetZone  uint16
	animation   uint16
	param4      uint8
	hide        uint8
	fadeOut     uint8
	param7      uint8
	fadeOutTime uint8
	unknown     uint8 // this changes whether or not the destination zone's name displays during the loading screen. Seems to always be 9 (=hidden) when going to an instance and certain zones, 0 otherwise.
	padding     uint16
}

/**
* Structural representation of the packet sent by the server
* to trigger content finder events
*
* See https://gist.github.com/Minoost/c35843c4c8a7a931f31fdaac9bce64c2
 */
type CFNotify struct {
	state1 uint32 // 3 = cancelled, 4 = duty ready
	state2 uint32 // if state1 == 3, state2 is cancelled reason

	param1 uint32 // usually classJobId
	param2 uint32 // usually flag
	param3 uint32 // usually languages, sometimes join in progress timestamp

	param4   uint16 // usually roulette id
	contents [5]uint16
}

/**
* Structural representation of the packet sent by the server
* to update contents available in duty finder or raid finder
*
* Do note that this packet has to come early in login phase (around initui packet)
* or it won't be applied until you reconnect
 */
type CFAvailableContents struct {
	contents [0x48]uint8
}

/**
* Structural representation of the packet sent by the server
* to update adventure in needs in duty roulette
 */
type CFPlayerInNeed struct {
	// Ordered by roulette id
	inNeeds [0x10]uint8
}

/**
* Structural representation of the packet sent by the server
* to update duty info in general
 */
type CFDutyInfo struct {
	penaltyTime uint8
	unknown     [7]uint8
}

type CFRegisterDuty struct {
	unknown0   uint32 // 0x301
	rouletteId uint8  // if it's a daily roulette
	unknown1   uint8  // 0xDB
	contentId  uint16
}

type CFMemberStatus struct {
	contentId     uint16
	unknown1      uint16
	status        uint8
	currentTank   uint8
	currentDps    uint8
	currentHealer uint8
	estimatedTime uint8
	unknown2      [3]uint8
	unknown3      uint32
}

type EorzeaTimeOffset struct {
	timestamp uint64
}

/**
* Structural representation of the packet sent by the server
* to set the gear show/hide status of a character
 */
type EquipDisplayFlags struct {
	bitmask uint8
}

/**
* Structural representation of the packet sent by the server
* to mount a player
 */
type Mount struct {
	id uint32
}

/**
* Structural representation of the packet sent by the server
* to mount a player
 */
type DirectorVars struct {
	/*! DirectorType | ContentId */
	m_directorId uint32
	/*! currect sequence */
	m_sequence uint8
	/*! current branch */
	m_branch uint8
	/*! raw storage for flags/vars */
	m_unionData [10]uint8
	/*! unknown */
	u20 uint16
	u22 uint16
	u24 uint16
	u28 uint16
}

type DirectorPopUp struct {
	directorId    uint32
	pad1          [2]uint16
	sourceActorId uint64
	/*!
	* 2 = green text in log
	 */
	flags       uint8
	pad2        [3]uint8
	bNPCName    uint32
	textId      uint32
	popupTimeMs uint32
	pad3        [4]uint32
}

type ActorGauge struct {
	classJobId uint8
	data       [15]uint8 // depends on classJobId
}

type PerformNote struct {
	data [32]uint8
}

type HousingUpdateLandFlagsSlot struct {
	slotType uint32 // Note: Changed "type" to "slotType"
	unknown  uint32
	flagSet  LandFlagSet
}

type HousingLandFlags struct {
	freeCompanyHouse LandFlagSet // 00
	unkown1          uint64
	privateHouse     LandFlagSet // 24
	unkown2          uint64
	apartment        LandFlagSet // 48
	unkown3          uint64
	sharedHouse      [2]LandFlagSet //72
	unkown4          uint64
	unkownHouse      LandFlagSet
	unkown5          uint64
}

//Structs
type LandStruct struct {
	plotSize    uint8     //0
	houseState  uint8     // 2
	flags       uint8     // bit1 -> hasPublicAccess; bit2 -> isPersonalHouse
	iconAddIcon uint8     // 6
	fcId        uint32    //8
	fcIcon      uint32    // 12
	fcIconColor uint32    // 16
	housePart   [8]uint16 // 34
	houseColour [8]uint8  // 36
}

type LandUpdate struct {
	landIdent LandIdent
	land      LandStruct
}

type LandPriceUpdate struct {
	price    uint32
	timeLeft uint32
}

type LandInfoSign struct {
	landIdent      LandIdent
	uint64         // ither contentId or fcId
	unknow1        uint32
	houseIconAdd   uint8
	houseSize      uint8
	houseType      uint8
	estateName     [23]byte
	estateGreeting [193]byte
	ownerName      [31]byte
	fcTag          [7]byte
	tag            [3]uint8
}

type LandRename struct {
	landIdent LandIdent
	houseName [20]byte
	padding   uint32
}

type LandUpdateHouseName struct {
	unknown   [3]uint32
	houseName [20]byte
	unknown2  [2]uint32
}

type LandSetMap struct {
	u1          uint8
	subdivision uint8
	u3          uint8
	landInfo    [30]struct {
		status    uint8
		size      uint8
		isPrivate uint8
	}
	padding [3]uint8
}

type LandSetInitialize struct {
	landIdent   LandIdent
	unknown1    uint8
	subInstance uint8 //  (default
	unknown3    uint8
	unknown4    uint8
	unknown5    uint8
	unknown6    uint8
	unknown7    uint8
	unknown8    uint8
	land        [30]LandStruct
}

type YardObjectSpawn struct {
	landId      uint8
	objectArray uint8
	unknown1    uint16
	object      HousingObject
}

type HousingObjectMove struct {
	itemRotation uint16
	objectArray  uint8
	landId       uint8
	pos          FFXIVARR_POSITION3
	unknown1     uint16
	unknown2     uint16
	unknown3     uint16
}

type HousingObjectInitialize struct {
	landIdent LandIdent
	/*!
	* when this is 2, actrl 0x400 will hide the additional quarters door
	* if it's any other value, it will stay there regardless
	 */
	u1          int8 //Outdoor -1 / Indoor 0 - probably indicator
	packetNum   uint8
	packetTotal uint8
	u2          uint8 //Outdoor 0 / Indoor 100(?)
	object      [100]HousingObject
	unknown4    uint32 //unused
}

type HousingInternalObjectSpawn struct {
	containerId     uint16
	containerOffset uint8
	pad1            uint8

	object HousingObject
}

type HousingIndoorInitialize struct {
	u1          uint16
	u2          uint16
	u3          uint16
	u4          uint16
	indoorItems [10]uint32
}

type HousingWardInfo struct {
	landIdent LandIdent

	houseInfoEntry [60]struct {
		housePrice      uint32
		infoFlags       uint8
		houseAppeal     [3]uint8
		estateOwnerName [30]byte
	}
}

type HousingEstateGreeting struct {
	landIdent LandIdent
	message   [200]byte
}

type HousingShowEstateGuestAccess struct {
	unknown [2]uint32
	ident   LandIdent
}

/**
* Structural representation of the packet sent by the server
* to show the current shared estate settings
 */
type SharedEstateSettingsResponse struct {
	entry [3]struct {
		contentId   uint64
		permissions uint8
		name        [0x20]byte
		padding     [0x7]byte
	}
}

type MSQTrackerProgress struct {
	id      uint32
	padding uint32
}

type MSQTrackerComplete struct {
	id       uint32
	padding1 uint32
	padding2 uint64
	padding3 uint64
	padding4 uint64 // last 4 bytes is but who cares
}

type ObjectSpawn struct {
	spawnIndex    uint8
	objKind       uint8
	state         uint8
	unknown3      uint8
	objId         uint32
	actorId       uint32
	levelId       uint32
	unknown10     uint32
	someActorId14 uint32
	gimmickId     uint32
	scale         float32
	unknown20a    int16
	rotation      uint16
	unknown24a    int16
	unknown24b    int16
	unknown28a    uint16
	unknown28c    int16
	housingLink   uint32
	position      FFXIVARR_POSITION3
	unknown3C     int16
	unknown3E     int16
}

type ObjectDespawn struct {
	spawnIndex uint8
	padding    [7]uint8
}

type DuelChallenge struct {
	otherClassJobId uint8
	otherLevel      uint8 // class job level
	challengeByYou  uint8 // 0 if the other challenges you, 1 if you challenges the other.
	otherItemLevel  uint8

	otherActorId uint32

	otherName [32]byte
}
