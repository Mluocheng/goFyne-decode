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

type MyUrlGroup struct {
	*canvas.Line
	*canvas.Text
	*fyne.Container
}

// https%3A%2F%2Flzltool.com%2FEscape%2FUrlEncode%3Fdata%3DURL%E5%86%85%E5%AE%B9%26encodeing%3Dutf-8
func UrlGroup() (*fyne.Container, *widget.Entry, *widget.Entry) {
	line := canvas.NewLine(color.White)
	groupLabel := canvas.NewText("URL", color.White)
	var decodeURIItem, decodeURIInput = base.Item("decodeURI:")
	var decodeURIComponentItem, decodeURIComponentInput = base.Item("decodeURIComponent:")
	grid := container.New(layout.NewGridLayout(3), decodeURIItem, decodeURIComponentItem)

	boxUrl := container.NewVBox(line, groupLabel, grid)
	return boxUrl, decodeURIInput, decodeURIComponentInput
}
