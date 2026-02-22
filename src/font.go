package src

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf)) //font file
	if err != nil {
		log.Fatal(err)
	}
	fontFace = s
	// Code from https://ebitengine.org/en/examples/font.html
}
