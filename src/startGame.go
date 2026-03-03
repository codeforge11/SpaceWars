package src

import (
	"image"
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
	fontFace  *text.GoTextFaceSource
	gameLevel int
)

func (g *Game) Update() error {
	// Turn on/off debug mode
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.DebugMode = !g.DebugMode
	}

	switch gameLevel {
	case 0:
		{
			if ebiten.IsKeyPressed(ebiten.KeySpace) {
				gameLevel = 1
			}
		}
	case 1:
		{

			g.Count++
			g.pointsNumber++

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

			if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				g.shoot()
			}

			if g.DebugMode {
				if inpututil.IsKeyJustPressed(ebiten.KeyP) {
					gameLevel = 0
				}
				if inpututil.IsKeyJustPressed(ebiten.KeyO) {
					gameLevel = 1
				}
				if inpututil.IsKeyJustPressed(ebiten.KeyM) {
					gameLevel = 4
				}
				if inpututil.IsKeyJustPressed(ebiten.KeyN) {
					gameLevel = 3
				}
			}

			if len(g.Enemies) < 7 {
				g.createNewEnemy()
			}

			//Moving bullets to top (if its activeState)
			for i := range g.Bullets {
				if g.Bullets[i].BulletY < 0 {
					g.Bullets[i].activeState = false
				} else {
					g.Bullets[i].activeState = true
					g.Bullets[i].BulletY -= 5
				}
			}

			//Moving enemies to bottom (if its activeState)
			for i := range g.Enemies {
				if g.Enemies[i].EnemyY > 250 {
					g.Enemies[i].activeState = false
				} else {
					g.Enemies[i].activeState = true
					g.Enemies[i].EnemyY += 1
				}

			}

			if g.pointsNumber == (100 * 60) {
				gameLevel = 3
			}

		}
	case 3, 4:
		{
			if ebiten.IsKeyPressed(ebiten.KeySpace) {
				gameLevel = 1
			}
			if ebiten.IsKeyPressed(ebiten.KeyEscape) {
				gameLevel = 0
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Main game logic

	screen.Fill(color.RGBA{27, 18, 18, 255}) //Window background

	//Debugger
	if g.DebugMode {
		g.DebugView(screen)
	}

	switch gameLevel {
	case 0:
		{
			g.drawMenu(screen)
		}
	case 1:
		{

			//Player
			if g.PlayerImg != nil { //Draw player
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(g.PlayerX, g.PlayerY)
				screen.DrawImage(g.PlayerImg, op)
			}

			//Creating enemies
			activeEnemies := 0
			for i := range g.Enemies {
				if g.Enemies[i].activeState {
					g.Enemies[i].DrawEnemy(screen, g.EnemyImg)
					g.Enemies[activeEnemies] = g.Enemies[i]
					activeEnemies++
				}
			}
			g.Enemies = g.Enemies[:activeEnemies]

			//Shooting
			activeBullets := 0
			for i := range g.Bullets {
				if g.Bullets[i].activeState {
					g.Bullets[i].Draw(screen, g.BulletImg)
					g.Bullets[activeBullets] = g.Bullets[i]
					activeBullets++
				}
			}
			g.Bullets = g.Bullets[:activeBullets]

			g.drawPointsNumber(screen)
		}
	case 3:
		{
			g.endGameWin(screen)
			// g.Enemies = nil //Remove enemies
		}
	case 4:
		{
			g.endGameLose(screen)
			// g.Enemies = nil //Remove enemies
		}
	}
}

func StartGame() {

	gameLevel = 0

	Game := &Game{
		PlayerImg: PlayerCostume(),
		PlayerX:   screenWidth / 2, // Start position
		PlayerY:   screenHeight / 2,

		BulletImg: BulletCostume(),

		EnemyImg: EnemyCostume(),
		GameImg:  GameCostume(),

		pointsNumber: 0,
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2) // Set window size
	ebiten.SetWindowTitle("SpaceWars")                  // Title
	ebiten.SetWindowIcon([]image.Image{Game.GameImg})

	err := ebiten.RunGame(Game) //Start game
	if err != nil {
		log.Fatal(err)
	}
}
