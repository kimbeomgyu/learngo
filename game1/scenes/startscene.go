package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// StartScene is struct
type StartScene struct {
}

// Update is StartScene update
func (s *StartScene) Update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Start Scene")
	return nil
}
