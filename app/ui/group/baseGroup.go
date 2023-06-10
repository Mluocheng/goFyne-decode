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

type Group struct {
	*canvas.Line
	*canvas.Text
	*fyne.Container
}

func BaseGroup() (*fyne.Container, *widget.Entry,
	*widget.Entry, *widget.Entry, *widget.Entry,
	*widget.Entry, *widget.Entry, *widget.Entry,
	*widget.Entry, *widget.Entry) {
	line := canvas.NewLine(color.White)
	label := canvas.NewText("BASE", color.White)
	var base16Item, base16Input = base.Item("base16:", "68656c6c6f20776f726c64")
	var base32Item, base32Input = base.Item("base32:", "NBSWY3DPEB3W64TMMQ======")
	var base45Item, base45Input = base.Item("base45:", "+8D VD82EK4F.KEA2")
	var base58Item, base58Input = base.Item("base58:", "StV1DL6CwTryKyV")
	var base62Item, base62Input = base.Item("base62:", "AAwf93rvy4aWQVw")
	var base64Item, base64Input = base.Item("base64:", "aGVsbG8gd29ybGQ=")
	var base64UrlItem, base64UrlInput = base.Item("base64url:", "d3d3LmdvdWd1b3lpbi5jbg==")
	var base85Item, base85Input = base.Item("base85:", "BOu!rD]j7BEbo7")
	var base91Item, base91Input = base.Item("base91:", "TPwJh>Io2Tv!lE")
	base64Grid := container.New(layout.NewGridLayout(3),
		base16Item, base32Item, base45Item, base58Item,
		base62Item, base64Item, base64UrlItem, base85Item, base91Item,
	)

	// groupBase64 := Group{lineBase64, base64GroupLabel, base64Grid}
	// item := container.New(layout.NewFormLayout(), label, input)
	boxBase64 := container.NewVBox(line, label, base64Grid)
	return boxBase64, base16Input, base32Input, base45Input,
		base58Input, base62Input, base64Input, base64UrlInput, base85Input, base91Input
}
