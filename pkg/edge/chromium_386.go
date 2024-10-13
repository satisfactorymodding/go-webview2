//go:build windows
// +build windows

package edge

import (
	"unsafe"

	"github.com/wailsapp/go-webview2/internal/w32"
	"golang.org/x/sys/windows"
)

func (e *Chromium) SetSize(bounds w32.Rect) {
	if e.controller == nil {
		return
	}

	hr, _, _ := e.controller.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(e.controller)),
		uintptr(bounds.Left),
		uintptr(bounds.Top),
		uintptr(bounds.Right),
		uintptr(bounds.Bottom),
	)
	if windows.Handle(hr) != windows.S_OK {
		e.globalErrorCallback(windows.Errno(hr))
	}
}
