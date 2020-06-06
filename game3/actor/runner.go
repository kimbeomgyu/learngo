package actor

import (
	"image"
	"learngo/game3/animation"
	"learngo/game3/global"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type runnerState int

// Constants
const (
	Idle runnerState = iota
	Running
)

// Runner is struct
type Runner struct {
	X, Y      float64
	state     runnerState
	animation *animation.Handler
}

// NewRunner is make a new runner
func NewRunner(x, y float64) *Runner {
	r := &Runner{}
	r.X, r.Y = x, y
	r.state = Idle
	r.animation = animation.New()

	runnerImage, _, err := ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	sprites := make([]*ebiten.Image, global.IdleFrames)
	for i := 0; i < global.IdleFrames; i++ {
		sx := global.IdleX + global.FrameWidth*i
		sy := global.IdleY
		sprites[i] = runnerImage.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)

	}
	r.animation.Add("Idle", sprites, global.IdleAnimationSpeed)

	sprites = make([]*ebiten.Image, global.RunningFrames)
	for i := 0; i < global.RunningFrames; i++ {
		sx := global.RunningX + global.FrameWidth*i
		sy := global.RunningY
		sprites[i] = runnerImage.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)

	}
	r.animation.Add("Run", sprites, global.RunningAnimationSpeed)
	return r
}

// SetState is state change
func (r *Runner) SetState(state runnerState) {
	r.state = state
	switch r.state {
	case Idle:
		r.animation.Play("Idle")
	case Running:
		r.animation.Play("Run")
	}
}

// Update is Draw runner
func (r *Runner) Update(screen *ebiten.Image) {
	r.animation.Update(screen, r.X, r.Y)
}
