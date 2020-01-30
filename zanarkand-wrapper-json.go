package zanarkandwrapperjson

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"

	"github.com/ayyaruq/zanarkand"
	devices "github.com/ayyaruq/zanarkand/devices"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

// ActorControlType - Value-first version of sapphire.ActorControlType
var ActorControlType = reverseMap(sapphire.ActorControlType)

// ClientTriggerType - Value-first version of sapphire.ClientTriggerType
var ClientTriggerType = reverseMap(sapphire.ClientTriggerType)

// ServerLobbyIpcType - Value-first version of sapphire.ServerLobbyIpcType
var ServerLobbyIpcType = reverseMap(sapphire.ServerLobbyIpcType)

// ClientLobbyIpcType - Value-first version of sapphire.ClientLobbyIpcType
var ClientLobbyIpcType = reverseMap(sapphire.ClientLobbyIpcType)

// ServerZoneIpcType - Value-first version of sapphire.ServerZoneIpcType
var ServerZoneIpcType = reverseMap(sapphire.ServerZoneIpcType)

// ClientZoneIpcType - Value-first version of sapphire.ClientZoneIpcType
var ClientZoneIpcType = reverseMap(sapphire.ClientZoneIpcType)

// ServerChatIpcType - Value-first version of sapphire.ServerChatIpcType
var ServerChatIpcType = reverseMap(sapphire.ServerChatIpcType)

// ClientChatIpcType - Value-first version of sapphire.ClientChatIpcType
var ClientChatIpcType = reverseMap(sapphire.ClientChatIpcType)

// ServerLobbyIpcTypeCN - Value-first version of sapphire.ServerLobbyIpcTypeCN
var ServerLobbyIpcTypeCN = reverseMap(sapphire.ServerLobbyIpcTypeCN)

// ClientLobbyIpcTypeCN - Value-first version of sapphire.ClientLobbyIpcTypeCN
var ClientLobbyIpcTypeCN = reverseMap(sapphire.ClientLobbyIpcTypeCN)

// ServerZoneIpcTypeCN - Value-first version of sapphire.ServerZoneIpcTypeCN
var ServerZoneIpcTypeCN = reverseMap(sapphire.ServerZoneIpcTypeCN)

// ClientZoneIpcTypeCN - Value-first version of sapphire.ClientZoneIpcTypeCN
var ClientZoneIpcTypeCN = reverseMap(sapphire.ClientZoneIpcTypeCN)

// ServerChatIpcTypeCN - Value-first version of sapphire.ServerChatIpcTypeCN
var ServerChatIpcTypeCN = reverseMap(sapphire.ServerChatIpcTypeCN)

// ClientChatIpcTypeCN - Value-first version of sapphire.ClientChatIpcTypeCN
var ClientChatIpcTypeCN = reverseMap(sapphire.ClientChatIpcTypeCN)

// ServerLobbyIpcTypeKR - Value-first version of sapphire.ServerLobbyIpcTypeKR
var ServerLobbyIpcTypeKR = reverseMap(sapphire.ServerLobbyIpcTypeKR)

// ClientLobbyIpcTypeKR - Value-first version of sapphire.ClientLobbyIpcTypeKR
var ClientLobbyIpcTypeKR = reverseMap(sapphire.ClientLobbyIpcTypeKR)

// ServerZoneIpcType - Value-first version of sapphire.ServerZoneIpcTypeKR
var ServerZoneIpcTypeKR = reverseMap(sapphire.ServerZoneIpcTypeKR)

// ClientZoneIpcTypeKR - Value-first version of sapphire.ClientZoneIpcTypeKR
var ClientZoneIpcTypeKR = reverseMap(sapphire.ClientZoneIpcTypeKR)

// ServerChatIpcTypeKR - Value-first version of sapphire.ServerChatIpcTypeKR
var ServerChatIpcTypeKR = reverseMap(sapphire.ServerChatIpcTypeKR)

// ClientChatIpcTypeKR - Value-first version of sapphire.ClientChatIpcTypeKR
var ClientChatIpcTypeKR = reverseMap(sapphire.ClientChatIpcTypeKR)

func main() {
	os.Exit(goLikeMain())
}

func goLikeMain() int {
	// MonitorType doesn't apply, nor do ProcessID and IPIndex
	region := flag.String("-Region", "Global", "Sets the IPC version to Global/CN/KR.")
	port := flag.Uint64("-Port", 13346, "Sets the port for the IPC connection between this application and Node.js.")

	// Setup our control mechanism
	commander := make(chan string)
	go func(input chan string) {
		stdin := bufio.NewReader(os.Stdin)
		for {
			cmd, err := stdin.ReadString('\n')
			if err != nil {
				log.Printf("Error reading command code: %v", err)
				close(input)
			}
			input <- cmd
		}
	}(commander)

	// Get the default network device (probably)
	netIfaces, err := devices.ListDeviceNames(false, false)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	// Initialize a sniffer on the default network device
	sniffer, err := zanarkand.NewSniffer("", netIfaces[0])
	if err != nil {
		log.Fatal(err)
		return 1
	}

	// Cleanly shutdown when we feel like it
	defer func(sniffer *zanarkand.Sniffer) {
		if sniffer.Active() {
			sniffer.Stop()
			time.Sleep(500) // time to drain
		}
	}(sniffer)

	subscriber := zanarkand.NewGameEventSubscriber()

	// Control loop
	for {
		select {
		case command := <-commander:
			switch command {
			case "kill":
				return 0
			case "start":
				if !sniffer.Active() {
					go subscriber.Subscribe(sniffer)
				}
			case "stop":
				if sniffer.Active() {
					sniffer.Stop()
				}
			default:
				log.Println("Unknown command recieved: ", command)
			}
		case message := <-subscriber.Events:
			go parseMessage(message, *region, uint16(*port))
		}
	}
}

func reverseMap(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}
