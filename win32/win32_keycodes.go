package win32

import (
	//reflect"

	"fmt"

	"github.com/rvIceBreaker/goclack/keycodes"
	//"golang.org/x/exp/slices"
)

var WinVKey = newWinVKeys()

//Referenced from https://gist.github.com/Nordaj/cc6cadcbee1019a3299939d0f1b86296

//	LeftMouseBtn			=0x01, //Left mouse button
//	RightMouseBtn			=0x02, //Right mouse button
//	CtrlBrkPrcs				=0x03, //Control-break processing
//	MidMouseBtn				=0x04, //Middle mouse button
//	ThumbForward			=0x05, //Thumb button back on mouse aka X1
//	ThumbBack				=0x06, //Thumb button forward on mouse aka X2
//	BackSpace				=0x08, //Backspace key
//	Tab						=0x09, //Tab key
//	Clear					=0x0C, //Clear key
//	Enter					=0x0D, //Enter or Return key
//	Shift					=0x10, //Shift key
//	Control					=0x11, //Ctrl key
//	Alt						=0x12, //Alt key
//	Pause					=0x13, //Pause key
//	CapsLock				=0x14, //Caps lock key
//	Kana					=0x15, //Kana input mode
//	Hangeul					=0x15, //Hangeul input mode
//	Hangul					=0x15, //Hangul input mode
//	Junju					=0x17, //Junja input method
//	Final					=0x18, //Final input method
//	Hanja					=0x19, //Hanja input method
//	Kanji					=0x19, //Kanji input method
//	Escape					=0x1B, //Esc key
//	Convert					=0x1C, //IME convert
//	NonConvert				=0x1D, //IME Non convert
//	Accept					=0x1E, //IME accept
//	ModeChange				=0x1F, //IME mode change
//	Space					=0x20, //Space bar
//	PageUp					=0x21, //Page up key
//	PageDown				=0x22, //Page down key
//	End						=0x23, //End key
//	Home					=0x24, //Home key
//	LeftArrow				=0x25, //Left arrow key
//	UpArrow					=0x26, //Up arrow key
//	RightArrow				=0x27, //Right arrow key
//	DownArrow				=0x28, //Down arrow key
//	Select					=0x29, //Select key
//	Print					=0x2A, //Print key
//	Execute					=0x2B, //Execute key
//	PrintScreen				=0x2C, //Print screen key
//	Inser					=0x2D, //Insert key
//	Delete					=0x2E, //Delete key
//	Help					=0x2F, //Help key
//	Num0					=0x30, //Top row 0 key (Matches '0')
//	Num1					=0x31, //Top row 1 key (Matches '1')
//	Num2					=0x32, //Top row 2 key (Matches '2')
//	Num3					=0x33, //Top row 3 key (Matches '3')
//	Num4					=0x34, //Top row 4 key (Matches '4')
//	Num5					=0x35, //Top row 5 key (Matches '5')
//	Num6					=0x36, //Top row 6 key (Matches '6')
//	Num7					=0x37, //Top row 7 key (Matches '7')
//	Num8					=0x38, //Top row 8 key (Matches '8')
//	Num9					=0x39, //Top row 9 key (Matches '9')
//	A						=0x41, //A key (Matches 'A')
//	B						=0x42, //B key (Matches 'B')
//	C						=0x43, //C key (Matches 'C')
//	D						=0x44, //D key (Matches 'D')
//	E						=0x45, //E key (Matches 'E')
//	F						=0x46, //F key (Matches 'F')
//	G						=0x47, //G key (Matches 'G')
//	H						=0x48, //H key (Matches 'H')
//	I						=0x49, //I key (Matches 'I')
//	J						=0x4A, //J key (Matches 'J')
//	K						=0x4B, //K key (Matches 'K')
//	L						=0x4C, //L key (Matches 'L')
//	M						=0x4D, //M key (Matches 'M')
//	N						=0x4E, //N key (Matches 'N')
//	O						=0x4F, //O key (Matches 'O')
//	P						=0x50, //P key (Matches 'P')
//	Q						=0x51, //Q key (Matches 'Q')
//	R						=0x52, //R key (Matches 'R')
//	S						=0x53, //S key (Matches 'S')
//	T						=0x54, //T key (Matches 'T')
//	U						=0x55, //U key (Matches 'U')
//	V						=0x56, //V key (Matches 'V')
//	W						=0x57, //W key (Matches 'W')
//	X						=0x58, //X key (Matches 'X')
//	Y						=0x59, //Y key (Matches 'Y')
//	Z						=0x5A, //Z key (Matches 'Z')
//	LeftWin					=0x5B, //Left windows key
//	RightWin				=0x5C, //Right windows key
//	Apps					=0x5D, //Applications key
//	Sleep					=0x5F, //Computer sleep key
//	Numpad0					=0x60, //Numpad 0
//	Numpad1					=0x61, //Numpad 1
//	Numpad2					=0x62, //Numpad 2
//	Numpad3					=0x63, //Numpad 3
//	Numpad4					=0x64, //Numpad 4
//	Numpad5					=0x65, //Numpad 5
//	Numpad6					=0x66, //Numpad 6
//	Numpad7					=0x67, //Numpad 7
//	Numpad8					=0x68, //Numpad 8
//	Numpad9					=0x69, //Numpad 9
//	Multiply				=0x6A, //Multiply key
//	Add						=0x6B, //Add key
//	Separator				=0x6C, //Separator key
//	Subtract				=0x6D, //Subtract key
//	Decimal					=0x6E, //Decimal key
//	Divide					=0x6F, //Divide key
//	F1						=0x70, //F1
//	F2						=0x71, //F2
//	F3						=0x72, //F3
//	F4						=0x73, //F4
//	F5						=0x74, //F5
//	F6						=0x75, //F6
//	F7						=0x76, //F7
//	F8						=0x77, //F8
//	F9						=0x78, //F9
//	F10						=0x79, //F10
//	F11						=0x7A, //F11
//	F12						=0x7B, //F12
//	F13						=0x7C, //F13
//	F14						=0x7D, //F14
//	F15						=0x7E, //F15
//	F16						=0x7F, //F16
//	F17						=0x80, //F17
//	F18						=0x81, //F18
//	F19						=0x82, //F19
//	F20						=0x83, //F20
//	F21						=0x84, //F21
//	F22						=0x85, //F22
//	F23						=0x86, //F23
//	F24						=0x87, //F24
//	NavigationView			=0x88, //reserved
//	NavigationMenu			=0x89, //reserved
//	NavigationUp			=0x8A, //reserved
//	NavigationDown			=0x8B, //reserved
//	NavigationLeft			=0x8C, //reserved
//	NavigationRight			=0x8D, //reserved
//	NavigationAccept		=0x8E, //reserved
//	NavigationCancel		=0x8F, //reserved
//	NumLock					=0x90, //Num lock key
//	ScrollLock				=0x91, //Scroll lock key
//	NumpadEqual				=0x92, //Numpad =
//	FJ_Jisho				=0x92, //Dictionary key
//	FJ_Masshou				=0x93, //Unregister word key
//	FJ_Touroku				=0x94, //Register word key
//	FJ_Loya					=0x95, //Left OYAYUBI key
//	FJ_Roya					=0x96, //Right OYAYUBI key
//	LeftShift				=0xA0, //Left shift key
//	RightShift				=0xA1, //Right shift key
//	LeftCtrl				=0xA2, //Left control key
//	RightCtrl				=0xA3, //Right control key
//	LeftMenu				=0xA4, //Left menu key
//	RightMenu				=0xA5, //Right menu
//	BrowserBack				=0xA6, //Browser back button
//	BrowserForward			=0xA7, //Browser forward button
//	BrowserRefresh			=0xA8, //Browser refresh button
//	BrowserStop				=0xA9, //Browser stop button
//	BrowserSearch			=0xAA, //Browser search button
//	BrowserFavorites		=0xAB, //Browser favorites button
//	BrowserHome				=0xAC, //Browser home button
//	VolumeMute				=0xAD, //Volume mute button
//	VolumeDown				=0xAE, //Volume down button
//	VolumeUp				=0xAF, //Volume up button
//	NextTrack				=0xB0, //Next track media button
//	PrevTrack				=0xB1, //Previous track media button
//	Stop					=0xB2, //Stop media button
//	PlayPause				=0xB3, //Play/pause media button
//	Mail					=0xB4, //Launch mail button
//	MediaSelect				=0xB5, //Launch media select button
//	App1					=0xB6, //Launch app 1 button
//	App2					=0xB7, //Launch app 2 button
//	OEM1					=0xBA, //;: key for US or misc keys for others
//	Plus					=0xBB, //Plus key
//	Comma					=0xBC, //Comma key
//	Minus					=0xBD, //Minus key
//	Period					=0xBE, //Period key
//	OEM2					=0xBF, //? for US or misc keys for others
//	OEM3					=0xC0, //~ for US or misc keys for others
//	Gamepad_A				=0xC3, //Gamepad A button
//	Gamepad_B				=0xC4, //Gamepad B button
//	Gamepad_X				=0xC5, //Gamepad X button
//	Gamepad_Y				=0xC6, //Gamepad Y button
//	GamepadRightBumper		=0xC7, //Gamepad right bumper
//	GamepadLeftBumper		=0xC8, //Gamepad left bumper
//	GamepadLeftTrigger		=0xC9, //Gamepad left trigger
//	GamepadRightTrigger		=0xCA, //Gamepad right trigger
//	GamepadDPadUp			=0xCB, //Gamepad DPad up
//	GamepadDPadDown			=0xCC, //Gamepad DPad down
//	GamepadDPadLeft			=0xCD, //Gamepad DPad left
//	GamepadDPadRight		=0xCE, //Gamepad DPad right
//	GamepadMenu				=0xCF, //Gamepad menu button
//	GamepadView				=0xD0, //Gamepad view button
//	GamepadLeftStickBtn		=0xD1, //Gamepad left stick button
//	GamepadRightStickBtn	=0xD2, //Gamepad right stick button
//	GamepadLeftStickUp		=0xD3, //Gamepad left stick up
//	GamepadLeftStickDown	=0xD4, //Gamepad left stick down
//	GamepadLeftStickRight	=0xD5, //Gamepad left stick right
//	GamepadLeftStickLeft	=0xD6, //Gamepad left stick left
//	GamepadRightStickUp		=0xD7, //Gamepad right stick up
//	GamepadRightStickDown	=0xD8, //Gamepad right stick down
//	GamepadRightStickRight	=0xD9, //Gamepad right stick right
//	GamepadRightStickLeft	=0xDA, //Gamepad right stick left
//	OEM4					=0xDB, //[ for US or misc keys for others
//	OEM5					=0xDC, //\ for US or misc keys for others
//	OEM6					=0xDD, //] for US or misc keys for others
//	OEM7					=0xDE, //' for US or misc keys for others
//	OEM8					=0xDF, //Misc keys for others
//	OEMAX					=0xE1, //AX key on Japanese AX keyboard
//	OEM102					=0xE2, //"<>" or "\|" on RT 102-key keyboard
//	ICOHelp					=0xE3, //Help key on ICO
//	ICO00					=0xE4, //00 key on ICO
//	ProcessKey				=0xE5, //Process key input method
//	OEMCLEAR				=0xE6, //OEM specific
//	Packet					=0xE7, //IDK man try to google it
//	OEMReset				=0xE9, //OEM reset button
//	OEMJump					=0xEA, //OEM jump button
//	OEMPA1					=0xEB, //OEM PA1 button
//	OEMPA2					=0xEC, //OEM PA2 button
//	OEMPA3					=0xED, //OEM PA3 button
//	OEMWSCtrl				=0xEE, //OEM WS Control button
//	OEMCusel				=0xEF, //OEM CUSEL button
//	OEMAttn					=0xF0, //OEM ATTN button
//	OEMFinish				=0xF1, //OEM finish button
//	OEMCopy					=0xF2, //OEM copy button
//	OEMAuto					=0xF3, //OEM auto button
//	OEMEnlw					=0xF4, //OEM ENLW
//	OEMBackTab				=0xF5, //OEM back tab
//	Attn					=0xF6, //Attn
//	CrSel					=0xF7, //CrSel
//	ExSel					=0xF8, //ExSel
//	EraseEOF				=0xF9, //Erase EOF key
//	Play					=0xFA, //Play key
//	Zoom					=0xFB, //Zoom key
//	NoName					=0xFC, //No name
//	PA1						=0xFD, //PA1 key
//	OEMClear				=0xFE, //OEM Clear key

