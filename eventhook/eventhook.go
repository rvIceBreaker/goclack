package eventhook

import (
	"runtime"

	"github.com/rvIceBreaker/goclack/eventhandler"
	"github.com/rvIceBreaker/goclack/win32"
)



type EventHook struct {
	Handler eventhandler.EventHandler
	Callbacks []func(*eventhandler.KeyEvent)
}

func NewEventHook() *EventHook {
	e := EventHook{}
	e.init()

	return &e
}

func (e *EventHook) init() {
	//Determine platform, hook events appropriately
	switch runtime.GOOS {
		case "windows":
			e.Handler = win32.NewWin32EventHandler()

	}
}

func (e *EventHook) RegisterKeyboardHook() error {
	return e.Handler.RegisterKeyboardHook()
}

func (e *EventHook) UnregisterKeyboardHook() error {
	return e.Handler.UnregisterKeyboardHook()
}

func (e *EventHook) RegisterKeyboardCallback(cb func(*eventhandler.KeyEvent)) {
	e.Callbacks = append(e.Callbacks, cb)
}

func (e *EventHook) ListenForEvents() {
	e.Handler.ListenForEvents()
	e.implEventListen()
}

func (e *EventHook) implEventListen() {
	for ev := range e.Handler.GetKeyEvents() {
		e.implPreKeyboardEvent(ev)
		for _,cb := range e.Callbacks {
			cb(ev)
		}
	}
}

func (e *EventHook) implPreKeyboardEvent(ev *eventhandler.KeyEvent) {

}

func (e *EventHook) Teardown() {
	e.Handler.Teardown()
}