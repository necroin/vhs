//go:build !windows

package remote_desktop_input

func KeyboardDown(keyCode uint16) {}

func KeyboardUp(keyCode uint16) {}

func MouseMove(x int32, y int32) {}

func MouseLeftDown() {}

func MouseLeftUp() {}

func MouseRightClick() {}

func MouseWheel(x int32, y int32, delta int32) {}
