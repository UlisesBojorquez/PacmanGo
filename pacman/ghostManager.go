package pacman

import (
	"math"

	pacmanimages "github.com/UlisesBojorquez/PacmanGo/images"
	"github.com/hajimehoshi/ebiten"
)

type ghostManager struct {
	ghosts              []*ghost
	images              map[elem][8]*ebiten.Image
	vulnerabilityImages [5]*ebiten.Image
	eaten               int //number of ghost eaten
}

func newGhostManager() *ghostManager {
	gm := &ghostManager{}
	gm.images = make(map[elem][8]*ebiten.Image)
	gm.loadImages()
	return gm
}

func (gm *ghostManager) loadImages() {
	gm.images[blinkyElem] = loadGhostImages(pacmanimages.BlinkyImages)
	gm.images[clydeElem] = loadGhostImages(pacmanimages.ClydeImages)
	gm.images[inkyElem] = loadGhostImages(pacmanimages.InkyImages)
	gm.images[pinkyElem] = loadGhostImages(pacmanimages.PinkyImages)
	copy(gm.vulnerabilityImages[:], loadImages(pacmanimages.VulnerabilityImages[:]))
}

func (gm *ghostManager) addGhost(y, x int, e elem) {
	gm.ghosts = append(gm.ghosts, newGhost(y, x, e))
}

func (gm *ghostManager) draw(screen *ebiten.Image) {
	for i := 0; i < len(gm.ghosts); i++ {
		go gm.drawAGhost(screen, i)
	}
}

// Draw a the slected ghost to the given screen
func (gm *ghostManager) drawAGhost(screen *ebiten.Image, i int) {
	g := gm.ghosts[i]
	imgs, _ := gm.images[g.kind]
	images := make([]*ebiten.Image, 13)
	copy(images, imgs[:])
	copy(images[8:], gm.vulnerabilityImages[:])
	g.draw(screen, images)
}

func loadGhostImages(g [8][]byte) [8]*ebiten.Image {
	var arr [8]*ebiten.Image
	copy(arr[:], loadImages(g[:]))
	return arr
}

func (gm *ghostManager) move(m [][]elem, pp pos) {
	for i := 0; i < len(gm.ghosts); i++ {
		go gm.moveAGhost(m, pp, i) // Move a ghost by using concurrency
	}
}

// This function moves a single ghost
func (gm *ghostManager) moveAGhost(m [][]elem, pp pos, i int) {
	g := gm.ghosts[i]
	if !g.isMoving() {
		g.findNextMove(m, pp)
	}
	g.move()
}

func (g *ghost) endMove() {
	g.prevPos = g.curPos
	g.curPos = g.nextPos
	g.stepsLength = pos{0, 0}
	g.steps = 0
}

/*COLLISION*/
func (gm *ghostManager) makeVulnerable() {
	gm.eaten = 0
	for i := 0; i < len(gm.ghosts); i++ {
		go gm.makeGhostVulnerable(i)
	}
}

// Make a given ghost vulnerable
func (gm *ghostManager) makeGhostVulnerable(i int) {
	gm.ghosts[i].makeVulnerable()
}

func (gm *ghostManager) detectCollision(yPosPlayer, xPosPlayer float64, cb func(bool, float64, float64)) {
	for i := 0; i < len(gm.ghosts); i++ {
		go gm.detectGhostCollision(yPosPlayer, xPosPlayer, cb, i)
	}
}

// Detect a colision for a single ghost
func (gm *ghostManager) detectGhostCollision(yPosPlayer, xPosPlayer float64, cb func(bool, float64, float64), i int) {
	g := gm.ghosts[i]
	yPosGhost, xPosGhost := g.screenPos()
	if math.Abs(yPosPlayer-yPosGhost) < 32 && math.Abs(xPosPlayer-xPosGhost) < 32 {
		if !g.isVulnerable() {
			cb(false, 0, 0)
			return
		}
		gm.eaten++
		g.makeEaten()
		g.resetGhost()
		cb(true, yPosGhost, xPosGhost)
	}
}

//Check if the ghost still playing or the player has won
func (gm *ghostManager) resetGhostManager(won bool) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		if !won {
			g.resetGhost()
		} else {
			g.killGhost()
		}

	}
}

/*REINIT*/
func (gm *ghostManager) reinit() {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		g.reinit()
	}
}
