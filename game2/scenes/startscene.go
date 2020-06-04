package scenes

import (
	"image/color"
	"learngo/game2/font"
	"learngo/game2/global"
	"learngo/game2/scenemanager"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// StartScene is struct
type StartScene struct {
	startImg *ebiten.Image
}

// Startup is StartScene Startup
func (s *StartScene) Startup() {
	var err error
	s.startImg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

// Update is StartScene update
func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)

	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2, global.ScrrenHeight/2, 2, color.Black)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
