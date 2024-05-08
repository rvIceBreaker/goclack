package win32

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rvIceBreaker/goclack/eventhandler"
)

//Windows hooking referenced from https://gist.github.com/obonyojimmy/52d836a1b31e2fc914d19a81bd2e0a1b

/*const (
	PM_NOREMOVE = 0x000
	PM_REMOVE = 0x001
	PM_NOYIELD = 0x002
	NULL = 0
)*/

type (
	tDWORD uint32
	tWPARAM uintptr
	tLPARAM uintptr
	tLRESULT uintptr
	tHANDLE uintptr
	tHINSTANCE tHANDLE
	tHHOOK tHANDLE
	tHWND tHANDLE
)

type tHOOKPROC func(int, tWPARAM, tLPARAM) tLRESULT

type tKBDLLHOOKSTRUCT struct {
	VkCode tDWORD
	ScanCode tDWORD
	Flags tDWORD
	Time tDWORD
	DwExtraInfo uintptr
}

type tPOINT struct {
	X, Y int32
}

type tMSG struct {
	Hwnd tHWND
	Message uint32
	WParam uintptr
	LParam uintptr
	Time uint32
	Pt tPOINT
}

type Win32EventHandler struct {
	user32 *syscall.DLL
	procSetWindowsHookEx *syscall.Proc
	procLowLevelKeyboard *syscall.Proc
	procCallNextHookEx *syscall.Proc
	procUnhookWindowsHookEx *syscall.Proc
	procGetMessage *syscall.Proc
	procTranslateMessage *syscall.Proc
	procDispatchMessage *syscall.Proc

	hooks map[int]win32EventHook

	KeyEvents chan *eventhandler.KeyEvent
	eventhandler.EventHandler
}

type win32EventHook struct {
	HookType *VWinHook
	Handle tHHOOK
}

func NewWin32EventHandler() (*Win32EventHandler) {
	e := Win32EventHandler{}
	e.init()
	return &e
}

func (e *Win32EventHandler) init() {
	e.user32 = syscall.MustLoadDLL("user32")
	e.procSetWindowsHookEx = e.user32.MustFindProc("SetWindowsHookExW")
	e.procCallNextHookEx = e.user32.MustFindProc("CallNextHookEx")
	e.procUnhookWindowsHookEx = e.user32.MustFindProc("UnhookWindowsHookEx")
	e.procGetMessage = e.user32.MustFindProc("GetMessageW")
	e.procTranslateMessage = e.user32.MustFindProc("TranslateMessage")
	e.procDispatchMessage = e.user32.MustFindProc("DispatchMessageW")

	e.KeyEvents = make(chan *eventhandler.KeyEvent)

	e.hooks = make(map[int]win32EventHook)
}

func (e *Win32EventHandler) GetHookForType(hookType *VWinHook) (tHHOOK, error) {
	hook,ok := e.hooks[hookType.ID]

	if (ok) {
		return hook.Handle, nil
	} else {
		return 1, fmt.Errorf("Hook does not exist for %s", hookType.Name)
	}
}

func (e *Win32EventHandler) iSetWindowsHookEx(idHook *VWinHook, lpfn tHOOKPROC, hMod tHINSTANCE, dwThreadId tDWORD) tHHOOK {
	ret, _, _ := e.procSetWindowsHookEx.Call(
		uintptr(idHook.ID),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)

	return tHHOOK(ret)
}

func (e *Win32EventHandler) iCallNextHookEx(hhk tHHOOK, nCode int, wParam tWPARAM, lParam tLPARAM) tLRESULT {
	ret, _, _ := e.procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)

	return tLRESULT(ret)
}

func (e *Win32EventHandler) iUnhookWindowsHookEx(hhk tHHOOK) bool {
	ret, _, _ := e.procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)

	return ret != 0
}

func (e *Win32EventHandler) iGetMessage(msg *tMSG, hwnd tHWND, msgFilterMin uint32, msgFilterMax uint32) int {
	ret, _, _ := e.procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))

	return int(ret)
}

func (e *Win32EventHandler) iTranslateMessage(msg *tMSG) bool {
	ret, _, _ := e.procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

func (e *Win32EventHandler) iDispatchMessage(msg *tMSG) uintptr {
	ret, _, _ := e.procDispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))
	return ret
}

func (e *Win32EventHandler) RegisterKeyboardHook() error {
	_,err := e.GetHookForType(WinHook.WH_KEYBOARD_LL)
	if err == nil {
		return fmt.Errorf("Hook already exists for %s", WinHook.WH_KEYBOARD_LL.Name)
	}

	hnd := e.iSetWindowsHookEx(WinHook.WH_KEYBOARD_LL, (tHOOKPROC)(e.handleKeyboardHook), 0, 0)

	hook := win32EventHook{
		HookType: WinHook.WH_KEYBOARD_LL,
		Handle: hnd,
	}
	e.hooks[hook.HookType.ID] = hook
	return nil
}

func (e *Win32EventHandler) UnregisterKeyboardHook() error {
	hook,err := e.GetHookForType(WinHook.WH_KEYBOARD_LL)
	if err != nil {
		return err
	}
	e.iUnhookWindowsHookEx(hook)
	return nil
}

func (e *Win32EventHandler) handleKeyboardHook(nCode int, wparam tWPARAM, lparam tLPARAM) tLRESULT {
	if nCode == 0 {
		switch wparam {
			case tWPARAM(WinMsg.WM_KEYDOWN.ID):
				kbdstruct := (*tKBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
				code := byte(kbdstruct.VkCode)
				vkey,err := win32TranslateKeycode(code)
				if err == nil {
					e.KeyEvents <- &eventhandler.KeyEvent{
						Key: vkey,
						Type: eventhandler.KEYEVENT_DOWN,
					}
				}
			case tWPARAM(WinMsg.WM_KEYUP.ID):
				kbdstruct := (*tKBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
				code := byte(kbdstruct.VkCode)
				vkey,err := win32TranslateKeycode(code)
				if err == nil {
					e.KeyEvents <- &eventhandler.KeyEvent{
						Key: vkey,
						Type: eventhandler.KEYEVENT_UP,
					}
				}
		}
	}
	hook,_ := e.GetHookForType(WinHook.WH_KEYBOARD_LL) //We shouldn't be here unless the hook exists

	return e.iCallNextHookEx(hook, nCode, wparam, lparam)
}

func (e *Win32EventHandler) GetKeyEvents() (<-chan *eventhandler.KeyEvent) {
	return e.KeyEvents
}

func (e *Win32EventHandler) ListenForEvents() {
	go e.implListenForEvents()
}

func (e *Win32EventHandler) implListenForEvents() error {
	var msg tMSG
	for e.iGetMessage(&msg, 0,0,0) != 0	{

	}
	return nil
}

func (e *Win32EventHandler) Teardown() {
	for _,hook := range e.hooks {
		e.iUnhookWindowsHookEx(hook.Handle)
	}
}