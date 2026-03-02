package src

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) DebugView(screen *ebiten.Image) {
	msg := fmt.Sprintf(
		"Welcome in debug mode\n"+
			"Seconds from startup : %d\n"+
			"FPS: %.2f\n"+
			"Player cords: X: %.2f, Y: %.2f \n"+
			"Active bullets: %d\n"+
			"Active enemies: %d\n"+
			"Game level: %d\n"+
			"Creator: codeforge11",
		g.Count/60,
		ebiten.ActualFPS(),
		g.PlayerX,
		g.PlayerY,
		len(g.Bullets),
		len(g.Enemies),
		gameLevel,
	)

	ebitenutil.DebugPrint(screen, msg)
}
