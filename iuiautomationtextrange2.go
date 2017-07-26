package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextRange2 struct {
	ole.IUIAutomationTextRange
}

type IUIAutomationTextRange2Vtbl struct {
	ole.IUIAutomationTextRangeVtbl
	ShowContextMenu uintptr
}

func (v *IUIAutomationTextRange2) VTable() *IUIAutomationTextRange2Vtbl {
	return (*IUIAutomationTextRange2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextRange2) ShowContextMenu() error {
	return ole.NewError(ole.E_NOTIMPL)
}
