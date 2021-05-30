package main

import (
	"log"

	//__"image/png"

	"github.com/UlisesBojorquez/PacmanGo/pacman"
	"github.com/hajimehoshi/ebiten"
)

func main() {

	g := pacman.NewGame() //create new game, pacman is the dir and NewGame is in game.go

	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 1, "Pacman"); err != nil { //2
		log.Fatal(err)
	}

}

/*
package main

import (
	"log"

	"github.com/UlisesBojorquez/pacman"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{} //Implements ebitent.Game for the ebiten interface

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	/*ebiten.SetWindowSize(640, 480)                  //assign the size of thr screen
	ebiten.SetWindowTitle("Hello, World!")          //assign the title
	if err := ebiten.RunGame(&Game{}); err != nil { //run the game
		log.Fatal(err)
	}

}*/
