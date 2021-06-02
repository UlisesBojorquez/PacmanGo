package pacman

import (
	"math"

	pacimages "github.com/UlisesBojorquez/PacmanGo/images"
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
	gm.images[blinkyElem] = loadGhostImages(pacimages.BlinkyImages)
	gm.images[clydeElem] = loadGhostImages(pacimages.ClydeImages)
	gm.images[inkyElem] = loadGhostImages(pacimages.InkyImages)
	gm.images[pinkyElem] = loadGhostImages(pacimages.PinkyImages)
	copy(gm.vulnerabilityImages[:], loadImages(pacimages.VulnerabilityImages[:]))
}

func (gm *ghostManager) addGhost(y, x int, e elem) {
	gm.ghosts = append(gm.ghosts, newGhost(y, x, e))
}

func (gm *ghostManager) draw(screen *ebiten.Image) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		imgs, _ := gm.images[g.kind]
		images := make([]*ebiten.Image, 13)
		copy(images, imgs[:])
		copy(images[8:], gm.vulnerabilityImages[:])
		g.draw(screen, images)
	}
}

func loadGhostImages(g [8][]byte) [8]*ebiten.Image {
	var arr [8]*ebiten.Image
	copy(arr[:], loadImages(g[:]))
	return arr
}

func (gm *ghostManager) move(m [][]elem, pp pos) {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		if !g.isMoving() {
			g.findNextMove(m, pp)
		}
		g.move()
	}

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
		gm.ghosts[i].makeVulnerable()
	}
}

func (gm *ghostManager) detectCollision(yPosPlayer, xPosPlayer float64, cb func(bool, float64, float64)) {
	for i := 0; i < len(gm.ghosts); i++ {
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
}

func (gm *ghostManager) resetGhostManager() {
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		g.resetGhost()

	}
}
