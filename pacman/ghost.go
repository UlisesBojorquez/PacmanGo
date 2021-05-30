package pacman

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ghost struct {
	kind        elem
	currentImg  int
	curPos      pos
	nextPos     pos
	prevPos     pos
	speed       int
	stepsLength pos
	steps       int
	direction   input
	vision      int
}

func newGhost(y, x int, k elem) *ghost {
	return &ghost{
		kind:        k,
		curPos:      pos{y, x},
		prevPos:     pos{y, x},
		nextPos:     pos{y, x},
		stepsLength: pos{},
		speed:       4,
		vision:      getVision(k),
	}
}

//return the actual image ghost
func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {
	return imgs[g.currentImg]
}

//draw the ghost
func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	x := float64(g.curPos.x*stageBlocSize + g.stepsLength.x)
	y := float64(g.curPos.y*stageBlocSize + g.stepsLength.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.image(imgs), op)
}

/*MOVEMENT PART*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *ghost) move() {
	switch g.direction {
	case up:
		g.stepsLength.y -= g.speed
	case right:
		g.stepsLength.x += g.speed
	case down:
		g.stepsLength.y += g.speed
	case left:
		g.stepsLength.x -= g.speed
	}

	if g.steps%4 == 0 {
		g.updateImage()
	}
	g.steps++

	if g.steps == 8 {
		g.endMove()
	}
}

func (g *ghost) updateImage() {
	switch g.direction {
	case up:
		if g.currentImg == 6 {
			g.currentImg = 7
		} else {
			g.currentImg = 6
		}
	case right:
		if g.currentImg == 0 {
			g.currentImg = 1
		} else {
			g.currentImg = 0
		}
	case down:
		if g.currentImg == 2 {
			g.currentImg = 3
		} else {
			g.currentImg = 2
		}
	case left:
		if g.currentImg == 4 {
			g.currentImg = 5
		} else {
			g.currentImg = 4
		}
	}
}

func (g *ghost) findNextMove(m [][]elem, pac pos) {

	switch g.localisePlayer(m, pac) {
	case up:
		g.direction = up
	case right:
		g.direction = right
	case down:
		g.direction = down
	case left:
		g.direction = left
	default:

		for _, v := range rand.Perm(5) {
			if v == 0 {
				continue
			}
			dir := input(v)
			np := addPosDirection(dir, g.curPos)
			if canMove(m, np) && np != g.prevPos {
				g.direction = dir
				g.nextPos = np
				return
			}
		}

		g.direction = oppDir(g.direction)
	}
	g.nextPos = addPosDirection(g.direction, g.curPos)
}

func (g *ghost) localisePlayer(m [][]elem, pac pos) input {

	maxY := len(m)
	maxX := len(m[0])

	// up
	if g.curPos.x == pac.x && g.curPos.y > pac.y {
		for y, v := g.curPos.y-1, 1; y >= 0 && v <= g.vision && !isWall(m[y][g.curPos.x]); y, v = y-1, v+1 {
			if y == pac.y {
				return up
			}
		}
	}

	// down
	if g.curPos.x == pac.x && g.curPos.y < pac.y {
		for y, v := g.curPos.y+1, 1; y < maxY && v <= g.vision && !isWall(m[y][g.curPos.x]); y, v = y+1, v+1 {
			if y == pac.y {
				return down
			}
		}
	}

	// right
	if g.curPos.y == pac.y && g.curPos.x < pac.x {
		for x, v := g.curPos.x+1, 1; x < maxX && v <= g.vision && !isWall(m[g.curPos.y][x]); x, v = x+1, v+1 {
			if x == pac.x {
				return right
			}
		}
	}

	// left
	if g.curPos.y == pac.y && g.curPos.x > pac.x {
		for x, v := g.curPos.x-1, 1; x >= 0 && v <= g.vision && !isWall(m[g.curPos.y][x]); x, v = x-1, v+1 {
			if x == pac.x {
				return left
			}
		}
	}

	return 0
}

func getVision(e elem) int {
	switch e {
	case pinkyElem:
		return 10
	case inkyElem:
		return 15
	case blinkyElem:
		return 50
	case clydeElem:
		return 60
	default:
		return 0
	}
}

//detect if the ghost is moving
func (g *ghost) isMoving() bool {
	if g.steps > 0 {
		return true
	}
	return false
}
