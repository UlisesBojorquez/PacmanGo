package pacman

import (
	pacimages "github.com/UlisesBojorquez/PacmanGo/images"
	"github.com/hajimehoshi/ebiten"
)

type player struct {
	images      [8]*ebiten.Image
	currentImg  int
	curPos      pos
	prevPos     pos
	nextPos     pos
	speed       int //speed of the player
	stepsLength pos
	steps       int //each step the player does increment it
	direction   input
	score       int
	lost        bool
	won         bool
	initialPos  pos
	lives       int
}

func newPlayer(y, x int) *player {
	p := &player{}
	p.loadImages()
	p.curPos = pos{y, x}
	p.prevPos = pos{y, x}
	p.nextPos = pos{y, x}
	p.initialPos = pos{y, x}
	p.lives = 3
	return p
}

func (p *player) loadImages() {
	copy(p.images[:], loadImages(pacimages.PlayerImages[:]))
}

func (p *player) image() *ebiten.Image {
	return p.images[p.currentImg]
}

func (p *player) draw(screen *ebiten.Image) {

	if p.lost {
		return
	}

	x := float64(p.curPos.x*stageBlocSize + p.stepsLength.x)
	y := float64(p.curPos.y*stageBlocSize + p.stepsLength.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(p.image(), op)
}

func (p *player) move(m [][]elem, direction input) {

	//no move and no direction
	if !p.isMoving() && direction == 0 {
		return
	}

	//new direction
	if !p.isMoving() && direction != 0 {
		if !canMove(m, addPosDirection(direction, p.curPos)) {
			return
		}
		p.updateDirection(direction)
	}

	// adjust the speed
	if p.steps <= 1 || p.steps >= 6 {
		p.speed = 4
	} else {
		p.speed = 5
	}
	// move (update the coordinates)
	switch p.direction {
	case up:
		p.stepsLength.y -= p.speed
	case right:
		p.stepsLength.x += p.speed
	case down:
		p.stepsLength.y += p.speed
	case left:
		p.stepsLength.x -= p.speed
	}

	if p.steps > 5 {
		p.updateImage(false)
	} else {
		p.updateImage(true)
	}

	p.steps++

	if p.steps >= 7 {
		p.endMove()
	}
}

func (p *player) updateImage(openMouth bool) {
	switch p.direction {
	case up:
		if openMouth {
			p.currentImg = 7
		} else {
			p.currentImg = 6
		}
	case right:
		if openMouth {
			p.currentImg = 1
		} else {
			p.currentImg = 0
		}
	case down:
		if openMouth {
			p.currentImg = 3
		} else {
			p.currentImg = 2
		}
	case left:
		if openMouth {
			p.currentImg = 5
		} else {
			p.currentImg = 4
		}
	}
}

func (p *player) isMoving() bool {
	if p.steps > 0 {
		return true
	}
	return false
}

func (p *player) updateDirection(d input) {
	p.stepsLength = pos{0, 0}
	p.direction = d
	p.nextPos = addPosDirection(d, p.curPos)
	p.prevPos = p.curPos
}

func (p *player) endMove() {
	p.curPos = p.nextPos
	p.stepsLength = pos{0, 0}
	p.steps = 0
}

func (p *player) screenPos() (y, x float64) {
	x = float64(p.curPos.x*stageBlocSize + p.stepsLength.x)
	y = float64(p.curPos.y*stageBlocSize + p.stepsLength.y)
	return
}

func (p *player) resetPlayer() {
	p.curPos, p.prevPos, p.nextPos = p.initialPos, p.initialPos, p.initialPos
	p.currentImg = 0
	p.lost = false
	p.stepsLength = pos{0, 0}
	p.steps = 0
}

/*REINIT*/
func (p *player) reinit() {
	p.lives = 3
	p.score = 0
	p.resetPlayer()
	p.won = false
}
