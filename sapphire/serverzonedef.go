package sapphire

// ActorControl - Structural representation of the packet sent by the server
// to update certain player details / status
type ActorControl struct {
	Category uint16 `json:"category"`
	Padding  uint16 `json:"padding"`
	Param1   uint32 `json:"param1"`
	Param2   uint32 `json:"param2"`
	Param3   uint32 `json:"param3"`
	Param4   uint32 `json:"param4"`
	Padding1 uint32 `json:"padding1"`
}

// ActorControlSelf - Structural representation of the packet sent by the server
// to update certain player details / status
type ActorControlSelf struct {
	Category uint16 `json:"category"`
	Padding  uint16 `json:"padding"`
	Param1   uint32 `json:"param1"`
	Param2   uint32 `json:"param2"`
	Param3   uint32 `json:"param3"`
	Param4   uint32 `json:"param4"`
	Param5   uint32 `json:"param5"`
	Param6   uint32 `json:"param6"`
	Padding1 uint32 `json:"padding1"`
}

// EffectResult - Structural representation of the packet sent by the server
// to add a status effect
type EffectResult struct {
	GlobalSequence   uint32 `json:"globalSequence"`
	ActorID          uint32 `json:"actor_id"`
	CurrentHp        uint32 `json:"current_hp"`
	MaxHp            uint32 `json:"max_hp"`
	CurrentMp        uint16 `json:"current_mp"`
	Unknown1         uint8  `json:"unknown1"`
	ClassID          uint8  `json:"classId"`
	ShieldPercentage uint8  `json:"shieldPercentage"`
	EntryCount       uint8  `json:"entryCount"`
	Unknown2         uint16 `json:"unknown2"`

	StatusEntries [4]struct {
		Index         uint8   `json:"index"` // which position do i display this
		Unknown3      uint8   `json:"unknown3"`
		ID            uint16  `json:"id"`
		Param         uint16  `json:"param"`
		Unknown4      uint16  `json:"unknown4"` // Sort this out (old right half of power/param property)
		Duration      float32 `json:"duration"`
		SourceActorID uint32  `json:"sourceActorId"`
	} `json:"statusEntries"`
}

type UpdateClassInfo struct {
	ClassID        uint8  `json:"classId"`
	Level1         uint8  `json:"level1"`
	Level          uint16 `json:"level"`
	NextLevelIndex uint32 `json:"nextLevelIndex"`
	CurrentExp     uint32 `json:"currentExp"`
	RestedExp      uint32 `json:"restedExp"`
}

// InitZone - Structural representation of the packet sent by the server
// to initialize a zone for the player
type InitZone struct {
	ServerID                 uint16 `json:"serverId"`
	ZoneID                   uint16 `json:"zoneId"`
	Unknown1                 uint16 `json:"unknown1"`
	ContentfinderConditionID uint16 `json:"contentFinderConditionId"`
	Unknown3                 uint32 `json:"unknown3"`
	Unknown4                 uint32 `json:"unknown4"`
	WeatherID                uint8  `json:"weatherId"`
	Bitmask                  uint8  `json:"bitmask"`
	Bitmask1                 uint8  `json:"bitmask1"`
	// Bitmask1 findings
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

	Unknown5             uint8              `json:"unknown5"`
	Unknown8             uint32             `json:"unknown8"`
	FestivalID           uint16             `json:"festivalId"`
	AdditionalFestivalID uint16             `json:"additionalFestivalId"`
	Unknown9             uint32             `json:"unknown9"`
	Unknown10            uint32             `json:"unknown10"`
	Unknown11            uint32             `json:"unknown11"`
	Unknown12            [4]uint32          `json:"unknown12"`
	Unknown13            [3]uint32          `json:"unknown13"`
	Pos                  FFXIVARR_POSITION3 `json:"pos"`
	Unknown14            [3]uint32          `json:"unknown14"`
	Unknown15            uint32             `json:"unknown15"`
}

