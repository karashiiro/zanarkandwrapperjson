package sapphire

type Ping struct {
	/* 0000 */ timeInMilliseconds uint64
	/* 0008 */ unknown_8 [0x38]uint8
}

type Init struct {
	unknown  uint64
	charId   uint32
	unknown1 uint32
}
