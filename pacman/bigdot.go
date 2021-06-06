package pacman

import (
	"container/list"

	pacmanimages "github.com/UlisesBojorquez/PacmanGo/images"
	"github.com/hajimehoshi/ebiten"
)

type bigDotManager struct {
	dots     *list.List
	initDots *list.List
	images   [2]*ebiten.Image
	count    int //used for animation
}

func newBigDotManager() *bigDotManager {
	bd := &bigDotManager{}
	bd.dots = list.New()
	bd.initDots = list.New()
	bd.loadImages()
	return bd
}

//Load the two images from bigdot, there are two dots due to animation
func (b *bigDotManager) loadImages() {
	b.images[0] = loadImage(pacmanimages.BigDot1_png)
	b.images[1] = loadImage(pacmanimages.BigDot2_png)
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

/*COLLISION*/
func (b *bigDotManager) detectCollision(m [][]elem, p pos, cb func()) {
	if m[p.y][p.x] == bigDotElem {
		cb()
	}
}

func (b *bigDotManager) delete(p pos) {
	for e := b.dots.Front(); e != nil; e = e.Next() {
		v := e.Value.(pos)
		if v.x == p.x && v.y == p.y {
			b.initDots.PushBack(b.dots.Remove(e).(pos))
			return
		}
	}
}

/*REINIT*/

func (b *bigDotManager) reinit(m [][]elem) {
	e := b.initDots.Front()
	for {
		if e == nil {
			break
		}
		v := e.Value.(pos)
		cur := e
		e = e.Next()
		b.dots.PushBack(b.initDots.Remove(cur))
		m[v.y][v.x] = bigDotElem
	}
}

func (b *bigDotManager) empty() bool {
	if b.dots.Len() == 0 {
		return true
	}
	return false
}
