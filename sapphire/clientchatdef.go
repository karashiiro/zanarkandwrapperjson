package sapphire

type Tell struct {
	ContentId      uint64
	WorldId        uint16
	Flags          uint8
	ReceipientName [32]byte
	Msg            [1029]byte
}

type TellErrNotFound struct {
	ReceipientName [32]byte
}

type FreeCompanyEvent struct {
	Unknown  uint16
	Unknown1 uint16
	Unknown2 uint16
	Unknown3 uint16
	Unknown4 uint16
	Padding  [6]byte
	EventID  uint8
	/*
	 * 0X0F Login
	 * 0X10 Logout
	 */
	Padding1  uint8
	Padding2  [6]byte
	Unknown5  uint16
	Parameter [46]byte
	/**
	 * EventID  | parameter usage
	 * 0X0F       FC name
	 * 0X10       FC name
	 */
	Parameter1 [32]byte
	/**
	 * EventID  | parameter1 usage
	 * 0X0F       byteacter name
	 * 0X10       byteacter name
	 */
}
