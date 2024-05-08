package eventhandler

import "github.com/rvIceBreaker/goclack/keycodes"

type KeyEventType int

const (
	KEYEVENT_UP KeyEventType = iota
	KEYEVENT_DOWN
)

type KeyEvent struct {
	Key          *keycodes.VKey
	Type KeyEventType
}