package src

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	defaultImagesSources = "./src/imgs/"

	PlaneImage = "plane.png"
)

func PlayerCostume() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(defaultImagesSources + PlaneImage)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
