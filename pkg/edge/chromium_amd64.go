//go:build windows
// +build windows

package edge

import (
	"github.com/wailsapp/go-webview2/internal/w32"
	"golang.org/x/sys/windows"
	"unsafe"
)

func (e *Chromium) SetSize(bounds w32.Rect) {
	if e.controller == nil {
		return
	}

	hr, _, _ := e.controller.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(e.controller)),
		uintptr(unsafe.Pointer(&bounds)),
	)
	if windows.Handle(hr) != windows.S_OK {
		e.globalErrorCallback(windows.Errno(hr))
	}
}
