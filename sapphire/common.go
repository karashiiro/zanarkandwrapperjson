package sapphire

type EffectEntry struct {
	EffectType uint8
	Param0     uint8
	Param1     uint8
	/*!
	 * @Brief Shows an additional percentage in the battle log
	 *
	 * Has no effect on what is shown and stored in value
	 */
	Param2          uint8
	ValueMultiplier uint8 // This multiplies whatever value is in the 'value' param by 10. Possibly a workaround for big numbers
	Flags           uint8
	Value           int16
}

type HousingObject struct {
	ItemId   uint32
	Padding  uint32 // was itemrotation + unknown/pad, looks unused now
	Rotation float32
	Pos      FFXIVARR_POSITION3
}

type LandIdent struct {
	LandId          uint16
	WardNum         uint16
	TerritoryTypeId uint16
	WorldId         uint16
}

type LandFlagSet struct {
	LandIdent LandIdent
	LandFlags uint32
	Unkown1   uint32
}

type QuestActive struct {
	QuestId   uint16
	Sequence  uint8
	Flags     uint8
	Padding   uint8
	BitFlag48 uint8
	BitFlag40 uint8
	BitFlag32 uint8
	BitFlag24 uint8
	BitFlag16 uint8
	BitFlag8  uint8
	Padding1  uint8
}

type StatusEffect struct {
	Effect_id     uint16
	Param         uint16
	Duration      float32
	SourceActorId uint32
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
