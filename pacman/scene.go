package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//pacimages "github.com/kgosse/pacmanresources/images"
)

type scene struct{}

//Create a new Scene
func newScene() *scene {
	s := &scene{} //create the structure pointer scene

	return s //return the pointer structure scene
}

//it works to show things in the screen
func (s *scene) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() { // when IsDrawingSkipped is true, the rendered result is not adopted.
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello World") // show in the screen what we see
	return nil
}
