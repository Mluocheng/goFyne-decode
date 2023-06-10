package main

import (
	"gin-decode/app/ui"
	"gin-decode/app/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	utils.SetFont() // 设置字体
	a := app.New()
	w := a.NewWindow("decode 1.0")
	w.Resize(fyne.NewSize(1200, 600))

	content := ui.NewContent()
	w.SetContent(content.Render())
	w.SetFixedSize(true) // 设置窗口不可缩放
	w.ShowAndRun()
}
