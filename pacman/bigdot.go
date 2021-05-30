package pacman

import (
	"container/list"

	pacimages "github.com/UlisesBojorquez/PacmanGo/images"
	"github.com/hajimehoshi/ebiten"
)

type bigDotManager struct {
	dots   *list.List
	images [2]*ebiten.Image
	count  int //used for animation
}

func newBigDotManager() *bigDotManager {
	bd := &bigDotManager{}
	bd.dots = list.New()
	bd.loadImages()
	return bd
}

//Load the two images from bigdot, there are two dots dou to animation
func (b *bigDotManager) loadImages() {
	b.images[0] = loadImage(pacimages.BigDot1_png)
	b.images[1] = loadImage(pacimages.BigDot2_png)
}

func (b *bigDotManager) add(y, x int) {
	b.dots.PushBack(pos{y, x})
}

func (b *bigDotManager) draw(sc *ebiten.Image) {
	//Animation part
	b.count++
	var img *ebiten.Image
	if b.count%40 == 0 {
		img = b.images[1]
	} else {
		img = b.images[0]
	}

	for e := b.dots.Front(); e != nil; e = e.Next() {
		d := e.Value.(pos)
		x := float64(d.x * stageBlocSize)
		y := float64(d.y * stageBlocSize)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		sc.DrawImage(img, op)
	}
}
