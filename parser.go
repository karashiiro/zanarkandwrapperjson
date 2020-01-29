package zanarkandwrapperjson

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"strings"
	"net/http"

	"github.com/ayyaruq/zanarkand"
)

func parsePacket(frame *zanarkand.Frame, region string, port uint16) {
	ipcPacket := new(IpcPacket)
	ipcPacket.Metadata = frame
	ipcPacket.PacketSize = binary.LittleEndian.Uint32(frame.Body[PacketSize : PacketSize+4])
	ipcPacket.SourceActor = binary.LittleEndian.Uint32(frame.Body[SourceActor : SourceActor+4])
	ipcPacket.TargetActor = binary.LittleEndian.Uint32(frame.Body[TargetActor : TargetActor+4])
	ipcPacket.SegmentType = binary.LittleEndian.Uint16(frame.Body[SegmentType : SegmentType+2])
	ipcPacket.Opcode = binary.LittleEndian.Uint16(frame.Body[IpcType : IpcType+2])
	ipcPacket.ServerID = binary.LittleEndian.Uint16(frame.Body[ServerID : ServerID+2])
	ipcPacket.Timestamp = binary.LittleEndian.Uint32(frame.Body[Timestamp : Timestamp+2])

	var type string
	var ok bool

	switch frame.Connection {
	case 0:
		type, ok = ServerLobbyIpcType[ipcPacket.Opcode]
		if (!ok) {
			type, ok = ClientLobbyIpcType[ipcPacket.Opcode]
		}
		if (!ok) {
			type = "unknown"
		}
		break
	case 1:
		type, ok = ServerZoneIpcType[ipcPacket.Opcode]
		if (!ok) {
			type, ok = ClientZoneIpcType[ipcPacket.Opcode]
		}
		if (!ok) {
			type = "unknown"
		}
		break
	case 2:
		type, ok = ServerChatIpcType[ipcPacket.Opcode]
		if (!ok) {
			type, ok = ClientChatIpcType[ipcPacket.Opcode]
		}
		if (!ok) {
			type = "unknown"
		}
		break
	}

	type = strings.ToLower(type[0]) + type[1:]

	if type[0:12] == "actorControl" {
		// ActorControlCategory
	} else if type[0:13] == "clientTrigger" {
		// ClientTriggerCategory
	}

	serializePacket(ipcPacket, region, port)
}

func serializePacket(packet *IpcPacket, region string, port uint16) {
	// Use strings.Builder instead
	json := "{"
	json += "\"type\":\"" + packet.Type + "\","
	json += "\"opcode\":\"" + fmt.Sprint(packet.Opcode) + "\","
	json += "\"region\":\"" + region + "\","
	json += "\"connection\":null,"                                                   // Initialized field, but undefined and unnecessary
	json += "\"connectionType\":\"" + fmt.Sprint(packet.Metadata.Connection) + "\"," // New and important
	json += "\"epoch\":\"" + fmt.Sprint(packet.Metadata.Timestamp.Unix()) + "\","
	json += "\"packetSize\":\"" + fmt.Sprint(packet.PacketSize) + "\","
	json += "\"segmentType\":\"" + fmt.Sprint(packet.SegmentType) + "\","
	if packet.SegmentType == 3 {
		json += "\"sourceActorSessionID\":\"" + fmt.Sprint(packet.SourceActor) + "\","
		json += "\"targetActorSessionID\":\"" + fmt.Sprint(packet.TargetActor) + "\","
		json += "\"serverID\":\"" + fmt.Sprint(packet.ServerID) + "\","
		json += "\"timestamp\":\"" + fmt.Sprint(packet.Timestamp) + "\","

		// To cut down on data transfer a bit, we trim this. The useful data before this is parsed by now anyways.
		packet.Metadata.Body = packet.Metadata.Body[IpcData:]

		if packet.ActorControlCategory != "" {
			json += "\"superType\":\"actorControl\","
			json += "\"subType\":\"" + packet.ActorControlCategory + "\","
		} else if packet.ClientTriggerCategory != "" {
			json += "\"superType\":\"clientTrigger\","
			json += "\"subType\":\"" + packet.ClientTriggerCategory + "\","
		}
	}
	json += "\"data\":["
	for i := 0; i < len(packet.Metadata.Body)-1; i++ {
		json += fmt.Sprint(packet.Metadata.Body[i]) + ","
	}
	json += fmt.Sprint(packet.Metadata.Body[len(packet.Metadata.Body)])
	json += "]}"

	var buf bytes.Buffer
	buf.WriteString(json)
	_, err := http.Post("http://localhost:"+fmt.Sprint(port), "application/json", &buf)
	if err != nil {
		log.Println(err)
	}
}
