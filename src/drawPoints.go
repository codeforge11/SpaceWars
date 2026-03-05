package src

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) drawPointsNumber(screen *ebiten.Image) {
	//Draw points number
	msg := fmt.Sprintf("Points:%d", (g.pointsNumber / 60))

	op := &text.DrawOptions{}
	op.GeoM.Translate(215, 280)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: fontFace,
		Size:   15, //Text font size
	}, op)
}
