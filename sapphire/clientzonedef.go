package sapphire

import "encoding/json"

type InventoryModifyHandler struct {
	Seq           uint32
	Action        uint8
	Pad0006       [6]uint8
	FromContainer uint16
	Pad000E       [2]uint8
	FromSlot      uint8
	Pad0011       [15]uint8
	ToContainer   uint16
	Pad0022       [2]uint8
	ToSlot        uint8
	Pad0025       [3]uint8
	SplitCount    uint32
}

// MarshalJSON override for InventoryModifyHandler
func (i *InventoryModifyHandler) MarshalJSON() ([]byte, error) {
	var actionType string
	action := uint32(i.Action)
	inventoryOperationID := DynamicConstants.ByKeys["InventoryOperationBaseValue"]
	if action == inventoryOperationID {
		actionType = "discard"
	} else if action == inventoryOperationID+1 {
		actionType = "move"
	} else if action == inventoryOperationID+2 {
		actionType = "swap"
	} else if action == inventoryOperationID+5 {
		actionType = "merge"
	} else if action == inventoryOperationID+10 {
		actionType = "split"
	}
	return json.Marshal(&struct {
		Seq           uint32    `json:"seq"`
		Action        string    `json:"action"`
		Pad0006       [6]uint8  `json:"pad_0006"`
		FromContainer uint16    `json:"fromContainer"`
		Pad000E       [2]uint8  `json:"pad_000E"`
		FromSlot      uint8     `json:"fromSlot"`
		Pad0011       [15]uint8 `json:"pad_0011"`
		ToContainer   uint16    `json:"toContainer"`
		Pad0022       [2]uint8  `json:"pad_0022"`
		ToSlot        uint8     `json:"toSlot"`
		Pad0025       [3]uint8  `json:"pad_0025"`
		SplitCount    uint32    `json:"splitCount"`
	}{
		Seq:           i.Seq,
		Action:        actionType,
		Pad0006:       i.Pad0006,
		FromContainer: i.FromContainer,
		Pad000E:       i.Pad000E,
		FromSlot:      i.FromSlot,
		Pad0011:       i.Pad0011,
		ToContainer:   i.ToContainer,
		Pad0022:       i.Pad0022,
		ToSlot:        i.ToSlot,
		Pad0025:       i.Pad0025,
		SplitCount:    i.SplitCount,
	})
}
