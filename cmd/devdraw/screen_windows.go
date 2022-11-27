//go:build windows
// +build windows

//
package main

import (
	"image"
	"syscall"
	"unsafe"
)

var (
	user32              = syscall.MustLoadDLL("user32.dll")
	setCursorPos        = user32.MustFindProc("SetCursorPos")
	getWindowRect       = user32.MustFindProc("GetWindowRect")
	getForegroundWindow = user32.MustFindProc("GetForegroundWindow")
)

type rect struct {
	left   int32
	top    int32
	right  int32
	bottom int32
}

func setMouse(p image.Point) {
	var r rect = rect{0, 0, 0, 0}
	hWnd, _, _ := getForegroundWindow.Call()
	if hWnd == 0 {
		return
	}
	ret, _, _ := getWindowRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&r)))
	if ret != 0 {
		_, _, _ = setCursorPos.Call(uintptr(p.X+int(r.left)+4), uintptr(p.Y+int(r.top)+32))
	}
}
