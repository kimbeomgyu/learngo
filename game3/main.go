package main

import (
	"learngo/game3/global"
	"learngo/game3/scenemanager"
	"learngo/game3/scenes"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	var err error
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	err = ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScrrenHeight, 2.0, "Monalisa")

	if err != nil {
		log.Fatalf("Ebiten run error %v", err)
	}
}
