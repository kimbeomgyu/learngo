package main

import (
	"learngo/game1/global"
	"learngo/game1/scenemanager"
	"learngo/game1/scenes"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	var err error

	scenemanager.SetScene(&scenes.StartScene{})

	err = ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScrrenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error %v", err)
	}
}
