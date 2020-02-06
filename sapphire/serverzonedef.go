package sapphire

/**
* Structural representation of the packet sent by the server as response
* to a ping packet
*/
struct Ping
{
/* 0000 */ uint64_t timeInMilliseconds uint64;
/* 0008 */ uint8_t unknown_8[0x38] uint8;
};

/**
* Structural representation of the packet sent by the server as response
* to a ping packet
*/
struct Init
{
uint64_t unknown uint64;
uint32_t charId uint32;
uint32_t unknown1 uint32;
};

/**
* Structural representation of the packet sent by the server
* carrying chat messages
*/
struct Chat
{
/* 0000 */  uint8_t padding[14]uint8; //Maybe this is SubCode, or some kind of talker ID...
Common::ChatType chatType uint16
char name[32]byte;
char msg[1012]byte;
};

struct ChatBanned
{
uint8_t padding[4]uint8; // I was not sure reinterpreting ZST is valid behavior in C++.
// client doesn't care about the data (zero sized) for this opcode anyway.
};

/**
* Structural representation of the packet sent by the server
* to show a list of worlds for world visit
*/
struct WorldVisitList
{
	world[16]struct
{
	uint16_t id uint16; // this is the id of the world from lobby
	uint16_t status uint16; // 1 = available (this is what retail sends) | 2+ = unavailable (this will need to be checked with retail if it's exactly 2 or not since it does not actually lock the option)
};
};

/**
* Structural representation of the packet sent by the server
* carrying chat messages
*/
struct Logout
{
uint32_t flags1 uint32;
uint32_t flags2 uint32;
};

/**
* Structural representation of the packet sent by the server
* sent to show the play time
*/
struct PlayTime
{
uint32_t playTimeInMinutes uint32;
uint32_t padding uint32;
};


/**
* Structural representation of the packet sent by the server
* with a list of players ( party list | friend list | search results )
*/
struct PlayerEntry
{
uint64_t contentId uint64;
uint8_t bytes[12]uint8;
uint16_t zoneId uint16;
uint16_t zoneId1 uint16;
char bytes1[8]byte;
uint64_t onlineStatusMask uint64;
uint8_t classJob uint8;
uint8_t padding uint8;
uint8_t level uint8;
uint8_t padding1 uint8;
uint16_t padding2 uint16;
uint8_t one uint8;
char name[0x20]byte;
char fcTag[9]byte;
};

struct SocialList
{
uint32_t padding uint32;
uint32_t padding1 uint32;
uint32_t padding2 uint32;
uint8_t socialType uint8; // Changed name from "type" to "socialType"
uint8_t sequence uint8;
uint16_t padding3 uint16;

entries[10]PlayerEntry;
};

struct ExamineSearchInfo
{
uint32_t unknown uint32;
uint16_t unknown1 uint16;
uint16_t unknown2 uint16;
char padding[16]byte;
uint32_t unknown3 uint32;
uint16_t unknown4 uint16;
uint16_t unknown5 uint16;
uint16_t unknown6 uint16;
uint8_t worldId uint8;
char searchMessage[193]byte;
char fcName[24]byte;
uint8_t unknown7 uint8;
uint16_t padding1 uint16;
levelEntries[CLASSJOB_TOTAL]struct
{
	uint16_t id uint16;
	uint16_t level uint16;
};
};

struct SetSearchInfo
{
uint64_t onlineStatusFlags uint64;
uint64_t unknown uint64;
uint32_t unknown1 uint32;
uint8_t padding uint8;
uint8_t selectRegion uint8;
char searchMessage[193]byte;
uint8_t padding2 uint8;
};

struct InitSearchInfo
{
uint64_t onlineStatusFlags uint64;
uint64_t unknown uint64;
uint8_t unknown1 uint8;
uint8_t selectRegion uint8;
char searchMessage[193]byte;
char padding[5]byte;
};

struct ExamineSearchComment
{
uint32_t charId uint32;
// packet only has 196 bytes after the charid
// likely utf8
char searchComment[195]byte;
char padding byte;
};

/**
* Structural representation of the packet sent by the server
* to display a server notice message
*/
struct ServerNoticeShort
{
// these are actually display flags
/* 0000 */ uint8_t padding uint8;
// 0 = chat log
// 2 = nothing
// 4 = on screen message
// 5 = on screen message + chat log
char message[538]byte;
};

/**
* Structural representation of the packet sent by the server
* to display a server notice message
*/
struct ServerNotice
{
// these are actually display flags
/* 0000 */ uint8_t padding uint8;
// 0 = chat log
// 2 = nothing
// 4 = on screen message
// 5 = on screen message + chat log
char message[775]byte;
};

struct SetOnlineStatus
{
uint64_t onlineStatusFlags uint64;
};

struct BlackList
{
	entry[20]struct
{
	uint64_t contentId uint64;
	char name[32]byte;
};
uint8_t padding uint8;
uint8_t padding1 uint8;
uint16_t sequence uint16;
uint32_t padding2 uint32;
};

struct LogMessage
{
uint32_t field_0 uint32;
uint32_t field_4 uint32;
uint32_t field_8 uint32;
uint32_t field_12 uint32;
uint32_t category uint32;
uint32_t logMessage uint32;
uint8_t field_24 uint8;
uint8_t field_25 uint8;
uint8_t field_26[32]uint8;
uint32_t field_58 uint32;
};

struct LinkshellList
{
	entry[8]struct
{
	uint64_t lsId uint64;
	uint64_t unknownId uint64;
	uint8_t unknown uint8;
	uint8_t rank uint8;
	uint16_t padding uint16;
	uint8_t lsName[20]uint8;
	uint8_t unk[16]uint8;
};
};

/**
* Structural representation of the packet sent by the server
* to send a list of mail the player has
*/
struct ReqMoogleMailList
{
	letter[5]struct
{
	char unk[0x8]byte;
	uint32_t timeStamp uint32; // The time the mail was sent (this also seems to be used as a Id)
	char unk1[0x30]byte; // This should be items, gil, etc for the letter
	uint8_t read bool; // 0 = false | 1 = true
	uint8_t letterType uint8; // 0 = Friends | 1 = Rewards | 2 = GM // Changed from "type" to "letterType"
	uint8_t unk2 uint8;
	char senderName[0x20]byte; // The name of the sender
	char summary[0x3C]byte; // The start of the full letter text
	char padding2[0x5]byte;
};
char unk3[0x08];
};

/**
* Structural representation of the packet sent by the server
* to show the mail delivery notification
*/
struct MailLetterNotification
{
uint32_t sendbackCount uint32; // The amount of letters sent back since you ran out of room (moogle dialog changes based on this)
uint16_t friendLetters uint16; // The amount of letters in the friends section of the letterbox
uint16_t unreadCount uint16; // The amount of unreads in the letterbox (this is the number that shows up)
uint16_t rewardLetters uint16; // The amount of letters in the rewards section of the letterbox
uint8_t isGmLetter uint8; // Makes the letter notification flash red
uint8_t isSupportDesk uint8; // After setting this to 1 we can no longer update mail notifications (more research needed on the support desk)
char unk2[0x4]byte; // This has probs something to do with the support desk (inquiry id?)
};

