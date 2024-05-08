package keycodes

var VKeys = newVKeys()

type EVKey struct {
	KEY_BackSpace     *VKey //		=0x08, //Backspace key
	KEY_Tab           *VKey //				=0x09, //Tab key
	KEY_Clear         *VKey //			=0x0C, //Clear key
	KEY_Enter         *VKey //			=0x0D, //Enter or Return key
	KEY_Shift         *VKey //			=0x10, //Shift key
	KEY_Control       *VKey //			=0x11, //Ctrl key
	KEY_Alt           *VKey //				=0x12, //Alt key
	KEY_Pause         *VKey //			=0x13, //Pause key
	KEY_CapsLock      *VKey //		=0x14, //Caps lock key
	KEY_Space         *VKey //			=0x20, //Space bar
	KEY_PageUp        *VKey //			=0x21, //Page up key
	KEY_PageDown      *VKey //		=0x22, //Page down key
	KEY_End           *VKey //				=0x23, //End key
	KEY_Home          *VKey //			=0x24, //Home key
	KEY_LeftArrow     *VKey //		=0x25, //Left arrow key
	KEY_UpArrow       *VKey //			=0x26, //Up arrow key
	KEY_RightArrow    *VKey //		=0x27, //Right arrow key
	KEY_DownArrow     *VKey //		=0x28, //Down arrow key
	KEY_PrintScreen   *VKey //		=0x2C, //Print screen key
	KEY_Inser         *VKey //			=0x2D, //Insert key
	KEY_Delete        *VKey //			=0x2E, //Delete key
	KEY_Num0          *VKey //			=0x30, //Top row 0 key (Matches '0')
	KEY_Num1          *VKey //			=0x31, //Top row 1 key (Matches '1')
	KEY_Num2          *VKey //			=0x32, //Top row 2 key (Matches '2')
	KEY_Num3          *VKey //			=0x33, //Top row 3 key (Matches '3')
	KEY_Num4          *VKey //			=0x34, //Top row 4 key (Matches '4')
	KEY_Num5          *VKey //			=0x35, //Top row 5 key (Matches '5')
	KEY_Num6          *VKey //			=0x36, //Top row 6 key (Matches '6')
	KEY_Num7          *VKey //			=0x37, //Top row 7 key (Matches '7')
	KEY_Num8          *VKey //			=0x38, //Top row 8 key (Matches '8')
	KEY_Num9          *VKey //			=0x39, //Top row 9 key (Matches '9')
	KEY_A             *VKey //				=0x41, //A key (Matches 'A')
	KEY_B             *VKey //				=0x42, //B key (Matches 'B')
	KEY_C             *VKey //				=0x43, //C key (Matches 'C')
	KEY_D             *VKey //				=0x44, //D key (Matches 'D')
	KEY_E             *VKey //				=0x45, //E key (Matches 'E')
	KEY_F             *VKey //				=0x46, //F key (Matches 'F')
	KEY_G             *VKey //				=0x47, //G key (Matches 'G')
	KEY_H             *VKey //				=0x48, //H key (Matches 'H')
	KEY_I             *VKey //				=0x49, //I key (Matches 'I')
	KEY_J             *VKey //				=0x4A, //J key (Matches 'J')
	KEY_K             *VKey //				=0x4B, //K key (Matches 'K')
	KEY_L             *VKey //				=0x4C, //L key (Matches 'L')
	KEY_M             *VKey //				=0x4D, //M key (Matches 'M')
	KEY_N             *VKey //				=0x4E, //N key (Matches 'N')
	KEY_O             *VKey //				=0x4F, //O key (Matches 'O')
	KEY_P             *VKey //				=0x50, //P key (Matches 'P')
	KEY_Q             *VKey //				=0x51, //Q key (Matches 'Q')
	KEY_R             *VKey //				=0x52, //R key (Matches 'R')
	KEY_S             *VKey //				=0x53, //S key (Matches 'S')
	KEY_T             *VKey //				=0x54, //T key (Matches 'T')
	KEY_U             *VKey //				=0x55, //U key (Matches 'U')
	KEY_V             *VKey //				=0x56, //V key (Matches 'V')
	KEY_W             *VKey //				=0x57, //W key (Matches 'W')
	KEY_X             *VKey //				=0x58, //X key (Matches 'X')
	KEY_Y             *VKey //				=0x59, //Y key (Matches 'Y')
	KEY_Z             *VKey //				=0x5A, //Z key (Matches 'Z')
	KEY_LeftWin       *VKey //			=0x5B, //Left windows key
	KEY_RightWin      *VKey //		=0x5C, //Right windows key
	KEY_NUM_0         *VKey //			=0x60, //Numpad 0
	KEY_NUM_1         *VKey //			=0x61, //Numpad 1
	KEY_NUM_2         *VKey //			=0x62, //Numpad 2
	KEY_NUM_3         *VKey //			=0x63, //Numpad 3
	KEY_NUM_4         *VKey //			=0x64, //Numpad 4
	KEY_NUM_5         *VKey //			=0x65, //Numpad 5
	KEY_NUM_6         *VKey //			=0x66, //Numpad 6
	KEY_NUM_7         *VKey //			=0x67, //Numpad 7
	KEY_NUM_8         *VKey //			=0x68, //Numpad 8
	KEY_NUM_9         *VKey //			=0x69, //Numpad 9
	KEY_NUM_Multiply  *VKey //		=0x6A, //Multiply key
	KEY_NUM_Add       *VKey //				=0x6B, //Add key
	KEY_NUM_Separator *VKey //		=0x6C, //Separator key
	KEY_NUM_Subtract  *VKey //		=0x6D, //Subtract key
	KEY_NUM_Decimal   *VKey //			=0x6E, //Decimal key
	KEY_NUM_Divide    *VKey //			=0x6F, //Divide key
	KEY_F1            *VKey //				=0x70, //F1
	KEY_F2            *VKey //				=0x71, //F2
	KEY_F3            *VKey //				=0x72, //F3
	KEY_F4            *VKey //				=0x73, //F4
	KEY_F5            *VKey //				=0x74, //F5
	KEY_F6            *VKey //				=0x75, //F6
	KEY_F7            *VKey //				=0x76, //F7
	KEY_F8            *VKey //				=0x77, //F8
	KEY_F9            *VKey //				=0x78, //F9
	KEY_F10           *VKey //				=0x79, //F10
	KEY_F11           *VKey //				=0x7A, //F11
	KEY_F12           *VKey //				=0x7B, //F12
	KEY_F13           *VKey //				=0x7C, //F13
	KEY_F14           *VKey //				=0x7D, //F14
	KEY_F15           *VKey //				=0x7E, //F15
	KEY_F16           *VKey //				=0x7F, //F16
	KEY_F17           *VKey //				=0x80, //F17
	KEY_F18           *VKey //				=0x81, //F18
	KEY_F19           *VKey //				=0x82, //F19
	KEY_F20           *VKey //				=0x83, //F20
	KEY_F21           *VKey //				=0x84, //F21
	KEY_F22           *VKey //				=0x85, //F22
	KEY_F23           *VKey //				=0x86, //F23
	KEY_F24           *VKey //				=0x87, //F24
	KEY_LeftShift     *VKey //		=0xA0, //Left shift key
	KEY_RightShift    *VKey //		=0xA1, //Right shift key
	KEY_LeftCtrl      *VKey //		=0xA2, //Left control key
	KEY_RightCtrl     *VKey //		=0xA3, //Right control key
	KEY_LeftMenu      *VKey //		=0xA4, //Left menu key
	KEY_RightMenu     *VKey //		=0xA5, //Right menu
	KEY_OEM1          *VKey //			=0xBA, //;: key for US or misc keys for others
	KEY_Plus          *VKey //			=0xBB, //Plus key
	KEY_Comma         *VKey //			=0xBC, //Comma key
	KEY_Minus         *VKey //			=0xBD, //Minus key
	KEY_Period        *VKey //			=0xBE, //Period key
	KEY_OEM2          *VKey //			=0xBF, //? for US or misc keys for others
	KEY_OEM3          *VKey //			=0xC0, //~ for US or misc keys for others
	KEY_OEM4          *VKey //			=0xDB, //[ for US or misc keys for others
	KEY_OEM5          *VKey //			=0xDC, //\ for US or misc keys for others
	KEY_OEM6          *VKey //			=0xDD, //] for US or misc keys for others
	KEY_OEM7          *VKey //			=0xDE, //' for US or misc keys for others
	KEY_OEM102        *VKey //			=0xE2, //"<>" or "\|" on RT 102-key keyboard
}

