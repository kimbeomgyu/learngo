package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	bg *ebiten.Image
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(bg, nil)
	return nil
}

func main() {
	var err error

	bg, _, err = ebitenutil.NewImageFromFile("./12janggi.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	err = ebiten.Run(update, 600, 400, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error %v", err)
	}
}
