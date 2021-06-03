package pacman

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"

	"github.com/UlisesBojorquez/PacmanGo/fonts"
	pacimages "github.com/UlisesBojorquez/PacmanGo/images"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const (
	keyText     = "KEYS"
	rText       = "r: Restart"
	hText       = "aswd: Move"
	livesText   = "LIVES"
	scoreText   = "SCORE"
	restartText = "R: Restart"
	moveText    = "←↓↑→: Move"
	pauseText   = "P: pause"
)

var (
	arialbdFontTitle font.Face
	arialbdFontBody  font.Face
	gold             = color.RGBA{255, 204, 0, 255}
)

type textManager struct {
	titleFF              font.Face
	bodyFF               font.Face
	entranceFF           font.Face
	keyX, livesX, scoreX int
	titleY               int
	count                int
	entrance             bool
	gameOverImage        *ebiten.Image
	gameOverAlpha        float64
	winImage             *ebiten.Image
	winAlpha             float64
}

func newTextManager(w, h int) *textManager {
	tm := &textManager{}
	tt, err := truetype.Parse(fonts.Arialbd_ttf)
	if err != nil {
		log.Fatal(err)
	}

	tm.titleFF = truetype.NewFace(tt, &truetype.Options{
		Size: 24,
	})
	tm.bodyFF = truetype.NewFace(tt, &truetype.Options{
		Size: 14,
	})
	tm.entranceFF = truetype.NewFace(tt, &truetype.Options{
		Size: 70,
	})

	tm.scoreX = w - 5*stageBlocSize
	tm.keyX = 20
	tm.livesX = w/2 - 2*stageBlocSize
	tm.titleY = h + 25

	tm.gameOverImage = loadImage(pacimages.GameOver_png[:])
	tm.winImage = loadImage(pacimages.Congrats_png[:])
	return tm
}

func (tm *textManager) draw(screen *ebiten.Image, score, lives int, won bool) {

	text.Draw(screen, keyText, tm.titleFF, tm.keyX, tm.titleY, gold)
	text.Draw(screen, rText, tm.bodyFF, tm.keyX, tm.titleY+stageBlocSize, gold)
	text.Draw(screen, hText, tm.bodyFF, tm.keyX, tm.titleY+2*stageBlocSize, gold)
	text.Draw(screen, moveText, tm.bodyFF, tm.keyX, tm.titleY+3*stageBlocSize, gold)
	text.Draw(screen, livesText, tm.titleFF, tm.livesX, tm.titleY, gold)
	text.Draw(screen, strconv.Itoa(lives), tm.titleFF, tm.livesX, tm.titleY+1*stageBlocSize-9, gold)
	text.Draw(screen, scoreText, tm.titleFF, tm.scoreX, tm.titleY, gold)
	text.Draw(screen, strconv.Itoa(score), tm.titleFF, tm.scoreX, tm.titleY+1*stageBlocSize-9, gold)

	if lives == 0 {
		tm.showGameOverImage(screen)

	} else if won {
		tm.showWinImage(screen)
	}

	tm.showEntranceAnimation(screen)
}

func (tm *textManager) entranceAnimation(b bool) {
	if b {
		tm.count = 0
	}
	tm.entrance = b
}

func (tm *textManager) showEntranceAnimation(screen *ebiten.Image) {
	if !tm.entrance {
		return
	}
	tm.count++

	if tm.count <= 60 {
		text.Draw(screen, "3", tm.entranceFF, 9*stageBlocSize, 5*stageBlocSize, gold)
	} else if tm.count <= 120 {
		text.Draw(screen, "2", tm.entranceFF, 9*stageBlocSize, 5*stageBlocSize, gold)
	} else if tm.count <= 180 {
		text.Draw(screen, "1", tm.entranceFF, 9*stageBlocSize, 5*stageBlocSize, gold)
	} else if tm.count <= 240 {
		text.Draw(screen, "GO!", tm.entranceFF, 7*stageBlocSize, 5*stageBlocSize, gold)
	} else {
		tm.entranceAnimation(false)
	}
}

func (tm *textManager) reinit() {
	tm.gameOverAlpha = 0
	tm.winAlpha = 0
}

func (tm *textManager) showGameOverImage(screen *ebiten.Image) {
	tm.gameOverAlpha += 0.01
	if tm.gameOverAlpha > 1 {
		tm.gameOverAlpha = 1
	}
	x := float64(3 * stageBlocSize)
	y := float64(4 * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.ColorM.Scale(1, 1, 1, tm.gameOverAlpha)
	screen.DrawImage(tm.gameOverImage, op)
}

func (tm *textManager) showWinImage(screen *ebiten.Image) {
	tm.winAlpha += 0.01
	if tm.winAlpha > 1 {
		tm.winAlpha = 1
	}
	x := float64(8)
	y := float64(4 * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.ColorM.Scale(1, 1, 1, tm.winAlpha)
	screen.DrawImage(tm.winImage, op)
}
