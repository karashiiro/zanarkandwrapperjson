package sapphire

type RetainerList struct {
	Padding [0x210]uint8
}

type ServiceIdInfo struct {
	Seq                uint64
	Padding            uint8
	NumServiceAccounts uint8
	U1                 uint8
	U2                 uint8
	Padding1           uint32

	ServiceAccount [8]struct {
		Id      uint32
		Unknown uint32
		Index   uint32
		Name    [0x44]byte
	}
}

type ServerList struct {
	Seq        uint64
	Final      uint16
	Offset     uint16
	NumServers uint32
	Padding    uint32
	Padding1   uint32

	Server [6]struct {
		Id       uint16
		Index    uint16
		Flags    uint32 // 0x02 = World not accepting new characters
		Padding1 uint32
		Icon     uint32 // 2 = bonus XP star
		Padding2 uint32
		Name     [0x40]byte
	}
}

type CharList struct {
	Seq               uint64
	Counter           uint8 // current packet count * 4, count * 4 +1 on last packet.
	NumInPacket       uint8 // always 2??
	Padding           uint16
	Unknown1          uint8
	Unknown2          uint8
	Unknown3          uint8
	Unknown4          uint8 // 0x80 in case of last packet
	Unknown5          [7]uint32
	Unknown6          uint8 // 0x80 in case of last packet
	VeteranRank       uint8
	Unknown7          uint8
	Padding1          uint8
	DaysSubscribed    uint32
	RemainingDays     uint32
	DaysToNextRank    uint32
	MaxCharOnWorld    uint16
	Unknown8          uint16
	EntitledExpansion uint32
	Padding2          uint32
	Padding3          uint32
	Padding4          uint32

	CharaDetails [2]struct {
		UniqueId       uint32
		Padding        uint32
		ContentId      uint64
		Index          uint32
		Padding2       uint32
		ServerId       uint16
		ServerId1      uint16
		Unknown        [9]uint8
		NameChara      [32]byte
		NameServer     [32]byte
		NameServer1    [32]byte
		CharDetailJson [1051]byte
	}
}

type EnterWorld struct {
	Seq       uint64
	CharId    uint32
	Padding   uint32
	ContentId uint64
	Padding2  uint32
	Sid       [66]byte
	Port      uint16
	Host      [48]byte
	Padding3  uint64
	Padding4  uint64
}

type CharCreate struct {
	Seq        uint64
	Unknown    uint8
	Unknown_2  uint8
	CharType   uint8 // Note: Changed from "type" to "charType" since type's a keyword
	Padding    uint8
	Unknown_3  uint32
	Unknown_4  uint32
	Unknown_5  uint32
	Unknown_6  uint64
	Unknown_61 uint64
	Unknown_62 uint64
	Unknown_63 uint64
	Content_id uint64
	Unknown_7  uint16
	Unknown_8  uint16
	Unknown_9  uint32
	Unknown_10 uint16
	Unknown_11 [11]uint8
	Name       [32]byte
	World      [32]byte
	World2     [32]byte
	Unknown_12 [0x953]uint8
}

type LobbyError struct {
	Seq        uint64
	Error_id   uint32
	Param      uint32
	Message_id uint16
	Message    [516]byte
}