// PlayerSetup - Structural representation of the packet sent by the server to initialize
// the client UI upon initial connection.
type PlayerSetup struct {
	// plain C types for a bit until the packet is actually fixed.
	// makes conversion between different editors easier.
	ContentID                       uint64                 `json:"contentId"`
	Unknown8                        uint32                 `json:"unknown8"`
	UnknownC                        uint32                 `json:"unknownC"`
	CharID                          uint32                 `json:"charId"`
	RestedExp                       uint32                 `json:"restedExp"`
	CompanionCurrentExp             uint32                 `json:"companionCurrentExp"`
	Unknown1C                       uint32                 `json:"unknown1C"`
	FishCaught                      uint32                 `json:"fishCaught"`
	UseBaitCatalogID                uint32                 `json:"useBaitCatalogId"`
	Unknown28                       uint32                 `json:"unknown28"`
	UnknownPvp2C                    uint16                 `json:"unknownPvp2C"`
	Unknown3                        uint16                 `json:"unknown3"`
	PvpFrontlineOverallCampaigns    uint32                 `json:"pvpFrontlineOverallCampaigns"`
	UnknownTimestamp34              uint32                 `json:"unknownTimestamp34"`
	UnknownTimestamp38              uint32                 `json:"unknownTimestamp38"`
	Unknown3C                       uint32                 `json:"unknown3C"`
	Unknown40                       uint32                 `json:"unknown40"`
	Unknown44                       uint32                 `json:"unknown44"`
	CompanionTimePassed             float32                `json:"companionTimePassed"`
	Unknown4C                       uint32                 `json:"unknown4C"`
	Unknown50                       uint16                 `json:"unknown50"`
	UnknownPvp52                    [4]uint16              `json:"unknownPvp52"`
	PlayerCommendations             uint16                 `json:"playerCommendations"`
	Unknown5C                       uint16                 `json:"unknown5C"`
	Unknown5E                       uint16                 `json:"unknown5E"`
	PvpFrontlineWeeklyCampaigns     uint16                 `json:"pvpFrontlineWeeklyCampaigns"`
	EnhancedAnimaGlassProgress      uint16                 `json:"enhancedAnimaGlassProgress"`
	Unknown64                       [4]uint16              `json:"unknown64"`
	PvpRivalWingsTotalMatches       uint16                 `json:"pvpRivalWingsTotalMatches"`
	PvpRivalWingsTotalVictories     uint16                 `json:"pvpRivalWingsTotalVictories"`
	PvpRivalWingsWeeklyMatches      uint16                 `json:"pvpRivalWingsWeeklyMatches"`
	PvpRivalWingsWeeklyVictories    uint16                 `json:"pvpRivalWingsWeeklyVictories"`
	MaxLevel                        uint8                  `json:"maxLevel"`
	Expansion                       uint8                  `json:"expansion"`
	Unknown76                       uint8                  `json:"unknown76"`
	Unknown77                       uint8                  `json:"unknown77"`
	Race                            uint8                  `json:"race"`
	Tribe                           uint8                  `json:"tribe"`
	Gender                          uint8                  `json:"gender"`
	CurrentJob                      uint8                  `json:"currentJob"`
	CurrentClass                    uint8                  `json:"currentClass"`
	Deity                           uint8                  `json:"deity"`
	NamedayMonth                    uint8                  `json:"namedayMonth"`
	NamedayDay                      uint8                  `json:"namedayDay"`
	CityState                       uint8                  `json:"cityState"`
	Homepoint                       uint8                  `json:"homepoint"`
	Unknown82                       uint8                  `json:"unknown82"`
	PetHotBar                       uint8                  `json:"petHotBar"`
	CompanionRank                   uint8                  `json:"companionRank"`
	CompanionStars                  uint8                  `json:"companionStars"`
	CompanionSp                     uint8                  `json:"companionSp"`
	CompanionUnk86                  uint8                  `json:"companionUnk86"`
	CompanionColor                  uint8                  `json:"companionColor"`
	CompanionFavoFeed               uint8                  `json:"companionFavoFeed"`
	Unknown89                       uint8                  `json:"unknown89"`
	Unknown8A                       [4]uint8               `json:"unknown8A"`
	HasRelicBook                    uint8                  `json:"hasRelicBook"`
	RelicBookID                     uint8                  `json:"relicBookId"`
	Unknown90                       [4]uint8               `json:"unknown90"`
	CraftingMasterMask              uint8                  `json:"craftingMasterMask"`
	Unknown95                       [9]uint8               `json:"unknown95"`
	Unknown9F                       [2]uint8               `json:"unknown9F"`
	UnknownA1                       [3]uint8               `json:"unknownA1"`
	Exp                             [CLASSJOB_SLOTS]uint32 `json:"exp"`
	Unknown108                      uint32                 `json:"unknown108"`
	PvpTotalExp                     uint32                 `json:"pvpTotalExp"`
	UnknownPvp110                   uint32                 `json:"unknownPvp110"`
	PvpExp                          uint32                 `json:"pvpExp"`
	PvpFrontlineOverallRanks        [3]uint32              `json:"pvpFrontlineOverallRanks"`
	Levels                          [CLASSJOB_SLOTS]uint16 `json:"levels"`
	Unknown15C                      [9]uint16              `json:"unknown15C"`
	U1                              uint16                 `json:"u1"`
	U2                              uint16                 `json:"u2"`
	Unknown112                      [23]uint16             `json:"unknown112"`
	FishingRecordsFish              [26]uint16             `json:"fishingRecordsFish"`
	BeastExp                        [11]uint16             `json:"beastExp"`
	Unknown1EA                      [5]uint16              `json:"unknown1EA"`
	PvpFrontlineWeeklyRanks         [3]uint16              `json:"pvpFrontlineWeeklyRanks"`
	UnknownMask1FA                  [4]uint16              `json:"unknownMask1FA"`
	CompanionName                   [21]uint8              `json:"companionName"`
	CompanionDefRank                uint8                  `json:"companionDefRank"`
	CompanionAttRank                uint8                  `json:"companionAttRank"`
	CompanionHealRank               uint8                  `json:"companionHealRank"`
	U19                             [8]uint8               `json:"u19"`
	MountGuideMask                  [22]uint8              `json:"mountGuideMask"`
	Name                            [32]byte               `json:"name"`
	UnknownOword                    [16]uint8              `json:"unknownOword"`
	UnknownOw                       uint8                  `json:"unknownOw"`
	UnlockBitmask                   [64]uint8              `json:"unlockBitmask"`
	Aetheryte                       [21]uint8              `json:"aetheryte"`
	Discovery                       [445]uint8             `json:"discovery"`
	Howto                           [34]uint8              `json:"howto"`
	Minions                         [45]uint8              `json:"minions"`
	ChocoboTaxiMask                 [10]uint8              `json:"chocoboTaxiMask"`
	WatchedCutscenes                [124]uint8             `json:"watchedCutscenes"`
	CompanionBardingMask            [10]uint8              `json:"companionBardingMask"`
	CompanionEquippedHead           uint8                  `json:"companionEquippedHead"`
	CompanionEquippedBody           uint8                  `json:"companionEquippedBody"`
	CompanionEquippedLegs           uint8                  `json:"companionEquippedLegs"`
	Unknown52A                      [4]uint8               `json:"unknown52A"`
	UnknownMask52E                  [11]uint8              `json:"unknownMask52E"`
	FishingGuideMask                [105]uint8             `json:"fishingGuideMask"`
	FishingSpotVisited              [31]uint8              `json:"fishingSpotVisited"`
	Unknown59A                      [27]uint8              `json:"unknown59A"`
	Unknown5A9                      [7]uint8               `json:"unknown5A9"`
	BeastRank                       [11]uint8              `json:"beastRank"`
	UnknownPvp5AB                   [11]uint8              `json:"unknownPvp5AB"`
	Unknown5B9                      [5]uint8               `json:"unknown5B9"`
	Pose                            uint8                  `json:"pose"`
	Unknown5B91                     uint8                  `json:"unknown5B91"`
	ChallengeLogComplete            [9]uint8               `json:"challengeLogComplete"`
	WeaponPose                      uint8                  `json:"weaponPose"`
	UnknownMask673                  [10]uint8              `json:"unknownMask673"`
	UnknownMask5DD                  [28]uint8              `json:"unknownMask5DD"`
	RelicCompletion                 [12]uint8              `json:"relicCompletion"`
	SightseeingMask                 [26]uint8              `json:"sightseeingMask"`
	HuntingMarkMask                 [55]uint8              `json:"huntingMarkMask"`
	TripleTriadCards                [32]uint8              `json:"tripleTriadCards"`
	U12                             [11]uint8              `json:"u12"`
	U13                             uint8                  `json:"u13"`
	AetherCurrentMask               [22]uint8              `json:"aetherCurrentMask"`
	U10                             [3]uint8               `json:"u10"`
	OrchestrionMask                 [40]uint8              `json:"orchestrionMask"`
	HallOfNoviceCompletion          [3]uint8               `json:"hallOfNoviceCompletion"`
	AnimaCompletion                 [11]uint8              `json:"animaCompletion"`
	U14                             [16]uint8              `json:"u14"`
	U15                             [13]uint8              `json:"u15"`
	UnlockedRaids                   [28]uint8              `json:"unlockedRaids"`
	UnlockedDungeons                [18]uint8              `json:"unlockedDungeons"`
	UnlockedGuildhests              [10]uint8              `json:"unlockedGuildhests"`
	UnlockedTrials                  [8]uint8               `json:"unlockedTrials"`
	UnlockedPvp                     [5]uint8               `json:"unlockedPvp"`
	ClearedRaids                    [28]uint8              `json:"clearedRaids"`
	ClearedDungeons                 [18]uint8              `json:"clearedDungeons"`
	ClearedGuildhests               [10]uint8              `json:"clearedGuildhests"`
	ClearedTrials                   [8]uint8               `json:"clearedTrials"`
	ClearedPvp                      [5]uint8               `json:"clearedPvp"`
	FishingRecordsFishWeight        [26]uint16             `json:"fishingRecordsFishWeight"`
	ExploratoryMissionNextTimestamp uint32                 `json:"exploratoryMissionNextTimestamp"`
	PvpLevel                        uint8                  `json:"pvpLevel"`
}

