package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	frameWidth  = 640
	frameHeight = 480

	screenWidth  = 1280
	screenHeight = 960

	frameOX  = 0
	frameOY  = 32
	frameNum = 8
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("game-helloworld/image/d.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xfd, 0xd6, 0x92, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 0)
	// op.GeoM.Rotate(-100) // Rotate

	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	// i := (g.count / 5) % frameNum
	// sx, sy := frameOX+i*frameWidth, frameOY
	op.GeoM.Scale(1, 1)
	screen.DrawImage(img, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight // 화면 해상도
}

func main() {
	ebiten.SetWindowSize(frameWidth, frameHeight) // 화면 사이즈
	ebiten.SetWindowTitle("Hello world")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
