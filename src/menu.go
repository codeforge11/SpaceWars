package src

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) drawMenu(screen *ebiten.Image) {
	title := "SPACE WARS"
	miText := "CONTROL: WSAD / ARROWS  FIRE: SPACE"
	miniText := "Click space to start..."

	titleDrawOptions := &text.DrawOptions{}
	titleDrawOptions.GeoM.Translate(100, 20)
	titleDrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, title, &text.GoTextFace{
		Source: fontFace,
		Size:   18, //Text font size
	}, titleDrawOptions)

	miTextDrawOptions := &text.DrawOptions{}
	miTextDrawOptions.GeoM.Translate(15, 125)
	miTextDrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, miText, &text.GoTextFace{
		Source: fontFace,
		Size:   14, //Text font size
	}, miTextDrawOptions)

	miniTextDrawOptions := &text.DrawOptions{}
	miniTextDrawOptions.GeoM.Translate(90, 170)
	miniTextDrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, miniText, &text.GoTextFace{
		Source: fontFace,
		Size:   10, //Text font size
	}, miniTextDrawOptions)
}