type EWinVKey struct {
	KEY_BackSpace     *VWinVKey //		=0x08, //Backspace key
	KEY_Tab           *VWinVKey //				=0x09, //Tab key
	KEY_Clear         *VWinVKey //			=0x0C, //Clear key
	KEY_Enter         *VWinVKey //			=0x0D, //Enter or Return key
	KEY_Shift         *VWinVKey //			=0x10, //Shift key
	KEY_Control       *VWinVKey //			=0x11, //Ctrl key
	KEY_Alt           *VWinVKey //				=0x12, //Alt key
	KEY_Pause         *VWinVKey //			=0x13, //Pause key
	KEY_CapsLock      *VWinVKey //		=0x14, //Caps lock key
	KEY_Space         *VWinVKey //			=0x20, //Space bar
	KEY_PageUp        *VWinVKey //			=0x21, //Page up key
	KEY_PageDown      *VWinVKey //		=0x22, //Page down key
	KEY_End           *VWinVKey //				=0x23, //End key
	KEY_Home          *VWinVKey //			=0x24, //Home key
	KEY_LeftArrow     *VWinVKey //		=0x25, //Left arrow key
	KEY_UpArrow       *VWinVKey //			=0x26, //Up arrow key
	KEY_RightArrow    *VWinVKey //		=0x27, //Right arrow key
	KEY_DownArrow     *VWinVKey //		=0x28, //Down arrow key
	KEY_PrintScreen   *VWinVKey //		=0x2C, //Print screen key
	KEY_Inser         *VWinVKey //			=0x2D, //Insert key
	KEY_Delete        *VWinVKey //			=0x2E, //Delete key
	KEY_Num0          *VWinVKey //			=0x30, //Top row 0 key (Matches '0')
	KEY_Num1          *VWinVKey //			=0x31, //Top row 1 key (Matches '1')
	KEY_Num2          *VWinVKey //			=0x32, //Top row 2 key (Matches '2')
	KEY_Num3          *VWinVKey //			=0x33, //Top row 3 key (Matches '3')
	KEY_Num4          *VWinVKey //			=0x34, //Top row 4 key (Matches '4')
	KEY_Num5          *VWinVKey //			=0x35, //Top row 5 key (Matches '5')
	KEY_Num6          *VWinVKey //			=0x36, //Top row 6 key (Matches '6')
	KEY_Num7          *VWinVKey //			=0x37, //Top row 7 key (Matches '7')
	KEY_Num8          *VWinVKey //			=0x38, //Top row 8 key (Matches '8')
	KEY_Num9          *VWinVKey //			=0x39, //Top row 9 key (Matches '9')
	KEY_A             *VWinVKey //				=0x41, //A key (Matches 'A')
	KEY_B             *VWinVKey //				=0x42, //B key (Matches 'B')
	KEY_C             *VWinVKey //				=0x43, //C key (Matches 'C')
	KEY_D             *VWinVKey //				=0x44, //D key (Matches 'D')
	KEY_E             *VWinVKey //				=0x45, //E key (Matches 'E')
	KEY_F             *VWinVKey //				=0x46, //F key (Matches 'F')
	KEY_G             *VWinVKey //				=0x47, //G key (Matches 'G')
	KEY_H             *VWinVKey //				=0x48, //H key (Matches 'H')
	KEY_I             *VWinVKey //				=0x49, //I key (Matches 'I')
	KEY_J             *VWinVKey //				=0x4A, //J key (Matches 'J')
	KEY_K             *VWinVKey //				=0x4B, //K key (Matches 'K')
	KEY_L             *VWinVKey //				=0x4C, //L key (Matches 'L')
	KEY_M             *VWinVKey //				=0x4D, //M key (Matches 'M')
	KEY_N             *VWinVKey //				=0x4E, //N key (Matches 'N')
	KEY_O             *VWinVKey //				=0x4F, //O key (Matches 'O')
	KEY_P             *VWinVKey //				=0x50, //P key (Matches 'P')
	KEY_Q             *VWinVKey //				=0x51, //Q key (Matches 'Q')
	KEY_R             *VWinVKey //				=0x52, //R key (Matches 'R')
	KEY_S             *VWinVKey //				=0x53, //S key (Matches 'S')
	KEY_T             *VWinVKey //				=0x54, //T key (Matches 'T')
	KEY_U             *VWinVKey //				=0x55, //U key (Matches 'U')
	KEY_V             *VWinVKey //				=0x56, //V key (Matches 'V')
	KEY_W             *VWinVKey //				=0x57, //W key (Matches 'W')
	KEY_X             *VWinVKey //				=0x58, //X key (Matches 'X')
	KEY_Y             *VWinVKey //				=0x59, //Y key (Matches 'Y')
	KEY_Z             *VWinVKey //				=0x5A, //Z key (Matches 'Z')
	KEY_LeftWin       *VWinVKey //			=0x5B, //Left windows key
	KEY_RightWin      *VWinVKey //		=0x5C, //Right windows key
	KEY_NUM_0         *VWinVKey //			=0x60, //Numpad 0
	KEY_NUM_1         *VWinVKey //			=0x61, //Numpad 1
	KEY_NUM_2         *VWinVKey //			=0x62, //Numpad 2
	KEY_NUM_3         *VWinVKey //			=0x63, //Numpad 3
	KEY_NUM_4         *VWinVKey //			=0x64, //Numpad 4
	KEY_NUM_5         *VWinVKey //			=0x65, //Numpad 5
	KEY_NUM_6         *VWinVKey //			=0x66, //Numpad 6
	KEY_NUM_7         *VWinVKey //			=0x67, //Numpad 7
	KEY_NUM_8         *VWinVKey //			=0x68, //Numpad 8
	KEY_NUM_9         *VWinVKey //			=0x69, //Numpad 9
	KEY_NUM_Multiply  *VWinVKey //		=0x6A, //Multiply key
	KEY_NUM_Add       *VWinVKey //				=0x6B, //Add key
	KEY_NUM_Separator *VWinVKey //		=0x6C, //Separator key
	KEY_NUM_Subtract  *VWinVKey //		=0x6D, //Subtract key
	KEY_NUM_Decimal   *VWinVKey //			=0x6E, //Decimal key
	KEY_NUM_Divide    *VWinVKey //			=0x6F, //Divide key
	KEY_F1            *VWinVKey //				=0x70, //F1
	KEY_F2            *VWinVKey //				=0x71, //F2
	KEY_F3            *VWinVKey //				=0x72, //F3
	KEY_F4            *VWinVKey //				=0x73, //F4
	KEY_F5            *VWinVKey //				=0x74, //F5
	KEY_F6            *VWinVKey //				=0x75, //F6
	KEY_F7            *VWinVKey //				=0x76, //F7
	KEY_F8            *VWinVKey //				=0x77, //F8
	KEY_F9            *VWinVKey //				=0x78, //F9
	KEY_F10           *VWinVKey //				=0x79, //F10
	KEY_F11           *VWinVKey //				=0x7A, //F11
	KEY_F12           *VWinVKey //				=0x7B, //F12
	KEY_F13           *VWinVKey //				=0x7C, //F13
	KEY_F14           *VWinVKey //				=0x7D, //F14
	KEY_F15           *VWinVKey //				=0x7E, //F15
	KEY_F16           *VWinVKey //				=0x7F, //F16
	KEY_F17           *VWinVKey //				=0x80, //F17
	KEY_F18           *VWinVKey //				=0x81, //F18
	KEY_F19           *VWinVKey //				=0x82, //F19
	KEY_F20           *VWinVKey //				=0x83, //F20
	KEY_F21           *VWinVKey //				=0x84, //F21
	KEY_F22           *VWinVKey //				=0x85, //F22
	KEY_F23           *VWinVKey //				=0x86, //F23
	KEY_F24           *VWinVKey //				=0x87, //F24
	KEY_LeftShift     *VWinVKey //		=0xA0, //Left shift key
	KEY_RightShift    *VWinVKey //		=0xA1, //Right shift key
	KEY_LeftCtrl      *VWinVKey //		=0xA2, //Left control key
	KEY_RightCtrl     *VWinVKey //		=0xA3, //Right control key
	KEY_LeftMenu      *VWinVKey //		=0xA4, //Left menu key
	KEY_RightMenu     *VWinVKey //		=0xA5, //Right menu
	KEY_OEM1          *VWinVKey //			=0xBA, //;: key for US or misc keys for others
	KEY_Plus          *VWinVKey //			=0xBB, //Plus key
	KEY_Comma         *VWinVKey //			=0xBC, //Comma key
	KEY_Minus         *VWinVKey //			=0xBD, //Minus key
	KEY_Period        *VWinVKey //			=0xBE, //Period key
	KEY_OEM2          *VWinVKey //			=0xBF, //? for US or misc keys for others
	KEY_OEM3          *VWinVKey //			=0xC0, //~ for US or misc keys for others
	KEY_OEM4          *VWinVKey //			=0xDB, //[ for US or misc keys for others
	KEY_OEM5          *VWinVKey //			=0xDC, //\ for US or misc keys for others
	KEY_OEM6          *VWinVKey //			=0xDD, //] for US or misc keys for others
	KEY_OEM7          *VWinVKey //			=0xDE, //' for US or misc keys for others
	KEY_OEM102        *VWinVKey //			=0xE2, //"<>" or "\|" on RT 102-key keyboard
}

