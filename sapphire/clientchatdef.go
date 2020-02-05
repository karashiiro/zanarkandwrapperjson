package sapphire

type Tell struct {
	contentId      uint64
	worldId        uint16
	flags          uint8
	receipientName [32]byte
	msg            [1029]byte
}

type TellErrNotFound struct {
	receipientName [32]byte
}

type FreeCompanyEvent struct {
	unknown  uint16
	unknown1 uint16
	unknown2 uint16
	unknown3 uint16
	unknown4 uint16
	padding  [6]byte
	eventID  uint8
	/*
	 * 0x0F Login
	 * 0x10 Logout
	 */
	padding1  uint8
	padding2  [6]byte
	unknown5  uint16
	parameter [46]byte
	/**
	 * eventID  | parameter usage
	 * 0x0F       FC name
	 * 0x10       FC name
	 */
	parameter1 [32]byte
	/**
	 * eventID  | parameter1 usage
	 * 0x0F       byteacter name
	 * 0x10       byteacter name
	 */
}