struct FMarketTaxRates
{
uint32_t unknown1 uint32;
uint16_t padding1 uint16;
uint16_t unknown2 uint16;
uint32_t taxRate[TOWN_COUNT]uint32; // In the order of Common::Town
uint64_t unknown3 uint64;
};

struct FMarketBoardItemListingCount
{
uint32_t itemCatalogId uint32;
uint32_t unknown1 uint32; // does some shit if nonzero
uint16_t requestId uint16;
uint16_t quantity uint16; // high/low u8s read separately?
uint32_t unknown3 uint32;
};

struct MarketBoardItemListing
{
	listing[10]struct // 152 bytes each
{
	uint64_t listingId uint64;
	uint64_t retainerId uint64;
	uint64_t retainerOwnerId uint64;
	uint64_t artisanId uint64;
	uint32_t pricePerUnit uint32;
	uint32_t totalTax uint32;
	uint32_t itemQuantity uint32;
	uint32_t itemId uint32;
	uint16_t lastReviewTime uint16;
	uint16_t containerId uint16;
	uint32_t slotId uint32;
	uint16_t durability uint16;
	uint16_t spiritBond uint16;
	/**
	* auto materiaId = (i & 0xFF0) >> 4;
	* auto index = i & 0xF;
	* auto leftover = i >> 8;
	*/
	uint16_t materiaValue[5] uint16;
	uint16_t padding1 uint16;
	uint32_t padding2 uint32;
	char retainerName[32]byte;
	char playerName[32]byte;
	hq bool;
	uint8_t materiaCount uint8;
	uint8_t onMannequin uint8;
	Common::Town marketCity uint8
	uint16_t dyeId uint16;
	uint16_t padding3 uint16;
	uint32_t padding4 uint32;
}; // Multiple packets are sent if there are more than 10 search results.
uint8_t listingIndexEnd uint8;
uint8_t listingIndexStart uint8;
uint16_t requestId uint16;
char padding7[16]byte;
uint8_t unknown13 uint8;
uint16_t padding8 uint16;
uint8_t unknown14 uint8;
uint64_t padding9 uint64;
uint32_t unknown15 uint32;
uint32_t padding10 uint32;
};

struct MarketBoardItemListingHistory
{
uint32_t itemCatalogId uint32;
uint32_t itemCatalogId2 uint32;

listing[20]struct
{
	uint32_t salePrice uint32;
	uint32_t purchaseTime uint32;
	uint32_t quantity uint32;
	uint8_t isHq uint8;
	uint8_t padding uint8;
	uint8_t onMannequin uint8;

	char buyerName[33]byte;

	uint32_t itemCatalogId uint32;
}
};

struct MarketBoardSearchResult
{
	items[20] struct
{
	uint32_t itemCatalogId uint32;
	uint16_t quantity uint16;
	uint16_t demand uint16;
};

uint32_t itemIndexEnd uint32;
uint32_t padding1 uint32;
uint32_t itemIndexStart uint32;
uint32_t requestId uint32;
};

struct ExamineFreeCompanyInfo
{
char unknown[0x20]byte; // likely fc allegiance/icon/housing info etc
uint32_t charId uint32;
uint32_t fcTimeCreated uint32;
char unknown2[0x10]byte;
uint16_t unknown3 uint16;
char fcName[0x14]byte; // 20 char limit
uint16_t padding uint16;
char fcTag[0x05]byte; // 5 char tag limit
uint16_t padding2 uint16; // null terminator?
char fcLeader[0x20]byte; // leader name (32 bytes)
char fcSlogan[192]byte; // source: https://ffxiv.gamerescape.com/wiki/Free_Company (packet cap confirms this size also)
char padding3 byte; // null terminator?
char fcEstateProfile[20] byte; // todo: size needs confirmation
uint32_t padding4 uint32;
};

struct FreeCompanyUpdateShortMessage
{
uint32_t unknown uint32;
uint16_t unknown1 uint16;
uint16_t unknown2 uint16;
uint32_t unknown3 uint32;
uint32_t unknown5 uint32;
char shortMessage[104]byte;
};

struct StatusEffectList
{
uint8_t classId uint8;
uint8_t level1 uint8;
uint16_t level uint16;
uint32_t current_hp uint32;
uint32_t max_hp uint32;
uint16_t current_mp uint16;
uint16_t max_mp uint16;
uint16_t currentTp uint16;
uint8_t shieldPercentage uint8;
uint8_t unknown1 uint8;
effect[30] StatusEffect;
uint32_t padding uint32;
};

struct FFXIVGCAffiliation
{
uint8_t gcId uint8;
uint8_t gcRank[3]uint8;
};

/**
* Structural representation of the packet sent by the server
* add a status effect
*/
struct EffectResult
{
uint32_t globalSequence uint32;
uint32_t actor_id uint32;
uint32_t current_hp uint32;
uint32_t max_hp uint32;
uint16_t current_mp uint16;
uint16_t current_tp uint16;
uint16_t max_mp uint16;
uint8_t unknown1 uint8;
uint8_t classId uint8;
uint8_t shieldPercentage uint8;
uint8_t entryCount uint8;
uint16_t unknown2 uint16;

statusEntries[4]struct
{
	uint8_t index uint8; // which position do i display this
	uint8_t unknown3 uint8;
	uint16_t id uint16;
	uint16_t param uint16;
	uint16_t unknown4 uint16;    // Sort this out (old right half of power/param property)
	float duration float32;
	uint32_t sourceActorId uint32;
};

uint32_t unknown5 uint32;
};

/**
* Structural representation of the packet sent by the server
* to update certain player details / status
*/
struct ActorControl
{
/* 0000 */ uint16_t category uint16;
/* 0002 */ uint16_t padding uint16;
/* 0004 */ uint32_t param1 uint32;
/* 0008 */ uint32_t param2 uint32;
/* 000C */ uint32_t param3 uint32;
/* 0010 */ uint32_t param4 uint32;
/* 0014 */ uint32_t padding1 uint32;
};

/**
* Structural representation of the packet sent by the server
* to update certain player details / status
*/
struct ActorControlSelf
{
/* 0000 */ uint16_t category uint16;
/* 0002 */ uint16_t padding uint16;
/* 0004 */ uint32_t param1 uint32;
/* 0008 */ uint32_t param2 uint32;
/* 000C */ uint32_t param3 uint32;
/* 0010 */ uint32_t param4 uint32;
/* 0014 */ uint32_t param5 uint32;
/* 0018 */ uint32_t param6 uint32;
/* 0018 */ uint32_t padding1 uint32;
};

/**
* Structural representation of the packet sent by the server
* to update certain player details / status
*/
struct ActorControlTarget
{
/* 0000 */ uint16_t category uint16;
/* 0002 */ uint16_t padding uint16;
/* 0004 */ uint32_t param1 uint32;
/* 0008 */ uint32_t param2 uint32;
/* 000C */ uint32_t param3 uint32;
/* 0010 */ uint32_t param4 uint32;
/* 0014 */ uint32_t padding1 uint32;
/* 0018 */ uint64_t targetId uint64;
};

