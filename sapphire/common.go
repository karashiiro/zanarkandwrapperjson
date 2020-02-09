package sapphire

type EffectEntry struct {
	effectType uint8
	param0     uint8
	param1     uint8
	/*!
	 * @brief Shows an additional percentage in the battle log
	 *
	 * Has no effect on what is shown and stored in value
	 */
	param2          uint8
	valueMultiplier uint8 // This multiplies whatever value is in the 'value' param by 10. Possibly a workaround for big numbers
	flags           uint8
	value           int16
}

type HousingObject struct {
	itemId   uint32
	padding  uint32 // was itemrotation + unknown/pad, looks unused now
	rotation float32
	pos      FFXIVARR_POSITION3
}

type LandIdent struct {
	landId          uint16
	wardNum         uint16
	territoryTypeId uint16
	worldId         uint16
}

type LandFlagSet struct {
	landIdent LandIdent
	landFlags uint32
	unkown1   uint32
}

type QuestActive struct {
	questId   uint16
	sequence  uint8
	flags     uint8
	padding   uint8
	BitFlag48 uint8
	BitFlag40 uint8
	BitFlag32 uint8
	BitFlag24 uint8
	BitFlag16 uint8
	BitFlag8  uint8
	padding1  uint8
}

type StatusEffect struct {
	effect_id     uint16
	param         uint16
	duration      float32
	sourceActorId uint32
}

const MAX_DISPLAYED_ACTORS uint8 = 99
const MAX_DISPLAYED_EOBJS uint8 = 40

const INVALID_GAME_OBJECT_ID uint32 = 0xE0000000
const INVALID_GAME_OBJECT_ID64 uint64 = 0xE0000000

const MAX_PLAYER_LEVEL uint16 = 80
const CURRENT_EXPANSION_ID uint8 = 3

const CLASSJOB_TOTAL uint8 = 38
const CLASSJOB_SLOTS uint8 = 28

const TOWN_COUNT uint8 = 6
