package scenes

import (
	"learngo/game1/scenemanager"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// GameoverScene is struct
type GameoverScene struct {
	gameoverImg *ebiten.Image
}

// Startup is GameoverScene Startup
func (g *GameoverScene) Startup() {
	var err error
	g.gameoverImg, _, err = ebitenutil.NewImageFromFile("./images/gameover.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

// Update is GameoverScene update
func (g *GameoverScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(g.gameoverImg, nil)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&StartScene{})
	}
	return nil
}
