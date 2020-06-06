package scenes

import (
	"image"
	"learngo/game2/global"
	"learngo/game2/scenemanager"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// GameScene is GameScene
type GameScene struct {
	bgimg          *ebiten.Image
	subImages      [global.PuzzleColumns * global.PuzzleRows]*ebiten.Image
	board          [global.PuzzleColumns][global.PuzzleRows]int
	blankX, blankY int
}

var gameoverBoard [global.PuzzleColumns][global.PuzzleRows]int

// Startup is GameScene Startup
func (g *GameScene) Startup() {

	var err error
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	width := global.ScreenWidth / global.PuzzleColumns
	height := global.ScrrenHeight / global.PuzzleRows

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			g.subImages[j*global.PuzzleColumns+i] = g.bgimg.SubImage(image.Rect(i*width, j*height, i*width+width, j*height+height)).(*ebiten.Image)
		}
	}

	arr := make([]int, global.PuzzleColumns*global.PuzzleRows-1)
	idx := 0
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			if i == global.PuzzleColumns-1 && j == global.PuzzleRows-1 {
				continue
			}
			arr[j*global.PuzzleColumns+i] = idx
			idx++
		}
	}

	g.blankX = global.PuzzleColumns - 1
	g.blankY = global.PuzzleRows - 1

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			if i == g.blankX && j == g.blankY {
				g.board[i][j] = -1
				continue
			}
			idx := rand.Intn(len(arr))
			g.board[i][j] = arr[idx]
			arr = append(arr[:idx], arr[idx+1:]...)
		}
	}

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			if i == g.blankX && j == g.blankY {
				gameoverBoard[i][j] = -1
			} else {
				gameoverBoard[i][j] = j*global.PuzzleColumns + i
			}
		}
	}

}

// Update is GameScene update
func (g *GameScene) Update(screen *ebiten.Image) error {

	if g.board == gameoverBoard {
		scenemanager.SetScene(&StartScene{})
		return nil
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		if g.blankY > 0 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY-1]
			g.board[g.blankX][g.blankY-1] = -1
			g.blankY--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		if g.blankY < global.PuzzleRows-1 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY+1]
			g.board[g.blankX][g.blankY+1] = -1
			g.blankY++
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		if g.blankX > 0 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX-1][g.blankY]
			g.board[g.blankX-1][g.blankY] = -1
			g.blankX--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		if g.blankX < global.PuzzleColumns-1 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX+1][g.blankY]
			g.board[g.blankX+1][g.blankY] = -1
			g.blankX++
		}
	}

	width := global.ScreenWidth / global.PuzzleColumns
	height := global.ScrrenHeight / global.PuzzleRows

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {

			if g.board[i][j] == -1 {
				continue
			}

			x := i * width
			y := j * height

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.subImages[g.board[i][j]], opts)
		}
	}

	return nil
}
