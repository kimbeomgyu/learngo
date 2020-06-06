package bgscroller

import "github.com/hajimehoshi/ebiten"

// Scroller is struct
type Scroller struct {
	bgimg  *ebiten.Image
	speed  int
	frames int
}

// New is
func New(bgimg *ebiten.Image, speed int) *Scroller {
	return &Scroller{bgimg, speed, 0}
}

// Update is
func (s *Scroller) Update(screen *ebiten.Image) {
	s.frames++

	bgWidth, _ := s.bgimg.Size()

	op := &ebiten.DrawImageOptions{}
	backX := (s.frames / s.speed) % bgWidth

	op.GeoM.Translate(float64(-backX), 0)
	screen.DrawImage(s.bgimg, op)

	op.GeoM.Translate(float64(bgWidth), 0)
	screen.DrawImage(s.bgimg, op)

}
