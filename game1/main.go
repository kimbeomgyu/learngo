package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// GimulType is Gimul Type
type GimulType int

// 장기말
const (
	GimulNone GimulType = -1 + iota
	GimulGreenWang
	GimulGreenJa
	GimulGreenJang
	GimulGreenSang
	GimulRedWang
	GimulRedJa
	GimulRedJang
	GimulRedSang
	GimulMax
)

// TeamType is Team Type
type TeamType int

// 팀
const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
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
	board       [BoardWidth][BoardHeight]GimulType
	bgimg       *ebiten.Image
	gimulImgs   [GimulMax]*ebiten.Image
	selectedImg *ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType = TeamGreen
	gameover    bool
)

func getTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulGreenJa ||
		gimulType == GimulGreenJang ||
		gimulType == GimulGreenSang ||
		gimulType == GimulGreenWang {
		return TeamGreen
	}
	if gimulType == GimulRedJa ||
		gimulType == GimulRedJang ||
		gimulType == GimulRedSang ||
		gimulType == GimulRedWang {
		return TeamRed
	}
	return TeamNone
}

func update(screen *ebiten.Image) error {

	screen.DrawImage(bgimg, nil)
	if gameover {
		return nil
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridWidth, y/GridHeight
		if selected {
			if i == selectedX && j == selectedY {
				selected = false
			} else {
				// move
				move(selectedX, selectedY, i, j)
			}
		} else {
			if board[i][j] != GimulNone && currentTeam == getTeamType(board[i][j]) {
				selected = true
				selectedX, selectedY = i, j
			}
		}
	}

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j))

			switch board[i][j] {
			case GimulGreenWang:
				// Draw GreenWang
				screen.DrawImage(gimulImgs[GimulGreenWang], options)
			case GimulGreenJa:
				// Draw GreenJa
				screen.DrawImage(gimulImgs[GimulGreenJa], options)
			case GimulGreenJang:
				// Draw GreenJang
				screen.DrawImage(gimulImgs[GimulGreenJang], options)
			case GimulGreenSang:
				// Draw GreenSang
				screen.DrawImage(gimulImgs[GimulGreenSang], options)
			case GimulRedWang:
				// Draw RedWang
				screen.DrawImage(gimulImgs[GimulRedWang], options)
			case GimulRedJa:
				// Draw RedJa
				screen.DrawImage(gimulImgs[GimulRedJa], options)
			case GimulRedJang:
				// Draw RedJang
				screen.DrawImage(gimulImgs[GimulRedJang], options)
			case GimulRedSang:
				// Draw RedSang
				screen.DrawImage(gimulImgs[GimulRedSang], options)
			}
		}
	}

	if selected {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(GimulStartX+GridWidth*selectedX), float64(GimulStartY+GridHeight*selectedY))

		screen.DrawImage(selectedImg, options)
	}
	return nil
}

func move(prevX, prevY, tarX, tarY int) {
	if isMovable(prevX, prevY, tarX, tarY) {
		OnDie(board[tarX][tarY])
		board[prevX][prevY], board[tarX][tarY] = GimulNone, board[prevX][prevY]
		selected = false
		if currentTeam == TeamGreen {
			currentTeam = TeamRed
		} else {
			currentTeam = TeamGreen
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isMovable(prevX, prevY, tarX, tarY int) bool {
	if getTeamType(board[prevX][prevY]) == getTeamType(board[tarX][tarY]) {
		return false
	}

	if tarX < 0 || tarY < 0 {
		return false
	}

	if tarX >= BoardWidth || tarY >= BoardHeight {
		return false
	}

	switch board[prevX][prevY] {
	case GimulGreenJa, GimulRedJa:
		return prevX+1 == tarX && prevY == tarY
	case GimulGreenSang, GimulRedSang:
		return abs(prevX-tarX) == 1 && abs(prevY-tarY) == 1
	case GimulGreenJang, GimulRedJang:
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulGreenWang, GimulRedWang:
		return abs(prevX-tarX) == 1 || abs(prevY-tarY) == 1
	}
	return false
}

// OnDie calls when gimul is died
func OnDie(gimulType GimulType) {
	if gimulType == GimulGreenWang ||
		gimulType == GimulRedWang {
		gameover = true
	}
}

func main() {
	var err error

	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImgs[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	// Initialize board
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulNone
		}
	}
	board[0][0] = GimulGreenSang
	board[0][1] = GimulGreenWang
	board[0][2] = GimulGreenJang
	board[1][1] = GimulGreenJa

	board[3][0] = GimulRedSang
	board[3][1] = GimulRedWang
	board[3][2] = GimulRedJang
	board[2][1] = GimulRedJa

	err = ebiten.Run(update, ScreenWidth, ScrrenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error %v", err)
	}
}
