package pacman

import (
	"github.com/hajimehoshi/ebiten"
)

type ghost struct {
	kind       elem
	currentImg int
	curPos     pos
}

func newGhost(y, x int, k elem) *ghost {
	return &ghost{
		kind:   k,
		curPos: pos{y, x},
	}
}

//return the actual image ghost
func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {
	return imgs[g.currentImg]
}

//draw the ghost
func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	x := float64(g.curPos.x * stageBlocSize)
	y := float64(g.curPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.image(imgs), op)
}
