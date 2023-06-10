package base

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MyButton struct {
	widget.Button
}

// func NewMyButton(label string, onClick func()) *MyButton {
func Button(label string, onClick func()) *MyButton {

	button := widget.NewButton(label, onClick)
	myButton := &MyButton{
		Button: *button,
	}
	return myButton
}

func (b *MyButton) CreateRenderer() fyne.WidgetRenderer {
	return b.Button.CreateRenderer()
}
