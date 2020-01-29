package zanarkandwrapperjson

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/ayyaruq/zanarkand"
	devices "github.com/ayyaruq/zanarkand/devices"

	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

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

func main() {
	// Map out args
	args := sort.StringSlice(os.Args[1:])
	// MonitorType doesn't apply, nor do ProcessID and IPIndex
	regionIndex := args.Search("--Region")
	var region string
	if regionIndex != -1 {
		region = args[regionIndex+1]
	}
	portIndex := args.Search("--Port")
	var port uint16
	if portIndex != -1 {
		tmpPort, err0 := strconv.ParseUint(args[portIndex+1], 10, 2)
		port = uint16(tmpPort)
		if err0 != nil {
			log.Fatal(err0)
		}
	}

	// Get the default network device (probably)
	netIfaces, err0 := devices.ListDeviceNames(false, false)
	if err0 != nil {
		log.Fatal(err0)
	}

	// Initialize a sniffer on the default network device
	sniffer, err1 := zanarkand.NewSniffer("", netIfaces[0])
	if err1 != nil {
		log.Fatal(err1)
	}
	subscriber := zanarkand.NewGameEventSubscriber()

	// Run loops
	var input string
	var lastErr error
	go inputLoop(&input, &lastErr)
	go inputProcssingLoop(&input, &lastErr, sniffer, subscriber)

	parseLoop(subscriber, region, port)
}

// Take input instantly
func inputLoop(input *string, lastErr *error) {
	stdin := bufio.NewReader(os.Stdin)

	for {
		*input, *lastErr = stdin.ReadString('\n')
	}
}

// Process inputs every 200ms
func inputProcssingLoop(input *string, lastErr *error, sniffer *zanarkand.Sniffer, subscriber *zanarkand.GameEventSubscriber) {
	for *input != "kill" {
		if *lastErr != nil {
			log.Println(*lastErr)
			*lastErr = nil
		}

		if *input == "start" && !sniffer.Active() {
			go subscriber.Subscribe(sniffer)
		} else if *input == "stop" && sniffer.Active() {
			sniffer.Stop()
		}

		// More than enough time for an input loop users rarely if ever interact with
		time.Sleep(200)
	}

	if sniffer.Active() {
		sniffer.Stop()
	}

	os.Exit(0)
}

func parseLoop(subscriber *zanarkand.GameEventSubscriber, region string, port uint16) {
	for {
		message := <-subscriber.Events

		// Push the message to the parsing sequence
		parseMessage(message, region, port)
	}
}

func reverseMap(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}
