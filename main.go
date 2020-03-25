package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ayyaruq/zanarkand"
	"github.com/ayyaruq/zanarkand/devices"
	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

func main() {
	os.Exit(goLikeMain())
}

func goLikeMain() int {
	// MonitorType doesn't apply, nor do ProcessID or ParseAlgorithm
	region := flag.String("-Region", "Global", "Sets the IPC version to Global/CN/KR.")
	port := flag.Uint64("-Port", 13346, "Sets the port for the IPC connection between this application and Node.js.")
	networkDevice := net.ParseIP(*flag.String("-LocalIP", "", "Specifies a network to capture traffic on."))
	isDev := flag.Bool("-Dev", false, "Sets the developer mode, enabling raw data output.")
	flag.Parse()

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

	// Get the default network device
	// Looking for ranges:
	// 24-bit private block: 10.*.*.*
	// 20-bit private block: 172.16.*.*-172.31.*.*
	// 16-bit private block: 192.168.*.*
	// 16-bit APIPA block: 169.254.*.*
	_, subnet24, _ := net.ParseCIDR("10.0.0.0/8")
	_, subnet20, _ := net.ParseCIDR("172.16.0.0/12")
	_, subnet16, _ := net.ParseCIDR("192.168.0.0/16")
	_, APIPA16, _ := net.ParseCIDR("169.254.0.0/16")

	rNetIfaces, err := devices.ListDeviceNames(false, true)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	netIfaces, _ := devices.ListDeviceNames(false, false) // It seems safe to assume that if the call that includes this and more doesn't error, this will not error if called immediately after either
	netIfaceIdx := len(netIfaces) - 1
	for i, nif := range rNetIfaces {
		ip := net.ParseIP(regexp.MustCompile("\\d+\\.\\d+\\.\\d+\\.\\d+").FindString(nif)) // Lazy but it works and it only runs once
		if networkDevice != nil && networkDevice.Equal(ip) {
			netIfaceIdx = i
			break
		}
		if subnet24.Contains(ip) || subnet20.Contains(ip) || subnet16.Contains(ip) || APIPA16.Contains(ip) {
			netIfaceIdx = i
		}
	}

	// Initialize a sniffer on the network device
	sniffer, err := zanarkand.NewSniffer("pcap", netIfaces[netIfaceIdx])
	if err != nil {
		log.Fatal(err)
		return 1
	}

	// Cleanly shutdown when we feel like it
	defer func(sniffer *zanarkand.Sniffer) {
		if sniffer.Active {
			sniffer.Stop()
			time.Sleep(500) // time to drain
		}
	}(sniffer)

	subscriber := zanarkand.NewGameEventSubscriber()

	// Start thread pools
	port16 := uint16(*port)
	threadCount := 4 // goroutine count per pool
	go spawnThreads(&ServerZonePool, threadCount, region, &port16, false, isDev)
	go spawnThreads(&ServerLobbyPool, threadCount, region, &port16, false, isDev)
	go spawnThreads(&ServerChatPool, threadCount, region, &port16, false, isDev)
	go spawnThreads(&ServerUnknownPool, threadCount, region, &port16, false, isDev)
	go spawnThreads(&ClientZonePool, threadCount, region, &port16, true, isDev)
	go spawnThreads(&ClientLobbyPool, threadCount, region, &port16, true, isDev)
	go spawnThreads(&ClientChatPool, threadCount, region, &port16, true, isDev)
	go spawnThreads(&ClientUnknownPool, threadCount, region, &port16, true, isDev)

	// Get opcodes
	sapphire.LoadOpcodes(*region)

	// Control loop
	for {
		select {
		case input := <-commander:
			args := strings.Split(" ", strings.Trim(input, "\n\r "))
			command := args[0]
			args = args[1:]
			switch command {
			case "kill":
				return 0
			case "start":
				log.Println("Starting sniff job.")
				go subscriber.Subscribe(sniffer)
			case "stop":
				if sniffer.Active {
					sniffer.Stop()
				}
			case "update":
				log.Println("Downloading latest opcodes...")
				sapphire.LoadOpcodes(*region)
			case "switchregion":
				log.Println("Switching region to ", args[0], ".")
				sapphire.LoadOpcodes(args[0])
			default:
				log.Println("Unknown command recieved: \"", command, "\"")
			}
		case message := <-subscriber.IngressEvents:
			if _, ok := sapphire.ServerZoneIpcType.Values[message.Opcode]; ok {
				ServerZonePool.Put(message)
			} else if _, ok := sapphire.ServerLobbyIpcType.Values[message.Opcode]; ok {
				ServerLobbyPool.Put(message)
			} else if _, ok := sapphire.ServerChatIpcType.Values[message.Opcode]; ok {
				ServerChatPool.Put(message)
			} else {
				ServerUnknownPool.Put(message)
			}
		case message := <-subscriber.EgressEvents:
			if _, ok := sapphire.ClientZoneIpcType.Values[message.Opcode]; ok {
				ClientZonePool.Put(message)
			} else if _, ok := sapphire.ClientLobbyIpcType.Values[message.Opcode]; ok {
				ClientLobbyPool.Put(message)
			} else if _, ok := sapphire.ClientChatIpcType.Values[message.Opcode]; ok {
				ClientChatPool.Put(message)
			} else {
				ClientUnknownPool.Put(message)
			}
		}
	}
}