// PlayerSpawn - Structural representation of the packet sent by the server
// to spawn an actor
type PlayerSpawn struct {
	Title          uint16 `json:"title"`
	U1b            uint16 `json:"u1b"`
	CurrentWorldID uint16 `json:"currentWorldId"`
	HomeWorldID    uint16 `json:"homeWorldId"`

	GmRank       uint8 `json:"gmRank"`
	U3c          uint8 `json:"u3c"`
	U4           uint8 `json:"u4"`
	OnlineStatus uint8 `json:"onlineStatus"`

	Pose uint8 `json:"pose"`
	U5a  uint8 `json:"u5a"`
	U5b  uint8 `json:"u5b"`
	U5c  uint8 `json:"u5c"`

	TargetID        uint64 `json:"targetId"`
	U6              uint32 `json:"u6"`
	U7              uint32 `json:"u7"`
	MainWeaponModel uint64 `json:"mainWeaponModel"`
	SecWeaponModel  uint64 `json:"secWeaponModel"`
	CraftToolModel  uint64 `json:"craftToolModel"`

	U14             int32  `json:"u14"`
	U15             uint32 `json:"u15"`
	BNPCBase        uint32 `json:"bNPCBase"`
	BNPCName        uint32 `json:"bNPCName"`
	U18             uint32 `json:"u18"`
	U19             uint32 `json:"u19"`
	DirectorID      uint32 `json:"directorId"`
	OwnerID         uint32 `json:"ownerId"`
	U22             uint32 `json:"u22"`
	HPMax           uint32 `json:"hPMax"`
	HPCurr          uint32 `json:"hPCurr"`
	DisplayFlags    uint32 `json:"displayFlags"`
	FateID          uint16 `json:"fateID"`
	MPCurr          uint16 `json:"mPCurr"`
	MPMax           uint16 `json:"mPMax"`
	Unk             uint16 `json:"unk"` // == 0
	ModelChara      uint16 `json:"modelChara"`
	Rotation        uint16 `json:"rotation"`
	ActiveMinion    uint16 `json:"activeMinion"`
	SpawnIndex      uint8  `json:"spawnIndex"`
	State           uint8  `json:"state"`
	PersistentEmote uint8  `json:"persistentEmote"`
	ModelType       uint8  `json:"modelType"`
	Subtype         uint8  `json:"subtype"`
	Voice           uint8  `json:"voice"`
	U25c            uint16 `json:"u25c"`
	EnemyType       uint8  `json:"enemyType"`
	Level           uint8  `json:"level"`
	ClassJob        uint8  `json:"classJob"`
	U26d            uint8  `json:"u26d"`
	U27a            uint16 `json:"u27a"`
	CurrentMount    uint8  `json:"currentMount"`
	MountHead       uint8  `json:"mountHead"`
	MountBody       uint8  `json:"mountBody"`
	MountFeet       uint8  `json:"mountFeet"`
	MountColor      uint8  `json:"mountColor"`
	Scale           uint8  `json:"scale"`

	//uint32 elementalLevel one of these two field changed to 16bit
	//uint32 element
	ElementData [6]uint8 `json:"elementData"`

	Effect [30]StatusEffect   `json:"effect"`
	Pos    FFXIVARR_POSITION3 `json:"pos"`
	Models [10]uint32         `json:"models"`
	Name   [32]byte           `json:"name"`
	Look   [26]uint8          `json:"look"`
	FcTag  [6]byte            `json:"fcTag"`
	Unk30  uint32             `json:"unk30"`
}

