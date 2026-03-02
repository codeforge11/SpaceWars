package src

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Count     int
	DebugMode bool

	GameImg *ebiten.Image

	PlayerX   float64
	PlayerY   float64
	PlayerImg *ebiten.Image

	Bullets   []Bullet
	BulletImg *ebiten.Image

	pointsNumber int

	Enemies  []Enemy
	EnemyImg *ebiten.Image
}
