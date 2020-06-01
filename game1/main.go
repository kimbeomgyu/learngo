package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// GimulType is Gimul Type
type GimulType int

// 장기말
const (
	GimulTypeNone GimulType = -1 + iota
	GimulTypeGreenWang
	GimulTypeGreenJa
	GimulTypeGreenJang
	GimulTypeGreenSang
	GimulTypeRedWang
	GimulTypeRedJa
	GimulTypeRedJang
	GimulTypeRedSang
	GimulTypeMax
)

// location
const (
	ScreenWidth  = 480
	ScrrenHeight = 362
	GridWidth    = 116
	GridHeight   = 116
	BoardWidth   = 4
	BoardHeight  = 3
	GimulStartX  = 20
	GimulStartY  = 23
)

var (
	board     [BoardWidth][BoardHeight]GimulType
	bgimg     *ebiten.Image
	gimulImgs [GimulTypeMax]*ebiten.Image
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(bgimg, nil)

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j))

			switch board[i][j] {
			case GimulTypeGreenWang:
				// Draw GreenWang
				screen.DrawImage(gimulImgs[GimulTypeGreenWang], options)
			case GimulTypeGreenJa:
				// Draw GreenJa
				screen.DrawImage(gimulImgs[GimulTypeGreenJa], options)
			case GimulTypeGreenJang:
				// Draw GreenJang
				screen.DrawImage(gimulImgs[GimulTypeGreenJang], options)
			case GimulTypeGreenSang:
				// Draw GreenSang
				screen.DrawImage(gimulImgs[GimulTypeGreenSang], options)
			case GimulTypeRedWang:
				// Draw RedWang
				screen.DrawImage(gimulImgs[GimulTypeRedWang], options)
			case GimulTypeRedJa:
				// Draw RedJa
				screen.DrawImage(gimulImgs[GimulTypeRedJa], options)
			case GimulTypeRedJang:
				// Draw RedJang
				screen.DrawImage(gimulImgs[GimulTypeRedJang], options)
			case GimulTypeRedSang:
				// Draw RedSang
				screen.DrawImage(gimulImgs[GimulTypeRedSang], options)
			}
		}
	}
	return nil
}

func main() {
	var err error

	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	// Initialize board
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulTypeNone
		}
	}
	board[0][0] = GimulTypeGreenSang
	board[0][1] = GimulTypeGreenWang
	board[0][2] = GimulTypeGreenJang
	board[1][1] = GimulTypeGreenJa

	board[3][0] = GimulTypeRedSang
	board[3][1] = GimulTypeRedWang
	board[3][2] = GimulTypeRedJang
	board[2][1] = GimulTypeRedJa

	err = ebiten.Run(update, ScreenWidth, ScrrenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error %v", err)
	}
}
