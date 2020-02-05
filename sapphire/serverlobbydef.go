package sapphire

type RetainerList struct {
	padding [0x210]uint8
}

type ServiceIdInfo struct {
	seq                uint64
	padding            uint8
	numServiceAccounts uint8
	u1                 uint8
	u2                 uint8
	padding1           uint32

	serviceAccount [8]struct {
		id      uint32
		unknown uint32
		index   uint32
		name    [0x44]byte
	}
}

type ServerList struct {
	seq        uint64
	final      uint16
	offset     uint16
	numServers uint32
	padding    uint32
	padding1   uint32

	server [6]struct {
		id       uint16
		index    uint16
		flags    uint32 // 0x02 = World not accepting new characters
		padding1 uint32
		icon     uint32 // 2 = bonus XP star
		padding2 uint32
		name     [0x40]byte
	}
}

type CharList struct {
	seq               uint64
	counter           uint8 // current packet count * 4, count * 4 +1 on last packet.
	numInPacket       uint8 // always 2??
	padding           uint16
	unknown1          uint8
	unknown2          uint8
	unknown3          uint8
	unknown4          uint8 // 0x80 in case of last packet
	unknown5          [7]uint32
	unknown6          uint8 // 0x80 in case of last packet
	veteranRank       uint8
	unknown7          uint8
	padding1          uint8
	daysSubscribed    uint32
	remainingDays     uint32
	daysToNextRank    uint32
	maxCharOnWorld    uint16
	unknown8          uint16
	entitledExpansion uint32
	padding2          uint32
	padding3          uint32
	padding4          uint32

	charaDetails [2]struct {
		uniqueId       uint32
		padding        uint32
		contentId      uint64
		index          uint32
		padding2       uint32
		serverId       uint16
		serverId1      uint16
		unknown        [9]uint8
		nameChara      [32]byte
		nameServer     [32]byte
		nameServer1    [32]byte
		charDetailJson [1051]byte
	}
}

type EnterWorld struct {
	seq       uint64
	charId    uint32
	padding   uint32
	contentId uint64
	padding2  uint32
	sid       [66]byte
	port      uint16
	host      [48]byte
	padding3  uint64
	padding4  uint64
}

type CharCreate struct {
	seq        uint64
	unknown    uint8
	unknown_2  uint8
	charType   uint8 // Note: Changed from "type" to "charType" since type's a keyword
	padding    uint8
	unknown_3  uint32
	unknown_4  uint32
	unknown_5  uint32
	unknown_6  uint64
	unknown_61 uint64
	unknown_62 uint64
	unknown_63 uint64
	content_id uint64
	unknown_7  uint16
	unknown_8  uint16
	unknown_9  uint32
	unknown_10 uint16
	unknown_11 [11]uint8
	name       [32]byte
	world      [32]byte
	world2     [32]byte
	unknown_12 [0x953]uint8
}

type LobbyError struct {
	seq        uint64
	error_id   uint32
	param      uint32
	message_id uint16
	message    [516]byte
}
