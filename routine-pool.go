package main

import (
	"sync"

	"github.com/ayyaruq/zanarkand"
)

// ServerZonePool has water.
var ServerZonePool = sync.Pool{
	New: func() interface{} {
		return nil
	},
}

// ClientZonePool has water.
var ClientZonePool = sync.Pool{
	New: func() interface{} {
		return nil
	},
}

// LobbyPool has less-sanitary water.
var LobbyPool = sync.Pool{
	New: func() interface{} {
		return nil
	},
}

// ChatPool has water.
var ChatPool = sync.Pool{
	New: func() interface{} {
		return nil
	},
}

// UnknownPool has muck.
var UnknownPool = sync.Pool{
	New: func() interface{} {
		return nil
	},
}

func readPool(pool *sync.Pool, region *string, port *uint16, isDev *bool) {
	for {
		message := pool.Get().(*zanarkand.GameEventMessage)
		if message == nil {
			continue
		}
		go parseMessage(message, region, port, isDev)
	}
}