/**
* Structural representation of the packet sent by the server
* to update HP / MP / TP
*/
struct UpdateHpMpTp
{
/* 0000 */ uint32_t hp uint32;
/* 0004 */ uint16_t mp uint16;
/* 0006 */ uint16_t tp uint16;
/* 0008 */ uint16_t gp uint16;
/* 0010 */ uint16_t unknown_10 uint16;
/* 0012 */ uint32_t unknown_12 uint32;
};

struct Effect
{
uint64_t animationTargetId uint64; // who the animation targets

uint32_t actionId uint32; // what the casting player casts, shown in battle log/ui
/*!
	* @brief Zone sequence for the effect. Used to link effects that are split across multiple packets as one
	*/
uint32_t sequence uint32;

float animationLockTime float32; // maybe? doesn't seem to do anything
uint32_t someTargetId uint32; // always 0x0E000000?

/*!
	* @brief The cast sequence from the originating player. Should only be sent to the source, 0 for every other player.
	*
	* This needs to match the sequence sent from the player in the action start packet otherwise you'll cancel the
	* initial animation and start a new one once the packet arrives.
	*/
uint16_t sourceSequence uint16;
uint16_t rotation uint16;
uint16_t actionAnimationId uint16; // the animation that is played by the casting character
uint8_t variation uint8; // variation in the animation
effectDisplayType uint8;

uint8_t unknown20 uint8; // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
uint8_t effectCount uint8; // ignores effects if 0, otherwise parses all of them
uint16_t padding_21 uint16;

uint16_t padding_22[3]uint16;

uint8_t effects[8*8]uint8;

uint16_t padding_6A[3]uint16;

uint32_t effectTargetId uint32; // who the effect targets
uint32_t effectFlags uint32; // nonzero = effects do nothing, no battle log, no ui text - only shows animations

uint32_t padding_78 uint32;
};

struct AoeEffect8
{
uint64_t animationTargetId uint64; // who the animation targets

uint32_t actionId uint32; // what the casting player casts, shown in battle log/ui
uint32_t globalSequence uint32; // seems to only increment on retail?

float animationLockTime float32; // maybe? doesn't seem to do anything
uint32_t someTargetId uint32; // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

uint16_t sourceSequence uint16; // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
uint16_t rotation uint16;
uint16_t actionAnimationId uint16; // the animation that is played by the casting character
uint8_t variation uint8; // variation in the animation
effectDisplayType uint8;

uint8_t unknown20 uint8; // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
uint8_t effectCount uint8; // ignores effects if 0, otherwise parses all of them
uint16_t padding_21[3]uint16;
uint16_t padding uint16;

effects[8]struct
{
	entries[8]EffectEntry;
};

uint16_t padding_6A[3]uint16;

uint64_t effectTargetId[8]uint64;
uint16_t unkFlag[3]uint16; // all 0x7FFF
uint16_t unk[3]uint16;
};

struct AoeEffect16
{
uint64_t animationTargetId uint64; // who the animation targets

uint32_t actionId uint32; // what the casting player casts, shown in battle log/ui
uint32_t globalSequence uint32; // seems to only increment on retail?

float animationLockTime float32; // maybe? doesn't seem to do anything
uint32_t someTargetId uint32; // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

uint16_t sourceSequence uint16; // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
uint16_t rotation uint16;
uint16_t actionAnimationId uint16; // the animation that is played by the casting character
uint8_t variation uint8; // variation in the animation
effectDisplayType uint8;

uint8_t unknown20 uint8; // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
uint8_t effectCount uint8; // ignores effects if 0, otherwise parses all of them
uint16_t padding_21[3]uint16;
uint16_t padding uint16;

effects[16]struct
{
	entries[8]EffectEntry;
};

uint16_t padding_6A[3]uint16;

uint64_t effectTargetId[16]uint64;
uint16_t unkFlag[3]uint16; // all 0x7FFF
uint16_t unk[3]uint16;
};

struct AoeEffect24
{
uint64_t animationTargetId uint64; // who the animation targets

uint32_t actionId uint32; // what the casting player casts, shown in battle log/ui
uint32_t globalSequence uint32; // seems to only increment on retail?

float animationLockTime float32; // maybe? doesn't seem to do anything
uint32_t someTargetId uint32; // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

uint16_t sourceSequence uint16; // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
uint16_t rotation uint16;
uint16_t actionAnimationId uint16; // the animation that is played by the casting character
uint8_t variation uint8; // variation in the animation
effectDisplayType uint8;

uint8_t unknown20 uint8; // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
uint8_t effectCount uint8; // ignores effects if 0, otherwise parses all of them
uint16_t padding_21[3]uint16;
uint16_t padding uint16;

effects[24]struct
{
	entries[8]EffectEntry;
};

uint16_t padding_6A[3]uint16;

uint64_t effectTargetId[24]uint64;
uint16_t unkFlag[3]uint16; // all 0x7FFF
uint16_t unk[3]uint16;
};

struct AoeEffect32
{
uint64_t animationTargetId uint64; // who the animation targets

uint32_t actionId uint32; // what the casting player casts, shown in battle log/ui
uint32_t globalSequence uint32; // seems to only increment on retail?

float animationLockTime float32; // maybe? doesn't seem to do anything
uint32_t someTargetId uint32; // always 00 00 00 E0, 0x0E000000 is the internal def for INVALID TARGET ID

uint16_t sourceSequence uint16; // if 0, always shows animation, otherwise hides it. counts up by 1 for each animation skipped on a caster
uint16_t rotation uint16;
uint16_t actionAnimationId uint16; // the animation that is played by the casting character
uint8_t variation uint8; // variation in the animation
effectDisplayType uint8;

uint8_t unknown20 uint8; // is read by handler, runs code which gets the LODWORD of animationLockTime (wtf?)
uint8_t effectCount uint8; // ignores effects if 0, otherwise parses all of them
uint16_t padding_21[3]uint16;
uint16_t padding uint16;

effects[32]struct
{
	entries[8]EffectEntry;
};

uint16_t padding_6A[3]uint16;

uint64_t effectTargetId[32]uint64;
uint16_t unkFlag[3]uint16; // all 0x7FFF
uint16_t unk[3]uint16;
};

/**
* Structural representation of the packet sent by the server
* to spawn an actor
*/
struct PlayerSpawn
{
uint16_t title uint16;
uint16_t u1b uint16;
uint16_t currentWorldId uint16;
uint16_t homeWorldId uint16;

uint8_t gmRank uint8;
uint8_t u3c uint8;
uint8_t u4 uint8;
uint8_t onlineStatus uint8;

uint8_t pose uint8;
uint8_t u5a uint8;
uint8_t u5b uint8;
uint8_t u5c uint8;

uint64_t targetId uint64;
uint32_t u6 uint32;
uint32_t u7 uint32;
uint64_t mainWeaponModel uint64;
uint64_t secWeaponModel uint64;
uint64_t craftToolModel uint64;

uint32_t u14 uint32;
uint32_t u15 uint32;
uint32_t bNPCBase uint32;
uint32_t bNPCName uint32;
uint32_t u18 uint32;
uint32_t u19 uint32;
uint32_t directorId uint32;
uint32_t ownerId uint32;
uint32_t u22; uint32
uint32_t hPMax uint32;
uint32_t hPCur uint32r;
uint32_t displayFlags uint32;
uint16_t fateID uint16;
uint16_t mPCurr uint16;
uint16_t tPCurr uint16;
uint16_t mPMax uint16;
uint16_t tPMax uint16;
uint16_t modelChara uint16;
uint16_t rotation uint16;
uint16_t activeMinion uint16;
uint8_t spawnIndex uint8;
uint8_t state uint8;
uint8_t persistentEmote uint8;
uint8_t modelType uint8;
uint8_t subtype uint8;
uint8_t voice uint8;
uint16_t u25c uint16;
uint8_t enemyType uint8;
uint8_t level uint8;
uint8_t classJob uint8;
uint8_t u26d uint8;
uint16_t u27a uint16;
uint8_t currentMount uint8;
uint8_t mountHead uint8;
uint8_t mountBody uint8;
uint8_t mountFeet uint8;
uint8_t mountColor uint8;
uint8_t scale uint8;
uint32_t elementalLevel uint32;
uint32_t element uint32;
effect[30]StatusEffect;
pos FFXIVARR_POSITION3;
uint32_t models[10] uint32;
char name[32]byte;
uint8_t look[26]uint8;
char fcTag[6]byte;
uint32_t unk30 uint32;
};

