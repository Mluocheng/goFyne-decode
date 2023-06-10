// 用户输入

package base

import (
	"fyne.io/fyne/v2/widget"
)

func Input() *widget.Entry {
	inputEntry := widget.NewEntry()
	inputEntry.MultiLine = true
	inputEntry.SetPlaceHolder("输入密文")

	return inputEntry

}
