package animation

import "github.com/hajimehoshi/ebiten"

type animationInfo struct {
	sprites []*ebiten.Image
	speed   int
}

// Handler is struct
type Handler struct {
	animationInfo    map[string]animationInfo
	currentAnimation animationInfo
	lastIndex        int
	remainFrames     int
}

// New make a ne animation handler
func New() *Handler {
	h := &Handler{}
	h.animationInfo = make(map[string]animationInfo)
	return h
}

// Add is New Animation
func (h *Handler) Add(name string, sprites []*ebiten.Image, speed int) {
	h.animationInfo[name] = animationInfo{sprites, speed}
}

// Play is Play the Animation
func (h *Handler) Play(name string) {
	h.currentAnimation = h.animationInfo[name]
	h.lastIndex = 0
	h.remainFrames = h.currentAnimation.speed
}

// Update is draw animation frame
func (h *Handler) Update(screen *ebiten.Image, x, y float64) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(h.currentAnimation.sprites[h.lastIndex], &op)

	h.remainFrames--
	if h.remainFrames == 0 {
		h.lastIndex++
		if len(h.currentAnimation.sprites) == h.lastIndex {
			h.lastIndex = 0
		}
		h.remainFrames = h.currentAnimation.speed
	}
}
