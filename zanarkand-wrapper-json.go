package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ayyaruq/zanarkand"
	devices "github.com/ayyaruq/zanarkand/devices"
)

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

	// Get the default network device
	netIfaces, err := devices.ListDeviceNames(false, false)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	rNetIfaces, err := devices.ListDeviceNames(false, true)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	netIfaceIdx := len(netIfaces) - 1
	for i, nif := range rNetIfaces {
		nif = strings.TrimFunc(nif, func(c rune) bool {
			return c == '[' || c == ']'
		})
		// Looking for ranges:
		// 24-bit private block: 10.*.*.*
		// 20-bit private block: 172.16.*.*-172.31.*.*
		// 16-bit private block: 192.168.*.*
		ipAddress := regexp.MustCompile("( |\\[|\\])").Split(nif, 3) // Handles IPv6 addresses "DeviceName [ffff::ffff:ffff:ffff:ffff 0.0.0.0]"
		if len(ipAddress) >= 2 {
			ipAddress = regexp.MustCompile("\\.").Split(ipAddress[len(ipAddress)-1], 4)[0:2] // Discard third and fourth fields
		}
		if ipAddress[0] == "10" {
			netIfaceIdx = i
		} else if ipAddress[0] == "172" {
			net2, err := strconv.Atoi(ipAddress[1])
			if err != nil {
				continue
			}
			if net2 >= 16 && net2 <= 31 {
				netIfaceIdx = i
			}
		} else if ipAddress[0] == "192" && ipAddress[1] == "168" {
			netIfaceIdx = i
		}
	}

	// Initialize a sniffer on the network device
	sniffer, err := zanarkand.NewSniffer("", netIfaces[netIfaceIdx])
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
			switch strings.Trim(command, "\n\r ") {
			case "kill":
				return 0
			case "start":
				go subscriber.Subscribe(sniffer)
			case "stop":
				if sniffer.Active() {
					sniffer.Stop()
				}
			default:
				log.Println("Unknown command recieved: \"", command, "\"")
			}
		case message := <-subscriber.Events:
			go parseMessage(message, *region, uint16(*port))
		}
	}
}
