package scenes

import (
	"image/color"
	"learngo/game3/actor"
	"learngo/game3/font"
	"learngo/game3/global"
	"learngo/game3/scenemanager"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// StartScene is struct
type StartScene struct {
	bgImg  *ebiten.Image
	runner *actor.Runner
}

// Startup is StartScene Startup
func (s *StartScene) Startup() {
	s.runner = actor.NewRunner(0, global.ScrrenHeight/2)

	var err error
	s.bgImg, _, err = ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	s.runner.SetState(actor.Idle)
}

var frameCount = 0

// Update is StartScene update
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	screen.DrawImage(s.bgImg, nil)

	// Idle Animation
	s.runner.Update(screen)

	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2, global.ScrrenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
