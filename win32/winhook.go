package win32

//Welcome to a new reason to hate golang...
var WinHook = newWinHooks()

/*
	WH_KEYBOARD_LL = 13
	WH_KEYBOARD = 2
*/

type EWinHooks struct {
	WH_KEYBOARD_LL *VWinHook
	WH_KEYBOARD *VWinHook
}

type VWinHook struct {
	ID int
	Name string
}

func newWinHooks() (*EWinHooks) {
	return &EWinHooks{
		WH_KEYBOARD_LL: &VWinHook{ID: 13, Name: "WH_KEYBOARD_LL"},
		WH_KEYBOARD: &VWinHook{ID: 2, Name: "WH_KEYBOARD"},
	}
}