/**
* Structural representation of the packet sent by the server
* to spawn an actor
*/
struct NpcSpawn
{
uint32_t gimmickId uint32; // needs to be existing in the map, mob will snap to it
uint8_t u2b uint8;
uint8_t u2ab uint8;
uint8_t gmRank uint8;
uint8_t u3b uint8;

uint8_t aggressionMode uint8; // 1 passive, 2 aggressive
uint8_t onlineStatus uint8;
uint8_t u3c uint8;
uint8_t pose uint8;

uint32_t u4 uint32;

uint64_t targetId uint64;
uint32_t u6 uint32;
uint32_t u7 uint32;
uint64_t mainWeaponModel uint64;
uint64_t secWeaponModel uint64;
uint64_t craftToolModel uint64;

uint32_t u14 uint32;
uint32_t u15 uint32;
uint32_t bNPCBase uint32;
uint32_t bNPCName uint32;
uint32_t levelId uint32;
uint32_t u19 uint32;
uint32_t directorId uint32;
uint32_t spawnerId uint32;
uint32_t parentActorId uint32;
uint32_t hPMax uint32;
uint32_t hPCurr uint32;
uint32_t displayFlags uint32;
uint16_t fateID uint16;
uint16_t mPCurr uint16;
uint16_t tPCurr uint16;
uint16_t mPMax uint16;
uint16_t tPMax uint16;
uint16_t modelChara uint16;
uint16_t rotation uint16;
uint16_t activeMinion uint16;
uint8_t spawnIndex uint8;
uint8_t state uint8;
uint8_t persistantEmote uint8;
uint8_t modelType uint8;
uint8_t subtype uint8;
uint8_t voice uint8;
uint16_t u25c uint16;
uint8_t enemyType uint8;
uint8_t level uint8;
uint8_t classJob uint8;
uint8_t u26d uint8;
uint16_t u27a uint8;
uint8_t currentMount uint8;
uint8_t mountHead uint8;
uint8_t mountBody uint8;
uint8_t mountFeet uint8;
uint8_t mountColor uint8;
uint8_t scale uint8;
uint16_t elementalLevel uint16; // Eureka
uint16_t element uint16; // Eureka
uint32_t u30b uint32;
effect[30]StatusEffect;
pos FFXIVARR_POSITION3;
uint32_t models[10] uint32;
char name[32]byte;
uint8_t look[26]uint8;
char fcTag[6]byte;
uint32_t unk30 uint32;
uint32_t unk31 uint32;
uint8_t bNPCPartSlot uint8;
uint8_t unk32 uint8;
uint16_t unk33 uint16;
uint32_t unk34 uint32;
};

/**
* Structural representation of the packet sent by the server
* to show player movement
*/
struct ActorFreeSpawn
{
uint32_t spawnId uint32;
uint32_t actorId uint32;
};

/**
* Structural representation of the packet sent by the server
* to show player movement
*/
struct ActorMove
{
/* 0000 */ uint8_t headRotation uint8;
/* 0001 */ uint8_t rotation uint8;
/* 0002 */ uint8_t animationType uint8;
/* 0003 */ uint8_t animationState uint8;
/* 0004 */ uint8_t animationSpeed uint8;
/* 0005 */ uint8_t unknownRotation uint8;
/* 0006 */ uint16_t posX uint16;
/* 0008 */ uint16_t posY uint16;
/* 000a */ uint16_t posZ uint16;
/* 000C */ uint32_t unknown_12 uint32;
};

/**
* Structural representation of the packet sent by the server
* to set an actors position
*/
struct ActorSetPos
{
uint16_t r16 uint16;
uint8_t waitForLoad uint8;
uint8_t unknown1 uint8;
uint32_t unknown2 uint32;
float x float32;
float y float32;
float z float32;
uint32_t unknown3 uint32;

};


/**
* Structural representation of the packet sent by the server
* to start an actors casting
*/
struct ActorCast
{
uint16_t action_id uint16;
skillType uint8;
uint8_t unknown uint8;
uint32_t unknown_1 uint32; // action id or mount id
float cast_time float32;
uint32_t target_id uint32;
uint16_t rotation uint16;
uint16_t flag uint16; // 1 = interruptible blinking cast bar
uint32_t unknown_2 uint32;
uint16_t posX uint16;
uint16_t posY uint16;
uint16_t posZ uint16;
uint16_t unknown_3 uint16;
};

struct HateList
{
uint32_t numEntries uint32;
entry[32]struct
{
	uint32_t actorId uint32;
	uint8_t hatePercent uint8;
	uint8_t unknown uint8;
	uint16_t padding uint16;
};
uint32_t padding uint32;
};

struct HateRank
{
uint32_t numEntries uint32;
entry[32]struct
{
	uint32_t actorId uint32;
	uint32_t hateAmount uint32;
};
uint32_t padding uint32;
};

struct UpdateClassInfo
{
uint8_t classId uint8;
uint8_t level1 uint8;
uint16_t level uint16;
uint32_t nextLevelIndex uint32;
uint32_t currentExp uint32;
uint32_t restedExp uint32;
};

/**
* Structural representation of the packet sent by the server
* to send the titles available to the player
*/
struct PlayerTitleList
{
uint8_t titleList[48]uint8;
};

/**
* Structural representation of the packet sent by the server
* to initialize a zone for the player
*/
struct InitZone
{
uint16_t serverId uint16;
uint16_t zoneId uint16;
uint16_t unknown1 uint16;
uint16_t contentfinderConditionId uint16;
uint32_t unknown3 uint32;
uint32_t unknown4 uint32;
uint8_t weatherId uint8;
uint8_t bitmask uint8;
uint8_t bitmask1 uint8;
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

uint8_t unknown5 uint8;
uint32_t unknown8 uint32;
uint16_t festivalId uint16;
uint16_t additionalFestivalId uint16;
uint32_t unknown9 uint32;
uint32_t unknown10 uint32;
uint32_t unknown11 uint32;
uint32_t unknown12[4] uint32;
uint32_t unknown13[3] uint32;
pos FFXIVARR_POSITION3;
uint32_t unknown14[3] uint32;
uint32_t unknown15 uint32;
};


