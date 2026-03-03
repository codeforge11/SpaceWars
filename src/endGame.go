package src

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	title     string
	miniText1 string
	miniText2 string
)

func (g *Game) endGame(screen *ebiten.Image, endType bool) {
	g.Bullets = nil
	g.Enemies = nil
	g.pointsNumber = 0

	if endType {
		title = "YOU WIN"
	} else {
		title = "YOU LOSE :("
	}

	miniText1 = "Click space to restart..."
	miniText2 = "Click ESC to return to menu"

	titleDrawOptions := &text.DrawOptions{}
	titleDrawOptions.GeoM.Translate(50, 50)
	titleDrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, title, &text.GoTextFace{
		Source: fontFace,
		Size:   18, //Text font size
	}, titleDrawOptions)

	miniText1DrawOptions := &text.DrawOptions{}
	miniText1DrawOptions.GeoM.Translate(50, 150)
	miniText1DrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, miniText1, &text.GoTextFace{
		Source: fontFace,
		Size:   12, //Text font size
	}, miniText1DrawOptions)

	miniText2DrawOptions := &text.DrawOptions{}
	miniText2DrawOptions.GeoM.Translate(50, 200)
	miniText2DrawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, miniText2, &text.GoTextFace{
		Source: fontFace,
		Size:   12, //Text font size
	}, miniText2DrawOptions)

}
