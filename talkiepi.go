package talkiepi

import (
	"crypto/tls"

	"github.com/dchote/gumble/gumble"
	"github.com/dchote/gumble/gumbleopenal"
)

type Talkiepi struct {
	Config          *gumble.Config
	Client          *gumble.Client
	Address         string
	TLSConfig       tls.Config
	ConnectAttempts uint
	Stream          *gumbleopenal.Stream
	ChannelName     string
	IsConnected     bool
	IsTransmitting  bool
}