/**
* Structural representation of the packet sent by the server to initialize
* the client UI upon initial connection.
*/
struct PlayerSetup
{
// plain C types for a bit until the packet is actually fixed.
// makes conversion between different editors easier.
uint64_t contentId uint64;
unsigned int unknown8 uint32;
unsigned int unknownC uint32;
unsigned int charId uint32;
unsigned int restedExp uint32;
unsigned int companionCurrentExp uint32;
unsigned int unknown1C uint32;
unsigned int fishCaught uint32;
unsigned int useBaitCatalogId uint32;
unsigned int unknown28 uint32;
unsigned short unknownPvp2C uint16;
unsigned short unknown3 uint16;
unsigned int pvpFrontlineOverallCampaigns uint32;
unsigned int unknownTimestamp34 uint32;
unsigned int unknownTimestamp38 uint32;
unsigned int unknown3C uint32;
unsigned int unknown40 uint32;
unsigned int unknown44 uint32;
float companionTimePassed float32;
unsigned int unknown4C uint32;
unsigned short unknown50 uint16;
unsigned short unknownPvp52[4]uint16;
unsigned short playerCommendations uint16;
unsigned short unknown5C uint16;
unsigned short unknown5E uint16;
unsigned short pvpFrontlineWeeklyCampaigns uint16;
unsigned short enhancedAnimaGlassProgress uint16;
unsigned short unknown64[4] uint16;
unsigned short pvpRivalWingsTotalMatches uint16;
unsigned short pvpRivalWingsTotalVictories uint16;
unsigned short pvpRivalWingsWeeklyMatches uint16;
unsigned short pvpRivalWingsWeeklyVictories uint16;
unsigned char maxLevel uint8;
unsigned char expansion uint8;
unsigned char unknown76 uint8;
unsigned char unknown77 uint8;
unsigned char race uint8;
unsigned char tribe uint8;
unsigned char gender uint8;
unsigned char currentJob uint8;
unsigned char currentClass uint8;
unsigned char deity uint8;
unsigned char namedayMonth uint8;
unsigned char namedayDay uint8;
unsigned char cityState uint8;
unsigned char homepoint uint8;
unsigned char unknown82 uint8;
unsigned char petHotBar uint8;
unsigned char companionRank uint8;
unsigned char companionStars uint8;
unsigned char companionSp uint8;
unsigned char companionUnk86 uint8;
unsigned char companionColor uint8;
unsigned char companionFavoFeed uint8;
unsigned char unknown89 uint8;
unsigned char unknown8A[4] uint8;
unsigned char hasRelicBook uint8;
unsigned char relicBookId uint8;
unsigned char unknown90[4] uint8;
unsigned char craftingMasterMask uint8;
unsigned char unknown95[9] uint8;
unsigned char unknown9F[2] uint8;
unsigned char unknownA1[3] uint8;
unsigned int exp[CLASSJOB_SLOTS] uint32;
unsigned int unknown108 uint32;
unsigned int pvpTotalExp uint32;
unsigned int unknownPvp110 uint32;
unsigned int pvpExp uint32;
unsigned int pvpFrontlineOverallRanks[3] uint32;
unsigned short levels[CLASSJOB_SLOTS] uint16;
unsigned short unknown15C[9] uint16;
unsigned short u1 uint16;
unsigned short u2 uint16;
unsigned short unknown112[23] uint16;
unsigned short fishingRecordsFish[26] uint16;
unsigned short beastExp[11] uint16;
unsigned short unknown1EA[5] uint16;
unsigned short pvpFrontlineWeeklyRanks[3] uint16;
unsigned short unknownMask1FA[4] uint16;
unsigned char companionName[21] uint8;
unsigned char companionDefRank uint8;
unsigned char companionAttRank uint8;
unsigned char companionHealRank uint8;
unsigned char u19[8] uint8;
unsigned char mountGuideMask[22] uint8;
char name[32]byte;
unsigned char unknownOword[16] uint8;
unsigned char unknownOw uint8;
unsigned char unlockBitmask[64] uint8;
unsigned char aetheryte[21] uint8;
unsigned char discovery[445]; uint8
unsigned char howto[34] uint8;
unsigned char minions[45] uint8;
unsigned char chocoboTaxiMask[10] uint8;
unsigned char watchedCutscenes[124] uint8;
unsigned char companionBardingMask[10] uint8;
unsigned char companionEquippedHead uint8;
unsigned char companionEquippedBody uint8;
unsigned char companionEquippedLegs uint8;
unsigned char unknown52A[4] uint8;
unsigned char unknownMask52E[11] uint8;
unsigned char fishingGuideMask[105] uint8;
unsigned char fishingSpotVisited[31] uint8;
unsigned char unknown59A[27] uint8;
unsigned char unknown5A9[7] uint8;
unsigned char beastRank[11] uint8;
unsigned char unknownPvp5AB[11] uint8;
unsigned char unknown5B9[5] uint8;
unsigned char pose uint8;
unsigned char unknown5B91 uint8;
unsigned char challengeLogComplete[9] uint8;
unsigned char weaponPose uint8;
unsigned char unknownMask673[10] uint8;
unsigned char unknownMask5DD[28] uint8;
unsigned char relicCompletion[12] uint8;
unsigned char sightseeingMask[26] uint8;
unsigned char huntingMarkMask[55] uint8;
unsigned char tripleTriadCards[32] uint8;
unsigned char u12[11] uint8;
unsigned char u13 uint8;
unsigned char aetherCurrentMask[22] uint8;
unsigned char u10[3] uint8;
unsigned char orchestrionMask[40] uint8;
unsigned char hallOfNoviceCompletion[3] uint8;
unsigned char animaCompletion[11] uint8;
unsigned char u14[16] uint8;
unsigned char u15[13] uint8;
unsigned char unlockedRaids[28] uint8;
unsigned char unlockedDungeons[18] uint8;
unsigned char unlockedGuildhests[10] uint8;
unsigned char unlockedTrials[8] uint8;
unsigned char unlockedPvp[5] uint8;
unsigned char clearedRaids[28] uint8;
unsigned char clearedDungeons[18] uint8;
unsigned char clearedGuildhests[10] uint8;
unsigned char clearedTrials[8] uint8;
unsigned char clearedPvp[5] uint8;
unsigned short fishingRecordsFishWeight[26] uint16;
unsigned int exploratoryMissionNextTimestamp uint32;
unsigned char pvpLevel uint8;
};


/**
* Structural representation of the packet sent by the server
* to set a players stats
*/
struct PlayerStats
{
// order comes from baseparam order column
uint32_t strength uint32;
uint32_t dexterity uint32;
uint32_t vitality uint32;
uint32_t intelligence uint32;
uint32_t mind uint32;
uint32_t piety uint32;
uint32_t hp uint32;
uint32_t mp uint32;
uint32_t tp uint32;
uint32_t gp uint32;
uint32_t cp uint32;
uint32_t delay uint32;
uint32_t tenacity uint32;
uint32_t attackPower uint32;
uint32_t defense uint32;
uint32_t directHitRate uint32;
uint32_t evasion uint32;
uint32_t magicDefense uint32;
uint32_t criticalHit uint32;
uint32_t attackMagicPotency uint32;
uint32_t healingMagicPotency uint32;
uint32_t elementalBonus uint32;
uint32_t determination uint32;
uint32_t skillSpeed uint32;
uint32_t spellSpeed uint32;
uint32_t haste uint32;
uint32_t craftsmanship uint32;
uint32_t control uint32;
uint32_t gathering uint32;
uint32_t perception uint32;

// todo: what is here?
uint32_t unknown[26] uint32;
};