type VWinVKey struct {
	KeyCode uint8
	Key     *keycodes.VKey
}

func newWinVKeys() *EWinVKey {
	return &EWinVKey{
		KEY_BackSpace:     	&VWinVKey{KeyCode: 0x08, Key: keycodes.VKeys.KEY_BackSpace},
		KEY_Tab:           	&VWinVKey{KeyCode: 0x09, Key: keycodes.VKeys.KEY_Tab},
		KEY_Clear:         	&VWinVKey{KeyCode: 0x0C, Key: keycodes.VKeys.KEY_Clear},
		KEY_Enter:         	&VWinVKey{KeyCode: 0x0D, Key: keycodes.VKeys.KEY_Enter},
		KEY_Shift:         	&VWinVKey{KeyCode: 0x10, Key: keycodes.VKeys.KEY_Shift},
		KEY_Control:       	&VWinVKey{KeyCode: 0x11, Key: keycodes.VKeys.KEY_Control},
		KEY_Alt:           	&VWinVKey{KeyCode: 0x12, Key: keycodes.VKeys.KEY_Alt},
		KEY_Pause:         	&VWinVKey{KeyCode: 0x13, Key: keycodes.VKeys.KEY_Pause},
		KEY_CapsLock:      	&VWinVKey{KeyCode: 0x14, Key: keycodes.VKeys.KEY_CapsLock},
		KEY_Space:         	&VWinVKey{KeyCode: 0x20, Key: keycodes.VKeys.KEY_Space},
		KEY_PageUp:        	&VWinVKey{KeyCode: 0x21, Key: keycodes.VKeys.KEY_PageUp},
		KEY_PageDown:      	&VWinVKey{KeyCode: 0x22, Key: keycodes.VKeys.KEY_PageDown},
		KEY_End:           	&VWinVKey{KeyCode: 0x23, Key: keycodes.VKeys.KEY_End},
		KEY_Home:          	&VWinVKey{KeyCode: 0x24, Key: keycodes.VKeys.KEY_Home},
		KEY_LeftArrow:     	&VWinVKey{KeyCode: 0x25, Key: keycodes.VKeys.KEY_LeftArrow},
		KEY_UpArrow:       	&VWinVKey{KeyCode: 0x26, Key: keycodes.VKeys.KEY_UpArrow},
		KEY_RightArrow:    	&VWinVKey{KeyCode: 0x27, Key: keycodes.VKeys.KEY_RightArrow},
		KEY_DownArrow:     	&VWinVKey{KeyCode: 0x28, Key: keycodes.VKeys.KEY_DownArrow},
		KEY_PrintScreen:   	&VWinVKey{KeyCode: 0x2C, Key: keycodes.VKeys.KEY_PrintScreen},
		KEY_Inser:         	&VWinVKey{KeyCode: 0x2D, Key: keycodes.VKeys.KEY_Inser},
		KEY_Delete:        	&VWinVKey{KeyCode: 0x2E, Key: keycodes.VKeys.KEY_Delete},
		KEY_Num0:          	&VWinVKey{KeyCode: 0x30, Key: keycodes.VKeys.KEY_Num0},
		KEY_Num1:          	&VWinVKey{KeyCode: 0x31, Key: keycodes.VKeys.KEY_Num1},
		KEY_Num2:          	&VWinVKey{KeyCode: 0x32, Key: keycodes.VKeys.KEY_Num2},
		KEY_Num3:          	&VWinVKey{KeyCode: 0x33, Key: keycodes.VKeys.KEY_Num3},
		KEY_Num4:          	&VWinVKey{KeyCode: 0x34, Key: keycodes.VKeys.KEY_Num4},
		KEY_Num5:          	&VWinVKey{KeyCode: 0x35, Key: keycodes.VKeys.KEY_Num5},
		KEY_Num6:          	&VWinVKey{KeyCode: 0x36, Key: keycodes.VKeys.KEY_Num6},
		KEY_Num7:          	&VWinVKey{KeyCode: 0x37, Key: keycodes.VKeys.KEY_Num7},
		KEY_Num8:          	&VWinVKey{KeyCode: 0x38, Key: keycodes.VKeys.KEY_Num8},
		KEY_Num9:          	&VWinVKey{KeyCode: 0x39, Key: keycodes.VKeys.KEY_Num9},
		KEY_A:             	&VWinVKey{KeyCode: 0x41, Key: keycodes.VKeys.KEY_A},
		KEY_B:             	&VWinVKey{KeyCode: 0x42, Key: keycodes.VKeys.KEY_B},
		KEY_C:             	&VWinVKey{KeyCode: 0x43, Key: keycodes.VKeys.KEY_C},
		KEY_D:             	&VWinVKey{KeyCode: 0x44, Key: keycodes.VKeys.KEY_D},
		KEY_E:             	&VWinVKey{KeyCode: 0x45, Key: keycodes.VKeys.KEY_E},
		KEY_F:             	&VWinVKey{KeyCode: 0x46, Key: keycodes.VKeys.KEY_F},
		KEY_G:             	&VWinVKey{KeyCode: 0x47, Key: keycodes.VKeys.KEY_G},
		KEY_H:             	&VWinVKey{KeyCode: 0x48, Key: keycodes.VKeys.KEY_H},
		KEY_I:             	&VWinVKey{KeyCode: 0x49, Key: keycodes.VKeys.KEY_I},
		KEY_J:             	&VWinVKey{KeyCode: 0x4A, Key: keycodes.VKeys.KEY_J},
		KEY_K:             	&VWinVKey{KeyCode: 0x4B, Key: keycodes.VKeys.KEY_K},
		KEY_L:             	&VWinVKey{KeyCode: 0x4C, Key: keycodes.VKeys.KEY_L},
		KEY_M:             	&VWinVKey{KeyCode: 0x4D, Key: keycodes.VKeys.KEY_M},
		KEY_N:             	&VWinVKey{KeyCode: 0x4E, Key: keycodes.VKeys.KEY_N},
		KEY_O:             	&VWinVKey{KeyCode: 0x4F, Key: keycodes.VKeys.KEY_O},
		KEY_P:             	&VWinVKey{KeyCode: 0x50, Key: keycodes.VKeys.KEY_P},
		KEY_Q:             	&VWinVKey{KeyCode: 0x51, Key: keycodes.VKeys.KEY_Q},
		KEY_R:             	&VWinVKey{KeyCode: 0x52, Key: keycodes.VKeys.KEY_R},
		KEY_S:             	&VWinVKey{KeyCode: 0x53, Key: keycodes.VKeys.KEY_S},
		KEY_T:             	&VWinVKey{KeyCode: 0x54, Key: keycodes.VKeys.KEY_T},
		KEY_U:             	&VWinVKey{KeyCode: 0x55, Key: keycodes.VKeys.KEY_U},
		KEY_V:             	&VWinVKey{KeyCode: 0x56, Key: keycodes.VKeys.KEY_V},
		KEY_W:             	&VWinVKey{KeyCode: 0x57, Key: keycodes.VKeys.KEY_W},
		KEY_X:             	&VWinVKey{KeyCode: 0x58, Key: keycodes.VKeys.KEY_X},
		KEY_Y:             	&VWinVKey{KeyCode: 0x59, Key: keycodes.VKeys.KEY_Y},
		KEY_Z:             	&VWinVKey{KeyCode: 0x5A, Key: keycodes.VKeys.KEY_Z},
		KEY_LeftWin:       	&VWinVKey{KeyCode: 0x5B, Key: keycodes.VKeys.KEY_LeftWin},
		KEY_RightWin:      	&VWinVKey{KeyCode: 0x5C, Key: keycodes.VKeys.KEY_RightWin},
		KEY_NUM_0:         	&VWinVKey{KeyCode: 0x60, Key: keycodes.VKeys.KEY_NUM_0},
		KEY_NUM_1:         	&VWinVKey{KeyCode: 0x61, Key: keycodes.VKeys.KEY_NUM_1},
		KEY_NUM_2:         	&VWinVKey{KeyCode: 0x62, Key: keycodes.VKeys.KEY_NUM_2},
		KEY_NUM_3:         	&VWinVKey{KeyCode: 0x63, Key: keycodes.VKeys.KEY_NUM_3},
		KEY_NUM_4:         	&VWinVKey{KeyCode: 0x64, Key: keycodes.VKeys.KEY_NUM_4},
		KEY_NUM_5:         	&VWinVKey{KeyCode: 0x65, Key: keycodes.VKeys.KEY_NUM_5},
		KEY_NUM_6:         	&VWinVKey{KeyCode: 0x66, Key: keycodes.VKeys.KEY_NUM_6},
		KEY_NUM_7:         	&VWinVKey{KeyCode: 0x67, Key: keycodes.VKeys.KEY_NUM_7},
		KEY_NUM_8:         	&VWinVKey{KeyCode: 0x68, Key: keycodes.VKeys.KEY_NUM_8},
		KEY_NUM_9:         	&VWinVKey{KeyCode: 0x69, Key: keycodes.VKeys.KEY_NUM_9},
		KEY_NUM_Multiply:  	&VWinVKey{KeyCode: 0x6A, Key: keycodes.VKeys.KEY_NUM_Multiply},
		KEY_NUM_Add:       	&VWinVKey{KeyCode: 0x6B, Key: keycodes.VKeys.KEY_NUM_Add},
		KEY_NUM_Separator: 	&VWinVKey{KeyCode: 0x6C, Key: keycodes.VKeys.KEY_NUM_Separator},
		KEY_NUM_Subtract:  	&VWinVKey{KeyCode: 0x6D, Key: keycodes.VKeys.KEY_NUM_Subtract},
		KEY_NUM_Decimal:   	&VWinVKey{KeyCode: 0x6E, Key: keycodes.VKeys.KEY_NUM_Decimal},
		KEY_NUM_Divide:    	&VWinVKey{KeyCode: 0x6F, Key: keycodes.VKeys.KEY_NUM_Divide},
		KEY_F1:            	&VWinVKey{KeyCode: 0x70, Key: keycodes.VKeys.KEY_F1},
		KEY_F2:            	&VWinVKey{KeyCode: 0x71, Key: keycodes.VKeys.KEY_F2},
		KEY_F3:            	&VWinVKey{KeyCode: 0x72, Key: keycodes.VKeys.KEY_F3},
		KEY_F4:            	&VWinVKey{KeyCode: 0x73, Key: keycodes.VKeys.KEY_F4},
		KEY_F5:            	&VWinVKey{KeyCode: 0x74, Key: keycodes.VKeys.KEY_F5},
		KEY_F6:            	&VWinVKey{KeyCode: 0x75, Key: keycodes.VKeys.KEY_F6},
		KEY_F7:            	&VWinVKey{KeyCode: 0x76, Key: keycodes.VKeys.KEY_F7},
		KEY_F8:            	&VWinVKey{KeyCode: 0x77, Key: keycodes.VKeys.KEY_F8},
		KEY_F9:            	&VWinVKey{KeyCode: 0x78, Key: keycodes.VKeys.KEY_F9},
		KEY_F10:           	&VWinVKey{KeyCode: 0x79, Key: keycodes.VKeys.KEY_F10},
		KEY_F11:           	&VWinVKey{KeyCode: 0x7A, Key: keycodes.VKeys.KEY_F11},
		KEY_F12:           	&VWinVKey{KeyCode: 0x7B, Key: keycodes.VKeys.KEY_F12},
		KEY_F13:           	&VWinVKey{KeyCode: 0x7C, Key: keycodes.VKeys.KEY_F13},
		KEY_F14:           	&VWinVKey{KeyCode: 0x7D, Key: keycodes.VKeys.KEY_F14},
		KEY_F15:           	&VWinVKey{KeyCode: 0x7E, Key: keycodes.VKeys.KEY_F15},
		KEY_F16:           	&VWinVKey{KeyCode: 0x7F, Key: keycodes.VKeys.KEY_F16},
		KEY_F17:           	&VWinVKey{KeyCode: 0x80, Key: keycodes.VKeys.KEY_F17},
		KEY_F18:           	&VWinVKey{KeyCode: 0x81, Key: keycodes.VKeys.KEY_F18},
		KEY_F19:           	&VWinVKey{KeyCode: 0x82, Key: keycodes.VKeys.KEY_F19},
		KEY_F20:           	&VWinVKey{KeyCode: 0x83, Key: keycodes.VKeys.KEY_F20},
		KEY_F21:           	&VWinVKey{KeyCode: 0x84, Key: keycodes.VKeys.KEY_F21},
		KEY_F22:           	&VWinVKey{KeyCode: 0x85, Key: keycodes.VKeys.KEY_F22},
		KEY_F23:           	&VWinVKey{KeyCode: 0x86, Key: keycodes.VKeys.KEY_F23},
		KEY_F24:           	&VWinVKey{KeyCode: 0x87, Key: keycodes.VKeys.KEY_F24},
		KEY_LeftShift:     	&VWinVKey{KeyCode: 0xA0, Key: keycodes.VKeys.KEY_LeftShift},
		KEY_RightShift:    	&VWinVKey{KeyCode: 0xA1, Key: keycodes.VKeys.KEY_RightShift},
		KEY_LeftCtrl:      	&VWinVKey{KeyCode: 0xA2, Key: keycodes.VKeys.KEY_LeftCtrl},
		KEY_RightCtrl:     	&VWinVKey{KeyCode: 0xA3, Key: keycodes.VKeys.KEY_RightCtrl},
		KEY_LeftMenu:      	&VWinVKey{KeyCode: 0xA4, Key: keycodes.VKeys.KEY_LeftMenu},
		KEY_RightMenu:     	&VWinVKey{KeyCode: 0xA5, Key: keycodes.VKeys.KEY_RightMenu},
		KEY_OEM1:          	&VWinVKey{KeyCode: 0xBA, Key: keycodes.VKeys.KEY_OEM1},
		KEY_Plus:          	&VWinVKey{KeyCode: 0xBB, Key: keycodes.VKeys.KEY_Plus},
		KEY_Comma:         	&VWinVKey{KeyCode: 0xBC, Key: keycodes.VKeys.KEY_Comma},
		KEY_Minus:         	&VWinVKey{KeyCode: 0xBD, Key: keycodes.VKeys.KEY_Minus},
		KEY_Period:        	&VWinVKey{KeyCode: 0xBE, Key: keycodes.VKeys.KEY_Period},
		KEY_OEM2:          	&VWinVKey{KeyCode: 0xBF, Key: keycodes.VKeys.KEY_OEM2},
		KEY_OEM3:          	&VWinVKey{KeyCode: 0xC0, Key: keycodes.VKeys.KEY_OEM3},
		KEY_OEM4:          	&VWinVKey{KeyCode: 0xDB, Key: keycodes.VKeys.KEY_OEM4},
		KEY_OEM5:        	&VWinVKey{KeyCode: 0xDC, Key: keycodes.VKeys.KEY_OEM5},
		KEY_OEM6:          	&VWinVKey{KeyCode: 0xDD, Key: keycodes.VKeys.KEY_OEM6},
		KEY_OEM7:          	&VWinVKey{KeyCode: 0xDE, Key: keycodes.VKeys.KEY_OEM7},
		KEY_OEM102:        	&VWinVKey{KeyCode: 0xE2, Key: keycodes.VKeys.KEY_OEM102},
	}
}

