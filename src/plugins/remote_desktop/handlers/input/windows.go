package remote_desktop_input

import (
	"unsafe"
	"vhs/src/utils/winapi"
)

func KeyboardDown(keyCode uint16) {
	input := winapi.KEYBD_INPUT{}
	input.Type = winapi.INPUT_KEYBOARD
	input.Ki = winapi.KEYBDINPUT{
		WVk: keyCode,
	}

	winapi.SendInput(1, unsafe.Pointer(&input), unsafe.Sizeof(input))
}

func KeyboardUp(keyCode uint16) {
	input := winapi.KEYBD_INPUT{}
	input.Type = winapi.INPUT_KEYBOARD
	input.Ki = winapi.KEYBDINPUT{
		WVk:     keyCode,
		DwFlags: winapi.KEYEVENTF_KEYUP,
	}

	winapi.SendInput(1, unsafe.Pointer(&input), unsafe.Sizeof(input))
}

func MouseMove(x int32, y int32) {
	cx_screen, _ := winapi.GetSystemMetrics(winapi.SM_CXSCREEN)
	cy_screen, _ := winapi.GetSystemMetrics(winapi.SM_CYSCREEN)
	real_x := 65535 * x / cx_screen
	real_y := 65535 * y / cy_screen

	mouseInput := winapi.MOUSE_INPUT{}
	mouseInput.Type = winapi.INPUT_MOUSE
	mouseInput.Mi = winapi.MOUSEINPUT{
		Dx: int32(real_x),
		Dy: int32(real_y),
	}
	mouseInput.Mi.MouseData = 0
	mouseInput.Mi.DwExtraInfo = 0
	mouseInput.Mi.Time = 0

	mouseInput.Mi.DwFlags = winapi.MOUSEEVENTF_ABSOLUTE | winapi.MOUSEEVENTF_MOVE
	winapi.SendInput(1, unsafe.Pointer(&mouseInput), unsafe.Sizeof(mouseInput))
}

func MouseLeftDown() {
	mouseInput := winapi.MOUSE_INPUT{}
	mouseInput.Type = winapi.INPUT_MOUSE
	mouseInput.Mi = winapi.MOUSEINPUT{
		Dx:          0,
		Dy:          0,
		MouseData:   0,
		DwExtraInfo: 0,
		Time:        0,
	}

	mouseInput.Mi.DwFlags = winapi.MOUSEEVENTF_LEFTDOWN
	winapi.SendInput(1, unsafe.Pointer(&mouseInput), unsafe.Sizeof(mouseInput))
}

func MouseLeftUp() {
	mouseInput := winapi.MOUSE_INPUT{}
	mouseInput.Type = winapi.INPUT_MOUSE
	mouseInput.Mi = winapi.MOUSEINPUT{
		Dx:          0,
		Dy:          0,
		MouseData:   0,
		DwExtraInfo: 0,
		Time:        0,
	}
	mouseInput.Mi.DwFlags = winapi.MOUSEEVENTF_LEFTUP
	winapi.SendInput(1, unsafe.Pointer(&mouseInput), unsafe.Sizeof(mouseInput))
}

func MouseRightClick() {
	mouseInput := winapi.MOUSE_INPUT{}
	mouseInput.Type = winapi.INPUT_MOUSE
	mouseInput.Mi = winapi.MOUSEINPUT{
		Dx:          0,
		Dy:          0,
		MouseData:   0,
		DwExtraInfo: 0,
		Time:        0,
	}

	mouseInput.Mi.DwFlags = winapi.MOUSEEVENTF_RIGHTDOWN
	winapi.SendInput(1, unsafe.Pointer(&mouseInput), unsafe.Sizeof(mouseInput))

	mouseInput.Mi.DwFlags = winapi.MOUSEEVENTF_RIGHTUP
	winapi.SendInput(1, unsafe.Pointer(&mouseInput), unsafe.Sizeof(mouseInput))
}

func MouseWheel(x int32, y int32, delta int32) {
	mouseInput := winapi.MOUSE_INPUT{}
	mouseInput.Type = winapi.INPUT_MOUSE
	mouseInput.Mi = winapi.MOUSEINPUT{
		Dx:          x,
		Dy:          y,
		MouseData:   uint32(delta),
		DwExtraInfo: 0,
		Time:        0,
	}

	mouseInput.Mi.DwFlags = winapi.MOUSEEVENTF_WHEEL
	winapi.SendInput(1, unsafe.Pointer(&mouseInput), unsafe.Sizeof(mouseInput))
}