/**
* Structural representation of the packet sent by the server
* to set an actors current owner
*/
struct ActorOwner
{
uint8_t actorType uint8; // Note: Changed "type" to "actorType"
uint8_t padding[7] uint8;
uint32_t actorId uint32;
uint32_t actorId2 uint32;
};

/**
* Structural representation of the packet sent by the server
* to set a players state
*/
struct PlayerStateFlags
{
uint8_t flags[12]uint8;
uint32_t padding uint32;
};

/**
* Structural representation of the packet sent by the server
* containing current class information
*/
struct PlayerClassInfo
{
uint16_t classId uint32;
uint8_t unknown uint8;
uint8_t isSpecialist uint8;
uint16_t syncedLevel uint16;   // Locks actions, equipment, prob more. Player's current level (synced).
uint16_t classLevel uint16;  // Locks roles, prob more. Player's actual unsynced level.
uint32_t roleActions[10] uint32;
};

/**
* Structural representation of the packet sent by the server
* to update a players appearance
*/
struct ModelEquip
{
/* 0000 */ uint64_t mainWeapon uint64;
/* 0008 */ uint64_t offWeapon uint64;
/* 0010 */ uint8_t unk1 uint8;
/* 0011 */ uint8_t classJobId uint8;
/* 0012 */ uint8_t level uint8;
/* 0013 */ uint8_t unk2 uint8;
/* 0014 */ uint32_t models[10] uint32;
/* 003C */ uint32_t padding2 uint32;
};

struct Examine
{
uint8_t unkFlag1 uint8;
uint8_t unkFlag2 uint8;
char classJob byte;
char level byte;
uint16_t padding uint16;
uint16_t titleId uint16;
char grandCompany byte;
char grandCompanyRank byte;

char unknown[6]byte;
uint32_t u6_fromPSpawn uint32;
uint32_t u7_fromPSpawn uint32;
char padding1[8]byte;
uint64_t mainWeaponModel uint64;
uint64_t secWeaponModel uint64;
uint8_t unknown2 uint8;
uint16_t worldId uint16;
char unknown3[12]byte;
entries[14]struct ItemData
{
	uint32_t catalogId uint32;
	uint32_t appearanceCatalogId uint32;
	uint64_t crafterId uint64;
	uint8_t quality uint8;
	uint8_t unknown[3] uint8;
	materia[5]struct Materia
	{
	uint16_t materiaId uint16;
	uint16_t tier uint16;
	};
};
char name[32] byte;
char padding2 byte;
char unk3[16] byte;
char look[26] byte;
char padding3[5] byte;
uint32_t models[10] uint32;
char unknown4[200] byte;
};

struct CharaNameReq
{
uint64_t contentId uint64;
char name[32]byte;
};

/**
* Structural representation of the packet sent by the server
* to update a players appearance
*/
struct ItemInfo
{
uint32_t containerSequence uint32;
uint32_t unknown uint32;
uint16_t containerId uint16;
uint16_t slot uint16;
uint32_t quantity uint32;
uint32_t catalogId uint32;
uint32_t reservedFlag uint32;
uint64_t signatureId uint64;
uint8_t hqFlag uint8;
uint8_t unknown2 uint8;
uint16_t condition uint16;
uint16_t spiritBond uint16;
uint16_t stain uint16;
uint32_t glamourCatalogId uint32;
uint16_t materia1 uint16;
uint16_t materia2 uint16;
uint16_t materia3 uint16;
uint16_t materia4 uint16;
uint16_t materia5 uint16;
uint8_t tier1 uint8;
uint8_t tier2 uint8;
uint8_t tier3 uint8;
uint8_t tier4 uint8;
uint8_t tier5 uint8;
uint8_t padding uint8;
uint32_t unknown10 uint32;
};

/**
* Structural representation of the packet sent by the server
* to update a players appearance
*/
struct ContainerInfo
{
uint32_t containerSequence uint32;
uint32_t numItems uint32;
uint32_t containerId uint32;
uint32_t unknown uint32;
};

/**
* Structural representation of the packet sent by the server
* to update a players appearance
*/
struct CurrencyCrystalInfo
{
uint32_t containerSequence uint32;
uint16_t containerId uint16;
uint16_t slot uint16;
uint32_t quantity uint32;
uint32_t unknown uint32;
uint32_t catalogId uint32;
uint32_t unknown1 uint32;
uint32_t unknown2 uint32;
uint32_t unknown3 uint32;
};

struct InventoryTransactionFinish
{
uint32_t sequenceId uint32;
uint32_t sequenceId1 uint32;
uint64_t padding uint64;
};

struct InventoryTransaction
{
uint32_t sequence uint32;
uint8_t transactionType uint8; // Note: Changed "type" to "transactionType"
uint8_t padding uint8;
uint16_t padding1 uint16;
uint32_t ownerId uint32;
uint32_t storageId uint32;
uint16_t slotId uint16;
uint16_t padding2 uint16;
uint32_t stackSize uint32;
uint32_t catalogId uint32;
uint32_t someActorId uint32;
int32_t targetStorageId int32;
uint32_t padding3[3] uint32;
};


struct InventoryActionAck
{
uint32_t sequence uint32;
uint16_t actionType uint16; // Note: Changed "type" to "actionType"
uint16_t padding uint16;
uint32_t padding1 uint32;
uint32_t padding2 uint32;
};


/**
* Structural representation of the packet sent by the server
* to update a slot in the inventory
*/
struct UpdateInventorySlot
{
uint32_t sequence uint32;
uint32_t unknown uint32;
uint16_t containerId uint16;
uint16_t slot uint16;
uint32_t quantity uint32;
uint32_t catalogId uint32;
uint32_t reservedFlag uint32;
uint64_t signatureId uint64;
uint16_t hqFlag uint16;
uint16_t condition uint16;
uint16_t spiritBond uint16;
uint16_t color uint16;
uint32_t glamourCatalogId uint32;
uint16_t materia1 uint16;
uint16_t materia2 uint16;
uint16_t materia3 uint16;
uint16_t materia4 uint16;
uint16_t materia5 uint16;
uint8_t tier1 uint8;
uint8_t tier2 uint8;
uint8_t tier3 uint8;
uint8_t tier4 uint8;
uint8_t tier5 uint8;
uint8_t padding uint8;
uint32_t unknown10 uint32;
};

/**
* Structural representation of the packet sent by the server
* to start an event, not actually playing it, but registering
*/
struct EventStart
{
/* 0000 */ uint64_t actorId uint64;
/* 0008 */ uint32_t eventId uint32;
/* 000C */ uint8_t param1 uint8;
/* 000D */ uint8_t param2 uint8;
/* 000E */ uint16_t padding uint16;
/* 0010 */ uint32_t param3 uint32;
/* 0014 */ uint32_t padding1 uint32;
};

