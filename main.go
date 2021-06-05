package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"log"
	"os"

	pacman "github.com/UlisesBojorquez/PacmanGo/pacman"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	enemiesNumber := flag.Int("enemies", 4, "an int")
	flag.Parse()

	if *enemiesNumber < 1 {
		fmt.Println("Few Enemies: Select more than 0 and less than 5 enemies")
		os.Exit(0)

	} else if *enemiesNumber > 4 {
		fmt.Println("Too much enemies: Select more than 0 and less than 5 enemies")
		os.Exit(0)
	}

	g := pacman.NewGame(*enemiesNumber) //create new game, pacman is the dir and NewGame is in game.go

	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 1, "Pacman"); err != nil { //2
		log.Fatal(err)
	}
}
