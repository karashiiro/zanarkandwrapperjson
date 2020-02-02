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
	// Looking for ranges:
	// 24-bit private block: 10.*.*.*
	// 20-bit private block: 172.16.*.*-172.31.*.*
	// 16-bit private block: 192.168.*.*
	_, subnet24, _ := net.ParseCIDR("10.0.0.0/8")
	_, subnet20, _ := net.ParseCIDR("172.16.0.0/12")
	_, subnet16, _ := net.ParseCIDR("192.168.0.0/16")

	rNetIfaces, err := devices.ListDeviceNames(false, true)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	netIfaces, _ := devices.ListDeviceNames(false, false) // It seems safe to assume that if the call that includes this and more doesn't error, this will not error if called immediately after either
	netIfaceIdx := len(netIfaces) - 1
	for i, nif := range rNetIfaces {
		ip := net.ParseIP(regexp.MustCompile("\\d+\\.\\d+\\.\\d+\\.\\d+").FindString(nif)) // Lazy but it works and it only runs once
		if subnet24.Contains(ip) || subnet20.Contains(ip) || subnet16.Contains(ip) {
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

	/*go sniffer.Start()
	for {
		frame, _ := sniffer.NextFrame()
		if frame != nil {
			gross := make([]int, len(frame.Body))
			for i, b := range frame.Body {
				gross[i] = int(b)
			}
			str, _ := json.Marshal(gross)
			log.Println(string(str))
		}
	}*/

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