/**
* Structural representation of the packet sent by the server
* to fill a huntin log entry
*/
struct HuntingLogEntry
{
int32_t u0 int32; // -1 for all normal classes
uint8_t rank uint8; // starting from 0
uint8_t index uint8; // classes and gcs
uint8_t entries[10][4] uint8;
uint16_t pad uint16;
uint64_t completeFlags uint64; // 4 bit for each potential entry and the 5th bit for completion of the section
uint64_t pad1 uint64;
};

/**
* Structural representation of the packet sent by the server
* to play an event
*/
struct EventPlay
{
uint64_t actorId uint64;
uint32_t eventId uint32;
uint16_t scene uint16;
uint16_t padding uint16;
uint32_t flags uint32;
uint32_t param3 uint32;
uint8_t param4 uint8;
uint8_t padding1[3] uint8;
uint32_t param5 uint32;
uint8_t unknown[8] uint8;
};

/**
* Structural representation of the packet sent by the server
* to play an event
*/
struct DirectorPlayScene
{
uint64_t actorId uint64;
uint32_t eventId uint32;
uint16_t scene uint16;
uint16_t padding uint16;
uint32_t flags uint32;
uint32_t param3 uint32;
uint8_t param4 uint8;
uint8_t padding1[3] uint8;
uint32_t param5 uint32;
uint8_t unknown8[0x08] uint8;
uint8_t unknown[0x38] uint8;
};

/**
* Structural representation of the packet sent by the server
* to finish an event
*/
struct EventFinish
{
/* 0000 */ uint32_t eventId uint32;
/* 0004 */ uint8_t param1 uint8;
/* 0005 */ uint8_t param2 uint8;
/* 0006 */ uint16_t padding uint16;
/* 0008 */ uint32_t param3 uint32;
/* 000C */ uint32_t padding1 uint32;
};

struct EventPlayN
{
uint64_t actorId uint64;
uint32_t eventId uint32;
uint16_t scene uint16;
uint16_t padding uint16;
uint32_t sceneFlags uint32;
uint8_t paramCount uint8;
uint8_t padding2[3] uint8;
uint32_t params[1] uint32;
};

struct EventPlay255
{
uint64_t actorId uint64;
uint32_t eventId uint32;
uint16_t scene uint16;
uint16_t padding uint16;
uint32_t sceneFlags uint32;
uint8_t paramCount uint8;
uint8_t padding2[3] uint8;
uint32_t params[255] uint32;
};

/**
* Structural representation of the packet sent by the server
* to respond to a linkshell creation event
*/
struct EventLinkshell
{
uint32_t eventId uint32;
uint8_t scene uint8;
uint8_t param1 uint8;
uint8_t param2 uint8;
uint8_t param3 uint8;
uint32_t unknown1 uint32;
uint32_t unknown2 uint32;
uint32_t unknown3 uint32;
uint32_t unknown4 uint32;
};

/**
* Structural representation of the packet sent by the server
* to send the active quests
*/
struct QuestActiveList
{
Common::QuestActive activeQuests[30]; // Uh wtf
};

/**
* Structural representation of the packet sent by the server
* to send update a quest slot
*/
struct QuestUpdate
{
uint16_t slot uint16;
uint16_t padding uint16;
Common::QuestActive questInfo; // Uh wtf
};

/**
* Structural representation of the packet sent by the server
* to send the completed quests mask
*/
struct QuestCompleteList
{
uint8_t questCompleteMask[480]uint8;
uint8_t unknownCompleteMask[80]uint8;
};

/**
* Structural representation of the packet sent by the server
* to finish a quest
*/
struct QuestFinish
{
uint16_t questId uint16;
uint8_t flag1 uint8;
uint8_t flag2 uint8;
uint32_t padding uint32;
};

/**
* Structural representation of the packet sent by the server
* to send a quest message
* type 0 default
* type 1 icon
* type 5 status
*/
struct QuestMessage
{
/* 0000 */ uint32_t questId uint32;
/* 0000 */ uint8_t msgId uint8;
/* 0000 */ uint8_t questType uint8; // Note: Changed "type" to "questType"
/* 0000 */ uint16_t padding1 uint16;
/* 0000 */ uint32_t var1 uint32;
/* 0000 */ uint32_t var2 uint32;
};

struct QuestTracker
{
	entry[5]struct
{
	uint8_t active uint8;
	uint8_t questIndex uint8;
};
uint16_t padding[3] uint16;
};


struct WeatherChange
{
uint32_t weatherId uint32;
float delay float32;
};

/**
* Structural representation of the packet sent by the server
* to send a unviel a map
*/
struct Discovery
{
/* 0000 */ uint32_t mapPartId uint32;
/* 0004 */ uint32_t mapId uint32;
};

/**
* UNKOWN TYPE
*/
struct FFXIVARR_IPC_UNK322
{
/* 0000 */ uint8_t unk[8] uint8;
};

/**
* UNKOWN TYPE
*/
struct FFXIVARR_IPC_UNK320
{
/* 0000 */ uint8_t unk[0x38] uint8;
};

/**
* Structural representation of the packet sent by the server
* prepare zoning, showing screenmessage
*/
struct PrepareZoning
{
uint32_t logMessage uint32;
uint16_t targetZone uint16;
uint16_t animation uint16;
uint8_t param4 uint8;
uint8_t hideChar uint8;
uint8_t fadeOut uint8;
uint8_t param7 uint8;
uint8_t fadeOutTime uint8;
uint8_t unknown uint8; // this changes whether or not the destination zone's name displays during the loading screen. Seems to always be 9 (=hidden) when going to an instance and certain zones, 0 otherwise.
uint16_t padding uint16;
};

/**
* Structural representation of the packet sent by the server
* to trigger content finder events
*
* See https://gist.github.com/Minoost/c35843c4c8a7a931f31fdaac9bce64c2
*/
struct CFNotify
{
uint32_t state1 uint32; // 3 = cancelled, 4 = duty ready
uint32_t state2 uint32; // if state1 == 3, state2 is cancelled reason

uint32_t param1 uint32; // usually classJobId
uint32_t param2 uint32; // usually flag
uint32_t param3 uint32; // usually languages, sometimes join in progress timestamp

uint16_t param4 uint16; // usually roulette id
uint16_t contents[5] uint16;
};

/**
* Structural representation of the packet sent by the server
* to update contents available in duty finder or raid finder
*
* Do note that this packet has to come early in login phase (around initui packet)
* or it won't be applied until you reconnect
*/
struct CFAvailableContents
{
uint8_t contents[0x48]uint8;
};

/**
* Structural representation of the packet sent by the server
* to update adventure in needs in duty roulette
*/
struct CFPlayerInNeed
{
// Ordered by roulette id
uint8_t inNeeds[0x10]uint8;
};

/**
* Structural representation of the packet sent by the server
* to update duty info in general
*/
struct CFDutyInfo
{
uint8_t penaltyTime uint8;
uint8_t unknown[7]uint8;
};

struct CFRegisterDuty
{
uint32_t unknown0 uint32; // 0x301
uint8_t rouletteId uint8; // if it's a daily roulette
uint8_t unknown1 uint8; // 0xDB
uint16_t contentId uint16;
};


