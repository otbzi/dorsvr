package rtspserver

import "github.com/otbzi/dorsvr/livemedia"

type StreamServerState struct {
	subsession  livemedia.IServerMediaSubsession
	streamToken *livemedia.StreamState
}
