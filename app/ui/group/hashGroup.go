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

type MyHashGroup struct {
	*fyne.Container
	*widget.Entry
}

func HashGroup() (*fyne.Container, *widget.Entry, *widget.Entry, *widget.Entry,
	*widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry) {

	lineHash := canvas.NewLine(color.White)
	groupLabel := canvas.NewText("散列/哈希", color.White)
	var SHA224Item, SHA224Input = base.Item("SHA224:", "")
	var SHA256Item, SHA256Input = base.Item("SHA256:", "")
	var SHA384Item, SHA384Input = base.Item("SHA384:", "")
	var SHA251Item, SHA251Input = base.Item("SHA251:", "")
	var HmacSHA1Item, HmacSHA1Input = base.Item("HmacSHA1:", "")
	var HmacMD5Item, HmacMD5Input = base.Item("HmacMD5:", "")
	var HmacSHA224Item, HmacSHA224Input = base.Item("HmacSHA224:", "")
	grid := container.New(layout.NewGridLayout(3),
		SHA224Item, SHA256Item, SHA384Item, SHA251Item, HmacSHA1Item, HmacMD5Item, HmacSHA224Item)
	boxHash := container.NewVBox(lineHash, groupLabel, grid)
	return boxHash, SHA224Input, SHA256Input, SHA384Input, SHA251Input, HmacSHA1Input, HmacMD5Input, HmacSHA224Input
}