type NpcSpawn struct {
	GimmickID uint32 `json:"gimmickId"`
	U2b       uint8  `json:"u2b"`
	U2ab      uint8  `json:"u2ab"`
	GmRank    uint8  `json:"gmRank"`
	U3b       uint8  `json:"u3b"`

	AggressionMode uint8 `json:"aggressionMode"`
	OnlineStatus   uint8 `json:"onlineStatus"`
	U3c            uint8 `json:"u3c"`
	Pose           uint8 `json:"pose"`

	U4 uint8 `json:"u4"`

	TargetID        uint64 `json:"targetId"`
	U6              uint32 `json:"u6"`
	U7              uint32 `json:"u7"`
	MainWeaponModel uint64 `json:"mainWeaponModel"`
	SecWeaponModel  uint64 `json:"secWeaponModel"`
	CraftToolModel  uint64 `json:"craftToolModel"`

	U14             int32              `json:"u14"`
	U15             uint32             `json:"u15"`
	BNPCBase        uint32             `json:"bNPCBase"`
	BNPCName        uint32             `json:"bNPCName"`
	LevelID         uint32             `json:"levelId"`
	U19             uint32             `json:"u19"`
	DirectorID      uint32             `json:"directorId"`
	SpawnerID       uint32             `json:"spawnerId"`
	ParentActorID   uint32             `json:"parentActorId"`
	HPMax           uint32             `json:"hPMax"`
	HPCurr          uint32             `json:"hPCurr"`
	DisplayFlags    uint32             `json:"displayFlags"`
	FateID          uint16             `json:"fateID"`
	MPCurr          uint16             `json:"mPCurr"`
	Unknown1        uint16             `json:"unknown1"`
	Unknown2        uint16             `json:"unknown2"`
	ModelChara      uint16             `json:"modelChara"`
	Rotation        uint16             `json:"rotation"`
	ActiveMinion    uint16             `json:"activeMinion"`
	SpawnIndex      uint8              `json:"spawnIndex"`
	State           uint8              `json:"state"`
	PersistentEmote uint8              `json:"persistentEmote"`
	ModelType       uint8              `json:"modelType"`
	Subtype         uint8              `json:"subtype"`
	Voice           uint8              `json:"voice"`
	U25c            uint16             `json:"u25c"`
	EnemyType       uint8              `json:"enemyType"`
	Level           uint8              `json:"level"`
	ClassJob        uint8              `json:"classJob"`
	U26d            uint8              `json:"u26d"`
	U27a            uint16             `json:"u27a"`
	CurrentMount    uint8              `json:"currentMount"`
	MountHead       uint8              `json:"mountHead"`
	MountBody       uint8              `json:"mountBody"`
	MountFeet       uint8              `json:"mountFeet"`
	MountColor      uint8              `json:"mountColor"`
	Scale           uint8              `json:"scale"`
	ElementalLevel  uint16             `json:"elementalLevel"` // Eureka
	Element         uint16             `json:"element"`        // Eureka
	Effect          [30]StatusEffect   `json:"effect"`
	Pos             FFXIVARR_POSITION3 `json:"pos"`
	Models          [10]uint32         `json:"models"`
	Name            [32]byte           `json:"name"`
	Look            [26]uint8          `json:"look"`
	FcTag           [6]byte            `json:"fcTag"`
	Unk30           uint32             `json:"unk30"`
	Unk31           uint32             `json:"unk31"`
	BNPCPartSlot    uint8              `json:"bNPCPartSlot"`
	Unk32           uint8              `json:"unk32"`
	Unk33           uint16             `json:"unk33"`
	Unk34           uint32             `json:"unk34"`
}

