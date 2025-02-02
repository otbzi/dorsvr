package livemedia

import (
	sys "syscall"

	gs "github.com/otbzi/dorsvr/groupsock"
)

const maxPayloadSize = 1450

type BasicUDPSink struct {
	MediaSink
	gs             *gs.GroupSock
	maxPayloadSize uint
	outputBuffer   []byte
	nextSendTime   sys.Timeval
}

func NewBasicUDPSink(gs *gs.GroupSock) *BasicUDPSink {
	return &BasicUDPSink{
		gs:             gs,
		maxPayloadSize: maxPayloadSize,
		outputBuffer:   make([]byte, maxPayloadSize),
	}
}

func (s *BasicUDPSink) ContinuePlaying() {
}
