package src

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	EnemyX, EnemyY float64
	activeState    bool
	isHit          bool
}

const (
	enemyWidth  = 30.0
	enemyHeight = 30.0
)

func (g *Game) createNewEnemy() {
	newEnemy := Enemy{
		EnemyX: rand.Float64() * screenWidth,
		EnemyY: 0,
		isHit:  false,
	}
	g.Enemies = append(g.Enemies, newEnemy)
}

func (e *Enemy) DrawEnemy(screen *ebiten.Image, img *ebiten.Image) {
	enemyDrawOption := &ebiten.DrawImageOptions{}
	enemyDrawOption.GeoM.Translate(e.EnemyX, e.EnemyY)
	screen.DrawImage(img, enemyDrawOption)
}