// PlayerStats - Structural representation of the packet sent by the server
// to set a players stats
type PlayerStats struct {
	// order comes from baseparam order column
	Strength            uint32 `json:"strength"`
	Dexterity           uint32 `json:"dexterity"`
	Vitality            uint32 `json:"vitality"`
	Intelligence        uint32 `json:"intelligence"`
	Mind                uint32 `json:"mind"`
	Piety               uint32 `json:"piety"`
	Hp                  uint32 `json:"hp"`
	Mp                  uint32 `json:"mp"`
	Tp                  uint32 `json:"tp"`
	Gp                  uint32 `json:"gp"`
	Cp                  uint32 `json:"cp"`
	Delay               uint32 `json:"delay"`
	Tenacity            uint32 `json:"tenacity"`
	AttackPower         uint32 `json:"attackPower"`
	Defense             uint32 `json:"defense"`
	DirectHitRate       uint32 `json:"directHitRate"`
	Evasion             uint32 `json:"evasion"`
	MagicDefense        uint32 `json:"magicDefense"`
	CriticalHit         uint32 `json:"criticalHit"`
	AttackMagicPotency  uint32 `json:"attackMagicPotency"`
	HealingMagicPotency uint32 `json:"healingMagicPotency"`
	ElementalBonus      uint32 `json:"elementalBonus"`
	Determination       uint32 `json:"determination"`
	SkillSpeed          uint32 `json:"skillSpeed"`
	SpellSpeed          uint32 `json:"spellSpeed"`
	Haste               uint32 `json:"haste"`
	Craftsmanship       uint32 `json:"craftsmanship"`
	Control             uint32 `json:"control"`
	Gathering           uint32 `json:"gathering"`
	Perception          uint32 `json:"perception"`

	// todo: what is here?
	Unknown [26]uint32 `json:"unknown"`
}

