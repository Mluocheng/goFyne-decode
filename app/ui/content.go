package ui

import (
	"fmt"
	"gin-decode/app/ui/base"
	"gin-decode/app/ui/group"
	"gin-decode/app/utils"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MyContent struct {
	container fyne.CanvasObject
}

func NewContent() *MyContent {
	content := &MyContent{}

	// 输入框
	var inputEntry = base.Input()
	// 散列/哈希 组
	var boxHash, _, _, _, _, _, _, _ = group.HashGroup()
	// base64 组
	var boxBase64, base16Input, base32Input, base45Input, base58Input, base62Input, base64Input, base64UrlInput, base85Input, base91Input = group.BaseGroup()
	// url 组
	var boxUrl, decodeURIInput, decodeURIComponentInput = group.UrlGroup()
	// 字符集
	var boxCharacter, ansiInput, ucs2Input, utf8Input, utf16Input, utf32Input = group.CharacterGroup()

	// 重置函数
	resetInput := func() {
		str := ""
		base16Input.SetText(str)
		base32Input.SetText(str)
		base45Input.SetText(str)
		base58Input.SetText(str)
		base62Input.SetText(str)
		base64Input.SetText(str)
		base64UrlInput.SetText(str)
		base85Input.SetText(str)
		base91Input.SetText(str)

		decodeURIInput.SetText(str)
		decodeURIComponentInput.SetText(str)
		ansiInput.SetText(str)
		ucs2Input.SetText(str)
		utf8Input.SetText(str)
		utf16Input.SetText(str)
		utf32Input.SetText(str)
	}

	// 设置值方法
	setBaseValue := func(text string) {
		base16Value, _ := utils.Base16Decode(text)
		base32Value, _ := utils.Base32Decode(text)
		base45Value, _ := utils.Base45Decode(text)
		base58Value, _ := utils.Base58Decode(text)
		base62Value, _ := utils.Base62Decode(text)
		base64Value, _ := utils.Base64Decode(text)
		base64UrlValue, _ := utils.Base64UrlDecode(text)
		base85Value, _ := utils.Base85Decode(text)
		base91Value, _ := utils.Base91Decode(text)

		if len(base16Value) > 0 {
			base16Input.SetText(base16Value)
		}

		if len(base32Value) > 0 {
			base32Input.SetText(base32Value)
		}

		if len(base45Value) > 0 {
			base45Input.SetText(base45Value)
		}

		if len(base58Value) > 0 {
			base58Input.SetText(base58Value)
		}

		if len(base62Value) > 0 {
			base62Input.SetText(base62Value)
		}

		if len(base64Value) > 0 {
			base64Input.SetText(base64Value)
		}
		if len(base64UrlValue) > 0 {
			base64UrlInput.SetText(base64UrlValue)
		}

		if len(base85Value) > 0 {
			base85Input.SetText(base85Value)
		}
		if len(base91Value) > 0 {
			base91Input.SetText(base91Value)
		}
	}

	// 提交按钮
	button := base.Button("一键解码", func() {
		resetInput()

		text := inputEntry.Text

		urlValue, _ := utils.UrlDecode(text)
		Ansi2String := utils.Ansi2String(text)
		ucs2String := utils.Ucs2Tostring(text)
		utf8String := utils.Utf8Encode2String(text)
		utf16String := utils.Utf16Encode2String(text)
		utf32StringValue, _ := utils.UTF32ToString(strings.Trim(text, "\n"))

		if len(text) > 0 {
			setBaseValue(text)

			if len(urlValue) > 0 {
				decodeURIInput.SetText(urlValue)
				decodeURIComponentInput.SetText(urlValue)
			}

			if len(Ansi2String) > 0 {
				ansiInput.SetText(Ansi2String)
			}

			if len(ucs2String) > 0 {
				ucs2Input.SetText(ucs2String)
			}

			if len(utf8String) > 0 {
				utf8Input.SetText(utf8String)
			}
			if len(utf16String) > 0 {
				utf16Input.SetText(utf16String)
			}
			if len(utf32StringValue) > 0 {
				utf32Input.SetText(utf32StringValue)
			}

			fmt.Println("解码完成 -----")
		}
	})

	// 填充
	content.container = container.NewVBox(inputEntry,
		boxBase64,
		boxUrl,
		boxCharacter,
		boxHash,
		button)

	return content
}

func (c *MyContent) Render() fyne.CanvasObject {
	return c.container
}
