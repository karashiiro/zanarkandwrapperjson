package zanarkandwrapperjson

import (
	"bufio"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/ayyaruq/zanarkand"
	"github.com/ayyaruq/zanarkand/devices"
	"sapphire"
)

// I'm not reversing these manually.
var ServerLobbyIpcType := reverseMap(sapphire.ServerLobbyIpcType)
var ClientLobbyIpcType := reverseMap(sapphire.ClientLobbyIpcType)
var ServerZoneIpcType := reverseMap(sapphire.ServerZoneIpcType)
var ClientZoneIpcType := reverseMap(sapphire.ClientZoneIpcType)
var ServerChatIpcType := reverseMap(sapphire.ServerChatIpcType)
var ClientChatIpcType := reverseMap(sapphire.ClientChatIpcType)

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
	defaultNetIface, err0 := net.InterfaceByIndex(0)
	if err0 != nil {
		log.Fatal(err0)
	}

	// Initialize a sniffer on the default network device
	sniffer, err1 := zanarkand.NewSniffer("", defaultNetIface.Name)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Run goroutines
	var input string
	var lastErr error
	go inputLoop(&input, &lastErr)
	go inputProcssingLoop(&input, &lastErr, sniffer)
	go parseLoop(sniffer, region, port)
}

// Take input instantly
func inputLoop(input *string, lastErr *error) {
	stdin := bufio.NewReader(os.Stdin)

	for {
		*input, *lastErr = stdin.ReadString('\n') // In .NET this blocks the thread on its own, I imagine it's the same here
	}
}

// Process inputs every 200ms
func inputProcssingLoop(input *string, lastErr *error, sniffer *zanarkand.Sniffer) {
	for *input != "kill" {
		if *lastErr != nil {
			log.Println(*lastErr)
			*lastErr = nil
		}

		if *input == "start" && !sniffer.Active() {
			go sniffer.Start()
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

func parseLoop(sniffer *zanarkand.Sniffer, region string, port uint16) {
	for {
		frame, err := sniffer.NextFrame()
		if err != nil {
			log.Println(err)
		}

		// Push the frame to the parsing sequence
		parsePacket(frame, region, port)
	}
}

func reverseMap(m map[string]uint16) map[uint16]string {
	flip := make(map[uint16]string)
	for k, v := range m {
		flip[v] = k
	}
	return flip
}