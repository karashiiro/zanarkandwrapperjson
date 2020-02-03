package main

import "github.com/ayyaruq/zanarkand"

// IpcStructure - Struct of fields IPC packets can have
type IpcStructure struct {
	zanarkand.GameEventMessage
	Region    string `json:"region"`
	SubType   string `json:"subType"`
	SuperType string `json:"superType"`
	Type      string `json:"type"`
	IpcParameters
}

// IpcParameters - Holds any IPC struct to be serialized later
type IpcParameters interface{}
