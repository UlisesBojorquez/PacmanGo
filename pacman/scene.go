package pacman

import (
	"github.com/hajimehoshi/ebiten"

	pacmanimages "github.com/UlisesBojorquez/PacmanGo/images"
)

type scene struct {
	matrix        [][]elem //matrix stage
	wallSurface   *ebiten.Image
	images        map[elem]*ebiten.Image
	stage         *stage         //this is the map walls array
	dotManager    *dotManager    //this is the dot manager
	bigDotManager *bigDotManager //this is the big dot manager
	player        *player        //this is the player
	ghostManager  *ghostManager  //this is the ghost manager
	textManager   *textManager
	sounds        *sounds
	over          bool
	enemiesNumb   int //number of ghost desired by the player
}

//Create a new Scene
func newScene(st *stage, enemies int) *scene {
	s := &scene{} //create the structure pointer scene
	s.stage = st
	if s.stage == nil {
		s.stage = defaultStage //we assign the default stage for walls from stage.go
	}
	s.images = make(map[elem]*ebiten.Image) //we create the map from images
	s.dotManager = newDotManager()          //initialice the dot
	s.bigDotManager = newBigDotManager()    //initialice the bigdot
	s.ghostManager = newGhostManager()      //initialice the ghostmanager
	s.sounds = newSounds()
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])
	s.textManager = newTextManager(w*stageBlocSize, h*stageBlocSize) //initilice the textmanager
	s.loadImages()                                                   //initialice the image attribute
	s.createStage(enemies)                                           //initialice matrix of elems
	s.buildWallSurface()                                             //initialice wall surface, paint it
	s.enemiesNumb = enemies

	s.textManager.entranceAnimation(true)
	s.sounds.playEntrance()
	return s //return the pointer structure scene
}

func (s *scene) move(in input) {
	if s.player.lives > 0 {
		s.player.move(s.matrix, in) //player movement
	}
	if !s.over {
		s.ghostManager.move(s.matrix, s.player.curPos) //the ghost movement
	}

}

//it works to show things in the screen
func (s *scene) update(screen *ebiten.Image, in input) error {
	if ebiten.IsDrawingSkipped() { // when IsDrawingSkipped is true, the rendered result is not adopted.
		return nil
	}

	if in == rKey && !s.player.won {
		s.reinit()
	} else if !s.textManager.entrance {
		s.move(in) //movement
		s.detectCollision()
	}

	screen.Clear()

	screen.DrawImage(s.wallSurface, nil)
	s.dotManager.draw(screen)                                                //paint the dots on screen
	s.bigDotManager.draw(screen)                                             //paint the bigdots on screen
	s.player.draw(screen)                                                    //paint the player
	s.ghostManager.draw(screen)                                              //paint the ghosts
	s.textManager.draw(screen, s.player.score, s.player.lives, s.player.won) //paint the text
	s.sounds.playSiren()
	return nil
}

