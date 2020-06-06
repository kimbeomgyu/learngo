package scenes

import (
	"image"
	"image/color"
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
	runnerImage *ebiten.Image
	bgImg       *ebiten.Image
}

// Startup is StartScene Startup
func (s *StartScene) Startup() {
	var err error
	s.runnerImage, _, err = ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	s.bgImg, _, err = ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

var frameCount = 0

// Update is StartScene update
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	screen.DrawImage(s.bgImg, nil)

	// Idle Animation
	frameIdx := (frameCount / global.IdleAnimationSpeed) % global.IdleFrames
	sx := global.IdleX + global.FrameWidth*frameIdx
	sy := global.IdleY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, global.ScrrenHeight/2)
	subImg := s.runnerImage.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)

	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2, global.ScrrenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
