package group

import (
	"gin-decode/app/ui/base"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MyCharacterGroup struct {
	*canvas.Line
	*canvas.Text
	*fyne.Container
}

func CharacterGroup() (*fyne.Container, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry) {
	line := canvas.NewLine(color.White)
	groupLabel := canvas.NewText("字符集", color.White)
	var ansiItem, ansiInput = base.Item("ANSI:")
	var ucs2Item, ucs2Input = base.Item("Ucs2:")
	var utf8Item, utf8Input = base.Item("UTF-8:")
	var utf16Item, utf16Input = base.Item("UTF-16:")
	var utf32Item, utf32Input = base.Item("UTF-32:")
	grid := container.New(layout.NewGridLayout(3), ansiItem, ucs2Item, utf8Item, utf16Item, utf32Item)

	boxUrl := container.NewVBox(line, groupLabel, grid)
	return boxUrl, ansiInput, ucs2Input, utf8Input, utf16Input, utf32Input
}
