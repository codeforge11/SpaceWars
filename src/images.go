package src

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	defaultImagesSources = "./src/imgs/"

	PlaneImage  = "plane.png"
	BulletImage = "bullet.png"
)

func PlayerCostume() *ebiten.Image {
	playerImg, _, err := ebitenutil.NewImageFromFile(defaultImagesSources + PlaneImage)
	if err != nil {
		log.Fatal(err)
	}
	return playerImg
}

func BulletCostume() *ebiten.Image {
	bulletImg, _, err := ebitenutil.NewImageFromFile(defaultImagesSources + BulletImage)
	if err != nil {
		log.Fatal(err)
	}
	return bulletImg
}
