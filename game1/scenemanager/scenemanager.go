package scenemanager

import "github.com/hajimehoshi/ebiten"

// Scene is Scene
type Scene interface {
	Update(*ebiten.Image) error
}

type scenemanager struct {
	currentScene Scene
}

var manager *scenemanager

func init() {
	manager = &scenemanager{}
}

// Update is scenemanager update
func Update(screen *ebiten.Image) error {
	if manager.currentScene != nil {
		return manager.currentScene.Update(screen)
	}
	return nil
}

// SetScene is scenemanager setting scene
func SetScene(scene Scene) {
	manager.currentScene = scene
}
