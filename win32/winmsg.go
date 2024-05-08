package win32

//Welcome to a new reason to hate golang...
var WinMsg = newWinMsgs()

/*
	WM_KEYDOWN = 256
	WM_SYSKEYDOWN = 260
	WM_KEYUP = 257
	WM_SYSKEYUP = 261
	WM_KEYFIRST = 256
	WM_KEYLAST = 264
	WM_LBUTTONDOWN = 513
	WM_RBUTTONDOWN = 516
*/

type EWinMsgs struct {
	WM_KEYDOWN *VWinMsg
	WM_SYSKEYDOWN *VWinMsg
	WM_KEYUP *VWinMsg
	WM_SYSKEYUP *VWinMsg
	WM_KEYFIRST *VWinMsg
	WM_KEYLAST *VWinMsg
	WM_LBUTTONDOWN *VWinMsg
	WM_RBUTTONDOWN *VWinMsg
}

type VWinMsg struct {
	ID int
	Name string
}

func newWinMsgs() (*EWinMsgs) {
	return &EWinMsgs{
		WM_KEYDOWN: &VWinMsg{Name: "WM_KEYDOWN", ID: 256},
		WM_SYSKEYDOWN: &VWinMsg{Name: "WM_SYSKEYDOWN", ID: 260},
		WM_KEYUP: &VWinMsg{Name: "WM_KEYUP", ID: 257},
		WM_SYSKEYUP: &VWinMsg{Name: "WM_SYSKEYUP", ID: 261},
		WM_KEYFIRST: &VWinMsg{Name: "WM_KEYFIRST", ID: 256},
		WM_KEYLAST: &VWinMsg{Name: "WM_KEYLAST", ID: 264},
		WM_LBUTTONDOWN: &VWinMsg{Name: "WM_LBUTTONDOWN", ID: 513},
		WM_RBUTTONDOWN: &VWinMsg{Name: "WM_RBUTTONDOWN", ID: 516},
	}
}