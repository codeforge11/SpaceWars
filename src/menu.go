package src

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) drawMenu(screen *ebiten.Image) {
	title := "SPACE WARS"
	miniText := "Click space to start..."

	titleDrawOptions := &text.DrawOptions{}
	titleDrawOptions.GeoM.Translate(100, 50)
	titleDrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, title, &text.GoTextFace{
		Source: fontFace,
		Size:   18, //Text font size
	}, titleDrawOptions)

	miniTextDrawOptions := &text.DrawOptions{}
	miniTextDrawOptions.GeoM.Translate(100, 150)
	miniTextDrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, miniText, &text.GoTextFace{
		Source: fontFace,
		Size:   10, //Text font size
	}, miniTextDrawOptions)
}