// ItemInfo -
type ItemInfo struct {
	ContainerSequence uint32 `json:"containerSequence"`
	Unknown           uint32 `json:"unknown"`
	ContainerID       uint16 `json:"containerId"`
	Slot              uint16 `json:"slot"`
	Quantity          uint32 `json:"quantity"`
	CatalogID         uint32 `json:"catalogId"`
	ReservedFlag      uint32 `json:"reservedFlag"`
	SignatureID       uint64 `json:"signatureId"`
	HqFlag            uint8  `json:"hqFlag"`
	Unknown2          uint8  `json:"unknown2"`
	Condition         uint16 `json:"condition"`
	SpiritBond        uint16 `json:"spiritbond"`
	Stain             uint16 `json:"stain"`
	GlamourCatalogID  uint32 `json:"glamourCatalogId"`
	Materia1          uint16 `json:"materia1"`
	Materia2          uint16 `json:"materia2"`
	Materia3          uint16 `json:"materia3"`
	Materia4          uint16 `json:"materia4"`
	Materia5          uint16 `json:"materia5"`
	// Someone still needs to get around to PRing Sapphire with these
	Tier1     uint8  `json:"tier1"`
	Tier2     uint8  `json:"tier2"`
	Tier3     uint8  `json:"tier3"`
	Tier4     uint8  `json:"tier4"`
	Tier5     uint8  `json:"tier5"`
	Padding   uint8  `json:"padding"`
	Unknown10 uint32 `json:"unknown10"`
}

// UpdateInventorySlot - Structural representation of the packet sent by the server
// to update a slot in the inventory
type UpdateInventorySlot struct {
	Sequence         uint32 `json:"sequence"`
	Unknown          uint32 `json:"unknown"`
	ContainerID      uint16 `json:"containerId"`
	Slot             uint16 `json:"slot"`
	Quantity         uint32 `json:"quantity"`
	CatalogID        uint32 `json:"catalogId"`
	ReservedFlag     uint32 `json:"reservedFlag"`
	SignatureID      uint64 `json:"signatureId"`
	HqFlag           uint8  `json:"hqFlag"`
	Condition        uint16 `json:"condition"`
	SpiritBond       uint16 `json:"spiritbond"`
	Color            uint16 `json:"color"`
	GlamourCatalogID uint32 `json:"glamourCatalogId"`
	Materia1         uint16 `json:"materia1"`
	Materia2         uint16 `json:"materia2"`
	Materia3         uint16 `json:"materia3"`
	Materia4         uint16 `json:"materia4"`
	Materia5         uint16 `json:"materia5"`
	Tier1            uint8  `json:"tier1"`
	Tier2            uint8  `json:"tier2"`
	Tier3            uint8  `json:"tier3"`
	Tier4            uint8  `json:"tier4"`
	Tier5            uint8  `json:"tier5"`
	Padding          uint8  `json:"padding"`
	Unknown10        uint32 `json:"unknown10"`
}

