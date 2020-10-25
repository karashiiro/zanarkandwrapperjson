package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ayyaruq/zanarkand"
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
	dataPath := flag.String("DataPath", "", "Sets the download path for internet resources, such as the opcode store.")
	flag.Parse()

	// Set up our control mechanism
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

	networkInterface, err := GetNetworkInterface(networkDevice)
	if err != nil {
		log.Fatalln(err)
		return 1
	}

	// Initialize a sniffer on the network interface
	sniffer, err := zanarkand.NewSniffer("pcap", networkInterface)
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
	var connections []net.Conn
	log.Println("Server started on port: " + *port)
	go func() {
		http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			curConn, _, _, err := ws.UpgradeHTTP(r, w)
			i := len(connections) // Store the length of the slice, we'll use this to remove the connection by index later.
			connections = append(connections, curConn)
			if err != nil {
				log.Println(err)
			}
			log.Println("Connection established with", curConn.RemoteAddr().String())
			go func() {
				defer curConn.Close()

				for {
					data, _, err := wsutil.ReadClientData(curConn)
					if err != nil {
						log.Println(err)
						connections = remove(connections, i)
						break
					}

					message := string(data[:])
					commander <- message
				}
			}()
		}))
	}()

	// Get resources
	sapphire.LoadOpcodes(*region, *dataPath)
	sapphire.LoadConstants(*region, *dataPath)
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
				sapphire.LoadOpcodes(*region, *dataPath)
				sapphire.LoadConstants(*region, *dataPath)
			case "switchregion":
				log.Println("Switching region to ", args[0], ".")
				region = &args[0]
				sapphire.LoadOpcodes(*region, *dataPath)
				sapphire.LoadConstants(*region, *dataPath)
			default:
				log.Println("Unknown command recieved: \"", command, "\"")
			}
		case message := <-subscriber.IngressEvents:
			ipcStructure := ParseMessage(message, *region, Ingress, *isDev)
			go SerializePackout(ipcStructure, connections, *isDev)
		case message := <-subscriber.EgressEvents:
			ipcStructure := ParseMessage(message, *region, Egress, *isDev)
			go SerializePackout(ipcStructure, connections, *isDev)
		}
	}
}

func remove(s []net.Conn, i int) []net.Conn {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
