package scenes

import (
	"learngo/game3/actor"
	"learngo/game3/bgscroller"
	"learngo/game3/global"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// GameScene is GameScene
type GameScene struct {
	bgscroller *bgscroller.Scroller
	runner     *actor.Runner
}

// Startup is GameScene Startup
func (g *GameScene) Startup() {
	frameCount = 0

	g.runner = actor.NewRunner(0, global.ScrrenHeight/2)

	backImg, _, err := ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.bgscroller = bgscroller.New(backImg, global.BGScrollSpeed)
	g.runner.SetState(actor.Running)

}

// Update is GameScene update
func (g *GameScene) Update(screen *ebiten.Image) error {
	g.bgscroller.Update(screen)
	// Running Animation
	g.runner.Update(screen)
	return nil
}
