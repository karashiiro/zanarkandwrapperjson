package sapphire

// Bimap16 is a structure containing two maps, a by-key map and a by-value map
type Bimap16 struct {
	ByKeys   map[string]uint16
	ByValues map[uint16]string
}

// Bimap32 is a structure containing two maps, a by-key map and a by-value map
type Bimap32 struct {
	ByKeys   map[string]uint32
	ByValues map[uint32]string
}

func ReverseMap16(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}

func ReverseMap32(m map[string]uint32) map[uint32]string {
	flip := make(map[uint32]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}
