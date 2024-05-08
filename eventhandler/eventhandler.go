package eventhandler

type EventHandler interface {
	RegisterKeyboardHook() error
	UnregisterKeyboardHook() error
	ListenForEvents()
	GetKeyEvents() (<-chan *KeyEvent)
	Teardown()
}