struct CFMemberStatus
{
uint16_t contentId uint16;
uint16_t unknown1 uint16;
uint8_t status uint8;
uint8_t currentTank uint8;
uint8_t currentDps uint8;
uint8_t currentHealer uint8;
uint8_t estimatedTime uint8;
uint8_t unknown2[3] uint8;
uint32_t unknown3 uint32;
};

struct EorzeaTimeOffset
{
uint64_t timestamp uint64;
};

/**
* Structural representation of the packet sent by the server
* to set the gear show/hide status of a character
*/
struct EquipDisplayFlags
{
uint8_t bitmask uint8;
};

/**
* Structural representation of the packet sent by the server
* to mount a player
*/
struct Mount
{
uint32_t id uint32;
};

/**
* Structural representation of the packet sent by the server
* to mount a player
*/
struct DirectorVars
{
/*! DirectorType | ContentId */
uint32_t m_directorId;
/*! currect sequence */
uint8_t m_sequence;
/*! current branch */
uint8_t m_branch;
/*! raw storage for flags/vars */
uint8_t m_unionData[10];
/*! unknown */
uint16_t u20;
uint16_t u22;
uint16_t u24;
uint16_t u28;
};

struct DirectorPopUp
{
uint32_t directorId;
uint16_t pad1[2];
uint64_t sourceActorId;
/*!
	* 2 = green text in log
	*/
uint8_t flags;
uint8_t pad2[3];
uint32_t bNPCName;
uint32_t textId;
uint32_t popupTimeMs;
uint32_t pad3[4];
};


struct ActorGauge
{
uint8_t classJobId;
uint8_t data[15]; // depends on classJobId
};

struct PerformNote
{
uint8_t data[32];
};

struct HousingUpdateLandFlagsSlot
{
uint32_t type;
uint32_t unknown;
Common::LandFlagSet flagSet;
};

struct HousingLandFlags
{
Common::LandFlagSet freeCompanyHouse; // 00
uint64_t unkown1;
Common::LandFlagSet privateHouse; // 24
uint64_t unkown2;
Common::LandFlagSet apartment; // 48
uint64_t unkown3;
Common::LandFlagSet sharedHouse[2]; //72
uint64_t unkown4;
Common::LandFlagSet unkownHouse;
uint64_t unkown5;
};

//Structs
struct LandStruct
{
uint8_t plotSize; //0
uint8_t houseState; // 2
uint8_t flags; // bit1 -> hasPublicAccess; bit2 -> isPersonalHouse
uint8_t iconAddIcon; // 6
uint32_t fcId; //8
uint32_t fcIcon;// 12
uint32_t fcIconColor; // 16
uint16_t housePart[ 8 ]; // 34
uint8_t houseColour[ 8 ]; // 36
};

struct LandUpdate
{
Common::LandIdent landIdent;
LandStruct land;
};

struct LandPriceUpdate
{
uint32_t price;
uint32_t timeLeft;
};

struct LandInfoSign
{
Common::LandIdent landIdent;
uint64_t ownerId; // ither contentId or fcId
uint32_t unknow1;
uint8_t houseIconAdd;
uint8_t houseSize;
uint8_t houseType;
char estateName[23];
char estateGreeting[193];
char ownerName[31];
char fcTag[7];
uint8_t tag[3];
};

struct LandRename
{
Common::LandIdent landIdent;
char houseName[20];
uint32_t padding;
};

struct LandUpdateHouseName
{
uint32_t unknown[3];
char houseName[20];
uint32_t unknown2[2];
};

struct LandSetMap
{
uint8_t u1;
uint8_t subdivision;
uint8_t u3;
struct
{
	uint8_t status;
	uint8_t size;
	uint8_t isPrivate;
} landInfo[ 30 ];
uint8_t padding[ 3 ];
};

struct LandSetInitialize
{
Common::LandIdent landIdent;
uint8_t unknown1;
uint8_t subInstance; //  (default
uint8_t unknown3;
uint8_t unknown4;
uint8_t unknown5;
uint8_t unknown6;
uint8_t unknown7;
uint8_t unknown8;
LandStruct land[ 30 ];
};

struct YardObjectSpawn
{
uint8_t landId;
uint8_t objectArray;
uint16_t unknown1;
Common::HousingObject object;
};

struct HousingObjectMove
{
uint16_t itemRotation;
uint8_t objectArray;
uint8_t landId;
Common::FFXIVARR_POSITION3 pos;
uint16_t unknown1;
uint16_t unknown2;
uint16_t unknown3;
};

struct HousingObjectInitialize
{
Common::LandIdent landIdent;
/*!
	* when this is 2, actrl 0x400 will hide the additional quarters door
	* if it's any other value, it will stay there regardless
	*/
int8_t u1; //Outdoor -1 / Indoor 0 - probably indicator
uint8_t packetNum;
uint8_t packetTotal;
uint8_t u2; //Outdoor 0 / Indoor 100(?)
Common::HousingObject object[100];
uint32_t unknown4; //unused
};

struct HousingInternalObjectSpawn
{
uint16_t containerId;
uint8_t containerOffset;
uint8_t pad1;

Common::HousingObject object;
};

struct HousingIndoorInitialize
{
uint16_t u1;
uint16_t u2;
uint16_t u3;
uint16_t u4;
uint32_t indoorItems[10];
};


struct HousingWardInfo
{
Common::LandIdent landIdent;

struct HouseInfoEntry
{
	uint32_t housePrice;
	uint8_t infoFlags;
	Common::HousingAppeal houseAppeal[3];
	char estateOwnerName[30];
} houseInfoEntry[60];
};

struct HousingEstateGreeting
{
Common::LandIdent landIdent;
char message[200];
};

struct HousingShowEstateGuestAccess
{
uint32_t unknown[2];
Common::LandIdent ident;
};

/**
* Structural representation of the packet sent by the server
* to show the current shared estate settings
*/
struct SharedEstateSettingsResponse
{
struct playerEntry
{
	uint64_t contentId;
	uint8_t permissions;
	char name[0x20];
	char padding[0x7];
} entry[3];
};

struct MSQTrackerProgress
{
uint32_t id;
uint32_t padding;
};

struct MSQTrackerComplete
{
uint32_t id;
uint32_t padding1;
uint64_t padding2;
uint64_t padding3;
uint64_t padding4; // last 4 bytes is uint32_t but who cares
};

struct ObjectSpawn
{
uint8_t spawnIndex;
uint8_t objKind;
uint8_t state;
uint8_t unknown3;
uint32_t objId;
uint32_t actorId;
uint32_t levelId;
uint32_t unknown10;
uint32_t someActorId14;
uint32_t gimmickId;
float scale;
int16_t unknown20a;
uint16_t rotation;
int16_t unknown24a;
int16_t unknown24b;
uint16_t unknown28a;
int16_t unknown28c;
uint32_t housingLink;
Common::FFXIVARR_POSITION3 position;
int16_t unknown3C;
int16_t unknown3E;
};

struct ObjectDespawn
{
uint8_t spawnIndex;
uint8_t padding[7];
};

struct DuelChallenge
{
uint8_t otherClassJobId;
uint8_t otherLevel; // class job level
uint8_t challengeByYou; // 0 if the other challenges you, 1 if you challenges the other.
uint8_t otherItemLevel;

uint32_t otherActorId;

char otherName[32];
};