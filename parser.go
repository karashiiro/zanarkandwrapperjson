package zanarkandwrapperjson

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ayyaruq/zanarkand"
)

func parseMessage(message *zanarkand.GameEventMessage, region string, port uint16) {
	var ipcType string
	var ok bool

	ipcType, ok = ServerLobbyIpcType[message.Opcode]
	if !ok {
		ipcType, ok = ClientLobbyIpcType[message.Opcode]
	}
	if !ok {
		ipcType, ok = ServerZoneIpcType[message.Opcode]
	}
	if !ok {
		ipcType, ok = ClientZoneIpcType[message.Opcode]
	}
	if !ok {
		ipcType, ok = ServerChatIpcType[message.Opcode]
	}
	if !ok {
		ipcType, ok = ClientChatIpcType[message.Opcode]
	}
	if !ok {
		ipcType = "unknown"
	}

	ipcType = strings.ToLower(ipcType[0:1]) + ipcType[1:]

	var actorControlCategory string
	var clientTriggerCategory string
	if ipcType[0:12] == "actorControl" {
		// ActorControlCategory
	} else if ipcType[0:13] == "clientTrigger" {
		// ClientTriggerCategory
	}

	serializePacket(message, ipcType, actorControlCategory, clientTriggerCategory, region, port)
}

func serializePacket(message *zanarkand.GameEventMessage, ipcType string, actorControlCategory string, clientTriggerCategory string, region string, port uint16) {
	// Use strings.Builder instead
	json := "{"
	json += "\"type\":\"" + ipcType + "\","
	json += "\"opcode\":\"" + fmt.Sprint(message.Opcode) + "\","
	json += "\"region\":\"" + region + "\","
	json += "\"packetSize\":\"" + fmt.Sprint(message.Length) + "\","
	json += "\"segmentType\":\"" + fmt.Sprint(message.Segment) + "\","
	if message.Segment == 3 {
		json += "\"sourceActorSessionID\":\"" + fmt.Sprint(message.SourceActor) + "\","
		json += "\"targetActorSessionID\":\"" + fmt.Sprint(message.TargetActor) + "\","
		json += "\"serverID\":\"" + fmt.Sprint(message.ServerID) + "\","
		json += "\"timestamp\":\"" + fmt.Sprint(message.Timestamp) + "\","

		// To cut down on data transfer a bit, we trim this. The useful data before this is parsed by now anyways.
		message.Body = message.Body[0x20:]

		if actorControlCategory != "" {
			json += "\"superType\":\"actorControl\","
			json += "\"subType\":\"" + actorControlCategory + "\","
		} else if clientTriggerCategory != "" {
			json += "\"superType\":\"clientTrigger\","
			json += "\"subType\":\"" + clientTriggerCategory + "\","
		}
	}
	json += "\"data\":["
	for i := 0; i < len(message.Body)-1; i++ {
		json += fmt.Sprint(message.Body[i]) + ","
	}
	json += fmt.Sprint(message.Body[len(message.Body)])
	json += "]}"

	var buf bytes.Buffer
	buf.WriteString(json)
	_, err := http.Post("http://localhost:"+fmt.Sprint(port), "application/json", &buf)
	if err != nil {
		log.Println(err)
	}
}
