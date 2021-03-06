package pacman

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ghost struct {
	kind               elem
	currentImg         int
	curPos             pos
	nextPos            pos
	prevPos            pos
	speed              int
	stepsLength        pos
	steps              int
	direction          input
	vision             int
	countVulnerability int
	vulnerableMove     bool
	initialPos         pos
	eaten              bool
	killed             bool
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
		initialPos:  pos{y, x},
	}
}

//return the actual image ghost
func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {

	if g.isVulnerable() {
		i := g.currentImg + 8
		if i >= len(imgs) {
			i = 8
		}
		return imgs[i]
	}
	return imgs[g.currentImg]
}

//draw the ghost
func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	if g.killed {
		return
	}

	x := float64(g.curPos.x*stageBlocSize + g.stepsLength.x)
	y := float64(g.curPos.y*stageBlocSize + g.stepsLength.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.image(imgs), op)
}

func (g *ghost) screenPos() (y, x float64) {
	x = float64(g.curPos.x*stageBlocSize + g.stepsLength.x)
	y = float64(g.curPos.y*stageBlocSize + g.stepsLength.y)
	return
}

/*MOVEMENT PART*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *ghost) move() {
	if !g.killed {
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

		if g.vulnerableMove {
			g.countVulnerability++
			if g.steps == 16 {
				g.endMove()
				if g.countVulnerability >= 392 { //seconds 392/60 = 7seconds aprox
					g.endVulnerability()
				}
			}
			return
		}

		if g.steps == 8 {
			g.endMove()
		}
	}
}

/*UPDATING PART FORM THE IMAGES*/
func (g *ghost) updateImage() {

	if g.isVulnerable() {
		if g.countVulnerability <= 310 {
			if g.currentImg == 0 {
				g.currentImg = 1
			} else {
				g.currentImg = 0
			}
		} else {
			if g.currentImg == 2 {
				g.currentImg = 3
			} else {
				g.currentImg = 2
			}
		}
		return
	}

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

/*FINDING THE NEXT MOVE FOR THE GHOST*/
func (g *ghost) findNextMove(m [][]elem, pac pos) {

	if g.isVulnerable() {
		g.vulnerableMove = true
		g.speed = 2
	} else {
		g.speed = 4
	}

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
			np := addNextDirection(dir, g.curPos)
			if canMove(m, np) && np != g.prevPos {
				g.direction = dir
				g.nextPos = np
				return
			}
		}

		g.direction = oppDirection(g.direction)
	}
	g.nextPos = addNextDirection(g.direction, g.curPos)
}

/*LOCALIZATION OF THE PLAYER*/
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

func (g *ghost) isVulnerable() bool {
	if g.countVulnerability > 0 {
		return true
	}
	return false
}

func (g *ghost) endVulnerability() {
	g.vulnerableMove = false
	g.countVulnerability = 0
	g.eaten = false
}

/*COLLISION*/
func (g *ghost) makeVulnerable() {
	g.countVulnerability = 1
}

func (g *ghost) resetGhost() {
	g.prevPos, g.curPos, g.nextPos = g.initialPos, g.initialPos, g.initialPos
	g.stepsLength = pos{}
	g.currentImg = 0
	g.direction = 0
	g.steps = 0
}

func (g *ghost) makeEaten() {
	g.eaten = true
}

func (g *ghost) isEaten() bool {
	return g.eaten
}

func (g *ghost) reinit() {
	g.resetGhost()
	g.killed = false
	g.speed = 4
	g.endVulnerability()
}

func (g *ghost) killGhost() {
	p := pos{0, 0}
	g.prevPos, g.curPos, g.nextPos = p, p, p
	g.killed = true
}
