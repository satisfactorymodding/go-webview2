//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Controller4Vtbl struct {
	IUnknownVtbl
	GetAllowExternalDrop ComProc
	PutAllowExternalDrop ComProc
}

type ICoreWebView2Controller4 struct {
	Vtbl *ICoreWebView2Controller4Vtbl
}

func (i *ICoreWebView2Controller4) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Controller4() *ICoreWebView2Controller4 {
	var result *ICoreWebView2Controller4

	iidICoreWebView2Controller4 := NewGUID("{97d418d5-a426-4e49-a151-e1a10f327d9e}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Controller4)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Controller4) GetAllowExternalDrop() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetAllowExternalDrop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2Controller4) PutAllowExternalDrop(value bool) error {

	hr, _, err := i.Vtbl.PutAllowExternalDrop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}