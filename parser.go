package zanarkandwrapperjson

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ayyaruq/zanarkand"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

var actorControl uint16 = sapphire.ServerLobbyIpcType["ActorControl"]
var actorControlSelf uint16 = sapphire.ServerLobbyIpcType["ActorControlSelf"]
var actorControlTarget uint16 = sapphire.ServerLobbyIpcType["ActorControlTarget"]
var clientTrigger uint16 = sapphire.ServerLobbyIpcType["ClientTrigger"]

func parseMessage(message *zanarkand.GameEventMessage, region string, port uint16) {
	ipcType := getPacketType(message.Opcode, region)

	ipcType = strings.ToLower(ipcType[0:1]) + ipcType[1:]

	var actorControlCategory string
	var clientTriggerCategory string
	if message.Opcode == actorControl || message.Opcode == actorControlSelf || message.Opcode == actorControlTarget {
		actorControlCategory = ActorControlType[binary.LittleEndian.Uint16(message.Body[IpcData:IpcData+2])]
	} else if message.Opcode == clientTrigger {
		clientTriggerCategory = ClientTriggerType[binary.LittleEndian.Uint16(message.Body[IpcData:IpcData+2])]
	}

	serializePacket(message, ipcType, actorControlCategory, clientTriggerCategory, region, port)
}

func serializePacket(message *zanarkand.GameEventMessage, ipcType string, actorControlCategory string, clientTriggerCategory string, region string, port uint16) {
	var outputBase OutputBase
	var ipcBase IpcBase
	var ipcActorClientControl IpcActorClientControl

	outputBase.Type = ipcType
	outputBase.Opcode = message.Opcode
	outputBase.Region = region
	outputBase.PacketSize = message.Length
	outputBase.SegmentType = message.Segment

	if message.Segment == 3 {
		ipcBase.OutputBase = outputBase
		ipcBase.SourceActor = message.SourceActor
		ipcBase.TargetActor = message.TargetActor
		ipcBase.ServerID = message.ServerID
		ipcBase.Timestamp = message.Timestamp

		// To cut down on data transfer a bit, we trim this. The useful data before this is parsed by now anyways.
		message.Body = message.Body[IpcData:]

		if actorControlCategory != "" {
			ipcActorClientControl.IpcBase = ipcBase
			ipcActorClientControl.SuperType = "actorControl"
			ipcActorClientControl.SubType = actorControlCategory
		} else if clientTriggerCategory != "" {
			ipcActorClientControl.IpcBase = ipcBase
			ipcActorClientControl.SuperType = "clientTrigger"
			ipcActorClientControl.SubType = clientTriggerCategory
		}
	}

	outputBase.Body = message.Body

	var buf bytes.Buffer
	var bytes []byte
	if ipcActorClientControl.SubType != "" {
		bytes, _ = json.Marshal(ipcActorClientControl)
	} else if ipcBase.SourceActor != 0 {
		bytes, _ = json.Marshal(ipcBase)
	} else {
		bytes, _ = json.Marshal(outputBase)
	}
	buf.Write(bytes)
	_, err := http.Post("http://localhost:"+fmt.Sprint(port), "application/json", &buf)
	if err != nil {
		log.Println(err)
	}
}

func getPacketType(opcode uint16, region string) string {
	var ipcType string
	var ok bool
	if region == "Global" {
		ipcType, ok = ServerLobbyIpcType[opcode]
		if !ok {
			ipcType, ok = ClientLobbyIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ServerZoneIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ClientZoneIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ServerChatIpcType[opcode]
		}
		if !ok {
			ipcType, ok = ClientChatIpcType[opcode]
		}
	} else if region == "CN" {
		// NYI
		ipcType, ok = ServerLobbyIpcTypeCN[opcode]
		if !ok {
			ipcType, ok = ClientLobbyIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ServerZoneIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ClientZoneIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ServerChatIpcTypeCN[opcode]
		}
		if !ok {
			ipcType, ok = ClientChatIpcTypeCN[opcode]
		}
	} else if region == "KR" {
		// NYI
		ipcType, ok = ServerLobbyIpcTypeKR[opcode]
		if !ok {
			ipcType, ok = ClientLobbyIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ServerZoneIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ClientZoneIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ServerChatIpcTypeKR[opcode]
		}
		if !ok {
			ipcType, ok = ClientChatIpcTypeKR[opcode]
		}
	}

	if !ok {
		ipcType = "unknown"
	}

	return ipcType
}