func (s *scene) screenWidth() int {
	w := len(s.stage.matrix[0])
	return w * stageBlocSize
}
func (s *scene) screenHeight() int {
	h := len(s.stage.matrix)
	sizeH := ((h*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	return sizeH
}

func (s *scene) createStage(enemies int) {
	h := len(s.stage.matrix)     //altura
	w := len(s.stage.matrix[0])  //ancho
	s.matrix = make([][]elem, h) //we create the matrix with the number of rows
	for i := 0; i < h; i++ {
		s.matrix[i] = make([]elem, w)
		for j := 0; j < w; j++ {
			/*PART FOR THE BORDERS*/
			c := s.stage.matrix[i][j] - '0' //here we get the decimal representation for example char 3 is 51 in decimal and 0 is 58 as a result we have 3
			if c <= 9 {                     //used for numebers
				s.matrix[i][j] = elem(c)
			} else { //the rest of our constans
				s.matrix[i][j] = elem(s.stage.matrix[i][j] - 'a' + 10) //for example for 10 is char a is 97 in decimal minus char a which is 97 +10 give is 10
			}

			/*PART TO ADD THE REST*/
			switch s.matrix[i][j] {
			case dotElem:
				s.dotManager.add(i, j)
			case bigDotElem:
				s.bigDotManager.add(i, j)
			case playerElem:
				s.player = newPlayer(i, j)
			case blinkyElem:
				if enemies >= 1 {
					s.ghostManager.addGhost(i, j, blinkyElem)
				}
			case inkyElem:
				if enemies >= 2 {
					s.ghostManager.addGhost(i, j, inkyElem)
				}
			case pinkyElem:
				if enemies >= 3 {
					s.ghostManager.addGhost(i, j, pinkyElem)
				}
			case clydeElem:
				if enemies == 4 {
					s.ghostManager.addGhost(i, j, clydeElem)
				}
			}
		}
	}

}

func (s *scene) buildWallSurface() {
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])

	sizeW := ((w*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH := ((h*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	s.wallSurface, _ = ebiten.NewImage(sizeW, sizeH, ebiten.FilterDefault)

	for i := 0; i < sizeH/backgroundImageSize; i++ {
		y := float64(i * backgroundImageSize)
		for j := 0; j < sizeW/backgroundImageSize; j++ {
			op := &ebiten.DrawImageOptions{}
			x := float64(j * backgroundImageSize)
			op.GeoM.Translate(x, y)
			s.wallSurface.DrawImage(s.images[backgroundElem], op)
		}
	}

	for i := 0; i < h; i++ {
		y := float64(i * stageBlocSize)
		for j := 0; j < w; j++ {
			if !isWall(s.matrix[i][j]) {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			x := float64(j * stageBlocSize)
			op.GeoM.Translate(x, y)
			s.wallSurface.DrawImage(s.images[s.matrix[i][j]], op)
		}
	}
}

func (s *scene) loadImages() {
	for i := w0; i <= w24; i++ {
		s.images[i] = loadImage(pacmanimages.WallImages[i])
	}
	s.images[backgroundElem] = loadImage(pacmanimages.Background_png)
}

/*COLLISION*/

func (s *scene) detectCollision() {
	//detect the collision between pacman and dot
	s.dotManager.detectCollision(s.matrix, s.player.curPos, s.afterPacmanDotCollision)
	//detect collision between pacman and bid dots
	s.bigDotManager.detectCollision(s.matrix, s.player.curPos, s.afterPacmanBigDotCollision)
	//detect colission between pacman and ghosts
	yPosPlayer, xPosPlayer := s.player.screenPos()
	s.ghostManager.detectCollision(yPosPlayer, xPosPlayer, s.afterPacmanGhostCollision)

}

func (s *scene) afterPacmanDotCollision() {
	s.player.score += 10
	s.dotManager.delete(s.player.curPos)
	s.matrix[s.player.curPos.y][s.player.curPos.x] = empty
	if !s.over && s.won() {
		s.victory()
	}
}

func (s *scene) afterPacmanBigDotCollision() {
	s.player.score += 50
	s.bigDotManager.delete(s.player.curPos)
	s.matrix[s.player.curPos.y][s.player.curPos.x] = empty
	if !s.over && s.won() {
		s.victory()
		return
	}
	s.ghostManager.makeVulnerable()
	s.sounds.playWail()
}

func (s *scene) afterPacmanGhostCollision(vulnerable bool, y, x float64) {
	if vulnerable {
		eaten := s.ghostManager.eaten
		if eaten == 1 {
			s.player.score += 200
		} else if eaten == 2 {
			s.player.score += 400
		} else if eaten == 3 {
			s.player.score += 800
		} else {
			s.player.score += 1600
		}
		s.sounds.playEeatGhost()

	} else {
		if s.player.lives > 1 {
			s.player.lives--
			s.player.resetPlayer()
			s.ghostManager.resetGhostManager(false)
			s.sounds.playDeath()
		} else {
			if !s.player.lost {
				s.GameOver()
			}
		}

	}
}

/*GAMEOVER*/
func (s *scene) GameOver() {
	s.player.lives--
	s.player.lost = true
	s.player.curPos.x = 0
	s.player.curPos.y = 0
	s.sounds.playDeath()
}

func (s *scene) won() bool {
	if s.dotManager.empty() && s.bigDotManager.empty() {
		return true
	}
	return false
}
func (s *scene) victory() {
	s.over = true
	s.player.won = true
	s.ghostManager.resetGhostManager(true)
	s.sounds.pause()
	s.sounds.playApplause()
}

/*REINIT*/
func (s *scene) reinit() {
	s.dotManager.reinit(s.matrix)
	s.bigDotManager.reinit(s.matrix)
	s.player.reinit()
	s.textManager.reinit()
	s.over = false
	s.ghostManager.reinit()
	s.textManager.entranceAnimation(true)
	s.sounds.pause()
	s.sounds.playEntrance()
}
