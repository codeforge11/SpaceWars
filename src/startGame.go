package src

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 300
	screenHeight = 300
)

var (
	fontFace *text.GoTextFaceSource
)

func (g *Game) Update() error {
	g.Count++
	g.pointsNumber++

	// Turn on/off debug mode
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.DebugMode = !g.DebugMode
	}

	// Player movement
	speed := 2.0 //Player speed

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if !(g.PlayerX < 0) {
			g.PlayerX -= speed
		}

	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if !(g.PlayerX >= (screenWidth - 30)) {
			g.PlayerX += speed
		}

	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if !(g.PlayerY <= 0) {
			g.PlayerY -= speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if !(g.PlayerY >= (screenHeight - 50)) {
			g.PlayerY += speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.shoot()
	}

	//Moving bullets to top
	for i := range g.Bullets {
		g.Bullets[i].BulletY -= 5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Main game logic

	screen.Fill(color.RGBA{27, 18, 18, 255}) //Window background

	if g.PlayerImg != nil { //Draw player
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.PlayerX, g.PlayerY)
		screen.DrawImage(g.PlayerImg, op)
	}

	for i := range g.Bullets {
		g.Bullets[i].Draw(screen, g.BulletImg)
	}

	//Debugger
	if g.DebugMode {
		g.DebugView(screen)
	}

	msg := fmt.Sprintf("Points:%d", (g.pointsNumber / 60))

	op := &text.DrawOptions{}
	op.GeoM.Translate(230, 280)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: fontFace,
		Size:   15, //Text font size
	}, op)

}

func StartGame() {

	Game := &Game{
		PlayerImg: PlayerCostume(),
		PlayerX:   screenWidth / 2, // Start posizion
		PlayerY:   screenHeight / 2,

		BulletImg: BulletCostume(),

		pointsNumber: 0,
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2) // Set window size
	ebiten.SetWindowTitle("SpaceWars")                  // Title

	err := ebiten.RunGame(Game) //Start game
	if err != nil {
		log.Fatal(err)
	}
}
