package sapphire

import (
	"encoding/json"
	"log"
)

// DynamicConstants is a Bimap containing constants that change from patch to patch.
var DynamicConstants Bimap32

// DynamicConstantsJSON is a representation of the constants downloaded from the server.
type DynamicConstantsJSON struct {
	Global map[string]uint32
	CN     map[string]uint32
	KR     map[string]uint32
}

var constSource = "https://raw.githubusercontent.com/karashiiro/FFXIVOpcodes/master/dynamic-constants.min.json"

// LoadDynamicConstants loads constants from the source URL.
func LoadDynamicConstants(region string) {
	log.Println("Downloading dynamic constants...")

	DynamicConstants.ByKeys = make(map[string]uint32)

	// Download opcode JSON and marshal it
	constantFile, err := GetFile("dynamic-constants.json", constSource)

	var constantStore DynamicConstantsJSON
	err = json.NewDecoder(constantFile).Decode(&constantStore)
	if err != nil {
		log.Fatalln(err)
	}

	// Load the opcodes
	if region == "Global" {
		DynamicConstants.ByKeys = constantStore.Global
	} else if region == "CN" {
		DynamicConstants.ByKeys = constantStore.CN
	} else { // region == "KR"
		DynamicConstants.ByKeys = constantStore.KR
	}

	DynamicConstants.ByValues = ReverseMap32(DynamicConstants.ByKeys)

	log.Println("Done!")
}
