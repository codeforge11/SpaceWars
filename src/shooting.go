package src

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	BulletX, BulletY float64
	activeState      bool
	isHit            bool
}

func (g *Game) shoot() {
	newBullet := Bullet{
		BulletX: g.PlayerX + 10, // Player's center
		BulletY: g.PlayerY,
		isHit:   false,
	}
	g.Bullets = append(g.Bullets, newBullet)
}

func (b *Bullet) Draw(screen *ebiten.Image, img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.BulletX, b.BulletY)
	screen.DrawImage(img, op)
}
