package base

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Item(laebl string) (*fyne.Container, *widget.Entry) {
	label := canvas.NewText(laebl, color.White)
	// content := container.NewVBox(
	// 	label,
	// )
	// entry := widget.NewButton("Hover over me", func() {})
	// entry.OnMouseEntered  = func() {
	//     tooltip.Show()
	// }

	input := widget.NewEntry()
	item := container.New(layout.NewFormLayout(), label, input)
	return item, input
}
