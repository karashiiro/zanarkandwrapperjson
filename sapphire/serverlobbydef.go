package sapphire

type LobbyRetainerList struct {
	Padding [0x210]uint8 `json:"padding"`
}

type LobbyServiceAccountList struct {
	Seq                uint64 `json:"seq"`
	Padding            uint8  `json:"padding"`
	NumServiceAccounts uint8  `json:"numServiceAccounts"`
	U1                 uint8  `json:"u1"`
	U2                 uint8  `json:"u2"`
	Padding1           uint32 `json:"padding1"`
	ServiceAccount     [8]struct {
		ID      uint32     `json:"id"`
		Unknown uint32     `json:"unknown"`
		Index   uint32     `json:"index"`
		Name    [0x44]byte `json:"name"`
	} `json:"serviceAccount"`
}

type LobbyServerList struct {
	Seq        uint64 `json:"seq"`
	Final      uint16 `json:"final"`
	Offset     uint16 `json:"offset"`
	NumServers uint32 `json:"numServers"`
	Padding    uint32 `json:"padding"`
	Padding1   uint32 `json:"padding1"`
	Server     struct {
		ID       uint16     `json:"id"`
		Index    uint16     `json:"index"`
		Flags    uint32     `json:"flags"`
		Padding1 uint32     `json:"padding1"`
		Icon     uint32     `json:"icon"`
		Padding2 uint32     `json:"padding2"`
		Name     [0x40]byte `json:"name"`
	} `json:"server"`
}
