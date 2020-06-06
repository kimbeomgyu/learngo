package scenes

import (
	"image"
	"learngo/game3/global"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// GameScene is GameScene
type GameScene struct {
	runnerImage *ebiten.Image
	bgImg       *ebiten.Image
}

// Startup is GameScene Startup
func (g *GameScene) Startup() {
	frameCount = 0
	var err error
	g.runnerImage, _, err = ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.bgImg, _, err = ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

}

// Update is GameScene update
func (g *GameScene) Update(screen *ebiten.Image) error {
	frameCount++

	bgWidth, _ := g.bgImg.Size()

	op := &ebiten.DrawImageOptions{}
	backX := (frameCount / 2) % bgWidth

	op.GeoM.Translate(float64(-backX), 0)
	screen.DrawImage(g.bgImg, op)

	op.GeoM.Translate(float64(-bgWidth), 0)
	screen.DrawImage(g.bgImg, op)

	// Running Animation
	frameIdx := (frameCount / global.RunningAnimationSpeed) % global.RunningFrames
	sx := global.RunningX + global.FrameWidth*frameIdx
	sy := global.RunningY

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, global.ScrrenHeight/2)
	subImg := g.runnerImage.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)
	return nil
}
