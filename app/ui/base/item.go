package base

import (
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func Item(labelStr string, placeHolder string) (*fyne.Container, *widget.Entry) {
	label := widget.NewButton(labelStr, func() {
		err := clipboard.WriteAll(placeHolder)
		if err != nil {
			// 处理错误
			return
		}
		// ShowMessage2("复制成功", placeHolder)
	})
	input := widget.NewEntry()
	// input.SetPlaceHolder(placeHolder)
	// input.TextStyle({ co })

	item := container.New(layout.NewFormLayout(), label, input)
	return item, input
}

func ShowMessage2(tittle, text string) {
	user32dll, _ := syscall.LoadLibrary("user32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(tittle), IntPtr(0))
	defer syscall.FreeLibrary(user32dll)
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}
