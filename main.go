package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ayyaruq/zanarkand"
	"github.com/ayyaruq/zanarkand/devices"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/karashiiro/ZanarkandWrapperJSON/sapphire"
)

func main() {
	os.Exit(goLikeMain())
}

func goLikeMain() int {
	// MonitorType doesn't apply, nor do ProcessID or ParseAlgorithm
	region := flag.String("Region", "Global", "Sets the IPC version to Global/CN/KR.")
	port := flag.String("Port", "13346", "Sets the port for the IPC connection between this application and Node.js.")
	networkDevice := net.ParseIP(*flag.String("LocalIP", "", "Specifies a network device by IP, to capture traffic on."))
	isDev := flag.Bool("Dev", true, "Enables the developer mode, enabling raw data output.")
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

	// Initialize websocket
	var cnctns []net.Conn
	log.Println("Server started on port: " + *port)
	go func() {
		http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			curConn, _, _, err := ws.UpgradeHTTP(r, w)
			i := len(cnctns) // Store the length of the slice, we'll use this to remove the connection by index later.
			cnctns = append(cnctns, curConn)
			if err != nil {
				log.Println(err)
			}
			log.Println("Connection established with", curConn.RemoteAddr().String())
			go func() {
				defer curConn.Close()

				for {
					_, _, err := wsutil.ReadClientData(curConn)
					if err != nil {
						log.Println(err)
						cnctns = remove(cnctns, i)
						break
					}
				}
			}()
		}))
	}()

	// Get resources
	sapphire.LoadOpcodes(*region)
	sapphire.LoadDynamicConstants(*region)
	log.Println("Initialization complete!")

	// Control loop
	for {
		select {
		case input := <-commander:
			args := strings.Split(strings.Trim(input, "\n\r "), " ")
			command := args[0]
			args = args[1:]
			switch command {
			case "kill":
				log.Println("Exiting application...")
				return 0
			case "start":
				log.Println("Starting sniff job.")
				go subscriber.Subscribe(sniffer)
			case "stop":
				if sniffer.Active {
					log.Println("Stopping sniff job.")
					sniffer.Stop()
				}
			case "update":
				sapphire.LoadOpcodes(*region)
				sapphire.LoadDynamicConstants(*region)
			case "switchregion":
				log.Println("Switching region to ", args[0], ".")
				sapphire.LoadOpcodes(args[0])
				sapphire.LoadDynamicConstants(args[0])
			default:
				log.Println("Unknown command recieved: \"", command, "\"")
			}
		case message := <-subscriber.IngressEvents:
			go parseMessage(message, *region, cnctns, false, *isDev)
		case message := <-subscriber.EgressEvents:
			go parseMessage(message, *region, cnctns, true, *isDev)
		}
	}
}

func remove(s []net.Conn, i int) []net.Conn {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
