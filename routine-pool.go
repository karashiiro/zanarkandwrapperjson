package main

import (
	"sync"

	"github.com/ayyaruq/zanarkand"
)

// ServerZonePool has water.
var ServerZonePool = sync.Pool{}

// ClientZonePool has water.
var ClientZonePool = sync.Pool{}

// ServerLobbyPool has less-sanitary water.
var ServerLobbyPool = sync.Pool{}

// ClientLobbyPool has less-sanitary water.
var ClientLobbyPool = sync.Pool{}

// ServerChatPool has water.
var ServerChatPool = sync.Pool{}

// ClientChatPool has water.
var ClientChatPool = sync.Pool{}

// ServerUnknownPool has muck.
var ServerUnknownPool = sync.Pool{}

// ClientUnknownPool has muck.
var ClientUnknownPool = sync.Pool{}

func spawnThreads(pool *sync.Pool, count int, region *string, port *uint16, isDirectionEgress bool, isDev *bool) {
	for i := 0; i < count; i++ {
		go readPool(pool, region, port, isDirectionEgress, isDev)
	}
}

func readPool(pool *sync.Pool, region *string, port *uint16, isDirectionEgress bool, isDev *bool) {
	for {
		nilable := pool.Get()
		if nilable == nil {
			continue
		}
		message := nilable.(*zanarkand.GameEventMessage)
		parseMessage(message, region, port, isDirectionEgress, isDev)
	}
}