type VKey struct {
	RawName   string
	literalName string
	friendlyName string
}

func (vk *VKey) GetSimpleName() string {
	return vk.friendlyName
}
func (vk *VKey) GetLiteralName() string {
	return vk.literalName
}

func newVKeys() *EVKey {
	return &EVKey{
		KEY_BackSpace:     &VKey{literalName: "KEY_BackSpace", friendlyName: "BackSpace"},
		KEY_Tab:           &VKey{literalName: "KEY_Tab", friendlyName: "Tab"},
		KEY_Clear:         &VKey{literalName: "KEY_Clear", friendlyName: "Clear"},
		KEY_Enter:         &VKey{literalName: "KEY_Enter", friendlyName: "Enter"},
		KEY_Shift:         &VKey{literalName: "KEY_Shift", friendlyName: "Shift"},
		KEY_Control:       &VKey{literalName: "KEY_Control", friendlyName: "Control"},
		KEY_Alt:           &VKey{literalName: "KEY_Alt", friendlyName: "Alt"},
		KEY_Pause:         &VKey{literalName: "KEY_Pause", friendlyName: "Pause"},
		KEY_CapsLock:      &VKey{literalName: "KEY_CapsLock", friendlyName: "CapsLock"},
		KEY_Space:         &VKey{literalName: "KEY_Space", friendlyName: "Space"},
		KEY_PageUp:        &VKey{literalName: "KEY_PageUp", friendlyName: "PageUp"},
		KEY_PageDown:      &VKey{literalName: "KEY_PageDown", friendlyName: "PageDown"},
		KEY_End:           &VKey{literalName: "KEY_End", friendlyName: "End"},
		KEY_Home:          &VKey{literalName: "KEY_Home", friendlyName: "Home"},
		KEY_LeftArrow:     &VKey{literalName: "KEY_LeftArrow", friendlyName: "LeftArrow"},
		KEY_UpArrow:       &VKey{literalName: "KEY_UpArrow", friendlyName: "UpArrow"},
		KEY_RightArrow:    &VKey{literalName: "KEY_RightArrow", friendlyName: "RightArrow"},
		KEY_DownArrow:     &VKey{literalName: "KEY_DownArrow", friendlyName: "DownArrow"},
		KEY_PrintScreen:   &VKey{literalName: "KEY_PrintScreen", friendlyName: "PrintScreen"},
		KEY_Inser:         &VKey{literalName: "KEY_Inser", friendlyName: "Inser"},
		KEY_Delete:        &VKey{literalName: "KEY_Delete", friendlyName: "Delete"},
		KEY_Num0:          &VKey{literalName: "KEY_Num0", friendlyName: "Num0"},
		KEY_Num1:          &VKey{literalName: "KEY_Num1", friendlyName: "Num1"},
		KEY_Num2:          &VKey{literalName: "KEY_Num2", friendlyName: "Num2"},
		KEY_Num3:          &VKey{literalName: "KEY_Num3", friendlyName: "Num3"},
		KEY_Num4:          &VKey{literalName: "KEY_Num4", friendlyName: "Num4"},
		KEY_Num5:          &VKey{literalName: "KEY_Num5", friendlyName: "Num5"},
		KEY_Num6:          &VKey{literalName: "KEY_Num6", friendlyName: "Num6"},
		KEY_Num7:          &VKey{literalName: "KEY_Num7", friendlyName: "Num7"},
		KEY_Num8:          &VKey{literalName: "KEY_Num8", friendlyName: "Num8"},
		KEY_Num9:          &VKey{literalName: "KEY_Num9", friendlyName: "Num9"},
		KEY_A:             &VKey{literalName: "KEY_A", friendlyName: "A"},
		KEY_B:             &VKey{literalName: "KEY_B", friendlyName: "B"},
		KEY_C:             &VKey{literalName: "KEY_C", friendlyName: "C"},
		KEY_D:             &VKey{literalName: "KEY_D", friendlyName: "D"},
		KEY_E:             &VKey{literalName: "KEY_E", friendlyName: "E"},
		KEY_F:             &VKey{literalName: "KEY_F", friendlyName: "F"},
		KEY_G:             &VKey{literalName: "KEY_G", friendlyName: "G"},
		KEY_H:             &VKey{literalName: "KEY_H", friendlyName: "H"},
		KEY_I:             &VKey{literalName: "KEY_I", friendlyName: "I"},
		KEY_J:             &VKey{literalName: "KEY_J", friendlyName: "J"},
		KEY_K:             &VKey{literalName: "KEY_K", friendlyName: "K"},
		KEY_L:             &VKey{literalName: "KEY_L", friendlyName: "L"},
		KEY_M:             &VKey{literalName: "KEY_M", friendlyName: "M"},
		KEY_N:             &VKey{literalName: "KEY_N", friendlyName: "N"},
		KEY_O:             &VKey{literalName: "KEY_O", friendlyName: "O"},
		KEY_P:             &VKey{literalName: "KEY_P", friendlyName: "P"},
		KEY_Q:             &VKey{literalName: "KEY_Q", friendlyName: "Q"},
		KEY_R:             &VKey{literalName: "KEY_R", friendlyName: "R"},
		KEY_S:             &VKey{literalName: "KEY_S", friendlyName: "S"},
		KEY_T:             &VKey{literalName: "KEY_T", friendlyName: "T"},
		KEY_U:             &VKey{literalName: "KEY_U", friendlyName: "U"},
		KEY_V:             &VKey{literalName: "KEY_V", friendlyName: "V"},
		KEY_W:             &VKey{literalName: "KEY_W", friendlyName: "W"},
		KEY_X:             &VKey{literalName: "KEY_X", friendlyName: "X"},
		KEY_Y:             &VKey{literalName: "KEY_Y", friendlyName: "Y"},
		KEY_Z:             &VKey{literalName: "KEY_Z", friendlyName: "Z"},
		KEY_LeftWin:       &VKey{literalName: "KEY_LeftWin", friendlyName: "Win"},
		KEY_RightWin:      &VKey{literalName: "KEY_RightWin", friendlyName: "Win"},
		KEY_NUM_0:         &VKey{literalName: "KEY_NUM_0", friendlyName: "NUM_0"},
		KEY_NUM_1:         &VKey{literalName: "KEY_NUM_1", friendlyName: "NUM_1"},
		KEY_NUM_2:         &VKey{literalName: "KEY_NUM_2", friendlyName: "NUM_2"},
		KEY_NUM_3:         &VKey{literalName: "KEY_NUM_3", friendlyName: "NUM_3"},
		KEY_NUM_4:         &VKey{literalName: "KEY_NUM_4", friendlyName: "NUM_4"},
		KEY_NUM_5:         &VKey{literalName: "KEY_NUM_5", friendlyName: "NUM_5"},
		KEY_NUM_6:         &VKey{literalName: "KEY_NUM_6", friendlyName: "NUM_6"},
		KEY_NUM_7:         &VKey{literalName: "KEY_NUM_7", friendlyName: "NUM_7"},
		KEY_NUM_8:         &VKey{literalName: "KEY_NUM_8", friendlyName: "NUM_8"},
		KEY_NUM_9:         &VKey{literalName: "KEY_NUM_9", friendlyName: "NUM_9"},
		KEY_NUM_Multiply:  &VKey{literalName: "KEY_NUM_Multiply", friendlyName: "NUM_Multiply"},
		KEY_NUM_Add:       &VKey{literalName: "KEY_NUM_Add", friendlyName: "NUM_Add"},
		KEY_NUM_Separator: &VKey{literalName: "KEY_NUM_Separator", friendlyName: "NUM_Separator"},
		KEY_NUM_Subtract:  &VKey{literalName: "KEY_NUM_Subtract", friendlyName: "NUM_Subtract"},
		KEY_NUM_Decimal:   &VKey{literalName: "KEY_NUM_Decimal", friendlyName: "NUM_Decimal"},
		KEY_NUM_Divide:    &VKey{literalName: "KEY_NUM_Divide", friendlyName: "NUM_Divide"},
		KEY_F1:            &VKey{literalName: "KEY_F1", friendlyName: "F1"},
		KEY_F2:            &VKey{literalName: "KEY_F2", friendlyName: "F2"},
		KEY_F3:            &VKey{literalName: "KEY_F3", friendlyName: "F3"},
		KEY_F4:            &VKey{literalName: "KEY_F4", friendlyName: "F4"},
		KEY_F5:            &VKey{literalName: "KEY_F5", friendlyName: "F5"},
		KEY_F6:            &VKey{literalName: "KEY_F6", friendlyName: "F6"},
		KEY_F7:            &VKey{literalName: "KEY_F7", friendlyName: "F7"},
		KEY_F8:            &VKey{literalName: "KEY_F8", friendlyName: "F8"},
		KEY_F9:            &VKey{literalName: "KEY_F9", friendlyName: "F9"},
		KEY_F10:           &VKey{literalName: "KEY_F10", friendlyName: "F10"},
		KEY_F11:           &VKey{literalName: "KEY_F11", friendlyName: "F11"},
		KEY_F12:           &VKey{literalName: "KEY_F12", friendlyName: "F12"},
		KEY_F13:           &VKey{literalName: "KEY_F13", friendlyName: "F13"},
		KEY_F14:           &VKey{literalName: "KEY_F14", friendlyName: "F14"},
		KEY_F15:           &VKey{literalName: "KEY_F15", friendlyName: "F15"},
		KEY_F16:           &VKey{literalName: "KEY_F16", friendlyName: "F16"},
		KEY_F17:           &VKey{literalName: "KEY_F17", friendlyName: "F17"},
		KEY_F18:           &VKey{literalName: "KEY_F18", friendlyName: "F18"},
		KEY_F19:           &VKey{literalName: "KEY_F19", friendlyName: "F19"},
		KEY_F20:           &VKey{literalName: "KEY_F20", friendlyName: "F20"},
		KEY_F21:           &VKey{literalName: "KEY_F21", friendlyName: "F21"},
		KEY_F22:           &VKey{literalName: "KEY_F22", friendlyName: "F22"},
		KEY_F23:           &VKey{literalName: "KEY_F23", friendlyName: "F23"},
		KEY_F24:           &VKey{literalName: "KEY_F24", friendlyName: "F24"},
		KEY_LeftShift:     &VKey{literalName: "KEY_LeftShift", friendlyName: "Shift"},
		KEY_RightShift:    &VKey{literalName: "KEY_RightShift", friendlyName: "Shift"},
		KEY_LeftCtrl:      &VKey{literalName: "KEY_LeftCtrl", friendlyName: "Ctrl"},
		KEY_RightCtrl:     &VKey{literalName: "KEY_RightCtrl", friendlyName: "Ctrl"},
		KEY_LeftMenu:      &VKey{literalName: "KEY_LeftMenu", friendlyName: "Menu"},
		KEY_RightMenu:     &VKey{literalName: "KEY_RightMenu", friendlyName: "Menu"},
		KEY_OEM1:          &VKey{literalName: "KEY_OEM1", friendlyName: "OEM1"},
		KEY_Plus:          &VKey{literalName: "KEY_Plus", friendlyName: "Plus"},
		KEY_Comma:         &VKey{literalName: "KEY_Comma", friendlyName: "Comma"},
		KEY_Minus:         &VKey{literalName: "KEY_Minus", friendlyName: "Minus"},
		KEY_Period:        &VKey{literalName: "KEY_Period", friendlyName: "Period"},
		KEY_OEM2:          &VKey{literalName: "KEY_OEM2", friendlyName: "OEM2"},
		KEY_OEM3:          &VKey{literalName: "KEY_OEM3", friendlyName: "OEM3"},
		KEY_OEM4:          &VKey{literalName: "KEY_OEM4", friendlyName: "OEM4"},
		KEY_OEM5:          &VKey{literalName: "KEY_OEM5", friendlyName: "OEM5"},
		KEY_OEM6:          &VKey{literalName: "KEY_OEM6", friendlyName: "OEM6"},
		KEY_OEM7:          &VKey{literalName: "KEY_OEM7", friendlyName: "OEM7"},
		KEY_OEM102:        &VKey{literalName: "KEY_OEM102", friendlyName: "OEM102"},
	}
}