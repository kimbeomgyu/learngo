package main

import (
	"learngo/game2/global"
	"learngo/game2/scenemanager"
	"learngo/game2/scenes"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	var err error

	scenemanager.SetScene(&scenes.StartScene{})

	err = ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScrrenHeight, 1.0, "Monalisa")

	if err != nil {
		log.Fatalf("Ebiten run error %v", err)
	}
}