type InventoryTransaction struct {
	Sequence        uint32    `json:"sequence"`
	Type            uint16    `json:"type"`
	Padding1        uint16    `json:"padding1"`
	OwnerID         uint32    `json:"ownerId"`
	StorageID       uint32    `json:"storageId"`
	SlotID          uint16    `json:"slotId"`
	Padding2        uint16    `json:"padding2"`
	StackSize       uint32    `json:"stackSize"`
	CatalogID       uint32    `json:"catalogId"`
	SomeActorID     uint32    `json:"someActorId"`
	TargetStorageID int32     `json:"targetStorageId"`
	Padding3        [3]uint32 `json:"padding3"`
}

type CurrencyCrystalInfo struct {
	ContainerSequence uint32 `json:"containerSequence"`
	ContainerID       uint16 `json:"containerId"`
	Slot              uint16 `json:"slot"`
	Quantity          uint32 `json:"quantity"`
	Unknown           uint32 `json:"unknown"`
	CatalogID         uint32 `json:"catalogId"`
	Unknown1          uint32 `json:"unknown1"`
	Unknown2          uint32 `json:"unknown2"`
	Unknown3          uint32 `json:"unknown3"`
}

// EventStart - Structural representation of the packet sent by the server
// to start an event, not actually playing it, but registering
type EventStart struct {
	/* 0000 */ ActorID uint64 `json:"actorId"`
	/* 0008 */ EventID uint32 `json:"eventId"`
	/* 000C */ Param1 uint8 `json:"param1"`
	/* 000D */ Param2 uint8 `json:"param2"`
	/* 000E */ Padding uint16 `json:"padding"`
	/* 0010 */ Param3 uint32 `json:"param3"`
	/* 0014 */ Padding1 uint32 `json:"padding1"`
}

// EventFinish - Structural representation of the packet sent by the server
// to finish an event
type EventFinish struct {
	/* 0000 */ EventID uint32 `json:"eventId"`
	/* 0004 */ Param1 uint8 `json:"param1"`
	/* 0005 */ Param2 uint8 `json:"param2"`
	/* 0006 */ Padding uint16 `json:"padding"`
	/* 0008 */ Param3 uint32 `json:"param3"`
	/* 000C */ Padding1 uint32 `json:"padding1"`
}

// EventPlay - Structural representation of the packet sent by the server
// to play an event
type EventPlay struct {
	ActorID  uint64   `json:"actorId"`
	EventID  uint32   `json:"eventId"`
	Scene    uint16   `json:"scene"`
	Padding  uint16   `json:"padding"`
	Flags    uint32   `json:"flags"`
	Param3   uint32   `json:"param3"`
	Param4   uint8    `json:"param4"`
	Padding1 [3]uint8 `json:"padding1"`
	Param5   uint32   `json:"param5"`
	Unknown  [8]uint8 `json:"unknown"`
}

// EventPlay4 - Structural representation of the packet sent by the server
// to play an event
type EventPlay4 struct {
	ActorID    uint64    `json:"actorId"`
	EventID    uint32    `json:"eventId"`
	Scene      uint16    `json:"scene"`
	Padding    uint16    `json:"padding"`
	SceneFlags uint32    `json:"sceneFlags"`
	ParamCount uint8     `json:"paramCount"`
	Padding2   [3]uint8  `json:"padding2"`
	Params     [4]uint32 `json:"params"`
}

type SomeDirectorUnk4 struct {
	EventID        uint32 `json:"eventId"`
	Param1         uint32 `json:"param1"`
	ActionTimeline uint32 `json:"actionTimeline"`
	Param3         uint32 `json:"param3"`
}

