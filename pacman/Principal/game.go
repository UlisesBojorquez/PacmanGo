package Principal

import (
	"github.com/hajimehoshi/ebiten"
)

// Game holds all the pacman game data
type Game struct {
	scene *scene //contains a scene inside the game
	input input  //used to control the keys of the player
}

//position struct
type pos struct {
	y, x int
}

// NewGame is a Game constructor
func NewGame(enemies int) *Game {
	g := &Game{}
	g.scene = newScene(nil, enemies) //create an empty scene

	return g
}

// ScreenWidth returns the game screen width
func (g *Game) ScreenWidth() int {
	return g.scene.screenWidth()
}

// ScreenHeight returns the game screen height
func (g *Game) ScreenHeight() int {
	return g.scene.screenHeight()
}

// Update updates the screen
func (g *Game) Update(screen *ebiten.Image) error {
	g.input = keyPressed()
	return g.scene.update(screen, g.input)
}
