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
	var ansiItem, ansiInput = base.Item("ANSI:", "C4E3BAC3D1BDA3ACBAC3BFAAD0C4A3A1")
	var ucs2Item, ucs2Input = base.Item("Ucs2:", "4F60597D5440FF0C597D5F005FC3FF01")
	var utf8Item, utf8Input = base.Item("UTF-8:", "E4BDA0E5A5BDE59180EFBC8CE5A5BDE5BC80E5BF83EFBC81")
	var utf16Item, utf16Input = base.Item("UTF-16:", "\u0048\u0065\u006c\u006c\u006f\u002c\u0020\u0057\u006f\u0072\u006c\u0064\u0021")
	var utf32Item, utf32Input = base.Item("UTF-32:", "\x00\x00\x00H\x00\x00\x00e\x00\x00\x00l\x00\x00\x00l\x00\x00\x00o\x00\x00\x00,\x00\x00\x00 \x00\x00\x00W\x00\x00\x00o\x00\x00\x00r\x00\x00\x00l\x00\x00\x00d\x00\x00\x00!")
	grid := container.New(layout.NewGridLayout(3), ansiItem, ucs2Item, utf8Item, utf16Item, utf32Item)

	boxUrl := container.NewVBox(line, groupLabel, grid)
	return boxUrl, ansiInput, ucs2Input, utf8Input, utf16Input, utf32Input
}
