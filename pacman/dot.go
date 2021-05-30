package pacman

import (
	"container/list"

	pacimages "github.com/UlisesBojorquez/PacmanGo/images"
	"github.com/hajimehoshi/ebiten"
)

type dotManager struct {
	dots  *list.List //a list container to handle the dots
	image *ebiten.Image
}

//Constructor for the dots
func newDotManager() *dotManager {
	d := &dotManager{}
	d.dots = list.New()
	d.loadImage()
	return d
}

func (d *dotManager) loadImage() {
	d.image = loadImage(pacimages.Dot_png)
}

func (d *dotManager) draw(sc *ebiten.Image) {
	for e := d.dots.Front(); e != nil; e = e.Next() {
		v := e.Value.(pos)
		x := float64(v.x * stageBlocSize)
		y := float64(v.y * stageBlocSize)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		sc.DrawImage(d.image, op)
	}
}

func (d *dotManager) add(y, x int) {
	d.dots.PushBack(pos{y, x})
}
