package group

import (
	"gin-decode/app/ui/base"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/widget"
)

func MorseGroup() (*fyne.Container, *widget.Entry) {
	line := canvas.NewLine(color.White)
	label := canvas.NewText("MORSE", color.White)
	var morseItem, morseInput = base.Item("morse:", "-..|---|-.|--.|.-..|.")
	morseGrid := container.New(layout.NewGridLayout(3), morseItem)

	// groupBase64 := Group{lineBase64, base64GroupLabel, base64Grid}
	// item := container.New(layout.NewFormLayout(), label, input)
	boxMorse := container.NewVBox(line, label, morseGrid)
	return boxMorse, morseInput
}