func (e *VWinVKey) getVKey() (*keycodes.VKey) {
	return e.Key
}

//This sucks, but I couldn't figure out a way to use an enum-like structure and map to multiple values
//I was going down a rabbithole of rediculous reflection stuff to try and parse and convert the fields to a flat list
//If I'm going to define flat list of things anyway, might as well just do this...
func win32TranslateKeycode(keycode byte) (*keycodes.VKey, error) {
	switch keycode {
		case 0x08:
			return keycodes.VKeys.KEY_BackSpace, nil
		case 0x09:
			return keycodes.VKeys.KEY_Tab, nil
		case 0x0C:
			return keycodes.VKeys.KEY_Clear, nil
		case 0x0D:
			return keycodes.VKeys.KEY_Enter, nil
		case 0x10:
			return keycodes.VKeys.KEY_Shift, nil
		case 0x11:
			return keycodes.VKeys.KEY_Control, nil
		case 0x12:
			return keycodes.VKeys.KEY_Alt, nil
		case 0x13:
			return keycodes.VKeys.KEY_Pause, nil
		case 0x14:
			return keycodes.VKeys.KEY_CapsLock, nil
		case 0x20:
			return keycodes.VKeys.KEY_Space, nil
		case 0x21:
			return keycodes.VKeys.KEY_PageUp, nil
		case 0x22:
			return keycodes.VKeys.KEY_PageDown, nil
		case 0x23:
			return keycodes.VKeys.KEY_End, nil
		case 0x24:
			return keycodes.VKeys.KEY_Home, nil
		case 0x25:
			return keycodes.VKeys.KEY_LeftArrow, nil
		case 0x26:
			return keycodes.VKeys.KEY_UpArrow, nil
		case 0x27:
			return keycodes.VKeys.KEY_RightArrow, nil
		case 0x28:
			return keycodes.VKeys.KEY_DownArrow, nil
		case 0x2C:
			return keycodes.VKeys.KEY_PrintScreen, nil
		case 0x2D:
			return keycodes.VKeys.KEY_Inser, nil
		case 0x2E:
			return keycodes.VKeys.KEY_Delete, nil
		case 0x30:
			return keycodes.VKeys.KEY_Num0, nil
		case 0x31:
			return keycodes.VKeys.KEY_Num1, nil
		case 0x32:
			return keycodes.VKeys.KEY_Num2, nil
		case 0x33:
			return keycodes.VKeys.KEY_Num3, nil
		case 0x34:
			return keycodes.VKeys.KEY_Num4, nil
		case 0x35:
			return keycodes.VKeys.KEY_Num5, nil
		case 0x36:
			return keycodes.VKeys.KEY_Num6, nil
		case 0x37:
			return keycodes.VKeys.KEY_Num7, nil
		case 0x38:
			return keycodes.VKeys.KEY_Num8, nil
		case 0x39:
			return keycodes.VKeys.KEY_Num9, nil
		case 0x41:
			return keycodes.VKeys.KEY_A, nil
		case 0x42:
			return keycodes.VKeys.KEY_B, nil
		case 0x43:
			return keycodes.VKeys.KEY_C, nil
		case 0x44:
			return keycodes.VKeys.KEY_D, nil
		case 0x45:
			return keycodes.VKeys.KEY_E, nil
		case 0x46:
			return keycodes.VKeys.KEY_F, nil
		case 0x47:
			return keycodes.VKeys.KEY_G, nil
		case 0x48:
			return keycodes.VKeys.KEY_H, nil
		case 0x49:
			return keycodes.VKeys.KEY_I, nil
		case 0x4A:
			return keycodes.VKeys.KEY_J, nil
		case 0x4B:
			return keycodes.VKeys.KEY_K, nil
		case 0x4C:
			return keycodes.VKeys.KEY_L, nil
		case 0x4D:
			return keycodes.VKeys.KEY_M, nil
		case 0x4E:
			return keycodes.VKeys.KEY_N, nil
		case 0x4F:
			return keycodes.VKeys.KEY_O, nil
		case 0x50:
			return keycodes.VKeys.KEY_P, nil
		case 0x51:
			return keycodes.VKeys.KEY_Q, nil
		case 0x52:
			return keycodes.VKeys.KEY_R, nil
		case 0x53:
			return keycodes.VKeys.KEY_S, nil
		case 0x54:
			return keycodes.VKeys.KEY_T, nil
		case 0x55:
			return keycodes.VKeys.KEY_U, nil
		case 0x56:
			return keycodes.VKeys.KEY_V, nil
		case 0x57:
			return keycodes.VKeys.KEY_W, nil
		case 0x58:
			return keycodes.VKeys.KEY_X, nil
		case 0x59:
			return keycodes.VKeys.KEY_Y, nil
		case 0x5A:
			return keycodes.VKeys.KEY_Z, nil
		case 0x5B:
			return keycodes.VKeys.KEY_LeftWin, nil
		case 0x5C:
			return keycodes.VKeys.KEY_RightWin, nil
		case 0x60:
			return keycodes.VKeys.KEY_NUM_0, nil
		case 0x61:
			return keycodes.VKeys.KEY_NUM_1, nil
		case 0x62:
			return keycodes.VKeys.KEY_NUM_2, nil
		case 0x63:
			return keycodes.VKeys.KEY_NUM_3, nil
		case 0x64:
			return keycodes.VKeys.KEY_NUM_4, nil
		case 0x65:
			return keycodes.VKeys.KEY_NUM_5, nil
		case 0x66:
			return keycodes.VKeys.KEY_NUM_6, nil
		case 0x67:
			return keycodes.VKeys.KEY_NUM_7, nil
		case 0x68:
			return keycodes.VKeys.KEY_NUM_8, nil
		case 0x69:
			return keycodes.VKeys.KEY_NUM_9, nil
		case 0x6A:
			return keycodes.VKeys.KEY_NUM_Multiply, nil
		case 0x6B:
			return keycodes.VKeys.KEY_NUM_Add, nil
		case 0x6C:
			return keycodes.VKeys.KEY_NUM_Separator, nil
		case 0x6D:
			return keycodes.VKeys.KEY_NUM_Subtract, nil
		case 0x6E:
			return keycodes.VKeys.KEY_NUM_Decimal, nil
		case 0x6F:
			return keycodes.VKeys.KEY_NUM_Divide, nil
		case 0x70:
			return keycodes.VKeys.KEY_F1, nil
		case 0x71:
			return keycodes.VKeys.KEY_F2, nil
		case 0x72:
			return keycodes.VKeys.KEY_F3, nil
		case 0x73:
			return keycodes.VKeys.KEY_F4, nil
		case 0x74:
			return keycodes.VKeys.KEY_F5, nil
		case 0x75:
			return keycodes.VKeys.KEY_F6, nil
		case 0x76:
			return keycodes.VKeys.KEY_F7, nil
		case 0x77:
			return keycodes.VKeys.KEY_F8, nil
		case 0x78:
			return keycodes.VKeys.KEY_F9, nil
		case 0x79:
			return keycodes.VKeys.KEY_F10, nil
		case 0x7A:
			return keycodes.VKeys.KEY_F11, nil
		case 0x7B:
			return keycodes.VKeys.KEY_F12, nil
		case 0x7C:
			return keycodes.VKeys.KEY_F13, nil
		case 0x7D:
			return keycodes.VKeys.KEY_F14, nil
		case 0x7E:
			return keycodes.VKeys.KEY_F15, nil
		case 0x7F:
			return keycodes.VKeys.KEY_F16, nil
		case 0x80:
			return keycodes.VKeys.KEY_F17, nil
		case 0x81:
			return keycodes.VKeys.KEY_F18, nil
		case 0x82:
			return keycodes.VKeys.KEY_F19, nil
		case 0x83:
			return keycodes.VKeys.KEY_F20, nil
		case 0x84:
			return keycodes.VKeys.KEY_F21, nil
		case 0x85:
			return keycodes.VKeys.KEY_F22, nil
		case 0x86:
			return keycodes.VKeys.KEY_F23, nil
		case 0x87:
			return keycodes.VKeys.KEY_F24, nil
		case 0xA0:
			return keycodes.VKeys.KEY_LeftShift, nil
		case 0xA1:
			return keycodes.VKeys.KEY_RightShift, nil
		case 0xA2:
			return keycodes.VKeys.KEY_LeftCtrl, nil
		case 0xA3:
			return keycodes.VKeys.KEY_RightCtrl, nil
		case 0xA4:
			return keycodes.VKeys.KEY_LeftMenu, nil
		case 0xA5:
			return keycodes.VKeys.KEY_RightMenu, nil
		case 0xBA:
			return keycodes.VKeys.KEY_OEM1, nil
		case 0xBB:
			return keycodes.VKeys.KEY_Plus, nil
		case 0xBC:
			return keycodes.VKeys.KEY_Comma, nil
		case 0xBD:
			return keycodes.VKeys.KEY_Minus, nil
		case 0xBE:
			return keycodes.VKeys.KEY_Period, nil
		case 0xBF:
			return keycodes.VKeys.KEY_OEM2, nil
		case 0xC0:
			return keycodes.VKeys.KEY_OEM3, nil
		case 0xDB:
			return keycodes.VKeys.KEY_OEM4, nil
		case 0xDC:
			return keycodes.VKeys.KEY_OEM5, nil
		case 0xDD:
			return keycodes.VKeys.KEY_OEM6, nil
		case 0xDE:
			return keycodes.VKeys.KEY_OEM7, nil
		case 0xE2:
			return keycodes.VKeys.KEY_OEM102, nil
		default:
			return nil, fmt.Errorf("No value for code")
	}
}