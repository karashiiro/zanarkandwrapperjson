package main

import (
	"sync"

	"github.com/ayyaruq/zanarkand"
)

// ServerZonePool has water.
var ServerZonePool = sync.Pool{}

// ClientZonePool has water.
var ClientZonePool = sync.Pool{}

// LobbyPool has less-sanitary water.
var LobbyPool = sync.Pool{}

// ChatPool has water.
var ChatPool = sync.Pool{}

// UnknownPool has muck.
var UnknownPool = sync.Pool{}

func spawnThreads(pool *sync.Pool, count int, region *string, port *uint16, isDev *bool) {
	for i := 0; i < count; i++ {
		go readPool(pool, region, port, isDev)
	}
}

func readPool(pool *sync.Pool, region *string, port *uint16, isDev *bool) {
	for {
		message := pool.Get().(*zanarkand.GameEventMessage)
		if message == nil {
			continue
		}
		parseMessage(message, region, port, isDev)
	}
}