type MarketTaxRates struct {
	Unknown1 uint32             `json:"unknown1"`
	Padding1 uint16             `json:"padding1"`
	Unknown2 uint16             `json:"unknown2"`
	TaxRate  [TOWN_COUNT]uint32 // In the order of Common::Town
	Unknown3 uint64             `json:"unknown3"`
}

type MarketBoardItemListingCount struct {
	ItemCatalogID uint32 `json:"itemCatalogId"`
	Unknown1      uint32 `json:"unknown1"` // does some shit if nonzero
	RequestID     uint16 `json:"requestId"`
	Quantity      uint16 `json:"quantity"` // high/low u8s read separately?
	Unknown3      uint32 `json:"unknown3"`
}

type MarketBoardItemListing struct {
	Listing [10]struct // 152 bytes each
	{
		ListingID       uint64 `json:"listingId"`
		RetainerID      uint64 `json:"retainerId"`
		RetainerOwnerID uint64 `json:"retainerOwnerId"`
		ArtisanID       uint64 `json:"artisanId"`
		PricePerUnit    uint32 `json:"pricePerUnit"`
		TotalTax        uint32 `json:"totalTax"`
		ItemQuantity    uint32 `json:"itemQuantity"`
		ItemID          uint32 `json:"itemId"`
		LastReviewTime  uint16 `json:"lastReviewTime"`
		ContainerID     uint16 `json:"containerId"`
		SlotID          uint32 `json:"slotId"`
		Durability      uint16 `json:"durability"`
		SpiritBond      uint16 `json:"spiritBond"`
		/**
		 * auto materiaID = (i & 0xFF0) >> 4
		 * auto index = i & 0xF
		 * auto leftover = i >> 8
		 */
		MateriaValue [5]uint16 `json:"materiaValue"`
		Padding1     uint16    `json:"padding1"`
		Padding2     uint32    `json:"padding2"`
		RetainerName [32]byte  `json:"retainerName"`
		PlayerName   [32]byte  `json:"playerName"`
		Hq           bool      `json:"hq"`
		MateriaCount uint8     `json:"materiaCount"`
		OnMannequin  uint8     `json:"onMannequin"`
		MarketCity   uint8     `json:"marketCity"`
		DyeID        uint16    `json:"dyeID"`
		Padding3     uint16    `json:"padding3"`
		Padding4     uint32    `json:"padding4"`
	} `json:"listing"` // Multiple packets are sent if there are more than 10 search results.
	ListingIndexEnd   uint8    `json:"listingIndexEnd"`
	ListingIndexStart uint8    `json:"listingIndexStart"`
	RequestID         uint16   `json:"requestId"`
	Padding7          [16]byte `json:"padding7"`
	Unknown13         uint8    `json:"unknown13"`
	Padding8          uint16   `json:"padding8"`
	Unknown14         uint8    `json:"unknown14"`
	Padding9          uint64   `json:"padding9"`
	Unknown15         uint32   `json:"unknown15"`
	Padding10         uint32   `json:"padding10"`
}

type MarketBoardItemListingHistory struct {
	ItemCatalogID  uint32 `json:"itemCatalogId"`
	ItemCatalogID2 uint32 `json:"itemCatalogId2"`

	Listing [20]struct {
		SalePrice    uint32 `json:"salePrice"`
		PurchaseTime uint32 `json:"purchaseTime"`
		Quantity     uint32 `json:"quantity"`
		IsHq         uint8  `json:"isHq"`
		Padding      uint8  `json:"padding"`
		OnMannequin  uint8  `json:"onMannequin"`

		BuyerName [33]byte `json:"buyerName"`

		ItemCatalogID uint32 `json:"itemCatalogId"`
	}
}

type MarketBoardSearchResult struct {
	Items [20]struct {
		ItemCatalogID uint32 `json:"itemCatalogId"`
		Quantity      uint16 `json:"quantity"`
		Demand        uint16 `json:"demand"`
	}

	ListingIndexEnd   uint32 `json:"listingIndexEnd"`
	Padding1          uint32 `json:"padding1"`
	ListingIndexStart uint32 `json:"listingIndexStart"`
	RequestID         uint32 `json:"requestId"`
}

type WeatherChange struct {
	WeatherID uint32  `json:"weatherId"`
	Delay     float32 `json:"delay"`
}
