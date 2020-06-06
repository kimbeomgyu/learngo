package scenes

import (
	"image"
	"learngo/game3/global"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// StartScene is struct
type StartScene struct {
	runnerImage *ebiten.Image
}

// Startup is StartScene Startup
func (s *StartScene) Startup() {
	var err error
	s.runnerImage, _, err = ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

var frameCount = 0

// Update is StartScene update
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	frameIdx := (frameCount / 5) % global.RunningFrames
	sx := global.RunningX + global.FrameWidth*frameIdx
	sy := global.RunningY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(global.ScreenWidth/2, global.ScrrenHeight/2)
	subImg := s.runnerImage.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)

	return nil
}
