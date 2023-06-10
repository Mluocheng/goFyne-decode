package utils

import (
	"fmt"
	"os"

	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
)

func SetFont() {
	fontPath, err := findfont.Find("simkai.ttf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found 'simkai.ttf' in '%s'\n", fontPath)

	// load the font with the freetype library
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}
