package scenes

import (
	"learngo/game1/global"
	"learngo/game1/scenemanager"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// GameScene is GameScene
type GameScene struct {
	board       [global.BoardWidth][global.BoardHeight]GimulType
	bgimg       *ebiten.Image
	gimulImgs   [GimulMax]*ebiten.Image
	selectedImg *ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType
	gameover    bool
}

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

// Startup is GameScene Startup
func (g *GameScene) Startup() {
	var err error
	g.gameover = false
	g.currentTeam = TeamGreen

	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImgs[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	// Initialize board
	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {
			g.board[i][j] = GimulNone
		}
	}
	g.board[0][0] = GimulGreenSang
	g.board[0][1] = GimulGreenWang
	g.board[0][2] = GimulGreenJang
	g.board[1][1] = GimulGreenJa

	g.board[3][0] = GimulRedSang
	g.board[3][1] = GimulRedWang
	g.board[3][2] = GimulRedJang
	g.board[2][1] = GimulRedJa

}

// Update is GameScene update
func (g *GameScene) Update(screen *ebiten.Image) error {

	screen.DrawImage(g.bgimg, nil)
	if g.gameover {
		return nil
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/global.GridWidth, y/global.GridHeight
		if g.selected {
			if i == g.selectedX && j == g.selectedY {
				g.selected = false
			} else {
				// move
				g.move(g.selectedX, g.selectedY, i, j)
			}
		} else {
			if g.board[i][j] != GimulNone && g.currentTeam == getTeamType(g.board[i][j]) {
				g.selected = true
				g.selectedX, g.selectedY = i, j
			}
		}
	}

	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*i), float64(global.GimulStartY+global.GridHeight*j))

			switch g.board[i][j] {
			case GimulGreenWang:
				// Draw GreenWang
				screen.DrawImage(g.gimulImgs[GimulGreenWang], options)
			case GimulGreenJa:
				// Draw GreenJa
				screen.DrawImage(g.gimulImgs[GimulGreenJa], options)
			case GimulGreenJang:
				// Draw GreenJang
				screen.DrawImage(g.gimulImgs[GimulGreenJang], options)
			case GimulGreenSang:
				// Draw GreenSang
				screen.DrawImage(g.gimulImgs[GimulGreenSang], options)
			case GimulRedWang:
				// Draw RedWang
				screen.DrawImage(g.gimulImgs[GimulRedWang], options)
			case GimulRedJa:
				// Draw RedJa
				screen.DrawImage(g.gimulImgs[GimulRedJa], options)
			case GimulRedJang:
				// Draw RedJang
				screen.DrawImage(g.gimulImgs[GimulRedJang], options)
			case GimulRedSang:
				// Draw RedSang
				screen.DrawImage(g.gimulImgs[GimulRedSang], options)
			}
		}
	}

	if g.selected {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*g.selectedX), float64(global.GimulStartY+global.GridHeight*g.selectedY))

		screen.DrawImage(g.selectedImg, options)
	}
	return nil
}

func (g *GameScene) move(prevX, prevY, tarX, tarY int) {
	if g.isMovable(prevX, prevY, tarX, tarY) {
		g.OnDie(g.board[tarX][tarY])
		g.board[prevX][prevY], g.board[tarX][tarY] = GimulNone, g.board[prevX][prevY]
		g.selected = false
		if g.currentTeam == TeamGreen {
			g.currentTeam = TeamRed
		} else {
			g.currentTeam = TeamGreen
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (g *GameScene) isMovable(prevX, prevY, tarX, tarY int) bool {
	if getTeamType(g.board[prevX][prevY]) == getTeamType(g.board[tarX][tarY]) {
		return false
	}

	if tarX < 0 || tarY < 0 {
		return false
	}

	if tarX >= global.BoardWidth || tarY >= global.BoardHeight {
		return false
	}

	switch g.board[prevX][prevY] {
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
func (g *GameScene) OnDie(gimulType GimulType) {
	if gimulType == GimulGreenWang ||
		gimulType == GimulRedWang {
		scenemanager.SetScene(&GameoverScene{})
	}
}
