package sapphire

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
)

// Constants is a Bimap containing constants that change from patch to patch.
var Constants Bimap32

// ConstantsJSON is a representation of the constants downloaded from the server.
type ConstantsJSON struct {
	Global map[string]uint32
	CN     map[string]uint32
	KR     map[string]uint32
}

var constSource = "https://raw.githubusercontent.com/karashiiro/FFXIVOpcodes/master/constants.min.json"

// LoadConstants loads constants from the source URL.
func LoadConstants(region string, dataPath string) {
	log.Println("Downloading constants...")

	Constants.ByKeys = make(map[string]uint32)

	// Download opcode JSON and marshal it
	if !exists(dataPath) {
		os.MkdirAll(dataPath, 0664)
	}

	fileName := path.Join(dataPath, "constants.json")
	constantFile, err := GetFile(fileName, constSource)
	if err != nil {
		log.Fatalln(err)
	}

	go PollForUpdates(fileName, constSource)
	go WatchFile(fileName, func(newData io.Reader) {
		log.Println("Got new constants, reloading...")
		unmarshalConstants(newData, region)
		log.Println("Done!")
	})

	unmarshalConstants(constantFile, region)

	log.Println("Done!")
}

func unmarshalConstants(stream io.Reader, region string) {
	var constantStore ConstantsJSON
	err := json.NewDecoder(stream).Decode(&constantStore)
	if err != nil {
		log.Fatalln(err)
	}

	// Load the opcodes
	if region == "Global" {
		Constants.ByKeys = constantStore.Global
	} else if region == "CN" {
		Constants.ByKeys = constantStore.CN
	} else { // region == "KR"
		Constants.ByKeys = constantStore.KR
	}

	Constants.ByValues = ReverseMap32(Constants.ByKeys)
}
