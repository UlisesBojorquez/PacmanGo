package pacman

import (
	"github.com/hajimehoshi/ebiten"

	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//pacimages "github.com/kgosse/pacmanresources/images"
	//pacimages "github.com/UlisesBojorquez/pacmanresources/tree/master/images"
	pacimages "github.com/UlisesBojorquez/PacmanGo/images"
)

type scene struct {
	matrix        [][]elem //matrix stage
	wallSurface   *ebiten.Image
	images        map[elem]*ebiten.Image
	stage         *stage //this is the map walls array
	dotManager    *dotManager
	bigDotManager *bigDotManager
	player        *player
}

//Create a new Scene
func newScene(st *stage) *scene {
	s := &scene{} //create the structure pointer scene
	s.stage = st
	if s.stage == nil {
		s.stage = defaultStage //we assign the default stage for walls from stage.go
	}
	s.images = make(map[elem]*ebiten.Image) //we create the map from images
	s.dotManager = newDotManager()          //initialice the dot
	s.bigDotManager = newBigDotManager()    //initialice the bigdot
	s.loadImages()                          //initialice the image attribute
	s.createStage()                         //initialice matrix of elems
	s.buildWallSurface()                    //initialice wall surface, paint it

	return s //return the pointer structure scene
}

//it works to show things in the screen
func (s *scene) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() { // when IsDrawingSkipped is true, the rendered result is not adopted.
		return nil
	}
	screen.Clear()
	screen.DrawImage(s.wallSurface, nil)
	s.dotManager.draw(screen)    //paint the dots on screen
	s.bigDotManager.draw(screen) //paint the bigdots on screen
	s.player.draw(screen)
	//ebitenutil.DebugPrint(screen, "Hello World") // show in the screen what we see
	return nil
}

func (s *scene) screenWidth() int {
	w := len(s.stage.matrix[0])
	return w * stageBlocSize
}
func (s *scene) screenHeight() int {
	h := len(s.stage.matrix)
	return h * stageBlocSize
}

func (s *scene) createStage() {
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
			}
		}
	}

}

func (s *scene) buildWallSurface() {
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])

	sizeW := ((w*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH := ((h*stageBlocSize)/backgroundImageSize + 2) * backgroundImageSize
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
		s.images[i] = loadImage(pacimages.WallImages[i])
	}
	s.images[backgroundElem] = loadImage(pacimages.Background_png)
}
